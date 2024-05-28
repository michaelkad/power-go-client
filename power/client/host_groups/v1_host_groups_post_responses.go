// Code generated by go-swagger; DO NOT EDIT.

package host_groups

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

// V1HostGroupsPostReader is a Reader for the V1HostGroupsPost structure.
type V1HostGroupsPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1HostGroupsPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewV1HostGroupsPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewV1HostGroupsPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewV1HostGroupsPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV1HostGroupsPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewV1HostGroupsPostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewV1HostGroupsPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV1HostGroupsPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewV1HostGroupsPostGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/host-groups] v1.hostGroups.post", response, response.Code())
	}
}

// NewV1HostGroupsPostCreated creates a V1HostGroupsPostCreated with default headers values
func NewV1HostGroupsPostCreated() *V1HostGroupsPostCreated {
	return &V1HostGroupsPostCreated{}
}

/*
V1HostGroupsPostCreated describes a response with status code 201, with default header values.

Created
*/
type V1HostGroupsPostCreated struct {
	Payload *models.HostGroup
}

// IsSuccess returns true when this v1 host groups post created response has a 2xx status code
func (o *V1HostGroupsPostCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this v1 host groups post created response has a 3xx status code
func (o *V1HostGroupsPostCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 host groups post created response has a 4xx status code
func (o *V1HostGroupsPostCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 host groups post created response has a 5xx status code
func (o *V1HostGroupsPostCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 host groups post created response a status code equal to that given
func (o *V1HostGroupsPostCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the v1 host groups post created response
func (o *V1HostGroupsPostCreated) Code() int {
	return 201
}

func (o *V1HostGroupsPostCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/host-groups][%d] v1HostGroupsPostCreated %s", 201, payload)
}

func (o *V1HostGroupsPostCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/host-groups][%d] v1HostGroupsPostCreated %s", 201, payload)
}

func (o *V1HostGroupsPostCreated) GetPayload() *models.HostGroup {
	return o.Payload
}

func (o *V1HostGroupsPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HostGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1HostGroupsPostBadRequest creates a V1HostGroupsPostBadRequest with default headers values
func NewV1HostGroupsPostBadRequest() *V1HostGroupsPostBadRequest {
	return &V1HostGroupsPostBadRequest{}
}

/*
V1HostGroupsPostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type V1HostGroupsPostBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 host groups post bad request response has a 2xx status code
func (o *V1HostGroupsPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 host groups post bad request response has a 3xx status code
func (o *V1HostGroupsPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 host groups post bad request response has a 4xx status code
func (o *V1HostGroupsPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 host groups post bad request response has a 5xx status code
func (o *V1HostGroupsPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 host groups post bad request response a status code equal to that given
func (o *V1HostGroupsPostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the v1 host groups post bad request response
func (o *V1HostGroupsPostBadRequest) Code() int {
	return 400
}

func (o *V1HostGroupsPostBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/host-groups][%d] v1HostGroupsPostBadRequest %s", 400, payload)
}

func (o *V1HostGroupsPostBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/host-groups][%d] v1HostGroupsPostBadRequest %s", 400, payload)
}

func (o *V1HostGroupsPostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1HostGroupsPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1HostGroupsPostUnauthorized creates a V1HostGroupsPostUnauthorized with default headers values
func NewV1HostGroupsPostUnauthorized() *V1HostGroupsPostUnauthorized {
	return &V1HostGroupsPostUnauthorized{}
}

/*
V1HostGroupsPostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type V1HostGroupsPostUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 host groups post unauthorized response has a 2xx status code
func (o *V1HostGroupsPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 host groups post unauthorized response has a 3xx status code
func (o *V1HostGroupsPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 host groups post unauthorized response has a 4xx status code
func (o *V1HostGroupsPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 host groups post unauthorized response has a 5xx status code
func (o *V1HostGroupsPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 host groups post unauthorized response a status code equal to that given
func (o *V1HostGroupsPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the v1 host groups post unauthorized response
func (o *V1HostGroupsPostUnauthorized) Code() int {
	return 401
}

func (o *V1HostGroupsPostUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/host-groups][%d] v1HostGroupsPostUnauthorized %s", 401, payload)
}

func (o *V1HostGroupsPostUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/host-groups][%d] v1HostGroupsPostUnauthorized %s", 401, payload)
}

func (o *V1HostGroupsPostUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1HostGroupsPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1HostGroupsPostForbidden creates a V1HostGroupsPostForbidden with default headers values
func NewV1HostGroupsPostForbidden() *V1HostGroupsPostForbidden {
	return &V1HostGroupsPostForbidden{}
}

/*
V1HostGroupsPostForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type V1HostGroupsPostForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 host groups post forbidden response has a 2xx status code
func (o *V1HostGroupsPostForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 host groups post forbidden response has a 3xx status code
func (o *V1HostGroupsPostForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 host groups post forbidden response has a 4xx status code
func (o *V1HostGroupsPostForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 host groups post forbidden response has a 5xx status code
func (o *V1HostGroupsPostForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 host groups post forbidden response a status code equal to that given
func (o *V1HostGroupsPostForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the v1 host groups post forbidden response
func (o *V1HostGroupsPostForbidden) Code() int {
	return 403
}

func (o *V1HostGroupsPostForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/host-groups][%d] v1HostGroupsPostForbidden %s", 403, payload)
}

func (o *V1HostGroupsPostForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/host-groups][%d] v1HostGroupsPostForbidden %s", 403, payload)
}

func (o *V1HostGroupsPostForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1HostGroupsPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1HostGroupsPostConflict creates a V1HostGroupsPostConflict with default headers values
func NewV1HostGroupsPostConflict() *V1HostGroupsPostConflict {
	return &V1HostGroupsPostConflict{}
}

/*
V1HostGroupsPostConflict describes a response with status code 409, with default header values.

Conflict
*/
type V1HostGroupsPostConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 host groups post conflict response has a 2xx status code
func (o *V1HostGroupsPostConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 host groups post conflict response has a 3xx status code
func (o *V1HostGroupsPostConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 host groups post conflict response has a 4xx status code
func (o *V1HostGroupsPostConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 host groups post conflict response has a 5xx status code
func (o *V1HostGroupsPostConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 host groups post conflict response a status code equal to that given
func (o *V1HostGroupsPostConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the v1 host groups post conflict response
func (o *V1HostGroupsPostConflict) Code() int {
	return 409
}

func (o *V1HostGroupsPostConflict) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/host-groups][%d] v1HostGroupsPostConflict %s", 409, payload)
}

func (o *V1HostGroupsPostConflict) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/host-groups][%d] v1HostGroupsPostConflict %s", 409, payload)
}

func (o *V1HostGroupsPostConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1HostGroupsPostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1HostGroupsPostUnprocessableEntity creates a V1HostGroupsPostUnprocessableEntity with default headers values
func NewV1HostGroupsPostUnprocessableEntity() *V1HostGroupsPostUnprocessableEntity {
	return &V1HostGroupsPostUnprocessableEntity{}
}

/*
V1HostGroupsPostUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type V1HostGroupsPostUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 host groups post unprocessable entity response has a 2xx status code
func (o *V1HostGroupsPostUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 host groups post unprocessable entity response has a 3xx status code
func (o *V1HostGroupsPostUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 host groups post unprocessable entity response has a 4xx status code
func (o *V1HostGroupsPostUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 host groups post unprocessable entity response has a 5xx status code
func (o *V1HostGroupsPostUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 host groups post unprocessable entity response a status code equal to that given
func (o *V1HostGroupsPostUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the v1 host groups post unprocessable entity response
func (o *V1HostGroupsPostUnprocessableEntity) Code() int {
	return 422
}

func (o *V1HostGroupsPostUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/host-groups][%d] v1HostGroupsPostUnprocessableEntity %s", 422, payload)
}

func (o *V1HostGroupsPostUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/host-groups][%d] v1HostGroupsPostUnprocessableEntity %s", 422, payload)
}

func (o *V1HostGroupsPostUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1HostGroupsPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1HostGroupsPostInternalServerError creates a V1HostGroupsPostInternalServerError with default headers values
func NewV1HostGroupsPostInternalServerError() *V1HostGroupsPostInternalServerError {
	return &V1HostGroupsPostInternalServerError{}
}

/*
V1HostGroupsPostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type V1HostGroupsPostInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 host groups post internal server error response has a 2xx status code
func (o *V1HostGroupsPostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 host groups post internal server error response has a 3xx status code
func (o *V1HostGroupsPostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 host groups post internal server error response has a 4xx status code
func (o *V1HostGroupsPostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 host groups post internal server error response has a 5xx status code
func (o *V1HostGroupsPostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this v1 host groups post internal server error response a status code equal to that given
func (o *V1HostGroupsPostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the v1 host groups post internal server error response
func (o *V1HostGroupsPostInternalServerError) Code() int {
	return 500
}

func (o *V1HostGroupsPostInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/host-groups][%d] v1HostGroupsPostInternalServerError %s", 500, payload)
}

func (o *V1HostGroupsPostInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/host-groups][%d] v1HostGroupsPostInternalServerError %s", 500, payload)
}

func (o *V1HostGroupsPostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1HostGroupsPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1HostGroupsPostGatewayTimeout creates a V1HostGroupsPostGatewayTimeout with default headers values
func NewV1HostGroupsPostGatewayTimeout() *V1HostGroupsPostGatewayTimeout {
	return &V1HostGroupsPostGatewayTimeout{}
}

/*
V1HostGroupsPostGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. Request is still processing and taking longer than expected.
*/
type V1HostGroupsPostGatewayTimeout struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 host groups post gateway timeout response has a 2xx status code
func (o *V1HostGroupsPostGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 host groups post gateway timeout response has a 3xx status code
func (o *V1HostGroupsPostGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 host groups post gateway timeout response has a 4xx status code
func (o *V1HostGroupsPostGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 host groups post gateway timeout response has a 5xx status code
func (o *V1HostGroupsPostGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this v1 host groups post gateway timeout response a status code equal to that given
func (o *V1HostGroupsPostGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the v1 host groups post gateway timeout response
func (o *V1HostGroupsPostGatewayTimeout) Code() int {
	return 504
}

func (o *V1HostGroupsPostGatewayTimeout) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/host-groups][%d] v1HostGroupsPostGatewayTimeout %s", 504, payload)
}

func (o *V1HostGroupsPostGatewayTimeout) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/host-groups][%d] v1HostGroupsPostGatewayTimeout %s", 504, payload)
}

func (o *V1HostGroupsPostGatewayTimeout) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1HostGroupsPostGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
