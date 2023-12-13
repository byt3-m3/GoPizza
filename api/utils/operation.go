package utils

import (
	apiModels "GoPizza/api/models"
	"fmt"
	"github.com/google/uuid"
	"reflect"
)

func BuildCreateOrderOperation(customerID uuid.UUID) apiModels.Operation {
	return apiModels.Operation{
		Name:        "v1:create_order",
		Description: "Create Order for customer",
		Method:      "POST",
		Href:        "/v1/order",
		HasPayload:  true,
		Payload: apiModels.NewOrderRequest{
			CustomerID: customerID,
			StoreID:    uuid.UUID{},
		},
		RequestType: reflect.TypeOf(apiModels.NewOrderRequest{}).String(),
	}
}

func BuildGetCustomerOperation(customerID uuid.UUID) apiModels.Operation {
	return apiModels.Operation{
		Name:        "v1:get_customer",
		Description: "fetches required customer",
		Method:      "GET",
		Href:        fmt.Sprintf("/customer?id=%s", customerID.String()),
		Payload:     nil,
		RequestType: "nil",
		HasPayload:  false,
	}

}
