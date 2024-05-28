// Code generated by go-swagger; DO NOT EDIT.

package open_stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// ServiceBrokerOpenstacksServersGetReader is a Reader for the ServiceBrokerOpenstacksServersGet structure.
type ServiceBrokerOpenstacksServersGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceBrokerOpenstacksServersGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceBrokerOpenstacksServersGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewServiceBrokerOpenstacksServersGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewServiceBrokerOpenstacksServersGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewServiceBrokerOpenstacksServersGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewServiceBrokerOpenstacksServersGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewServiceBrokerOpenstacksServersGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /broker/v1/openstacks/{openstack_id}/servers/{pvm_instance_id}] serviceBroker.openstacks.servers.get", response, response.Code())
	}
}

// NewServiceBrokerOpenstacksServersGetOK creates a ServiceBrokerOpenstacksServersGetOK with default headers values
func NewServiceBrokerOpenstacksServersGetOK() *ServiceBrokerOpenstacksServersGetOK {
	return &ServiceBrokerOpenstacksServersGetOK{}
}

/*
ServiceBrokerOpenstacksServersGetOK describes a response with status code 200, with default header values.

OK
*/
type ServiceBrokerOpenstacksServersGetOK struct {
	Payload *models.HostPVMInstance
}

// IsSuccess returns true when this service broker openstacks servers get o k response has a 2xx status code
func (o *ServiceBrokerOpenstacksServersGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service broker openstacks servers get o k response has a 3xx status code
func (o *ServiceBrokerOpenstacksServersGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker openstacks servers get o k response has a 4xx status code
func (o *ServiceBrokerOpenstacksServersGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service broker openstacks servers get o k response has a 5xx status code
func (o *ServiceBrokerOpenstacksServersGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker openstacks servers get o k response a status code equal to that given
func (o *ServiceBrokerOpenstacksServersGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the service broker openstacks servers get o k response
func (o *ServiceBrokerOpenstacksServersGetOK) Code() int {
	return 200
}

func (o *ServiceBrokerOpenstacksServersGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /broker/v1/openstacks/{openstack_id}/servers/{pvm_instance_id}][%d] serviceBrokerOpenstacksServersGetOK %s", 200, payload)
}

func (o *ServiceBrokerOpenstacksServersGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /broker/v1/openstacks/{openstack_id}/servers/{pvm_instance_id}][%d] serviceBrokerOpenstacksServersGetOK %s", 200, payload)
}

func (o *ServiceBrokerOpenstacksServersGetOK) GetPayload() *models.HostPVMInstance {
	return o.Payload
}

func (o *ServiceBrokerOpenstacksServersGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HostPVMInstance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerOpenstacksServersGetBadRequest creates a ServiceBrokerOpenstacksServersGetBadRequest with default headers values
func NewServiceBrokerOpenstacksServersGetBadRequest() *ServiceBrokerOpenstacksServersGetBadRequest {
	return &ServiceBrokerOpenstacksServersGetBadRequest{}
}

/*
ServiceBrokerOpenstacksServersGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ServiceBrokerOpenstacksServersGetBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker openstacks servers get bad request response has a 2xx status code
func (o *ServiceBrokerOpenstacksServersGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker openstacks servers get bad request response has a 3xx status code
func (o *ServiceBrokerOpenstacksServersGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker openstacks servers get bad request response has a 4xx status code
func (o *ServiceBrokerOpenstacksServersGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker openstacks servers get bad request response has a 5xx status code
func (o *ServiceBrokerOpenstacksServersGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker openstacks servers get bad request response a status code equal to that given
func (o *ServiceBrokerOpenstacksServersGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the service broker openstacks servers get bad request response
func (o *ServiceBrokerOpenstacksServersGetBadRequest) Code() int {
	return 400
}

func (o *ServiceBrokerOpenstacksServersGetBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /broker/v1/openstacks/{openstack_id}/servers/{pvm_instance_id}][%d] serviceBrokerOpenstacksServersGetBadRequest %s", 400, payload)
}

func (o *ServiceBrokerOpenstacksServersGetBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /broker/v1/openstacks/{openstack_id}/servers/{pvm_instance_id}][%d] serviceBrokerOpenstacksServersGetBadRequest %s", 400, payload)
}

func (o *ServiceBrokerOpenstacksServersGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerOpenstacksServersGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerOpenstacksServersGetUnauthorized creates a ServiceBrokerOpenstacksServersGetUnauthorized with default headers values
func NewServiceBrokerOpenstacksServersGetUnauthorized() *ServiceBrokerOpenstacksServersGetUnauthorized {
	return &ServiceBrokerOpenstacksServersGetUnauthorized{}
}

/*
ServiceBrokerOpenstacksServersGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ServiceBrokerOpenstacksServersGetUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker openstacks servers get unauthorized response has a 2xx status code
func (o *ServiceBrokerOpenstacksServersGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker openstacks servers get unauthorized response has a 3xx status code
func (o *ServiceBrokerOpenstacksServersGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker openstacks servers get unauthorized response has a 4xx status code
func (o *ServiceBrokerOpenstacksServersGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker openstacks servers get unauthorized response has a 5xx status code
func (o *ServiceBrokerOpenstacksServersGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker openstacks servers get unauthorized response a status code equal to that given
func (o *ServiceBrokerOpenstacksServersGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the service broker openstacks servers get unauthorized response
func (o *ServiceBrokerOpenstacksServersGetUnauthorized) Code() int {
	return 401
}

func (o *ServiceBrokerOpenstacksServersGetUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /broker/v1/openstacks/{openstack_id}/servers/{pvm_instance_id}][%d] serviceBrokerOpenstacksServersGetUnauthorized %s", 401, payload)
}

func (o *ServiceBrokerOpenstacksServersGetUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /broker/v1/openstacks/{openstack_id}/servers/{pvm_instance_id}][%d] serviceBrokerOpenstacksServersGetUnauthorized %s", 401, payload)
}

func (o *ServiceBrokerOpenstacksServersGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerOpenstacksServersGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerOpenstacksServersGetForbidden creates a ServiceBrokerOpenstacksServersGetForbidden with default headers values
func NewServiceBrokerOpenstacksServersGetForbidden() *ServiceBrokerOpenstacksServersGetForbidden {
	return &ServiceBrokerOpenstacksServersGetForbidden{}
}

/*
ServiceBrokerOpenstacksServersGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ServiceBrokerOpenstacksServersGetForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker openstacks servers get forbidden response has a 2xx status code
func (o *ServiceBrokerOpenstacksServersGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker openstacks servers get forbidden response has a 3xx status code
func (o *ServiceBrokerOpenstacksServersGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker openstacks servers get forbidden response has a 4xx status code
func (o *ServiceBrokerOpenstacksServersGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker openstacks servers get forbidden response has a 5xx status code
func (o *ServiceBrokerOpenstacksServersGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker openstacks servers get forbidden response a status code equal to that given
func (o *ServiceBrokerOpenstacksServersGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the service broker openstacks servers get forbidden response
func (o *ServiceBrokerOpenstacksServersGetForbidden) Code() int {
	return 403
}

func (o *ServiceBrokerOpenstacksServersGetForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /broker/v1/openstacks/{openstack_id}/servers/{pvm_instance_id}][%d] serviceBrokerOpenstacksServersGetForbidden %s", 403, payload)
}

func (o *ServiceBrokerOpenstacksServersGetForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /broker/v1/openstacks/{openstack_id}/servers/{pvm_instance_id}][%d] serviceBrokerOpenstacksServersGetForbidden %s", 403, payload)
}

func (o *ServiceBrokerOpenstacksServersGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerOpenstacksServersGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerOpenstacksServersGetNotFound creates a ServiceBrokerOpenstacksServersGetNotFound with default headers values
func NewServiceBrokerOpenstacksServersGetNotFound() *ServiceBrokerOpenstacksServersGetNotFound {
	return &ServiceBrokerOpenstacksServersGetNotFound{}
}

/*
ServiceBrokerOpenstacksServersGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ServiceBrokerOpenstacksServersGetNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker openstacks servers get not found response has a 2xx status code
func (o *ServiceBrokerOpenstacksServersGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker openstacks servers get not found response has a 3xx status code
func (o *ServiceBrokerOpenstacksServersGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker openstacks servers get not found response has a 4xx status code
func (o *ServiceBrokerOpenstacksServersGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker openstacks servers get not found response has a 5xx status code
func (o *ServiceBrokerOpenstacksServersGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker openstacks servers get not found response a status code equal to that given
func (o *ServiceBrokerOpenstacksServersGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the service broker openstacks servers get not found response
func (o *ServiceBrokerOpenstacksServersGetNotFound) Code() int {
	return 404
}

func (o *ServiceBrokerOpenstacksServersGetNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /broker/v1/openstacks/{openstack_id}/servers/{pvm_instance_id}][%d] serviceBrokerOpenstacksServersGetNotFound %s", 404, payload)
}

func (o *ServiceBrokerOpenstacksServersGetNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /broker/v1/openstacks/{openstack_id}/servers/{pvm_instance_id}][%d] serviceBrokerOpenstacksServersGetNotFound %s", 404, payload)
}

func (o *ServiceBrokerOpenstacksServersGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerOpenstacksServersGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerOpenstacksServersGetInternalServerError creates a ServiceBrokerOpenstacksServersGetInternalServerError with default headers values
func NewServiceBrokerOpenstacksServersGetInternalServerError() *ServiceBrokerOpenstacksServersGetInternalServerError {
	return &ServiceBrokerOpenstacksServersGetInternalServerError{}
}

/*
ServiceBrokerOpenstacksServersGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ServiceBrokerOpenstacksServersGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker openstacks servers get internal server error response has a 2xx status code
func (o *ServiceBrokerOpenstacksServersGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker openstacks servers get internal server error response has a 3xx status code
func (o *ServiceBrokerOpenstacksServersGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker openstacks servers get internal server error response has a 4xx status code
func (o *ServiceBrokerOpenstacksServersGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this service broker openstacks servers get internal server error response has a 5xx status code
func (o *ServiceBrokerOpenstacksServersGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this service broker openstacks servers get internal server error response a status code equal to that given
func (o *ServiceBrokerOpenstacksServersGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the service broker openstacks servers get internal server error response
func (o *ServiceBrokerOpenstacksServersGetInternalServerError) Code() int {
	return 500
}

func (o *ServiceBrokerOpenstacksServersGetInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /broker/v1/openstacks/{openstack_id}/servers/{pvm_instance_id}][%d] serviceBrokerOpenstacksServersGetInternalServerError %s", 500, payload)
}

func (o *ServiceBrokerOpenstacksServersGetInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /broker/v1/openstacks/{openstack_id}/servers/{pvm_instance_id}][%d] serviceBrokerOpenstacksServersGetInternalServerError %s", 500, payload)
}

func (o *ServiceBrokerOpenstacksServersGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerOpenstacksServersGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
