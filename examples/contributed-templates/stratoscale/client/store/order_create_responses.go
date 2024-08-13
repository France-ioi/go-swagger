// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/France-ioi/go-swagger/examples/contributed-templates/stratoscale/models"
)

// OrderCreateReader is a Reader for the OrderCreate structure.
type OrderCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrderCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrderCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOrderCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /store/order] OrderCreate", response, response.Code())
	}
}

// NewOrderCreateOK creates a OrderCreateOK with default headers values
func NewOrderCreateOK() *OrderCreateOK {
	return &OrderCreateOK{}
}

/*
OrderCreateOK describes a response with status code 200, with default header values.

successful operation
*/
type OrderCreateOK struct {
	Payload *models.Order
}

// IsSuccess returns true when this order create o k response has a 2xx status code
func (o *OrderCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this order create o k response has a 3xx status code
func (o *OrderCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this order create o k response has a 4xx status code
func (o *OrderCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this order create o k response has a 5xx status code
func (o *OrderCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this order create o k response a status code equal to that given
func (o *OrderCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the order create o k response
func (o *OrderCreateOK) Code() int {
	return 200
}

func (o *OrderCreateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /store/order][%d] orderCreateOK %s", 200, payload)
}

func (o *OrderCreateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /store/order][%d] orderCreateOK %s", 200, payload)
}

func (o *OrderCreateOK) GetPayload() *models.Order {
	return o.Payload
}

func (o *OrderCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Order)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrderCreateBadRequest creates a OrderCreateBadRequest with default headers values
func NewOrderCreateBadRequest() *OrderCreateBadRequest {
	return &OrderCreateBadRequest{}
}

/*
OrderCreateBadRequest describes a response with status code 400, with default header values.

Invalid Order
*/
type OrderCreateBadRequest struct {
}

// IsSuccess returns true when this order create bad request response has a 2xx status code
func (o *OrderCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this order create bad request response has a 3xx status code
func (o *OrderCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this order create bad request response has a 4xx status code
func (o *OrderCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this order create bad request response has a 5xx status code
func (o *OrderCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this order create bad request response a status code equal to that given
func (o *OrderCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the order create bad request response
func (o *OrderCreateBadRequest) Code() int {
	return 400
}

func (o *OrderCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /store/order][%d] orderCreateBadRequest", 400)
}

func (o *OrderCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /store/order][%d] orderCreateBadRequest", 400)
}

func (o *OrderCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
