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

// PcloudNetworksDeleteReader is a Reader for the PcloudNetworksDelete structure.
type PcloudNetworksDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudNetworksDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudNetworksDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudNetworksDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudNetworksDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudNetworksDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudNetworksDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPcloudNetworksDeleteGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudNetworksDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}] pcloud.networks.delete", response, response.Code())
	}
}

// NewPcloudNetworksDeleteOK creates a PcloudNetworksDeleteOK with default headers values
func NewPcloudNetworksDeleteOK() *PcloudNetworksDeleteOK {
	return &PcloudNetworksDeleteOK{}
}

/*
PcloudNetworksDeleteOK describes a response with status code 200, with default header values.

OK
*/
type PcloudNetworksDeleteOK struct {
	Payload models.Object
}

// IsSuccess returns true when this pcloud networks delete o k response has a 2xx status code
func (o *PcloudNetworksDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud networks delete o k response has a 3xx status code
func (o *PcloudNetworksDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks delete o k response has a 4xx status code
func (o *PcloudNetworksDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud networks delete o k response has a 5xx status code
func (o *PcloudNetworksDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks delete o k response a status code equal to that given
func (o *PcloudNetworksDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud networks delete o k response
func (o *PcloudNetworksDeleteOK) Code() int {
	return 200
}

func (o *PcloudNetworksDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}][%d] pcloudNetworksDeleteOK  %+v", 200, o.Payload)
}

func (o *PcloudNetworksDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}][%d] pcloudNetworksDeleteOK  %+v", 200, o.Payload)
}

func (o *PcloudNetworksDeleteOK) GetPayload() models.Object {
	return o.Payload
}

func (o *PcloudNetworksDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksDeleteBadRequest creates a PcloudNetworksDeleteBadRequest with default headers values
func NewPcloudNetworksDeleteBadRequest() *PcloudNetworksDeleteBadRequest {
	return &PcloudNetworksDeleteBadRequest{}
}

/*
PcloudNetworksDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudNetworksDeleteBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks delete bad request response has a 2xx status code
func (o *PcloudNetworksDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks delete bad request response has a 3xx status code
func (o *PcloudNetworksDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks delete bad request response has a 4xx status code
func (o *PcloudNetworksDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud networks delete bad request response has a 5xx status code
func (o *PcloudNetworksDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks delete bad request response a status code equal to that given
func (o *PcloudNetworksDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud networks delete bad request response
func (o *PcloudNetworksDeleteBadRequest) Code() int {
	return 400
}

func (o *PcloudNetworksDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}][%d] pcloudNetworksDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudNetworksDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}][%d] pcloudNetworksDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudNetworksDeleteBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksDeleteUnauthorized creates a PcloudNetworksDeleteUnauthorized with default headers values
func NewPcloudNetworksDeleteUnauthorized() *PcloudNetworksDeleteUnauthorized {
	return &PcloudNetworksDeleteUnauthorized{}
}

/*
PcloudNetworksDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudNetworksDeleteUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks delete unauthorized response has a 2xx status code
func (o *PcloudNetworksDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks delete unauthorized response has a 3xx status code
func (o *PcloudNetworksDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks delete unauthorized response has a 4xx status code
func (o *PcloudNetworksDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud networks delete unauthorized response has a 5xx status code
func (o *PcloudNetworksDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks delete unauthorized response a status code equal to that given
func (o *PcloudNetworksDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud networks delete unauthorized response
func (o *PcloudNetworksDeleteUnauthorized) Code() int {
	return 401
}

func (o *PcloudNetworksDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}][%d] pcloudNetworksDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudNetworksDeleteUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}][%d] pcloudNetworksDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudNetworksDeleteUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksDeleteForbidden creates a PcloudNetworksDeleteForbidden with default headers values
func NewPcloudNetworksDeleteForbidden() *PcloudNetworksDeleteForbidden {
	return &PcloudNetworksDeleteForbidden{}
}

/*
PcloudNetworksDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudNetworksDeleteForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks delete forbidden response has a 2xx status code
func (o *PcloudNetworksDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks delete forbidden response has a 3xx status code
func (o *PcloudNetworksDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks delete forbidden response has a 4xx status code
func (o *PcloudNetworksDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud networks delete forbidden response has a 5xx status code
func (o *PcloudNetworksDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks delete forbidden response a status code equal to that given
func (o *PcloudNetworksDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud networks delete forbidden response
func (o *PcloudNetworksDeleteForbidden) Code() int {
	return 403
}

func (o *PcloudNetworksDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}][%d] pcloudNetworksDeleteForbidden  %+v", 403, o.Payload)
}

func (o *PcloudNetworksDeleteForbidden) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}][%d] pcloudNetworksDeleteForbidden  %+v", 403, o.Payload)
}

func (o *PcloudNetworksDeleteForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksDeleteNotFound creates a PcloudNetworksDeleteNotFound with default headers values
func NewPcloudNetworksDeleteNotFound() *PcloudNetworksDeleteNotFound {
	return &PcloudNetworksDeleteNotFound{}
}

/*
PcloudNetworksDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudNetworksDeleteNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks delete not found response has a 2xx status code
func (o *PcloudNetworksDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks delete not found response has a 3xx status code
func (o *PcloudNetworksDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks delete not found response has a 4xx status code
func (o *PcloudNetworksDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud networks delete not found response has a 5xx status code
func (o *PcloudNetworksDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks delete not found response a status code equal to that given
func (o *PcloudNetworksDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud networks delete not found response
func (o *PcloudNetworksDeleteNotFound) Code() int {
	return 404
}

func (o *PcloudNetworksDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}][%d] pcloudNetworksDeleteNotFound  %+v", 404, o.Payload)
}

func (o *PcloudNetworksDeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}][%d] pcloudNetworksDeleteNotFound  %+v", 404, o.Payload)
}

func (o *PcloudNetworksDeleteNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksDeleteGone creates a PcloudNetworksDeleteGone with default headers values
func NewPcloudNetworksDeleteGone() *PcloudNetworksDeleteGone {
	return &PcloudNetworksDeleteGone{}
}

/*
PcloudNetworksDeleteGone describes a response with status code 410, with default header values.

Gone
*/
type PcloudNetworksDeleteGone struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks delete gone response has a 2xx status code
func (o *PcloudNetworksDeleteGone) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks delete gone response has a 3xx status code
func (o *PcloudNetworksDeleteGone) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks delete gone response has a 4xx status code
func (o *PcloudNetworksDeleteGone) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud networks delete gone response has a 5xx status code
func (o *PcloudNetworksDeleteGone) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks delete gone response a status code equal to that given
func (o *PcloudNetworksDeleteGone) IsCode(code int) bool {
	return code == 410
}

// Code gets the status code for the pcloud networks delete gone response
func (o *PcloudNetworksDeleteGone) Code() int {
	return 410
}

func (o *PcloudNetworksDeleteGone) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}][%d] pcloudNetworksDeleteGone  %+v", 410, o.Payload)
}

func (o *PcloudNetworksDeleteGone) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}][%d] pcloudNetworksDeleteGone  %+v", 410, o.Payload)
}

func (o *PcloudNetworksDeleteGone) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksDeleteGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksDeleteInternalServerError creates a PcloudNetworksDeleteInternalServerError with default headers values
func NewPcloudNetworksDeleteInternalServerError() *PcloudNetworksDeleteInternalServerError {
	return &PcloudNetworksDeleteInternalServerError{}
}

/*
PcloudNetworksDeleteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudNetworksDeleteInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks delete internal server error response has a 2xx status code
func (o *PcloudNetworksDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks delete internal server error response has a 3xx status code
func (o *PcloudNetworksDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks delete internal server error response has a 4xx status code
func (o *PcloudNetworksDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud networks delete internal server error response has a 5xx status code
func (o *PcloudNetworksDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud networks delete internal server error response a status code equal to that given
func (o *PcloudNetworksDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud networks delete internal server error response
func (o *PcloudNetworksDeleteInternalServerError) Code() int {
	return 500
}

func (o *PcloudNetworksDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}][%d] pcloudNetworksDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudNetworksDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}][%d] pcloudNetworksDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudNetworksDeleteInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
