// Code generated by go-swagger; DO NOT EDIT.

package internal_power_v_s_locations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.ibm.com/power-iaas/service-broker/power/models"
)

// InternalV1PowervsLocationsTransitgatewayGetReader is a Reader for the InternalV1PowervsLocationsTransitgatewayGet structure.
type InternalV1PowervsLocationsTransitgatewayGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InternalV1PowervsLocationsTransitgatewayGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInternalV1PowervsLocationsTransitgatewayGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewInternalV1PowervsLocationsTransitgatewayGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewInternalV1PowervsLocationsTransitgatewayGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /internal/v1/powervs/locations/transit-gateway] internal.v1.powervs.locations.transitgateway.get", response, response.Code())
	}
}

// NewInternalV1PowervsLocationsTransitgatewayGetOK creates a InternalV1PowervsLocationsTransitgatewayGetOK with default headers values
func NewInternalV1PowervsLocationsTransitgatewayGetOK() *InternalV1PowervsLocationsTransitgatewayGetOK {
	return &InternalV1PowervsLocationsTransitgatewayGetOK{}
}

/*
InternalV1PowervsLocationsTransitgatewayGetOK describes a response with status code 200, with default header values.

OK
*/
type InternalV1PowervsLocationsTransitgatewayGetOK struct {
	Payload *models.TransitGatewayLocations
}

// IsSuccess returns true when this internal v1 powervs locations transitgateway get o k response has a 2xx status code
func (o *InternalV1PowervsLocationsTransitgatewayGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this internal v1 powervs locations transitgateway get o k response has a 3xx status code
func (o *InternalV1PowervsLocationsTransitgatewayGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 powervs locations transitgateway get o k response has a 4xx status code
func (o *InternalV1PowervsLocationsTransitgatewayGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this internal v1 powervs locations transitgateway get o k response has a 5xx status code
func (o *InternalV1PowervsLocationsTransitgatewayGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 powervs locations transitgateway get o k response a status code equal to that given
func (o *InternalV1PowervsLocationsTransitgatewayGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the internal v1 powervs locations transitgateway get o k response
func (o *InternalV1PowervsLocationsTransitgatewayGetOK) Code() int {
	return 200
}

func (o *InternalV1PowervsLocationsTransitgatewayGetOK) Error() string {
	return fmt.Sprintf("[GET /internal/v1/powervs/locations/transit-gateway][%d] internalV1PowervsLocationsTransitgatewayGetOK  %+v", 200, o.Payload)
}

func (o *InternalV1PowervsLocationsTransitgatewayGetOK) String() string {
	return fmt.Sprintf("[GET /internal/v1/powervs/locations/transit-gateway][%d] internalV1PowervsLocationsTransitgatewayGetOK  %+v", 200, o.Payload)
}

func (o *InternalV1PowervsLocationsTransitgatewayGetOK) GetPayload() *models.TransitGatewayLocations {
	return o.Payload
}

func (o *InternalV1PowervsLocationsTransitgatewayGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TransitGatewayLocations)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1PowervsLocationsTransitgatewayGetForbidden creates a InternalV1PowervsLocationsTransitgatewayGetForbidden with default headers values
func NewInternalV1PowervsLocationsTransitgatewayGetForbidden() *InternalV1PowervsLocationsTransitgatewayGetForbidden {
	return &InternalV1PowervsLocationsTransitgatewayGetForbidden{}
}

/*
InternalV1PowervsLocationsTransitgatewayGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type InternalV1PowervsLocationsTransitgatewayGetForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 powervs locations transitgateway get forbidden response has a 2xx status code
func (o *InternalV1PowervsLocationsTransitgatewayGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 powervs locations transitgateway get forbidden response has a 3xx status code
func (o *InternalV1PowervsLocationsTransitgatewayGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 powervs locations transitgateway get forbidden response has a 4xx status code
func (o *InternalV1PowervsLocationsTransitgatewayGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal v1 powervs locations transitgateway get forbidden response has a 5xx status code
func (o *InternalV1PowervsLocationsTransitgatewayGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 powervs locations transitgateway get forbidden response a status code equal to that given
func (o *InternalV1PowervsLocationsTransitgatewayGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the internal v1 powervs locations transitgateway get forbidden response
func (o *InternalV1PowervsLocationsTransitgatewayGetForbidden) Code() int {
	return 403
}

func (o *InternalV1PowervsLocationsTransitgatewayGetForbidden) Error() string {
	return fmt.Sprintf("[GET /internal/v1/powervs/locations/transit-gateway][%d] internalV1PowervsLocationsTransitgatewayGetForbidden  %+v", 403, o.Payload)
}

func (o *InternalV1PowervsLocationsTransitgatewayGetForbidden) String() string {
	return fmt.Sprintf("[GET /internal/v1/powervs/locations/transit-gateway][%d] internalV1PowervsLocationsTransitgatewayGetForbidden  %+v", 403, o.Payload)
}

func (o *InternalV1PowervsLocationsTransitgatewayGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1PowervsLocationsTransitgatewayGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1PowervsLocationsTransitgatewayGetInternalServerError creates a InternalV1PowervsLocationsTransitgatewayGetInternalServerError with default headers values
func NewInternalV1PowervsLocationsTransitgatewayGetInternalServerError() *InternalV1PowervsLocationsTransitgatewayGetInternalServerError {
	return &InternalV1PowervsLocationsTransitgatewayGetInternalServerError{}
}

/*
InternalV1PowervsLocationsTransitgatewayGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type InternalV1PowervsLocationsTransitgatewayGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 powervs locations transitgateway get internal server error response has a 2xx status code
func (o *InternalV1PowervsLocationsTransitgatewayGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 powervs locations transitgateway get internal server error response has a 3xx status code
func (o *InternalV1PowervsLocationsTransitgatewayGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 powervs locations transitgateway get internal server error response has a 4xx status code
func (o *InternalV1PowervsLocationsTransitgatewayGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this internal v1 powervs locations transitgateway get internal server error response has a 5xx status code
func (o *InternalV1PowervsLocationsTransitgatewayGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this internal v1 powervs locations transitgateway get internal server error response a status code equal to that given
func (o *InternalV1PowervsLocationsTransitgatewayGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the internal v1 powervs locations transitgateway get internal server error response
func (o *InternalV1PowervsLocationsTransitgatewayGetInternalServerError) Code() int {
	return 500
}

func (o *InternalV1PowervsLocationsTransitgatewayGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /internal/v1/powervs/locations/transit-gateway][%d] internalV1PowervsLocationsTransitgatewayGetInternalServerError  %+v", 500, o.Payload)
}

func (o *InternalV1PowervsLocationsTransitgatewayGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /internal/v1/powervs/locations/transit-gateway][%d] internalV1PowervsLocationsTransitgatewayGetInternalServerError  %+v", 500, o.Payload)
}

func (o *InternalV1PowervsLocationsTransitgatewayGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1PowervsLocationsTransitgatewayGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
