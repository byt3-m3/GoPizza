package handlers

import (
	apiModels "GoPizza/api/models"
	"GoPizza/api/utils"
	coreModels "GoPizza/models"
	"GoPizza/repo"
	"encoding/json"
	"github.com/byt3-m3/goutils/http_utils"
	"github.com/google/uuid"
	"github.com/pmoule/go2hal/hal"
	"log/slog"
	"net/http"
)

type CustomerEndpointHandlerInput struct {
	Curries      []*hal.LinkObject
	CustomerRepo repo.CustomerRepo
}

func CustomerEndpointHandler(writer http.ResponseWriter, request *http.Request, input *CustomerEndpointHandlerInput) {
	halResponseRoot := hal.NewResourceObject()
	halResponseRoot.AddCurieLinks(input.Curries)

	if request.Method == "GET" {
		customerId := request.URL.Query().Get("id")

		customer, err := input.CustomerRepo.Get(uuid.MustParse(customerId))
		if err != nil {
			resp := apiModels.ErrorResponse{Msg: "customer lookup error", Error: err.Error()}
			_, _ = http_utils.WriteJSONFromAny(writer, resp, http.StatusConflict)
			return
		}

		response := apiModels.CustomerResponse{
			Customer:   customer,
			Operations: []apiModels.Operation{},
		}

		if customer.AccountBalance > 0 {

			response.Operations = append(response.Operations, utils.BuildCreateOrderOperation(customer.ID))
		}

		responseBytes, _ := json.Marshal(response)

		_, _ = http_utils.WriteJSONFromAny(writer, responseBytes, http.StatusOK)
		return

	}

	if request.Method == "POST" {

		var requestModel apiModels.CreateCustomerRequest

		if err := json.NewDecoder(request.Body).Decode(&requestModel); err != nil {

		}

		customerID := uuid.New()
		newCustomer := coreModels.NewCustomer(coreModels.NewCustomerInput{
			ID:               customerID,
			Name:             requestModel.Name,
			BeginningBalance: requestModel.StartingBalance,
		})

		if err := input.CustomerRepo.Save(newCustomer); err != nil {
			slog.Error("unable to save customer", err)
		}

		resp := apiModels.CustomerResponse{
			Customer: newCustomer,
			Operations: []apiModels.Operation{

				utils.BuildGetCustomerOperation(customerID),
			},
		}

		responseBytes, err := json.Marshal(&resp)
		if err != nil {
			slog.Error("error marshaling response", err)
		}

		slog.Info("successfully created customer",
			slog.Any("customer", newCustomer),
		)
		_, _ = http_utils.WriteJSONFromBytes(writer, responseBytes, http.StatusOK)

	}

}
