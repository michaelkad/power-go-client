// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_snapshots

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.ibm.com/power-iaas/service-broker/power/models"
)

// PcloudCloudinstancesSnapshotsPutReader is a Reader for the PcloudCloudinstancesSnapshotsPut structure.
type PcloudCloudinstancesSnapshotsPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudCloudinstancesSnapshotsPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudCloudinstancesSnapshotsPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudCloudinstancesSnapshotsPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudCloudinstancesSnapshotsPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudCloudinstancesSnapshotsPutForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudCloudinstancesSnapshotsPutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudCloudinstancesSnapshotsPutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots/{snapshot_id}] pcloud.cloudinstances.snapshots.put", response, response.Code())
	}
}

// NewPcloudCloudinstancesSnapshotsPutOK creates a PcloudCloudinstancesSnapshotsPutOK with default headers values
func NewPcloudCloudinstancesSnapshotsPutOK() *PcloudCloudinstancesSnapshotsPutOK {
	return &PcloudCloudinstancesSnapshotsPutOK{}
}

/*
PcloudCloudinstancesSnapshotsPutOK describes a response with status code 200, with default header values.

OK
*/
type PcloudCloudinstancesSnapshotsPutOK struct {
	Payload models.Object
}

// IsSuccess returns true when this pcloud cloudinstances snapshots put o k response has a 2xx status code
func (o *PcloudCloudinstancesSnapshotsPutOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud cloudinstances snapshots put o k response has a 3xx status code
func (o *PcloudCloudinstancesSnapshotsPutOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances snapshots put o k response has a 4xx status code
func (o *PcloudCloudinstancesSnapshotsPutOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudinstances snapshots put o k response has a 5xx status code
func (o *PcloudCloudinstancesSnapshotsPutOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances snapshots put o k response a status code equal to that given
func (o *PcloudCloudinstancesSnapshotsPutOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud cloudinstances snapshots put o k response
func (o *PcloudCloudinstancesSnapshotsPutOK) Code() int {
	return 200
}

func (o *PcloudCloudinstancesSnapshotsPutOK) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots/{snapshot_id}][%d] pcloudCloudinstancesSnapshotsPutOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsPutOK) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots/{snapshot_id}][%d] pcloudCloudinstancesSnapshotsPutOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsPutOK) GetPayload() models.Object {
	return o.Payload
}

func (o *PcloudCloudinstancesSnapshotsPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesSnapshotsPutBadRequest creates a PcloudCloudinstancesSnapshotsPutBadRequest with default headers values
func NewPcloudCloudinstancesSnapshotsPutBadRequest() *PcloudCloudinstancesSnapshotsPutBadRequest {
	return &PcloudCloudinstancesSnapshotsPutBadRequest{}
}

/*
PcloudCloudinstancesSnapshotsPutBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudCloudinstancesSnapshotsPutBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances snapshots put bad request response has a 2xx status code
func (o *PcloudCloudinstancesSnapshotsPutBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances snapshots put bad request response has a 3xx status code
func (o *PcloudCloudinstancesSnapshotsPutBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances snapshots put bad request response has a 4xx status code
func (o *PcloudCloudinstancesSnapshotsPutBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances snapshots put bad request response has a 5xx status code
func (o *PcloudCloudinstancesSnapshotsPutBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances snapshots put bad request response a status code equal to that given
func (o *PcloudCloudinstancesSnapshotsPutBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud cloudinstances snapshots put bad request response
func (o *PcloudCloudinstancesSnapshotsPutBadRequest) Code() int {
	return 400
}

func (o *PcloudCloudinstancesSnapshotsPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots/{snapshot_id}][%d] pcloudCloudinstancesSnapshotsPutBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsPutBadRequest) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots/{snapshot_id}][%d] pcloudCloudinstancesSnapshotsPutBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsPutBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesSnapshotsPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesSnapshotsPutUnauthorized creates a PcloudCloudinstancesSnapshotsPutUnauthorized with default headers values
func NewPcloudCloudinstancesSnapshotsPutUnauthorized() *PcloudCloudinstancesSnapshotsPutUnauthorized {
	return &PcloudCloudinstancesSnapshotsPutUnauthorized{}
}

/*
PcloudCloudinstancesSnapshotsPutUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudCloudinstancesSnapshotsPutUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances snapshots put unauthorized response has a 2xx status code
func (o *PcloudCloudinstancesSnapshotsPutUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances snapshots put unauthorized response has a 3xx status code
func (o *PcloudCloudinstancesSnapshotsPutUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances snapshots put unauthorized response has a 4xx status code
func (o *PcloudCloudinstancesSnapshotsPutUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances snapshots put unauthorized response has a 5xx status code
func (o *PcloudCloudinstancesSnapshotsPutUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances snapshots put unauthorized response a status code equal to that given
func (o *PcloudCloudinstancesSnapshotsPutUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud cloudinstances snapshots put unauthorized response
func (o *PcloudCloudinstancesSnapshotsPutUnauthorized) Code() int {
	return 401
}

func (o *PcloudCloudinstancesSnapshotsPutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots/{snapshot_id}][%d] pcloudCloudinstancesSnapshotsPutUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsPutUnauthorized) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots/{snapshot_id}][%d] pcloudCloudinstancesSnapshotsPutUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsPutUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesSnapshotsPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesSnapshotsPutForbidden creates a PcloudCloudinstancesSnapshotsPutForbidden with default headers values
func NewPcloudCloudinstancesSnapshotsPutForbidden() *PcloudCloudinstancesSnapshotsPutForbidden {
	return &PcloudCloudinstancesSnapshotsPutForbidden{}
}

/*
PcloudCloudinstancesSnapshotsPutForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudCloudinstancesSnapshotsPutForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances snapshots put forbidden response has a 2xx status code
func (o *PcloudCloudinstancesSnapshotsPutForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances snapshots put forbidden response has a 3xx status code
func (o *PcloudCloudinstancesSnapshotsPutForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances snapshots put forbidden response has a 4xx status code
func (o *PcloudCloudinstancesSnapshotsPutForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances snapshots put forbidden response has a 5xx status code
func (o *PcloudCloudinstancesSnapshotsPutForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances snapshots put forbidden response a status code equal to that given
func (o *PcloudCloudinstancesSnapshotsPutForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud cloudinstances snapshots put forbidden response
func (o *PcloudCloudinstancesSnapshotsPutForbidden) Code() int {
	return 403
}

func (o *PcloudCloudinstancesSnapshotsPutForbidden) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots/{snapshot_id}][%d] pcloudCloudinstancesSnapshotsPutForbidden  %+v", 403, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsPutForbidden) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots/{snapshot_id}][%d] pcloudCloudinstancesSnapshotsPutForbidden  %+v", 403, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsPutForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesSnapshotsPutForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesSnapshotsPutNotFound creates a PcloudCloudinstancesSnapshotsPutNotFound with default headers values
func NewPcloudCloudinstancesSnapshotsPutNotFound() *PcloudCloudinstancesSnapshotsPutNotFound {
	return &PcloudCloudinstancesSnapshotsPutNotFound{}
}

/*
PcloudCloudinstancesSnapshotsPutNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudCloudinstancesSnapshotsPutNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances snapshots put not found response has a 2xx status code
func (o *PcloudCloudinstancesSnapshotsPutNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances snapshots put not found response has a 3xx status code
func (o *PcloudCloudinstancesSnapshotsPutNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances snapshots put not found response has a 4xx status code
func (o *PcloudCloudinstancesSnapshotsPutNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances snapshots put not found response has a 5xx status code
func (o *PcloudCloudinstancesSnapshotsPutNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances snapshots put not found response a status code equal to that given
func (o *PcloudCloudinstancesSnapshotsPutNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud cloudinstances snapshots put not found response
func (o *PcloudCloudinstancesSnapshotsPutNotFound) Code() int {
	return 404
}

func (o *PcloudCloudinstancesSnapshotsPutNotFound) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots/{snapshot_id}][%d] pcloudCloudinstancesSnapshotsPutNotFound  %+v", 404, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsPutNotFound) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots/{snapshot_id}][%d] pcloudCloudinstancesSnapshotsPutNotFound  %+v", 404, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsPutNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesSnapshotsPutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesSnapshotsPutInternalServerError creates a PcloudCloudinstancesSnapshotsPutInternalServerError with default headers values
func NewPcloudCloudinstancesSnapshotsPutInternalServerError() *PcloudCloudinstancesSnapshotsPutInternalServerError {
	return &PcloudCloudinstancesSnapshotsPutInternalServerError{}
}

/*
PcloudCloudinstancesSnapshotsPutInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudCloudinstancesSnapshotsPutInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances snapshots put internal server error response has a 2xx status code
func (o *PcloudCloudinstancesSnapshotsPutInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances snapshots put internal server error response has a 3xx status code
func (o *PcloudCloudinstancesSnapshotsPutInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances snapshots put internal server error response has a 4xx status code
func (o *PcloudCloudinstancesSnapshotsPutInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudinstances snapshots put internal server error response has a 5xx status code
func (o *PcloudCloudinstancesSnapshotsPutInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud cloudinstances snapshots put internal server error response a status code equal to that given
func (o *PcloudCloudinstancesSnapshotsPutInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud cloudinstances snapshots put internal server error response
func (o *PcloudCloudinstancesSnapshotsPutInternalServerError) Code() int {
	return 500
}

func (o *PcloudCloudinstancesSnapshotsPutInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots/{snapshot_id}][%d] pcloudCloudinstancesSnapshotsPutInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsPutInternalServerError) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots/{snapshot_id}][%d] pcloudCloudinstancesSnapshotsPutInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsPutInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesSnapshotsPutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
