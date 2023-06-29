// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_images

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudV2ImagesExportPostReader is a Reader for the PcloudV2ImagesExportPost structure.
type PcloudV2ImagesExportPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudV2ImagesExportPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPcloudV2ImagesExportPostAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudV2ImagesExportPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudV2ImagesExportPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudV2ImagesExportPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPcloudV2ImagesExportPostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudV2ImagesExportPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudV2ImagesExportPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/images/{image_id}/export] pcloud.v2.images.export.post", response, response.Code())
	}
}

// NewPcloudV2ImagesExportPostAccepted creates a PcloudV2ImagesExportPostAccepted with default headers values
func NewPcloudV2ImagesExportPostAccepted() *PcloudV2ImagesExportPostAccepted {
	return &PcloudV2ImagesExportPostAccepted{}
}

/*
PcloudV2ImagesExportPostAccepted describes a response with status code 202, with default header values.

Accepted, image export successfully added to the jobs queue
*/
type PcloudV2ImagesExportPostAccepted struct {
	Payload *models.JobReference
}

// IsSuccess returns true when this pcloud v2 images export post accepted response has a 2xx status code
func (o *PcloudV2ImagesExportPostAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud v2 images export post accepted response has a 3xx status code
func (o *PcloudV2ImagesExportPostAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 images export post accepted response has a 4xx status code
func (o *PcloudV2ImagesExportPostAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud v2 images export post accepted response has a 5xx status code
func (o *PcloudV2ImagesExportPostAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v2 images export post accepted response a status code equal to that given
func (o *PcloudV2ImagesExportPostAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the pcloud v2 images export post accepted response
func (o *PcloudV2ImagesExportPostAccepted) Code() int {
	return 202
}

func (o *PcloudV2ImagesExportPostAccepted) Error() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/images/{image_id}/export][%d] pcloudV2ImagesExportPostAccepted  %+v", 202, o.Payload)
}

func (o *PcloudV2ImagesExportPostAccepted) String() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/images/{image_id}/export][%d] pcloudV2ImagesExportPostAccepted  %+v", 202, o.Payload)
}

func (o *PcloudV2ImagesExportPostAccepted) GetPayload() *models.JobReference {
	return o.Payload
}

func (o *PcloudV2ImagesExportPostAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobReference)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2ImagesExportPostBadRequest creates a PcloudV2ImagesExportPostBadRequest with default headers values
func NewPcloudV2ImagesExportPostBadRequest() *PcloudV2ImagesExportPostBadRequest {
	return &PcloudV2ImagesExportPostBadRequest{}
}

/*
PcloudV2ImagesExportPostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudV2ImagesExportPostBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v2 images export post bad request response has a 2xx status code
func (o *PcloudV2ImagesExportPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v2 images export post bad request response has a 3xx status code
func (o *PcloudV2ImagesExportPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 images export post bad request response has a 4xx status code
func (o *PcloudV2ImagesExportPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud v2 images export post bad request response has a 5xx status code
func (o *PcloudV2ImagesExportPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v2 images export post bad request response a status code equal to that given
func (o *PcloudV2ImagesExportPostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud v2 images export post bad request response
func (o *PcloudV2ImagesExportPostBadRequest) Code() int {
	return 400
}

func (o *PcloudV2ImagesExportPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/images/{image_id}/export][%d] pcloudV2ImagesExportPostBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudV2ImagesExportPostBadRequest) String() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/images/{image_id}/export][%d] pcloudV2ImagesExportPostBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudV2ImagesExportPostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2ImagesExportPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2ImagesExportPostUnauthorized creates a PcloudV2ImagesExportPostUnauthorized with default headers values
func NewPcloudV2ImagesExportPostUnauthorized() *PcloudV2ImagesExportPostUnauthorized {
	return &PcloudV2ImagesExportPostUnauthorized{}
}

/*
PcloudV2ImagesExportPostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudV2ImagesExportPostUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v2 images export post unauthorized response has a 2xx status code
func (o *PcloudV2ImagesExportPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v2 images export post unauthorized response has a 3xx status code
func (o *PcloudV2ImagesExportPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 images export post unauthorized response has a 4xx status code
func (o *PcloudV2ImagesExportPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud v2 images export post unauthorized response has a 5xx status code
func (o *PcloudV2ImagesExportPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v2 images export post unauthorized response a status code equal to that given
func (o *PcloudV2ImagesExportPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud v2 images export post unauthorized response
func (o *PcloudV2ImagesExportPostUnauthorized) Code() int {
	return 401
}

func (o *PcloudV2ImagesExportPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/images/{image_id}/export][%d] pcloudV2ImagesExportPostUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudV2ImagesExportPostUnauthorized) String() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/images/{image_id}/export][%d] pcloudV2ImagesExportPostUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudV2ImagesExportPostUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2ImagesExportPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2ImagesExportPostNotFound creates a PcloudV2ImagesExportPostNotFound with default headers values
func NewPcloudV2ImagesExportPostNotFound() *PcloudV2ImagesExportPostNotFound {
	return &PcloudV2ImagesExportPostNotFound{}
}

/*
PcloudV2ImagesExportPostNotFound describes a response with status code 404, with default header values.

image id not found
*/
type PcloudV2ImagesExportPostNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v2 images export post not found response has a 2xx status code
func (o *PcloudV2ImagesExportPostNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v2 images export post not found response has a 3xx status code
func (o *PcloudV2ImagesExportPostNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 images export post not found response has a 4xx status code
func (o *PcloudV2ImagesExportPostNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud v2 images export post not found response has a 5xx status code
func (o *PcloudV2ImagesExportPostNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v2 images export post not found response a status code equal to that given
func (o *PcloudV2ImagesExportPostNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud v2 images export post not found response
func (o *PcloudV2ImagesExportPostNotFound) Code() int {
	return 404
}

func (o *PcloudV2ImagesExportPostNotFound) Error() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/images/{image_id}/export][%d] pcloudV2ImagesExportPostNotFound  %+v", 404, o.Payload)
}

func (o *PcloudV2ImagesExportPostNotFound) String() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/images/{image_id}/export][%d] pcloudV2ImagesExportPostNotFound  %+v", 404, o.Payload)
}

func (o *PcloudV2ImagesExportPostNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2ImagesExportPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2ImagesExportPostConflict creates a PcloudV2ImagesExportPostConflict with default headers values
func NewPcloudV2ImagesExportPostConflict() *PcloudV2ImagesExportPostConflict {
	return &PcloudV2ImagesExportPostConflict{}
}

/*
PcloudV2ImagesExportPostConflict describes a response with status code 409, with default header values.

Conflict, a conflict has prevented adding image export job
*/
type PcloudV2ImagesExportPostConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v2 images export post conflict response has a 2xx status code
func (o *PcloudV2ImagesExportPostConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v2 images export post conflict response has a 3xx status code
func (o *PcloudV2ImagesExportPostConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 images export post conflict response has a 4xx status code
func (o *PcloudV2ImagesExportPostConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud v2 images export post conflict response has a 5xx status code
func (o *PcloudV2ImagesExportPostConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v2 images export post conflict response a status code equal to that given
func (o *PcloudV2ImagesExportPostConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the pcloud v2 images export post conflict response
func (o *PcloudV2ImagesExportPostConflict) Code() int {
	return 409
}

func (o *PcloudV2ImagesExportPostConflict) Error() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/images/{image_id}/export][%d] pcloudV2ImagesExportPostConflict  %+v", 409, o.Payload)
}

func (o *PcloudV2ImagesExportPostConflict) String() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/images/{image_id}/export][%d] pcloudV2ImagesExportPostConflict  %+v", 409, o.Payload)
}

func (o *PcloudV2ImagesExportPostConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2ImagesExportPostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2ImagesExportPostUnprocessableEntity creates a PcloudV2ImagesExportPostUnprocessableEntity with default headers values
func NewPcloudV2ImagesExportPostUnprocessableEntity() *PcloudV2ImagesExportPostUnprocessableEntity {
	return &PcloudV2ImagesExportPostUnprocessableEntity{}
}

/*
PcloudV2ImagesExportPostUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudV2ImagesExportPostUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v2 images export post unprocessable entity response has a 2xx status code
func (o *PcloudV2ImagesExportPostUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v2 images export post unprocessable entity response has a 3xx status code
func (o *PcloudV2ImagesExportPostUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 images export post unprocessable entity response has a 4xx status code
func (o *PcloudV2ImagesExportPostUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud v2 images export post unprocessable entity response has a 5xx status code
func (o *PcloudV2ImagesExportPostUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v2 images export post unprocessable entity response a status code equal to that given
func (o *PcloudV2ImagesExportPostUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the pcloud v2 images export post unprocessable entity response
func (o *PcloudV2ImagesExportPostUnprocessableEntity) Code() int {
	return 422
}

func (o *PcloudV2ImagesExportPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/images/{image_id}/export][%d] pcloudV2ImagesExportPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudV2ImagesExportPostUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/images/{image_id}/export][%d] pcloudV2ImagesExportPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudV2ImagesExportPostUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2ImagesExportPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2ImagesExportPostInternalServerError creates a PcloudV2ImagesExportPostInternalServerError with default headers values
func NewPcloudV2ImagesExportPostInternalServerError() *PcloudV2ImagesExportPostInternalServerError {
	return &PcloudV2ImagesExportPostInternalServerError{}
}

/*
PcloudV2ImagesExportPostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudV2ImagesExportPostInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v2 images export post internal server error response has a 2xx status code
func (o *PcloudV2ImagesExportPostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v2 images export post internal server error response has a 3xx status code
func (o *PcloudV2ImagesExportPostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 images export post internal server error response has a 4xx status code
func (o *PcloudV2ImagesExportPostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud v2 images export post internal server error response has a 5xx status code
func (o *PcloudV2ImagesExportPostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud v2 images export post internal server error response a status code equal to that given
func (o *PcloudV2ImagesExportPostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud v2 images export post internal server error response
func (o *PcloudV2ImagesExportPostInternalServerError) Code() int {
	return 500
}

func (o *PcloudV2ImagesExportPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/images/{image_id}/export][%d] pcloudV2ImagesExportPostInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudV2ImagesExportPostInternalServerError) String() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/images/{image_id}/export][%d] pcloudV2ImagesExportPostInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudV2ImagesExportPostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2ImagesExportPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
