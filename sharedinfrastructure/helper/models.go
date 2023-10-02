package helper

import (
	"github.com/google/uuid"
)

type Response struct {
	Data interface{} `json:"data"`
}
type SuccessBody struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	Reference string `json:"reference"`
}

type AllSuccessBody struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Service interface{} `json:"service"`
}
type ErrorResponse struct {
	TimeStamp      string    `bson:"timestamp" json:"timestamp"`
	ErrorReference uuid.UUID `bson:"error_reference" json:"error_reference"`
	Errors         []ErrorBody
}
type ErrorBody struct {
	Code    string      `bson:"code" json:"code"`
	Message interface{} `bson:"message" json:"message"`
	Source  string      `bson:"source" json:"source"`
	Details string      `bson:"details" json:"details"`
}
