// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.ibm.com/power-iaas/service-broker/power/models"
)

// PcloudCloudinstancesJobsDeleteReader is a Reader for the PcloudCloudinstancesJobsDelete structure.
type PcloudCloudinstancesJobsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudCloudinstancesJobsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudCloudinstancesJobsDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudCloudinstancesJobsDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudCloudinstancesJobsDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudCloudinstancesJobsDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudCloudinstancesJobsDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPcloudCloudinstancesJobsDeleteConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudCloudinstancesJobsDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs/{job_id}] pcloud.cloudinstances.jobs.delete", response, response.Code())
	}
}

// NewPcloudCloudinstancesJobsDeleteOK creates a PcloudCloudinstancesJobsDeleteOK with default headers values
func NewPcloudCloudinstancesJobsDeleteOK() *PcloudCloudinstancesJobsDeleteOK {
	return &PcloudCloudinstancesJobsDeleteOK{}
}

/*
PcloudCloudinstancesJobsDeleteOK describes a response with status code 200, with default header values.

OK
*/
type PcloudCloudinstancesJobsDeleteOK struct {
	Payload models.Object
}

// IsSuccess returns true when this pcloud cloudinstances jobs delete o k response has a 2xx status code
func (o *PcloudCloudinstancesJobsDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud cloudinstances jobs delete o k response has a 3xx status code
func (o *PcloudCloudinstancesJobsDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances jobs delete o k response has a 4xx status code
func (o *PcloudCloudinstancesJobsDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudinstances jobs delete o k response has a 5xx status code
func (o *PcloudCloudinstancesJobsDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances jobs delete o k response a status code equal to that given
func (o *PcloudCloudinstancesJobsDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud cloudinstances jobs delete o k response
func (o *PcloudCloudinstancesJobsDeleteOK) Code() int {
	return 200
}

func (o *PcloudCloudinstancesJobsDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs/{job_id}][%d] pcloudCloudinstancesJobsDeleteOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudinstancesJobsDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs/{job_id}][%d] pcloudCloudinstancesJobsDeleteOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudinstancesJobsDeleteOK) GetPayload() models.Object {
	return o.Payload
}

func (o *PcloudCloudinstancesJobsDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesJobsDeleteBadRequest creates a PcloudCloudinstancesJobsDeleteBadRequest with default headers values
func NewPcloudCloudinstancesJobsDeleteBadRequest() *PcloudCloudinstancesJobsDeleteBadRequest {
	return &PcloudCloudinstancesJobsDeleteBadRequest{}
}

/*
PcloudCloudinstancesJobsDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudCloudinstancesJobsDeleteBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances jobs delete bad request response has a 2xx status code
func (o *PcloudCloudinstancesJobsDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances jobs delete bad request response has a 3xx status code
func (o *PcloudCloudinstancesJobsDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances jobs delete bad request response has a 4xx status code
func (o *PcloudCloudinstancesJobsDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances jobs delete bad request response has a 5xx status code
func (o *PcloudCloudinstancesJobsDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances jobs delete bad request response a status code equal to that given
func (o *PcloudCloudinstancesJobsDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud cloudinstances jobs delete bad request response
func (o *PcloudCloudinstancesJobsDeleteBadRequest) Code() int {
	return 400
}

func (o *PcloudCloudinstancesJobsDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs/{job_id}][%d] pcloudCloudinstancesJobsDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudinstancesJobsDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs/{job_id}][%d] pcloudCloudinstancesJobsDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudinstancesJobsDeleteBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesJobsDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesJobsDeleteUnauthorized creates a PcloudCloudinstancesJobsDeleteUnauthorized with default headers values
func NewPcloudCloudinstancesJobsDeleteUnauthorized() *PcloudCloudinstancesJobsDeleteUnauthorized {
	return &PcloudCloudinstancesJobsDeleteUnauthorized{}
}

/*
PcloudCloudinstancesJobsDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudCloudinstancesJobsDeleteUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances jobs delete unauthorized response has a 2xx status code
func (o *PcloudCloudinstancesJobsDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances jobs delete unauthorized response has a 3xx status code
func (o *PcloudCloudinstancesJobsDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances jobs delete unauthorized response has a 4xx status code
func (o *PcloudCloudinstancesJobsDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances jobs delete unauthorized response has a 5xx status code
func (o *PcloudCloudinstancesJobsDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances jobs delete unauthorized response a status code equal to that given
func (o *PcloudCloudinstancesJobsDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud cloudinstances jobs delete unauthorized response
func (o *PcloudCloudinstancesJobsDeleteUnauthorized) Code() int {
	return 401
}

func (o *PcloudCloudinstancesJobsDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs/{job_id}][%d] pcloudCloudinstancesJobsDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudinstancesJobsDeleteUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs/{job_id}][%d] pcloudCloudinstancesJobsDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudinstancesJobsDeleteUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesJobsDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesJobsDeleteForbidden creates a PcloudCloudinstancesJobsDeleteForbidden with default headers values
func NewPcloudCloudinstancesJobsDeleteForbidden() *PcloudCloudinstancesJobsDeleteForbidden {
	return &PcloudCloudinstancesJobsDeleteForbidden{}
}

/*
PcloudCloudinstancesJobsDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudCloudinstancesJobsDeleteForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances jobs delete forbidden response has a 2xx status code
func (o *PcloudCloudinstancesJobsDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances jobs delete forbidden response has a 3xx status code
func (o *PcloudCloudinstancesJobsDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances jobs delete forbidden response has a 4xx status code
func (o *PcloudCloudinstancesJobsDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances jobs delete forbidden response has a 5xx status code
func (o *PcloudCloudinstancesJobsDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances jobs delete forbidden response a status code equal to that given
func (o *PcloudCloudinstancesJobsDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud cloudinstances jobs delete forbidden response
func (o *PcloudCloudinstancesJobsDeleteForbidden) Code() int {
	return 403
}

func (o *PcloudCloudinstancesJobsDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs/{job_id}][%d] pcloudCloudinstancesJobsDeleteForbidden  %+v", 403, o.Payload)
}

func (o *PcloudCloudinstancesJobsDeleteForbidden) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs/{job_id}][%d] pcloudCloudinstancesJobsDeleteForbidden  %+v", 403, o.Payload)
}

func (o *PcloudCloudinstancesJobsDeleteForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesJobsDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesJobsDeleteNotFound creates a PcloudCloudinstancesJobsDeleteNotFound with default headers values
func NewPcloudCloudinstancesJobsDeleteNotFound() *PcloudCloudinstancesJobsDeleteNotFound {
	return &PcloudCloudinstancesJobsDeleteNotFound{}
}

/*
PcloudCloudinstancesJobsDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudCloudinstancesJobsDeleteNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances jobs delete not found response has a 2xx status code
func (o *PcloudCloudinstancesJobsDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances jobs delete not found response has a 3xx status code
func (o *PcloudCloudinstancesJobsDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances jobs delete not found response has a 4xx status code
func (o *PcloudCloudinstancesJobsDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances jobs delete not found response has a 5xx status code
func (o *PcloudCloudinstancesJobsDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances jobs delete not found response a status code equal to that given
func (o *PcloudCloudinstancesJobsDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud cloudinstances jobs delete not found response
func (o *PcloudCloudinstancesJobsDeleteNotFound) Code() int {
	return 404
}

func (o *PcloudCloudinstancesJobsDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs/{job_id}][%d] pcloudCloudinstancesJobsDeleteNotFound  %+v", 404, o.Payload)
}

func (o *PcloudCloudinstancesJobsDeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs/{job_id}][%d] pcloudCloudinstancesJobsDeleteNotFound  %+v", 404, o.Payload)
}

func (o *PcloudCloudinstancesJobsDeleteNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesJobsDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesJobsDeleteConflict creates a PcloudCloudinstancesJobsDeleteConflict with default headers values
func NewPcloudCloudinstancesJobsDeleteConflict() *PcloudCloudinstancesJobsDeleteConflict {
	return &PcloudCloudinstancesJobsDeleteConflict{}
}

/*
PcloudCloudinstancesJobsDeleteConflict describes a response with status code 409, with default header values.

Conflict
*/
type PcloudCloudinstancesJobsDeleteConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances jobs delete conflict response has a 2xx status code
func (o *PcloudCloudinstancesJobsDeleteConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances jobs delete conflict response has a 3xx status code
func (o *PcloudCloudinstancesJobsDeleteConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances jobs delete conflict response has a 4xx status code
func (o *PcloudCloudinstancesJobsDeleteConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances jobs delete conflict response has a 5xx status code
func (o *PcloudCloudinstancesJobsDeleteConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances jobs delete conflict response a status code equal to that given
func (o *PcloudCloudinstancesJobsDeleteConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the pcloud cloudinstances jobs delete conflict response
func (o *PcloudCloudinstancesJobsDeleteConflict) Code() int {
	return 409
}

func (o *PcloudCloudinstancesJobsDeleteConflict) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs/{job_id}][%d] pcloudCloudinstancesJobsDeleteConflict  %+v", 409, o.Payload)
}

func (o *PcloudCloudinstancesJobsDeleteConflict) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs/{job_id}][%d] pcloudCloudinstancesJobsDeleteConflict  %+v", 409, o.Payload)
}

func (o *PcloudCloudinstancesJobsDeleteConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesJobsDeleteConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesJobsDeleteInternalServerError creates a PcloudCloudinstancesJobsDeleteInternalServerError with default headers values
func NewPcloudCloudinstancesJobsDeleteInternalServerError() *PcloudCloudinstancesJobsDeleteInternalServerError {
	return &PcloudCloudinstancesJobsDeleteInternalServerError{}
}

/*
PcloudCloudinstancesJobsDeleteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudCloudinstancesJobsDeleteInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances jobs delete internal server error response has a 2xx status code
func (o *PcloudCloudinstancesJobsDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances jobs delete internal server error response has a 3xx status code
func (o *PcloudCloudinstancesJobsDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances jobs delete internal server error response has a 4xx status code
func (o *PcloudCloudinstancesJobsDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudinstances jobs delete internal server error response has a 5xx status code
func (o *PcloudCloudinstancesJobsDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud cloudinstances jobs delete internal server error response a status code equal to that given
func (o *PcloudCloudinstancesJobsDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud cloudinstances jobs delete internal server error response
func (o *PcloudCloudinstancesJobsDeleteInternalServerError) Code() int {
	return 500
}

func (o *PcloudCloudinstancesJobsDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs/{job_id}][%d] pcloudCloudinstancesJobsDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudinstancesJobsDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs/{job_id}][%d] pcloudCloudinstancesJobsDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudinstancesJobsDeleteInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesJobsDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
