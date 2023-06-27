// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_shared_processor_pools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudSharedprocessorpoolsGetallReader is a Reader for the PcloudSharedprocessorpoolsGetall structure.
type PcloudSharedprocessorpoolsGetallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudSharedprocessorpoolsGetallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudSharedprocessorpoolsGetallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudSharedprocessorpoolsGetallBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudSharedprocessorpoolsGetallUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudSharedprocessorpoolsGetallForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudSharedprocessorpoolsGetallNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudSharedprocessorpoolsGetallInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/shared-processor-pools] pcloud.sharedprocessorpools.getall", response, response.Code())
	}
}

// NewPcloudSharedprocessorpoolsGetallOK creates a PcloudSharedprocessorpoolsGetallOK with default headers values
func NewPcloudSharedprocessorpoolsGetallOK() *PcloudSharedprocessorpoolsGetallOK {
	return &PcloudSharedprocessorpoolsGetallOK{}
}

/*
PcloudSharedprocessorpoolsGetallOK describes a response with status code 200, with default header values.

OK
*/
type PcloudSharedprocessorpoolsGetallOK struct {
	Payload *models.SharedProcessorPools
}

// IsSuccess returns true when this pcloud sharedprocessorpools getall o k response has a 2xx status code
func (o *PcloudSharedprocessorpoolsGetallOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud sharedprocessorpools getall o k response has a 3xx status code
func (o *PcloudSharedprocessorpoolsGetallOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sharedprocessorpools getall o k response has a 4xx status code
func (o *PcloudSharedprocessorpoolsGetallOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud sharedprocessorpools getall o k response has a 5xx status code
func (o *PcloudSharedprocessorpoolsGetallOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sharedprocessorpools getall o k response a status code equal to that given
func (o *PcloudSharedprocessorpoolsGetallOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud sharedprocessorpools getall o k response
func (o *PcloudSharedprocessorpoolsGetallOK) Code() int {
	return 200
}

func (o *PcloudSharedprocessorpoolsGetallOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/shared-processor-pools][%d] pcloudSharedprocessorpoolsGetallOK  %+v", 200, o.Payload)
}

func (o *PcloudSharedprocessorpoolsGetallOK) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/shared-processor-pools][%d] pcloudSharedprocessorpoolsGetallOK  %+v", 200, o.Payload)
}

func (o *PcloudSharedprocessorpoolsGetallOK) GetPayload() *models.SharedProcessorPools {
	return o.Payload
}

func (o *PcloudSharedprocessorpoolsGetallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SharedProcessorPools)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSharedprocessorpoolsGetallBadRequest creates a PcloudSharedprocessorpoolsGetallBadRequest with default headers values
func NewPcloudSharedprocessorpoolsGetallBadRequest() *PcloudSharedprocessorpoolsGetallBadRequest {
	return &PcloudSharedprocessorpoolsGetallBadRequest{}
}

/*
PcloudSharedprocessorpoolsGetallBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudSharedprocessorpoolsGetallBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud sharedprocessorpools getall bad request response has a 2xx status code
func (o *PcloudSharedprocessorpoolsGetallBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud sharedprocessorpools getall bad request response has a 3xx status code
func (o *PcloudSharedprocessorpoolsGetallBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sharedprocessorpools getall bad request response has a 4xx status code
func (o *PcloudSharedprocessorpoolsGetallBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud sharedprocessorpools getall bad request response has a 5xx status code
func (o *PcloudSharedprocessorpoolsGetallBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sharedprocessorpools getall bad request response a status code equal to that given
func (o *PcloudSharedprocessorpoolsGetallBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud sharedprocessorpools getall bad request response
func (o *PcloudSharedprocessorpoolsGetallBadRequest) Code() int {
	return 400
}

func (o *PcloudSharedprocessorpoolsGetallBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/shared-processor-pools][%d] pcloudSharedprocessorpoolsGetallBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudSharedprocessorpoolsGetallBadRequest) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/shared-processor-pools][%d] pcloudSharedprocessorpoolsGetallBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudSharedprocessorpoolsGetallBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSharedprocessorpoolsGetallBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSharedprocessorpoolsGetallUnauthorized creates a PcloudSharedprocessorpoolsGetallUnauthorized with default headers values
func NewPcloudSharedprocessorpoolsGetallUnauthorized() *PcloudSharedprocessorpoolsGetallUnauthorized {
	return &PcloudSharedprocessorpoolsGetallUnauthorized{}
}

/*
PcloudSharedprocessorpoolsGetallUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudSharedprocessorpoolsGetallUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud sharedprocessorpools getall unauthorized response has a 2xx status code
func (o *PcloudSharedprocessorpoolsGetallUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud sharedprocessorpools getall unauthorized response has a 3xx status code
func (o *PcloudSharedprocessorpoolsGetallUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sharedprocessorpools getall unauthorized response has a 4xx status code
func (o *PcloudSharedprocessorpoolsGetallUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud sharedprocessorpools getall unauthorized response has a 5xx status code
func (o *PcloudSharedprocessorpoolsGetallUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sharedprocessorpools getall unauthorized response a status code equal to that given
func (o *PcloudSharedprocessorpoolsGetallUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud sharedprocessorpools getall unauthorized response
func (o *PcloudSharedprocessorpoolsGetallUnauthorized) Code() int {
	return 401
}

func (o *PcloudSharedprocessorpoolsGetallUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/shared-processor-pools][%d] pcloudSharedprocessorpoolsGetallUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudSharedprocessorpoolsGetallUnauthorized) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/shared-processor-pools][%d] pcloudSharedprocessorpoolsGetallUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudSharedprocessorpoolsGetallUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSharedprocessorpoolsGetallUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSharedprocessorpoolsGetallForbidden creates a PcloudSharedprocessorpoolsGetallForbidden with default headers values
func NewPcloudSharedprocessorpoolsGetallForbidden() *PcloudSharedprocessorpoolsGetallForbidden {
	return &PcloudSharedprocessorpoolsGetallForbidden{}
}

/*
PcloudSharedprocessorpoolsGetallForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudSharedprocessorpoolsGetallForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud sharedprocessorpools getall forbidden response has a 2xx status code
func (o *PcloudSharedprocessorpoolsGetallForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud sharedprocessorpools getall forbidden response has a 3xx status code
func (o *PcloudSharedprocessorpoolsGetallForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sharedprocessorpools getall forbidden response has a 4xx status code
func (o *PcloudSharedprocessorpoolsGetallForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud sharedprocessorpools getall forbidden response has a 5xx status code
func (o *PcloudSharedprocessorpoolsGetallForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sharedprocessorpools getall forbidden response a status code equal to that given
func (o *PcloudSharedprocessorpoolsGetallForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud sharedprocessorpools getall forbidden response
func (o *PcloudSharedprocessorpoolsGetallForbidden) Code() int {
	return 403
}

func (o *PcloudSharedprocessorpoolsGetallForbidden) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/shared-processor-pools][%d] pcloudSharedprocessorpoolsGetallForbidden  %+v", 403, o.Payload)
}

func (o *PcloudSharedprocessorpoolsGetallForbidden) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/shared-processor-pools][%d] pcloudSharedprocessorpoolsGetallForbidden  %+v", 403, o.Payload)
}

func (o *PcloudSharedprocessorpoolsGetallForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSharedprocessorpoolsGetallForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSharedprocessorpoolsGetallNotFound creates a PcloudSharedprocessorpoolsGetallNotFound with default headers values
func NewPcloudSharedprocessorpoolsGetallNotFound() *PcloudSharedprocessorpoolsGetallNotFound {
	return &PcloudSharedprocessorpoolsGetallNotFound{}
}

/*
PcloudSharedprocessorpoolsGetallNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudSharedprocessorpoolsGetallNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud sharedprocessorpools getall not found response has a 2xx status code
func (o *PcloudSharedprocessorpoolsGetallNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud sharedprocessorpools getall not found response has a 3xx status code
func (o *PcloudSharedprocessorpoolsGetallNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sharedprocessorpools getall not found response has a 4xx status code
func (o *PcloudSharedprocessorpoolsGetallNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud sharedprocessorpools getall not found response has a 5xx status code
func (o *PcloudSharedprocessorpoolsGetallNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sharedprocessorpools getall not found response a status code equal to that given
func (o *PcloudSharedprocessorpoolsGetallNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud sharedprocessorpools getall not found response
func (o *PcloudSharedprocessorpoolsGetallNotFound) Code() int {
	return 404
}

func (o *PcloudSharedprocessorpoolsGetallNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/shared-processor-pools][%d] pcloudSharedprocessorpoolsGetallNotFound  %+v", 404, o.Payload)
}

func (o *PcloudSharedprocessorpoolsGetallNotFound) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/shared-processor-pools][%d] pcloudSharedprocessorpoolsGetallNotFound  %+v", 404, o.Payload)
}

func (o *PcloudSharedprocessorpoolsGetallNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSharedprocessorpoolsGetallNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSharedprocessorpoolsGetallInternalServerError creates a PcloudSharedprocessorpoolsGetallInternalServerError with default headers values
func NewPcloudSharedprocessorpoolsGetallInternalServerError() *PcloudSharedprocessorpoolsGetallInternalServerError {
	return &PcloudSharedprocessorpoolsGetallInternalServerError{}
}

/*
PcloudSharedprocessorpoolsGetallInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudSharedprocessorpoolsGetallInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud sharedprocessorpools getall internal server error response has a 2xx status code
func (o *PcloudSharedprocessorpoolsGetallInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud sharedprocessorpools getall internal server error response has a 3xx status code
func (o *PcloudSharedprocessorpoolsGetallInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sharedprocessorpools getall internal server error response has a 4xx status code
func (o *PcloudSharedprocessorpoolsGetallInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud sharedprocessorpools getall internal server error response has a 5xx status code
func (o *PcloudSharedprocessorpoolsGetallInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud sharedprocessorpools getall internal server error response a status code equal to that given
func (o *PcloudSharedprocessorpoolsGetallInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud sharedprocessorpools getall internal server error response
func (o *PcloudSharedprocessorpoolsGetallInternalServerError) Code() int {
	return 500
}

func (o *PcloudSharedprocessorpoolsGetallInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/shared-processor-pools][%d] pcloudSharedprocessorpoolsGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudSharedprocessorpoolsGetallInternalServerError) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/shared-processor-pools][%d] pcloudSharedprocessorpoolsGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudSharedprocessorpoolsGetallInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSharedprocessorpoolsGetallInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
