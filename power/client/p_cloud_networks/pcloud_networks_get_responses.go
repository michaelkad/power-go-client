// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.ibm.com/power-iaas/service-broker/power/models"
)

// PcloudNetworksGetReader is a Reader for the PcloudNetworksGet structure.
type PcloudNetworksGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudNetworksGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudNetworksGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudNetworksGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudNetworksGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudNetworksGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudNetworksGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudNetworksGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}] pcloud.networks.get", response, response.Code())
	}
}

// NewPcloudNetworksGetOK creates a PcloudNetworksGetOK with default headers values
func NewPcloudNetworksGetOK() *PcloudNetworksGetOK {
	return &PcloudNetworksGetOK{}
}

/*
PcloudNetworksGetOK describes a response with status code 200, with default header values.

OK
*/
type PcloudNetworksGetOK struct {
	Payload *models.Network
}

// IsSuccess returns true when this pcloud networks get o k response has a 2xx status code
func (o *PcloudNetworksGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud networks get o k response has a 3xx status code
func (o *PcloudNetworksGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks get o k response has a 4xx status code
func (o *PcloudNetworksGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud networks get o k response has a 5xx status code
func (o *PcloudNetworksGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks get o k response a status code equal to that given
func (o *PcloudNetworksGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud networks get o k response
func (o *PcloudNetworksGetOK) Code() int {
	return 200
}

func (o *PcloudNetworksGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}][%d] pcloudNetworksGetOK  %+v", 200, o.Payload)
}

func (o *PcloudNetworksGetOK) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}][%d] pcloudNetworksGetOK  %+v", 200, o.Payload)
}

func (o *PcloudNetworksGetOK) GetPayload() *models.Network {
	return o.Payload
}

func (o *PcloudNetworksGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Network)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksGetBadRequest creates a PcloudNetworksGetBadRequest with default headers values
func NewPcloudNetworksGetBadRequest() *PcloudNetworksGetBadRequest {
	return &PcloudNetworksGetBadRequest{}
}

/*
PcloudNetworksGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudNetworksGetBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks get bad request response has a 2xx status code
func (o *PcloudNetworksGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks get bad request response has a 3xx status code
func (o *PcloudNetworksGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks get bad request response has a 4xx status code
func (o *PcloudNetworksGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud networks get bad request response has a 5xx status code
func (o *PcloudNetworksGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks get bad request response a status code equal to that given
func (o *PcloudNetworksGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud networks get bad request response
func (o *PcloudNetworksGetBadRequest) Code() int {
	return 400
}

func (o *PcloudNetworksGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}][%d] pcloudNetworksGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudNetworksGetBadRequest) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}][%d] pcloudNetworksGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudNetworksGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksGetUnauthorized creates a PcloudNetworksGetUnauthorized with default headers values
func NewPcloudNetworksGetUnauthorized() *PcloudNetworksGetUnauthorized {
	return &PcloudNetworksGetUnauthorized{}
}

/*
PcloudNetworksGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudNetworksGetUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks get unauthorized response has a 2xx status code
func (o *PcloudNetworksGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks get unauthorized response has a 3xx status code
func (o *PcloudNetworksGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks get unauthorized response has a 4xx status code
func (o *PcloudNetworksGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud networks get unauthorized response has a 5xx status code
func (o *PcloudNetworksGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks get unauthorized response a status code equal to that given
func (o *PcloudNetworksGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud networks get unauthorized response
func (o *PcloudNetworksGetUnauthorized) Code() int {
	return 401
}

func (o *PcloudNetworksGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}][%d] pcloudNetworksGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudNetworksGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}][%d] pcloudNetworksGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudNetworksGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksGetForbidden creates a PcloudNetworksGetForbidden with default headers values
func NewPcloudNetworksGetForbidden() *PcloudNetworksGetForbidden {
	return &PcloudNetworksGetForbidden{}
}

/*
PcloudNetworksGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudNetworksGetForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks get forbidden response has a 2xx status code
func (o *PcloudNetworksGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks get forbidden response has a 3xx status code
func (o *PcloudNetworksGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks get forbidden response has a 4xx status code
func (o *PcloudNetworksGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud networks get forbidden response has a 5xx status code
func (o *PcloudNetworksGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks get forbidden response a status code equal to that given
func (o *PcloudNetworksGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud networks get forbidden response
func (o *PcloudNetworksGetForbidden) Code() int {
	return 403
}

func (o *PcloudNetworksGetForbidden) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}][%d] pcloudNetworksGetForbidden  %+v", 403, o.Payload)
}

func (o *PcloudNetworksGetForbidden) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}][%d] pcloudNetworksGetForbidden  %+v", 403, o.Payload)
}

func (o *PcloudNetworksGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksGetNotFound creates a PcloudNetworksGetNotFound with default headers values
func NewPcloudNetworksGetNotFound() *PcloudNetworksGetNotFound {
	return &PcloudNetworksGetNotFound{}
}

/*
PcloudNetworksGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudNetworksGetNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks get not found response has a 2xx status code
func (o *PcloudNetworksGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks get not found response has a 3xx status code
func (o *PcloudNetworksGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks get not found response has a 4xx status code
func (o *PcloudNetworksGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud networks get not found response has a 5xx status code
func (o *PcloudNetworksGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks get not found response a status code equal to that given
func (o *PcloudNetworksGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud networks get not found response
func (o *PcloudNetworksGetNotFound) Code() int {
	return 404
}

func (o *PcloudNetworksGetNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}][%d] pcloudNetworksGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudNetworksGetNotFound) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}][%d] pcloudNetworksGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudNetworksGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksGetInternalServerError creates a PcloudNetworksGetInternalServerError with default headers values
func NewPcloudNetworksGetInternalServerError() *PcloudNetworksGetInternalServerError {
	return &PcloudNetworksGetInternalServerError{}
}

/*
PcloudNetworksGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudNetworksGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks get internal server error response has a 2xx status code
func (o *PcloudNetworksGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks get internal server error response has a 3xx status code
func (o *PcloudNetworksGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks get internal server error response has a 4xx status code
func (o *PcloudNetworksGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud networks get internal server error response has a 5xx status code
func (o *PcloudNetworksGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud networks get internal server error response a status code equal to that given
func (o *PcloudNetworksGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud networks get internal server error response
func (o *PcloudNetworksGetInternalServerError) Code() int {
	return 500
}

func (o *PcloudNetworksGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}][%d] pcloudNetworksGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudNetworksGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}][%d] pcloudNetworksGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudNetworksGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
