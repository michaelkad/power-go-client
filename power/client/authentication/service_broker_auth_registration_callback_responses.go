// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.ibm.com/power-iaas/service-broker/power/models"
)

// ServiceBrokerAuthRegistrationCallbackReader is a Reader for the ServiceBrokerAuthRegistrationCallback structure.
type ServiceBrokerAuthRegistrationCallbackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceBrokerAuthRegistrationCallbackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceBrokerAuthRegistrationCallbackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewServiceBrokerAuthRegistrationCallbackUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewServiceBrokerAuthRegistrationCallbackInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /auth/v1/callback-registration] serviceBroker.auth.registration.callback", response, response.Code())
	}
}

// NewServiceBrokerAuthRegistrationCallbackOK creates a ServiceBrokerAuthRegistrationCallbackOK with default headers values
func NewServiceBrokerAuthRegistrationCallbackOK() *ServiceBrokerAuthRegistrationCallbackOK {
	return &ServiceBrokerAuthRegistrationCallbackOK{}
}

/*
ServiceBrokerAuthRegistrationCallbackOK describes a response with status code 200, with default header values.

OK
*/
type ServiceBrokerAuthRegistrationCallbackOK struct {
	Payload *models.AccessToken
}

// IsSuccess returns true when this service broker auth registration callback o k response has a 2xx status code
func (o *ServiceBrokerAuthRegistrationCallbackOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service broker auth registration callback o k response has a 3xx status code
func (o *ServiceBrokerAuthRegistrationCallbackOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker auth registration callback o k response has a 4xx status code
func (o *ServiceBrokerAuthRegistrationCallbackOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service broker auth registration callback o k response has a 5xx status code
func (o *ServiceBrokerAuthRegistrationCallbackOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker auth registration callback o k response a status code equal to that given
func (o *ServiceBrokerAuthRegistrationCallbackOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the service broker auth registration callback o k response
func (o *ServiceBrokerAuthRegistrationCallbackOK) Code() int {
	return 200
}

func (o *ServiceBrokerAuthRegistrationCallbackOK) Error() string {
	return fmt.Sprintf("[GET /auth/v1/callback-registration][%d] serviceBrokerAuthRegistrationCallbackOK  %+v", 200, o.Payload)
}

func (o *ServiceBrokerAuthRegistrationCallbackOK) String() string {
	return fmt.Sprintf("[GET /auth/v1/callback-registration][%d] serviceBrokerAuthRegistrationCallbackOK  %+v", 200, o.Payload)
}

func (o *ServiceBrokerAuthRegistrationCallbackOK) GetPayload() *models.AccessToken {
	return o.Payload
}

func (o *ServiceBrokerAuthRegistrationCallbackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccessToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerAuthRegistrationCallbackUnauthorized creates a ServiceBrokerAuthRegistrationCallbackUnauthorized with default headers values
func NewServiceBrokerAuthRegistrationCallbackUnauthorized() *ServiceBrokerAuthRegistrationCallbackUnauthorized {
	return &ServiceBrokerAuthRegistrationCallbackUnauthorized{}
}

/*
ServiceBrokerAuthRegistrationCallbackUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ServiceBrokerAuthRegistrationCallbackUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker auth registration callback unauthorized response has a 2xx status code
func (o *ServiceBrokerAuthRegistrationCallbackUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker auth registration callback unauthorized response has a 3xx status code
func (o *ServiceBrokerAuthRegistrationCallbackUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker auth registration callback unauthorized response has a 4xx status code
func (o *ServiceBrokerAuthRegistrationCallbackUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker auth registration callback unauthorized response has a 5xx status code
func (o *ServiceBrokerAuthRegistrationCallbackUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker auth registration callback unauthorized response a status code equal to that given
func (o *ServiceBrokerAuthRegistrationCallbackUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the service broker auth registration callback unauthorized response
func (o *ServiceBrokerAuthRegistrationCallbackUnauthorized) Code() int {
	return 401
}

func (o *ServiceBrokerAuthRegistrationCallbackUnauthorized) Error() string {
	return fmt.Sprintf("[GET /auth/v1/callback-registration][%d] serviceBrokerAuthRegistrationCallbackUnauthorized  %+v", 401, o.Payload)
}

func (o *ServiceBrokerAuthRegistrationCallbackUnauthorized) String() string {
	return fmt.Sprintf("[GET /auth/v1/callback-registration][%d] serviceBrokerAuthRegistrationCallbackUnauthorized  %+v", 401, o.Payload)
}

func (o *ServiceBrokerAuthRegistrationCallbackUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerAuthRegistrationCallbackUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerAuthRegistrationCallbackInternalServerError creates a ServiceBrokerAuthRegistrationCallbackInternalServerError with default headers values
func NewServiceBrokerAuthRegistrationCallbackInternalServerError() *ServiceBrokerAuthRegistrationCallbackInternalServerError {
	return &ServiceBrokerAuthRegistrationCallbackInternalServerError{}
}

/*
ServiceBrokerAuthRegistrationCallbackInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ServiceBrokerAuthRegistrationCallbackInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker auth registration callback internal server error response has a 2xx status code
func (o *ServiceBrokerAuthRegistrationCallbackInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker auth registration callback internal server error response has a 3xx status code
func (o *ServiceBrokerAuthRegistrationCallbackInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker auth registration callback internal server error response has a 4xx status code
func (o *ServiceBrokerAuthRegistrationCallbackInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this service broker auth registration callback internal server error response has a 5xx status code
func (o *ServiceBrokerAuthRegistrationCallbackInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this service broker auth registration callback internal server error response a status code equal to that given
func (o *ServiceBrokerAuthRegistrationCallbackInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the service broker auth registration callback internal server error response
func (o *ServiceBrokerAuthRegistrationCallbackInternalServerError) Code() int {
	return 500
}

func (o *ServiceBrokerAuthRegistrationCallbackInternalServerError) Error() string {
	return fmt.Sprintf("[GET /auth/v1/callback-registration][%d] serviceBrokerAuthRegistrationCallbackInternalServerError  %+v", 500, o.Payload)
}

func (o *ServiceBrokerAuthRegistrationCallbackInternalServerError) String() string {
	return fmt.Sprintf("[GET /auth/v1/callback-registration][%d] serviceBrokerAuthRegistrationCallbackInternalServerError  %+v", 500, o.Payload)
}

func (o *ServiceBrokerAuthRegistrationCallbackInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerAuthRegistrationCallbackInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
