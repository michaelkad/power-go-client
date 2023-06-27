// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_s_a_p

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.ibm.com/power-iaas/service-broker/power/models"
)

// PcloudSapGetReader is a Reader for the PcloudSapGet structure.
type PcloudSapGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudSapGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudSapGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudSapGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudSapGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudSapGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudSapGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/sap/{sap_profile_id}] pcloud.sap.get", response, response.Code())
	}
}

// NewPcloudSapGetOK creates a PcloudSapGetOK with default headers values
func NewPcloudSapGetOK() *PcloudSapGetOK {
	return &PcloudSapGetOK{}
}

/*
PcloudSapGetOK describes a response with status code 200, with default header values.

OK
*/
type PcloudSapGetOK struct {
	Payload *models.SAPProfile
}

// IsSuccess returns true when this pcloud sap get o k response has a 2xx status code
func (o *PcloudSapGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud sap get o k response has a 3xx status code
func (o *PcloudSapGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sap get o k response has a 4xx status code
func (o *PcloudSapGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud sap get o k response has a 5xx status code
func (o *PcloudSapGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sap get o k response a status code equal to that given
func (o *PcloudSapGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud sap get o k response
func (o *PcloudSapGetOK) Code() int {
	return 200
}

func (o *PcloudSapGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/sap/{sap_profile_id}][%d] pcloudSapGetOK  %+v", 200, o.Payload)
}

func (o *PcloudSapGetOK) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/sap/{sap_profile_id}][%d] pcloudSapGetOK  %+v", 200, o.Payload)
}

func (o *PcloudSapGetOK) GetPayload() *models.SAPProfile {
	return o.Payload
}

func (o *PcloudSapGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SAPProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSapGetBadRequest creates a PcloudSapGetBadRequest with default headers values
func NewPcloudSapGetBadRequest() *PcloudSapGetBadRequest {
	return &PcloudSapGetBadRequest{}
}

/*
PcloudSapGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudSapGetBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud sap get bad request response has a 2xx status code
func (o *PcloudSapGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud sap get bad request response has a 3xx status code
func (o *PcloudSapGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sap get bad request response has a 4xx status code
func (o *PcloudSapGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud sap get bad request response has a 5xx status code
func (o *PcloudSapGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sap get bad request response a status code equal to that given
func (o *PcloudSapGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud sap get bad request response
func (o *PcloudSapGetBadRequest) Code() int {
	return 400
}

func (o *PcloudSapGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/sap/{sap_profile_id}][%d] pcloudSapGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudSapGetBadRequest) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/sap/{sap_profile_id}][%d] pcloudSapGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudSapGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSapGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSapGetUnauthorized creates a PcloudSapGetUnauthorized with default headers values
func NewPcloudSapGetUnauthorized() *PcloudSapGetUnauthorized {
	return &PcloudSapGetUnauthorized{}
}

/*
PcloudSapGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudSapGetUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud sap get unauthorized response has a 2xx status code
func (o *PcloudSapGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud sap get unauthorized response has a 3xx status code
func (o *PcloudSapGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sap get unauthorized response has a 4xx status code
func (o *PcloudSapGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud sap get unauthorized response has a 5xx status code
func (o *PcloudSapGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sap get unauthorized response a status code equal to that given
func (o *PcloudSapGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud sap get unauthorized response
func (o *PcloudSapGetUnauthorized) Code() int {
	return 401
}

func (o *PcloudSapGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/sap/{sap_profile_id}][%d] pcloudSapGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudSapGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/sap/{sap_profile_id}][%d] pcloudSapGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudSapGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSapGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSapGetNotFound creates a PcloudSapGetNotFound with default headers values
func NewPcloudSapGetNotFound() *PcloudSapGetNotFound {
	return &PcloudSapGetNotFound{}
}

/*
PcloudSapGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudSapGetNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud sap get not found response has a 2xx status code
func (o *PcloudSapGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud sap get not found response has a 3xx status code
func (o *PcloudSapGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sap get not found response has a 4xx status code
func (o *PcloudSapGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud sap get not found response has a 5xx status code
func (o *PcloudSapGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sap get not found response a status code equal to that given
func (o *PcloudSapGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud sap get not found response
func (o *PcloudSapGetNotFound) Code() int {
	return 404
}

func (o *PcloudSapGetNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/sap/{sap_profile_id}][%d] pcloudSapGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudSapGetNotFound) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/sap/{sap_profile_id}][%d] pcloudSapGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudSapGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSapGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSapGetInternalServerError creates a PcloudSapGetInternalServerError with default headers values
func NewPcloudSapGetInternalServerError() *PcloudSapGetInternalServerError {
	return &PcloudSapGetInternalServerError{}
}

/*
PcloudSapGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudSapGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud sap get internal server error response has a 2xx status code
func (o *PcloudSapGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud sap get internal server error response has a 3xx status code
func (o *PcloudSapGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sap get internal server error response has a 4xx status code
func (o *PcloudSapGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud sap get internal server error response has a 5xx status code
func (o *PcloudSapGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud sap get internal server error response a status code equal to that given
func (o *PcloudSapGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud sap get internal server error response
func (o *PcloudSapGetInternalServerError) Code() int {
	return 500
}

func (o *PcloudSapGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/sap/{sap_profile_id}][%d] pcloudSapGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudSapGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/sap/{sap_profile_id}][%d] pcloudSapGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudSapGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSapGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
