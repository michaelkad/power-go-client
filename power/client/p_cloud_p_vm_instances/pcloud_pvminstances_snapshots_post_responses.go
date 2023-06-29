// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_p_vm_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudPvminstancesSnapshotsPostReader is a Reader for the PcloudPvminstancesSnapshotsPost structure.
type PcloudPvminstancesSnapshotsPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPvminstancesSnapshotsPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPcloudPvminstancesSnapshotsPostAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudPvminstancesSnapshotsPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudPvminstancesSnapshotsPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudPvminstancesSnapshotsPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudPvminstancesSnapshotsPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPcloudPvminstancesSnapshotsPostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudPvminstancesSnapshotsPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPcloudPvminstancesSnapshotsPostGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/snapshots] pcloud.pvminstances.snapshots.post", response, response.Code())
	}
}

// NewPcloudPvminstancesSnapshotsPostAccepted creates a PcloudPvminstancesSnapshotsPostAccepted with default headers values
func NewPcloudPvminstancesSnapshotsPostAccepted() *PcloudPvminstancesSnapshotsPostAccepted {
	return &PcloudPvminstancesSnapshotsPostAccepted{}
}

/*
PcloudPvminstancesSnapshotsPostAccepted describes a response with status code 202, with default header values.

Accepted
*/
type PcloudPvminstancesSnapshotsPostAccepted struct {
	Payload *models.SnapshotCreateResponse
}

// IsSuccess returns true when this pcloud pvminstances snapshots post accepted response has a 2xx status code
func (o *PcloudPvminstancesSnapshotsPostAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud pvminstances snapshots post accepted response has a 3xx status code
func (o *PcloudPvminstancesSnapshotsPostAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances snapshots post accepted response has a 4xx status code
func (o *PcloudPvminstancesSnapshotsPostAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud pvminstances snapshots post accepted response has a 5xx status code
func (o *PcloudPvminstancesSnapshotsPostAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances snapshots post accepted response a status code equal to that given
func (o *PcloudPvminstancesSnapshotsPostAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the pcloud pvminstances snapshots post accepted response
func (o *PcloudPvminstancesSnapshotsPostAccepted) Code() int {
	return 202
}

func (o *PcloudPvminstancesSnapshotsPostAccepted) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/snapshots][%d] pcloudPvminstancesSnapshotsPostAccepted  %+v", 202, o.Payload)
}

func (o *PcloudPvminstancesSnapshotsPostAccepted) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/snapshots][%d] pcloudPvminstancesSnapshotsPostAccepted  %+v", 202, o.Payload)
}

func (o *PcloudPvminstancesSnapshotsPostAccepted) GetPayload() *models.SnapshotCreateResponse {
	return o.Payload
}

func (o *PcloudPvminstancesSnapshotsPostAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnapshotCreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesSnapshotsPostBadRequest creates a PcloudPvminstancesSnapshotsPostBadRequest with default headers values
func NewPcloudPvminstancesSnapshotsPostBadRequest() *PcloudPvminstancesSnapshotsPostBadRequest {
	return &PcloudPvminstancesSnapshotsPostBadRequest{}
}

/*
PcloudPvminstancesSnapshotsPostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudPvminstancesSnapshotsPostBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances snapshots post bad request response has a 2xx status code
func (o *PcloudPvminstancesSnapshotsPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances snapshots post bad request response has a 3xx status code
func (o *PcloudPvminstancesSnapshotsPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances snapshots post bad request response has a 4xx status code
func (o *PcloudPvminstancesSnapshotsPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances snapshots post bad request response has a 5xx status code
func (o *PcloudPvminstancesSnapshotsPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances snapshots post bad request response a status code equal to that given
func (o *PcloudPvminstancesSnapshotsPostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud pvminstances snapshots post bad request response
func (o *PcloudPvminstancesSnapshotsPostBadRequest) Code() int {
	return 400
}

func (o *PcloudPvminstancesSnapshotsPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/snapshots][%d] pcloudPvminstancesSnapshotsPostBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudPvminstancesSnapshotsPostBadRequest) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/snapshots][%d] pcloudPvminstancesSnapshotsPostBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudPvminstancesSnapshotsPostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesSnapshotsPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesSnapshotsPostUnauthorized creates a PcloudPvminstancesSnapshotsPostUnauthorized with default headers values
func NewPcloudPvminstancesSnapshotsPostUnauthorized() *PcloudPvminstancesSnapshotsPostUnauthorized {
	return &PcloudPvminstancesSnapshotsPostUnauthorized{}
}

/*
PcloudPvminstancesSnapshotsPostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudPvminstancesSnapshotsPostUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances snapshots post unauthorized response has a 2xx status code
func (o *PcloudPvminstancesSnapshotsPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances snapshots post unauthorized response has a 3xx status code
func (o *PcloudPvminstancesSnapshotsPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances snapshots post unauthorized response has a 4xx status code
func (o *PcloudPvminstancesSnapshotsPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances snapshots post unauthorized response has a 5xx status code
func (o *PcloudPvminstancesSnapshotsPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances snapshots post unauthorized response a status code equal to that given
func (o *PcloudPvminstancesSnapshotsPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud pvminstances snapshots post unauthorized response
func (o *PcloudPvminstancesSnapshotsPostUnauthorized) Code() int {
	return 401
}

func (o *PcloudPvminstancesSnapshotsPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/snapshots][%d] pcloudPvminstancesSnapshotsPostUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudPvminstancesSnapshotsPostUnauthorized) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/snapshots][%d] pcloudPvminstancesSnapshotsPostUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudPvminstancesSnapshotsPostUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesSnapshotsPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesSnapshotsPostForbidden creates a PcloudPvminstancesSnapshotsPostForbidden with default headers values
func NewPcloudPvminstancesSnapshotsPostForbidden() *PcloudPvminstancesSnapshotsPostForbidden {
	return &PcloudPvminstancesSnapshotsPostForbidden{}
}

/*
PcloudPvminstancesSnapshotsPostForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudPvminstancesSnapshotsPostForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances snapshots post forbidden response has a 2xx status code
func (o *PcloudPvminstancesSnapshotsPostForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances snapshots post forbidden response has a 3xx status code
func (o *PcloudPvminstancesSnapshotsPostForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances snapshots post forbidden response has a 4xx status code
func (o *PcloudPvminstancesSnapshotsPostForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances snapshots post forbidden response has a 5xx status code
func (o *PcloudPvminstancesSnapshotsPostForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances snapshots post forbidden response a status code equal to that given
func (o *PcloudPvminstancesSnapshotsPostForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud pvminstances snapshots post forbidden response
func (o *PcloudPvminstancesSnapshotsPostForbidden) Code() int {
	return 403
}

func (o *PcloudPvminstancesSnapshotsPostForbidden) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/snapshots][%d] pcloudPvminstancesSnapshotsPostForbidden  %+v", 403, o.Payload)
}

func (o *PcloudPvminstancesSnapshotsPostForbidden) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/snapshots][%d] pcloudPvminstancesSnapshotsPostForbidden  %+v", 403, o.Payload)
}

func (o *PcloudPvminstancesSnapshotsPostForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesSnapshotsPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesSnapshotsPostNotFound creates a PcloudPvminstancesSnapshotsPostNotFound with default headers values
func NewPcloudPvminstancesSnapshotsPostNotFound() *PcloudPvminstancesSnapshotsPostNotFound {
	return &PcloudPvminstancesSnapshotsPostNotFound{}
}

/*
PcloudPvminstancesSnapshotsPostNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudPvminstancesSnapshotsPostNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances snapshots post not found response has a 2xx status code
func (o *PcloudPvminstancesSnapshotsPostNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances snapshots post not found response has a 3xx status code
func (o *PcloudPvminstancesSnapshotsPostNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances snapshots post not found response has a 4xx status code
func (o *PcloudPvminstancesSnapshotsPostNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances snapshots post not found response has a 5xx status code
func (o *PcloudPvminstancesSnapshotsPostNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances snapshots post not found response a status code equal to that given
func (o *PcloudPvminstancesSnapshotsPostNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud pvminstances snapshots post not found response
func (o *PcloudPvminstancesSnapshotsPostNotFound) Code() int {
	return 404
}

func (o *PcloudPvminstancesSnapshotsPostNotFound) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/snapshots][%d] pcloudPvminstancesSnapshotsPostNotFound  %+v", 404, o.Payload)
}

func (o *PcloudPvminstancesSnapshotsPostNotFound) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/snapshots][%d] pcloudPvminstancesSnapshotsPostNotFound  %+v", 404, o.Payload)
}

func (o *PcloudPvminstancesSnapshotsPostNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesSnapshotsPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesSnapshotsPostConflict creates a PcloudPvminstancesSnapshotsPostConflict with default headers values
func NewPcloudPvminstancesSnapshotsPostConflict() *PcloudPvminstancesSnapshotsPostConflict {
	return &PcloudPvminstancesSnapshotsPostConflict{}
}

/*
PcloudPvminstancesSnapshotsPostConflict describes a response with status code 409, with default header values.

Conflict
*/
type PcloudPvminstancesSnapshotsPostConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances snapshots post conflict response has a 2xx status code
func (o *PcloudPvminstancesSnapshotsPostConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances snapshots post conflict response has a 3xx status code
func (o *PcloudPvminstancesSnapshotsPostConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances snapshots post conflict response has a 4xx status code
func (o *PcloudPvminstancesSnapshotsPostConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances snapshots post conflict response has a 5xx status code
func (o *PcloudPvminstancesSnapshotsPostConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances snapshots post conflict response a status code equal to that given
func (o *PcloudPvminstancesSnapshotsPostConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the pcloud pvminstances snapshots post conflict response
func (o *PcloudPvminstancesSnapshotsPostConflict) Code() int {
	return 409
}

func (o *PcloudPvminstancesSnapshotsPostConflict) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/snapshots][%d] pcloudPvminstancesSnapshotsPostConflict  %+v", 409, o.Payload)
}

func (o *PcloudPvminstancesSnapshotsPostConflict) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/snapshots][%d] pcloudPvminstancesSnapshotsPostConflict  %+v", 409, o.Payload)
}

func (o *PcloudPvminstancesSnapshotsPostConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesSnapshotsPostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesSnapshotsPostInternalServerError creates a PcloudPvminstancesSnapshotsPostInternalServerError with default headers values
func NewPcloudPvminstancesSnapshotsPostInternalServerError() *PcloudPvminstancesSnapshotsPostInternalServerError {
	return &PcloudPvminstancesSnapshotsPostInternalServerError{}
}

/*
PcloudPvminstancesSnapshotsPostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudPvminstancesSnapshotsPostInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances snapshots post internal server error response has a 2xx status code
func (o *PcloudPvminstancesSnapshotsPostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances snapshots post internal server error response has a 3xx status code
func (o *PcloudPvminstancesSnapshotsPostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances snapshots post internal server error response has a 4xx status code
func (o *PcloudPvminstancesSnapshotsPostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud pvminstances snapshots post internal server error response has a 5xx status code
func (o *PcloudPvminstancesSnapshotsPostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud pvminstances snapshots post internal server error response a status code equal to that given
func (o *PcloudPvminstancesSnapshotsPostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud pvminstances snapshots post internal server error response
func (o *PcloudPvminstancesSnapshotsPostInternalServerError) Code() int {
	return 500
}

func (o *PcloudPvminstancesSnapshotsPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/snapshots][%d] pcloudPvminstancesSnapshotsPostInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPvminstancesSnapshotsPostInternalServerError) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/snapshots][%d] pcloudPvminstancesSnapshotsPostInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPvminstancesSnapshotsPostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesSnapshotsPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesSnapshotsPostGatewayTimeout creates a PcloudPvminstancesSnapshotsPostGatewayTimeout with default headers values
func NewPcloudPvminstancesSnapshotsPostGatewayTimeout() *PcloudPvminstancesSnapshotsPostGatewayTimeout {
	return &PcloudPvminstancesSnapshotsPostGatewayTimeout{}
}

/*
PcloudPvminstancesSnapshotsPostGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout
*/
type PcloudPvminstancesSnapshotsPostGatewayTimeout struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances snapshots post gateway timeout response has a 2xx status code
func (o *PcloudPvminstancesSnapshotsPostGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances snapshots post gateway timeout response has a 3xx status code
func (o *PcloudPvminstancesSnapshotsPostGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances snapshots post gateway timeout response has a 4xx status code
func (o *PcloudPvminstancesSnapshotsPostGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud pvminstances snapshots post gateway timeout response has a 5xx status code
func (o *PcloudPvminstancesSnapshotsPostGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud pvminstances snapshots post gateway timeout response a status code equal to that given
func (o *PcloudPvminstancesSnapshotsPostGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the pcloud pvminstances snapshots post gateway timeout response
func (o *PcloudPvminstancesSnapshotsPostGatewayTimeout) Code() int {
	return 504
}

func (o *PcloudPvminstancesSnapshotsPostGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/snapshots][%d] pcloudPvminstancesSnapshotsPostGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PcloudPvminstancesSnapshotsPostGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/snapshots][%d] pcloudPvminstancesSnapshotsPostGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PcloudPvminstancesSnapshotsPostGatewayTimeout) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesSnapshotsPostGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
