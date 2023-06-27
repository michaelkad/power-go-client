// Code generated by go-swagger; DO NOT EDIT.

package internal_power_v_s_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.ibm.com/power-iaas/service-broker/power/models"
)

// InternalV1PowervsInstancesGetReader is a Reader for the InternalV1PowervsInstancesGet structure.
type InternalV1PowervsInstancesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InternalV1PowervsInstancesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInternalV1PowervsInstancesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewInternalV1PowervsInstancesGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewInternalV1PowervsInstancesGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /internal/v1/powervs/instances] internal.v1.powervs.instances.get", response, response.Code())
	}
}

// NewInternalV1PowervsInstancesGetOK creates a InternalV1PowervsInstancesGetOK with default headers values
func NewInternalV1PowervsInstancesGetOK() *InternalV1PowervsInstancesGetOK {
	return &InternalV1PowervsInstancesGetOK{}
}

/*
InternalV1PowervsInstancesGetOK describes a response with status code 200, with default header values.

OK
*/
type InternalV1PowervsInstancesGetOK struct {
	Payload *models.PowerVSInstances
}

// IsSuccess returns true when this internal v1 powervs instances get o k response has a 2xx status code
func (o *InternalV1PowervsInstancesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this internal v1 powervs instances get o k response has a 3xx status code
func (o *InternalV1PowervsInstancesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 powervs instances get o k response has a 4xx status code
func (o *InternalV1PowervsInstancesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this internal v1 powervs instances get o k response has a 5xx status code
func (o *InternalV1PowervsInstancesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 powervs instances get o k response a status code equal to that given
func (o *InternalV1PowervsInstancesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the internal v1 powervs instances get o k response
func (o *InternalV1PowervsInstancesGetOK) Code() int {
	return 200
}

func (o *InternalV1PowervsInstancesGetOK) Error() string {
	return fmt.Sprintf("[GET /internal/v1/powervs/instances][%d] internalV1PowervsInstancesGetOK  %+v", 200, o.Payload)
}

func (o *InternalV1PowervsInstancesGetOK) String() string {
	return fmt.Sprintf("[GET /internal/v1/powervs/instances][%d] internalV1PowervsInstancesGetOK  %+v", 200, o.Payload)
}

func (o *InternalV1PowervsInstancesGetOK) GetPayload() *models.PowerVSInstances {
	return o.Payload
}

func (o *InternalV1PowervsInstancesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerVSInstances)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1PowervsInstancesGetForbidden creates a InternalV1PowervsInstancesGetForbidden with default headers values
func NewInternalV1PowervsInstancesGetForbidden() *InternalV1PowervsInstancesGetForbidden {
	return &InternalV1PowervsInstancesGetForbidden{}
}

/*
InternalV1PowervsInstancesGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type InternalV1PowervsInstancesGetForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 powervs instances get forbidden response has a 2xx status code
func (o *InternalV1PowervsInstancesGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 powervs instances get forbidden response has a 3xx status code
func (o *InternalV1PowervsInstancesGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 powervs instances get forbidden response has a 4xx status code
func (o *InternalV1PowervsInstancesGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal v1 powervs instances get forbidden response has a 5xx status code
func (o *InternalV1PowervsInstancesGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 powervs instances get forbidden response a status code equal to that given
func (o *InternalV1PowervsInstancesGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the internal v1 powervs instances get forbidden response
func (o *InternalV1PowervsInstancesGetForbidden) Code() int {
	return 403
}

func (o *InternalV1PowervsInstancesGetForbidden) Error() string {
	return fmt.Sprintf("[GET /internal/v1/powervs/instances][%d] internalV1PowervsInstancesGetForbidden  %+v", 403, o.Payload)
}

func (o *InternalV1PowervsInstancesGetForbidden) String() string {
	return fmt.Sprintf("[GET /internal/v1/powervs/instances][%d] internalV1PowervsInstancesGetForbidden  %+v", 403, o.Payload)
}

func (o *InternalV1PowervsInstancesGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1PowervsInstancesGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1PowervsInstancesGetInternalServerError creates a InternalV1PowervsInstancesGetInternalServerError with default headers values
func NewInternalV1PowervsInstancesGetInternalServerError() *InternalV1PowervsInstancesGetInternalServerError {
	return &InternalV1PowervsInstancesGetInternalServerError{}
}

/*
InternalV1PowervsInstancesGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type InternalV1PowervsInstancesGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 powervs instances get internal server error response has a 2xx status code
func (o *InternalV1PowervsInstancesGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 powervs instances get internal server error response has a 3xx status code
func (o *InternalV1PowervsInstancesGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 powervs instances get internal server error response has a 4xx status code
func (o *InternalV1PowervsInstancesGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this internal v1 powervs instances get internal server error response has a 5xx status code
func (o *InternalV1PowervsInstancesGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this internal v1 powervs instances get internal server error response a status code equal to that given
func (o *InternalV1PowervsInstancesGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the internal v1 powervs instances get internal server error response
func (o *InternalV1PowervsInstancesGetInternalServerError) Code() int {
	return 500
}

func (o *InternalV1PowervsInstancesGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /internal/v1/powervs/instances][%d] internalV1PowervsInstancesGetInternalServerError  %+v", 500, o.Payload)
}

func (o *InternalV1PowervsInstancesGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /internal/v1/powervs/instances][%d] internalV1PowervsInstancesGetInternalServerError  %+v", 500, o.Payload)
}

func (o *InternalV1PowervsInstancesGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1PowervsInstancesGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
