// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.ibm.com/power-iaas/service-broker/power/models"
)

// PcloudEventsGetReader is a Reader for the PcloudEventsGet structure.
type PcloudEventsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudEventsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudEventsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudEventsGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudEventsGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudEventsGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudEventsGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/events/{event_id}] pcloud.events.get", response, response.Code())
	}
}

// NewPcloudEventsGetOK creates a PcloudEventsGetOK with default headers values
func NewPcloudEventsGetOK() *PcloudEventsGetOK {
	return &PcloudEventsGetOK{}
}

/*
PcloudEventsGetOK describes a response with status code 200, with default header values.

OK
*/
type PcloudEventsGetOK struct {
	Payload *models.Event
}

// IsSuccess returns true when this pcloud events get o k response has a 2xx status code
func (o *PcloudEventsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud events get o k response has a 3xx status code
func (o *PcloudEventsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud events get o k response has a 4xx status code
func (o *PcloudEventsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud events get o k response has a 5xx status code
func (o *PcloudEventsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud events get o k response a status code equal to that given
func (o *PcloudEventsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud events get o k response
func (o *PcloudEventsGetOK) Code() int {
	return 200
}

func (o *PcloudEventsGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/events/{event_id}][%d] pcloudEventsGetOK  %+v", 200, o.Payload)
}

func (o *PcloudEventsGetOK) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/events/{event_id}][%d] pcloudEventsGetOK  %+v", 200, o.Payload)
}

func (o *PcloudEventsGetOK) GetPayload() *models.Event {
	return o.Payload
}

func (o *PcloudEventsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Event)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudEventsGetBadRequest creates a PcloudEventsGetBadRequest with default headers values
func NewPcloudEventsGetBadRequest() *PcloudEventsGetBadRequest {
	return &PcloudEventsGetBadRequest{}
}

/*
PcloudEventsGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudEventsGetBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud events get bad request response has a 2xx status code
func (o *PcloudEventsGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud events get bad request response has a 3xx status code
func (o *PcloudEventsGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud events get bad request response has a 4xx status code
func (o *PcloudEventsGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud events get bad request response has a 5xx status code
func (o *PcloudEventsGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud events get bad request response a status code equal to that given
func (o *PcloudEventsGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud events get bad request response
func (o *PcloudEventsGetBadRequest) Code() int {
	return 400
}

func (o *PcloudEventsGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/events/{event_id}][%d] pcloudEventsGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudEventsGetBadRequest) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/events/{event_id}][%d] pcloudEventsGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudEventsGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudEventsGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudEventsGetForbidden creates a PcloudEventsGetForbidden with default headers values
func NewPcloudEventsGetForbidden() *PcloudEventsGetForbidden {
	return &PcloudEventsGetForbidden{}
}

/*
PcloudEventsGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudEventsGetForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud events get forbidden response has a 2xx status code
func (o *PcloudEventsGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud events get forbidden response has a 3xx status code
func (o *PcloudEventsGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud events get forbidden response has a 4xx status code
func (o *PcloudEventsGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud events get forbidden response has a 5xx status code
func (o *PcloudEventsGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud events get forbidden response a status code equal to that given
func (o *PcloudEventsGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud events get forbidden response
func (o *PcloudEventsGetForbidden) Code() int {
	return 403
}

func (o *PcloudEventsGetForbidden) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/events/{event_id}][%d] pcloudEventsGetForbidden  %+v", 403, o.Payload)
}

func (o *PcloudEventsGetForbidden) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/events/{event_id}][%d] pcloudEventsGetForbidden  %+v", 403, o.Payload)
}

func (o *PcloudEventsGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudEventsGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudEventsGetNotFound creates a PcloudEventsGetNotFound with default headers values
func NewPcloudEventsGetNotFound() *PcloudEventsGetNotFound {
	return &PcloudEventsGetNotFound{}
}

/*
PcloudEventsGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudEventsGetNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud events get not found response has a 2xx status code
func (o *PcloudEventsGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud events get not found response has a 3xx status code
func (o *PcloudEventsGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud events get not found response has a 4xx status code
func (o *PcloudEventsGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud events get not found response has a 5xx status code
func (o *PcloudEventsGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud events get not found response a status code equal to that given
func (o *PcloudEventsGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud events get not found response
func (o *PcloudEventsGetNotFound) Code() int {
	return 404
}

func (o *PcloudEventsGetNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/events/{event_id}][%d] pcloudEventsGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudEventsGetNotFound) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/events/{event_id}][%d] pcloudEventsGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudEventsGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudEventsGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudEventsGetInternalServerError creates a PcloudEventsGetInternalServerError with default headers values
func NewPcloudEventsGetInternalServerError() *PcloudEventsGetInternalServerError {
	return &PcloudEventsGetInternalServerError{}
}

/*
PcloudEventsGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudEventsGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud events get internal server error response has a 2xx status code
func (o *PcloudEventsGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud events get internal server error response has a 3xx status code
func (o *PcloudEventsGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud events get internal server error response has a 4xx status code
func (o *PcloudEventsGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud events get internal server error response has a 5xx status code
func (o *PcloudEventsGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud events get internal server error response a status code equal to that given
func (o *PcloudEventsGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud events get internal server error response
func (o *PcloudEventsGetInternalServerError) Code() int {
	return 500
}

func (o *PcloudEventsGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/events/{event_id}][%d] pcloudEventsGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudEventsGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/events/{event_id}][%d] pcloudEventsGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudEventsGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudEventsGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
