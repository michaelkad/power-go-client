// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_tenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.ibm.com/power-iaas/service-broker/power/models"
)

// PcloudTenantsPutReader is a Reader for the PcloudTenantsPut structure.
type PcloudTenantsPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudTenantsPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudTenantsPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudTenantsPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudTenantsPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudTenantsPutUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudTenantsPutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /pcloud/v1/tenants/{tenant_id}] pcloud.tenants.put", response, response.Code())
	}
}

// NewPcloudTenantsPutOK creates a PcloudTenantsPutOK with default headers values
func NewPcloudTenantsPutOK() *PcloudTenantsPutOK {
	return &PcloudTenantsPutOK{}
}

/*
PcloudTenantsPutOK describes a response with status code 200, with default header values.

OK
*/
type PcloudTenantsPutOK struct {
	Payload *models.Tenant
}

// IsSuccess returns true when this pcloud tenants put o k response has a 2xx status code
func (o *PcloudTenantsPutOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud tenants put o k response has a 3xx status code
func (o *PcloudTenantsPutOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud tenants put o k response has a 4xx status code
func (o *PcloudTenantsPutOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud tenants put o k response has a 5xx status code
func (o *PcloudTenantsPutOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud tenants put o k response a status code equal to that given
func (o *PcloudTenantsPutOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud tenants put o k response
func (o *PcloudTenantsPutOK) Code() int {
	return 200
}

func (o *PcloudTenantsPutOK) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/tenants/{tenant_id}][%d] pcloudTenantsPutOK  %+v", 200, o.Payload)
}

func (o *PcloudTenantsPutOK) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/tenants/{tenant_id}][%d] pcloudTenantsPutOK  %+v", 200, o.Payload)
}

func (o *PcloudTenantsPutOK) GetPayload() *models.Tenant {
	return o.Payload
}

func (o *PcloudTenantsPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tenant)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTenantsPutBadRequest creates a PcloudTenantsPutBadRequest with default headers values
func NewPcloudTenantsPutBadRequest() *PcloudTenantsPutBadRequest {
	return &PcloudTenantsPutBadRequest{}
}

/*
PcloudTenantsPutBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudTenantsPutBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud tenants put bad request response has a 2xx status code
func (o *PcloudTenantsPutBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud tenants put bad request response has a 3xx status code
func (o *PcloudTenantsPutBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud tenants put bad request response has a 4xx status code
func (o *PcloudTenantsPutBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud tenants put bad request response has a 5xx status code
func (o *PcloudTenantsPutBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud tenants put bad request response a status code equal to that given
func (o *PcloudTenantsPutBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud tenants put bad request response
func (o *PcloudTenantsPutBadRequest) Code() int {
	return 400
}

func (o *PcloudTenantsPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/tenants/{tenant_id}][%d] pcloudTenantsPutBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudTenantsPutBadRequest) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/tenants/{tenant_id}][%d] pcloudTenantsPutBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudTenantsPutBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudTenantsPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTenantsPutUnauthorized creates a PcloudTenantsPutUnauthorized with default headers values
func NewPcloudTenantsPutUnauthorized() *PcloudTenantsPutUnauthorized {
	return &PcloudTenantsPutUnauthorized{}
}

/*
PcloudTenantsPutUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudTenantsPutUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud tenants put unauthorized response has a 2xx status code
func (o *PcloudTenantsPutUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud tenants put unauthorized response has a 3xx status code
func (o *PcloudTenantsPutUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud tenants put unauthorized response has a 4xx status code
func (o *PcloudTenantsPutUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud tenants put unauthorized response has a 5xx status code
func (o *PcloudTenantsPutUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud tenants put unauthorized response a status code equal to that given
func (o *PcloudTenantsPutUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud tenants put unauthorized response
func (o *PcloudTenantsPutUnauthorized) Code() int {
	return 401
}

func (o *PcloudTenantsPutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/tenants/{tenant_id}][%d] pcloudTenantsPutUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudTenantsPutUnauthorized) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/tenants/{tenant_id}][%d] pcloudTenantsPutUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudTenantsPutUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudTenantsPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTenantsPutUnprocessableEntity creates a PcloudTenantsPutUnprocessableEntity with default headers values
func NewPcloudTenantsPutUnprocessableEntity() *PcloudTenantsPutUnprocessableEntity {
	return &PcloudTenantsPutUnprocessableEntity{}
}

/*
PcloudTenantsPutUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudTenantsPutUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud tenants put unprocessable entity response has a 2xx status code
func (o *PcloudTenantsPutUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud tenants put unprocessable entity response has a 3xx status code
func (o *PcloudTenantsPutUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud tenants put unprocessable entity response has a 4xx status code
func (o *PcloudTenantsPutUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud tenants put unprocessable entity response has a 5xx status code
func (o *PcloudTenantsPutUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud tenants put unprocessable entity response a status code equal to that given
func (o *PcloudTenantsPutUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the pcloud tenants put unprocessable entity response
func (o *PcloudTenantsPutUnprocessableEntity) Code() int {
	return 422
}

func (o *PcloudTenantsPutUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/tenants/{tenant_id}][%d] pcloudTenantsPutUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudTenantsPutUnprocessableEntity) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/tenants/{tenant_id}][%d] pcloudTenantsPutUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudTenantsPutUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudTenantsPutUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTenantsPutInternalServerError creates a PcloudTenantsPutInternalServerError with default headers values
func NewPcloudTenantsPutInternalServerError() *PcloudTenantsPutInternalServerError {
	return &PcloudTenantsPutInternalServerError{}
}

/*
PcloudTenantsPutInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudTenantsPutInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud tenants put internal server error response has a 2xx status code
func (o *PcloudTenantsPutInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud tenants put internal server error response has a 3xx status code
func (o *PcloudTenantsPutInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud tenants put internal server error response has a 4xx status code
func (o *PcloudTenantsPutInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud tenants put internal server error response has a 5xx status code
func (o *PcloudTenantsPutInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud tenants put internal server error response a status code equal to that given
func (o *PcloudTenantsPutInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud tenants put internal server error response
func (o *PcloudTenantsPutInternalServerError) Code() int {
	return 500
}

func (o *PcloudTenantsPutInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/tenants/{tenant_id}][%d] pcloudTenantsPutInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudTenantsPutInternalServerError) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/tenants/{tenant_id}][%d] pcloudTenantsPutInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudTenantsPutInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudTenantsPutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
