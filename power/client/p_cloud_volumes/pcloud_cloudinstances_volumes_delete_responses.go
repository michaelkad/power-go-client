// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.ibm.com/power-iaas/service-broker/power/models"
)

// PcloudCloudinstancesVolumesDeleteReader is a Reader for the PcloudCloudinstancesVolumesDelete structure.
type PcloudCloudinstancesVolumesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudCloudinstancesVolumesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudCloudinstancesVolumesDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudCloudinstancesVolumesDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudCloudinstancesVolumesDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudCloudinstancesVolumesDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudCloudinstancesVolumesDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPcloudCloudinstancesVolumesDeleteGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudCloudinstancesVolumesDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}] pcloud.cloudinstances.volumes.delete", response, response.Code())
	}
}

// NewPcloudCloudinstancesVolumesDeleteOK creates a PcloudCloudinstancesVolumesDeleteOK with default headers values
func NewPcloudCloudinstancesVolumesDeleteOK() *PcloudCloudinstancesVolumesDeleteOK {
	return &PcloudCloudinstancesVolumesDeleteOK{}
}

/*
PcloudCloudinstancesVolumesDeleteOK describes a response with status code 200, with default header values.

OK
*/
type PcloudCloudinstancesVolumesDeleteOK struct {
	Payload models.Object
}

// IsSuccess returns true when this pcloud cloudinstances volumes delete o k response has a 2xx status code
func (o *PcloudCloudinstancesVolumesDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud cloudinstances volumes delete o k response has a 3xx status code
func (o *PcloudCloudinstancesVolumesDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances volumes delete o k response has a 4xx status code
func (o *PcloudCloudinstancesVolumesDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudinstances volumes delete o k response has a 5xx status code
func (o *PcloudCloudinstancesVolumesDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances volumes delete o k response a status code equal to that given
func (o *PcloudCloudinstancesVolumesDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud cloudinstances volumes delete o k response
func (o *PcloudCloudinstancesVolumesDeleteOK) Code() int {
	return 200
}

func (o *PcloudCloudinstancesVolumesDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}][%d] pcloudCloudinstancesVolumesDeleteOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudinstancesVolumesDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}][%d] pcloudCloudinstancesVolumesDeleteOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudinstancesVolumesDeleteOK) GetPayload() models.Object {
	return o.Payload
}

func (o *PcloudCloudinstancesVolumesDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesVolumesDeleteBadRequest creates a PcloudCloudinstancesVolumesDeleteBadRequest with default headers values
func NewPcloudCloudinstancesVolumesDeleteBadRequest() *PcloudCloudinstancesVolumesDeleteBadRequest {
	return &PcloudCloudinstancesVolumesDeleteBadRequest{}
}

/*
PcloudCloudinstancesVolumesDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudCloudinstancesVolumesDeleteBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances volumes delete bad request response has a 2xx status code
func (o *PcloudCloudinstancesVolumesDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances volumes delete bad request response has a 3xx status code
func (o *PcloudCloudinstancesVolumesDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances volumes delete bad request response has a 4xx status code
func (o *PcloudCloudinstancesVolumesDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances volumes delete bad request response has a 5xx status code
func (o *PcloudCloudinstancesVolumesDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances volumes delete bad request response a status code equal to that given
func (o *PcloudCloudinstancesVolumesDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud cloudinstances volumes delete bad request response
func (o *PcloudCloudinstancesVolumesDeleteBadRequest) Code() int {
	return 400
}

func (o *PcloudCloudinstancesVolumesDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}][%d] pcloudCloudinstancesVolumesDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudinstancesVolumesDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}][%d] pcloudCloudinstancesVolumesDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudinstancesVolumesDeleteBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesVolumesDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesVolumesDeleteUnauthorized creates a PcloudCloudinstancesVolumesDeleteUnauthorized with default headers values
func NewPcloudCloudinstancesVolumesDeleteUnauthorized() *PcloudCloudinstancesVolumesDeleteUnauthorized {
	return &PcloudCloudinstancesVolumesDeleteUnauthorized{}
}

/*
PcloudCloudinstancesVolumesDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudCloudinstancesVolumesDeleteUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances volumes delete unauthorized response has a 2xx status code
func (o *PcloudCloudinstancesVolumesDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances volumes delete unauthorized response has a 3xx status code
func (o *PcloudCloudinstancesVolumesDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances volumes delete unauthorized response has a 4xx status code
func (o *PcloudCloudinstancesVolumesDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances volumes delete unauthorized response has a 5xx status code
func (o *PcloudCloudinstancesVolumesDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances volumes delete unauthorized response a status code equal to that given
func (o *PcloudCloudinstancesVolumesDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud cloudinstances volumes delete unauthorized response
func (o *PcloudCloudinstancesVolumesDeleteUnauthorized) Code() int {
	return 401
}

func (o *PcloudCloudinstancesVolumesDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}][%d] pcloudCloudinstancesVolumesDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudinstancesVolumesDeleteUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}][%d] pcloudCloudinstancesVolumesDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudinstancesVolumesDeleteUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesVolumesDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesVolumesDeleteForbidden creates a PcloudCloudinstancesVolumesDeleteForbidden with default headers values
func NewPcloudCloudinstancesVolumesDeleteForbidden() *PcloudCloudinstancesVolumesDeleteForbidden {
	return &PcloudCloudinstancesVolumesDeleteForbidden{}
}

/*
PcloudCloudinstancesVolumesDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudCloudinstancesVolumesDeleteForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances volumes delete forbidden response has a 2xx status code
func (o *PcloudCloudinstancesVolumesDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances volumes delete forbidden response has a 3xx status code
func (o *PcloudCloudinstancesVolumesDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances volumes delete forbidden response has a 4xx status code
func (o *PcloudCloudinstancesVolumesDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances volumes delete forbidden response has a 5xx status code
func (o *PcloudCloudinstancesVolumesDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances volumes delete forbidden response a status code equal to that given
func (o *PcloudCloudinstancesVolumesDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud cloudinstances volumes delete forbidden response
func (o *PcloudCloudinstancesVolumesDeleteForbidden) Code() int {
	return 403
}

func (o *PcloudCloudinstancesVolumesDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}][%d] pcloudCloudinstancesVolumesDeleteForbidden  %+v", 403, o.Payload)
}

func (o *PcloudCloudinstancesVolumesDeleteForbidden) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}][%d] pcloudCloudinstancesVolumesDeleteForbidden  %+v", 403, o.Payload)
}

func (o *PcloudCloudinstancesVolumesDeleteForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesVolumesDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesVolumesDeleteNotFound creates a PcloudCloudinstancesVolumesDeleteNotFound with default headers values
func NewPcloudCloudinstancesVolumesDeleteNotFound() *PcloudCloudinstancesVolumesDeleteNotFound {
	return &PcloudCloudinstancesVolumesDeleteNotFound{}
}

/*
PcloudCloudinstancesVolumesDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudCloudinstancesVolumesDeleteNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances volumes delete not found response has a 2xx status code
func (o *PcloudCloudinstancesVolumesDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances volumes delete not found response has a 3xx status code
func (o *PcloudCloudinstancesVolumesDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances volumes delete not found response has a 4xx status code
func (o *PcloudCloudinstancesVolumesDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances volumes delete not found response has a 5xx status code
func (o *PcloudCloudinstancesVolumesDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances volumes delete not found response a status code equal to that given
func (o *PcloudCloudinstancesVolumesDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud cloudinstances volumes delete not found response
func (o *PcloudCloudinstancesVolumesDeleteNotFound) Code() int {
	return 404
}

func (o *PcloudCloudinstancesVolumesDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}][%d] pcloudCloudinstancesVolumesDeleteNotFound  %+v", 404, o.Payload)
}

func (o *PcloudCloudinstancesVolumesDeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}][%d] pcloudCloudinstancesVolumesDeleteNotFound  %+v", 404, o.Payload)
}

func (o *PcloudCloudinstancesVolumesDeleteNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesVolumesDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesVolumesDeleteGone creates a PcloudCloudinstancesVolumesDeleteGone with default headers values
func NewPcloudCloudinstancesVolumesDeleteGone() *PcloudCloudinstancesVolumesDeleteGone {
	return &PcloudCloudinstancesVolumesDeleteGone{}
}

/*
PcloudCloudinstancesVolumesDeleteGone describes a response with status code 410, with default header values.

Gone
*/
type PcloudCloudinstancesVolumesDeleteGone struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances volumes delete gone response has a 2xx status code
func (o *PcloudCloudinstancesVolumesDeleteGone) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances volumes delete gone response has a 3xx status code
func (o *PcloudCloudinstancesVolumesDeleteGone) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances volumes delete gone response has a 4xx status code
func (o *PcloudCloudinstancesVolumesDeleteGone) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances volumes delete gone response has a 5xx status code
func (o *PcloudCloudinstancesVolumesDeleteGone) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances volumes delete gone response a status code equal to that given
func (o *PcloudCloudinstancesVolumesDeleteGone) IsCode(code int) bool {
	return code == 410
}

// Code gets the status code for the pcloud cloudinstances volumes delete gone response
func (o *PcloudCloudinstancesVolumesDeleteGone) Code() int {
	return 410
}

func (o *PcloudCloudinstancesVolumesDeleteGone) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}][%d] pcloudCloudinstancesVolumesDeleteGone  %+v", 410, o.Payload)
}

func (o *PcloudCloudinstancesVolumesDeleteGone) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}][%d] pcloudCloudinstancesVolumesDeleteGone  %+v", 410, o.Payload)
}

func (o *PcloudCloudinstancesVolumesDeleteGone) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesVolumesDeleteGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesVolumesDeleteInternalServerError creates a PcloudCloudinstancesVolumesDeleteInternalServerError with default headers values
func NewPcloudCloudinstancesVolumesDeleteInternalServerError() *PcloudCloudinstancesVolumesDeleteInternalServerError {
	return &PcloudCloudinstancesVolumesDeleteInternalServerError{}
}

/*
PcloudCloudinstancesVolumesDeleteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudCloudinstancesVolumesDeleteInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances volumes delete internal server error response has a 2xx status code
func (o *PcloudCloudinstancesVolumesDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances volumes delete internal server error response has a 3xx status code
func (o *PcloudCloudinstancesVolumesDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances volumes delete internal server error response has a 4xx status code
func (o *PcloudCloudinstancesVolumesDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudinstances volumes delete internal server error response has a 5xx status code
func (o *PcloudCloudinstancesVolumesDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud cloudinstances volumes delete internal server error response a status code equal to that given
func (o *PcloudCloudinstancesVolumesDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud cloudinstances volumes delete internal server error response
func (o *PcloudCloudinstancesVolumesDeleteInternalServerError) Code() int {
	return 500
}

func (o *PcloudCloudinstancesVolumesDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}][%d] pcloudCloudinstancesVolumesDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudinstancesVolumesDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}][%d] pcloudCloudinstancesVolumesDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudinstancesVolumesDeleteInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesVolumesDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
