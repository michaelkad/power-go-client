// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_networks

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

// PcloudNetworksPostReader is a Reader for the PcloudNetworksPost structure.
type PcloudNetworksPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudNetworksPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudNetworksPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewPcloudNetworksPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudNetworksPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudNetworksPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudNetworksPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudNetworksPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPcloudNetworksPostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudNetworksPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudNetworksPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 550:
		result := NewPcloudNetworksPostStatus550()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/networks] pcloud.networks.post", response, response.Code())
	}
}

// NewPcloudNetworksPostOK creates a PcloudNetworksPostOK with default headers values
func NewPcloudNetworksPostOK() *PcloudNetworksPostOK {
	return &PcloudNetworksPostOK{}
}

/*
PcloudNetworksPostOK describes a response with status code 200, with default header values.

OK
*/
type PcloudNetworksPostOK struct {
	Payload *models.Network
}

// IsSuccess returns true when this pcloud networks post o k response has a 2xx status code
func (o *PcloudNetworksPostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud networks post o k response has a 3xx status code
func (o *PcloudNetworksPostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks post o k response has a 4xx status code
func (o *PcloudNetworksPostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud networks post o k response has a 5xx status code
func (o *PcloudNetworksPostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks post o k response a status code equal to that given
func (o *PcloudNetworksPostOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud networks post o k response
func (o *PcloudNetworksPostOK) Code() int {
	return 200
}

func (o *PcloudNetworksPostOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/networks][%d] pcloudNetworksPostOK %s", 200, payload)
}

func (o *PcloudNetworksPostOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/networks][%d] pcloudNetworksPostOK %s", 200, payload)
}

func (o *PcloudNetworksPostOK) GetPayload() *models.Network {
	return o.Payload
}

func (o *PcloudNetworksPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Network)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksPostCreated creates a PcloudNetworksPostCreated with default headers values
func NewPcloudNetworksPostCreated() *PcloudNetworksPostCreated {
	return &PcloudNetworksPostCreated{}
}

/*
PcloudNetworksPostCreated describes a response with status code 201, with default header values.

Created
*/
type PcloudNetworksPostCreated struct {
	Payload *models.Network
}

// IsSuccess returns true when this pcloud networks post created response has a 2xx status code
func (o *PcloudNetworksPostCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud networks post created response has a 3xx status code
func (o *PcloudNetworksPostCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks post created response has a 4xx status code
func (o *PcloudNetworksPostCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud networks post created response has a 5xx status code
func (o *PcloudNetworksPostCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks post created response a status code equal to that given
func (o *PcloudNetworksPostCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the pcloud networks post created response
func (o *PcloudNetworksPostCreated) Code() int {
	return 201
}

func (o *PcloudNetworksPostCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/networks][%d] pcloudNetworksPostCreated %s", 201, payload)
}

func (o *PcloudNetworksPostCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/networks][%d] pcloudNetworksPostCreated %s", 201, payload)
}

func (o *PcloudNetworksPostCreated) GetPayload() *models.Network {
	return o.Payload
}

func (o *PcloudNetworksPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Network)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksPostBadRequest creates a PcloudNetworksPostBadRequest with default headers values
func NewPcloudNetworksPostBadRequest() *PcloudNetworksPostBadRequest {
	return &PcloudNetworksPostBadRequest{}
}

/*
PcloudNetworksPostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudNetworksPostBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks post bad request response has a 2xx status code
func (o *PcloudNetworksPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks post bad request response has a 3xx status code
func (o *PcloudNetworksPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks post bad request response has a 4xx status code
func (o *PcloudNetworksPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud networks post bad request response has a 5xx status code
func (o *PcloudNetworksPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks post bad request response a status code equal to that given
func (o *PcloudNetworksPostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud networks post bad request response
func (o *PcloudNetworksPostBadRequest) Code() int {
	return 400
}

func (o *PcloudNetworksPostBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/networks][%d] pcloudNetworksPostBadRequest %s", 400, payload)
}

func (o *PcloudNetworksPostBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/networks][%d] pcloudNetworksPostBadRequest %s", 400, payload)
}

func (o *PcloudNetworksPostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksPostUnauthorized creates a PcloudNetworksPostUnauthorized with default headers values
func NewPcloudNetworksPostUnauthorized() *PcloudNetworksPostUnauthorized {
	return &PcloudNetworksPostUnauthorized{}
}

/*
PcloudNetworksPostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudNetworksPostUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks post unauthorized response has a 2xx status code
func (o *PcloudNetworksPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks post unauthorized response has a 3xx status code
func (o *PcloudNetworksPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks post unauthorized response has a 4xx status code
func (o *PcloudNetworksPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud networks post unauthorized response has a 5xx status code
func (o *PcloudNetworksPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks post unauthorized response a status code equal to that given
func (o *PcloudNetworksPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud networks post unauthorized response
func (o *PcloudNetworksPostUnauthorized) Code() int {
	return 401
}

func (o *PcloudNetworksPostUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/networks][%d] pcloudNetworksPostUnauthorized %s", 401, payload)
}

func (o *PcloudNetworksPostUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/networks][%d] pcloudNetworksPostUnauthorized %s", 401, payload)
}

func (o *PcloudNetworksPostUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksPostForbidden creates a PcloudNetworksPostForbidden with default headers values
func NewPcloudNetworksPostForbidden() *PcloudNetworksPostForbidden {
	return &PcloudNetworksPostForbidden{}
}

/*
PcloudNetworksPostForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudNetworksPostForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks post forbidden response has a 2xx status code
func (o *PcloudNetworksPostForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks post forbidden response has a 3xx status code
func (o *PcloudNetworksPostForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks post forbidden response has a 4xx status code
func (o *PcloudNetworksPostForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud networks post forbidden response has a 5xx status code
func (o *PcloudNetworksPostForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks post forbidden response a status code equal to that given
func (o *PcloudNetworksPostForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud networks post forbidden response
func (o *PcloudNetworksPostForbidden) Code() int {
	return 403
}

func (o *PcloudNetworksPostForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/networks][%d] pcloudNetworksPostForbidden %s", 403, payload)
}

func (o *PcloudNetworksPostForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/networks][%d] pcloudNetworksPostForbidden %s", 403, payload)
}

func (o *PcloudNetworksPostForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksPostNotFound creates a PcloudNetworksPostNotFound with default headers values
func NewPcloudNetworksPostNotFound() *PcloudNetworksPostNotFound {
	return &PcloudNetworksPostNotFound{}
}

/*
PcloudNetworksPostNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudNetworksPostNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks post not found response has a 2xx status code
func (o *PcloudNetworksPostNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks post not found response has a 3xx status code
func (o *PcloudNetworksPostNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks post not found response has a 4xx status code
func (o *PcloudNetworksPostNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud networks post not found response has a 5xx status code
func (o *PcloudNetworksPostNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks post not found response a status code equal to that given
func (o *PcloudNetworksPostNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud networks post not found response
func (o *PcloudNetworksPostNotFound) Code() int {
	return 404
}

func (o *PcloudNetworksPostNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/networks][%d] pcloudNetworksPostNotFound %s", 404, payload)
}

func (o *PcloudNetworksPostNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/networks][%d] pcloudNetworksPostNotFound %s", 404, payload)
}

func (o *PcloudNetworksPostNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksPostConflict creates a PcloudNetworksPostConflict with default headers values
func NewPcloudNetworksPostConflict() *PcloudNetworksPostConflict {
	return &PcloudNetworksPostConflict{}
}

/*
PcloudNetworksPostConflict describes a response with status code 409, with default header values.

Conflict
*/
type PcloudNetworksPostConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks post conflict response has a 2xx status code
func (o *PcloudNetworksPostConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks post conflict response has a 3xx status code
func (o *PcloudNetworksPostConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks post conflict response has a 4xx status code
func (o *PcloudNetworksPostConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud networks post conflict response has a 5xx status code
func (o *PcloudNetworksPostConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks post conflict response a status code equal to that given
func (o *PcloudNetworksPostConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the pcloud networks post conflict response
func (o *PcloudNetworksPostConflict) Code() int {
	return 409
}

func (o *PcloudNetworksPostConflict) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/networks][%d] pcloudNetworksPostConflict %s", 409, payload)
}

func (o *PcloudNetworksPostConflict) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/networks][%d] pcloudNetworksPostConflict %s", 409, payload)
}

func (o *PcloudNetworksPostConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksPostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksPostUnprocessableEntity creates a PcloudNetworksPostUnprocessableEntity with default headers values
func NewPcloudNetworksPostUnprocessableEntity() *PcloudNetworksPostUnprocessableEntity {
	return &PcloudNetworksPostUnprocessableEntity{}
}

/*
PcloudNetworksPostUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudNetworksPostUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks post unprocessable entity response has a 2xx status code
func (o *PcloudNetworksPostUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks post unprocessable entity response has a 3xx status code
func (o *PcloudNetworksPostUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks post unprocessable entity response has a 4xx status code
func (o *PcloudNetworksPostUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud networks post unprocessable entity response has a 5xx status code
func (o *PcloudNetworksPostUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks post unprocessable entity response a status code equal to that given
func (o *PcloudNetworksPostUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the pcloud networks post unprocessable entity response
func (o *PcloudNetworksPostUnprocessableEntity) Code() int {
	return 422
}

func (o *PcloudNetworksPostUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/networks][%d] pcloudNetworksPostUnprocessableEntity %s", 422, payload)
}

func (o *PcloudNetworksPostUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/networks][%d] pcloudNetworksPostUnprocessableEntity %s", 422, payload)
}

func (o *PcloudNetworksPostUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksPostInternalServerError creates a PcloudNetworksPostInternalServerError with default headers values
func NewPcloudNetworksPostInternalServerError() *PcloudNetworksPostInternalServerError {
	return &PcloudNetworksPostInternalServerError{}
}

/*
PcloudNetworksPostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudNetworksPostInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks post internal server error response has a 2xx status code
func (o *PcloudNetworksPostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks post internal server error response has a 3xx status code
func (o *PcloudNetworksPostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks post internal server error response has a 4xx status code
func (o *PcloudNetworksPostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud networks post internal server error response has a 5xx status code
func (o *PcloudNetworksPostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud networks post internal server error response a status code equal to that given
func (o *PcloudNetworksPostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud networks post internal server error response
func (o *PcloudNetworksPostInternalServerError) Code() int {
	return 500
}

func (o *PcloudNetworksPostInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/networks][%d] pcloudNetworksPostInternalServerError %s", 500, payload)
}

func (o *PcloudNetworksPostInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/networks][%d] pcloudNetworksPostInternalServerError %s", 500, payload)
}

func (o *PcloudNetworksPostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksPostStatus550 creates a PcloudNetworksPostStatus550 with default headers values
func NewPcloudNetworksPostStatus550() *PcloudNetworksPostStatus550 {
	return &PcloudNetworksPostStatus550{}
}

/*
PcloudNetworksPostStatus550 describes a response with status code 550, with default header values.

Workspace Status Error
*/
type PcloudNetworksPostStatus550 struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks post status550 response has a 2xx status code
func (o *PcloudNetworksPostStatus550) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks post status550 response has a 3xx status code
func (o *PcloudNetworksPostStatus550) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks post status550 response has a 4xx status code
func (o *PcloudNetworksPostStatus550) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud networks post status550 response has a 5xx status code
func (o *PcloudNetworksPostStatus550) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud networks post status550 response a status code equal to that given
func (o *PcloudNetworksPostStatus550) IsCode(code int) bool {
	return code == 550
}

// Code gets the status code for the pcloud networks post status550 response
func (o *PcloudNetworksPostStatus550) Code() int {
	return 550
}

func (o *PcloudNetworksPostStatus550) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/networks][%d] pcloudNetworksPostStatus550 %s", 550, payload)
}

func (o *PcloudNetworksPostStatus550) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/networks][%d] pcloudNetworksPostStatus550 %s", 550, payload)
}

func (o *PcloudNetworksPostStatus550) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksPostStatus550) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
