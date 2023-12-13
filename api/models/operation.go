package models

type Operation struct {
	Name        string
	Description string
	Method      string
	Href        string
	HasPayload  bool
	Payload     interface{}
	RequestType string
}
