// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_p_vm_instances

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

// PcloudV2PvminstancesGetallReader is a Reader for the PcloudV2PvminstancesGetall structure.
type PcloudV2PvminstancesGetallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudV2PvminstancesGetallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudV2PvminstancesGetallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudV2PvminstancesGetallBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudV2PvminstancesGetallUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudV2PvminstancesGetallForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudV2PvminstancesGetallNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPcloudV2PvminstancesGetallRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudV2PvminstancesGetallInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/pvm-instances] pcloud.v2.pvminstances.getall", response, response.Code())
	}
}

// NewPcloudV2PvminstancesGetallOK creates a PcloudV2PvminstancesGetallOK with default headers values
func NewPcloudV2PvminstancesGetallOK() *PcloudV2PvminstancesGetallOK {
	return &PcloudV2PvminstancesGetallOK{}
}

/*
PcloudV2PvminstancesGetallOK describes a response with status code 200, with default header values.

OK
*/
type PcloudV2PvminstancesGetallOK struct {
	Payload *models.PVMInstancesV2
}

// IsSuccess returns true when this pcloud v2 pvminstances getall o k response has a 2xx status code
func (o *PcloudV2PvminstancesGetallOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud v2 pvminstances getall o k response has a 3xx status code
func (o *PcloudV2PvminstancesGetallOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 pvminstances getall o k response has a 4xx status code
func (o *PcloudV2PvminstancesGetallOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud v2 pvminstances getall o k response has a 5xx status code
func (o *PcloudV2PvminstancesGetallOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v2 pvminstances getall o k response a status code equal to that given
func (o *PcloudV2PvminstancesGetallOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud v2 pvminstances getall o k response
func (o *PcloudV2PvminstancesGetallOK) Code() int {
	return 200
}

func (o *PcloudV2PvminstancesGetallOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudV2PvminstancesGetallOK %s", 200, payload)
}

func (o *PcloudV2PvminstancesGetallOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudV2PvminstancesGetallOK %s", 200, payload)
}

func (o *PcloudV2PvminstancesGetallOK) GetPayload() *models.PVMInstancesV2 {
	return o.Payload
}

func (o *PcloudV2PvminstancesGetallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PVMInstancesV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2PvminstancesGetallBadRequest creates a PcloudV2PvminstancesGetallBadRequest with default headers values
func NewPcloudV2PvminstancesGetallBadRequest() *PcloudV2PvminstancesGetallBadRequest {
	return &PcloudV2PvminstancesGetallBadRequest{}
}

/*
PcloudV2PvminstancesGetallBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudV2PvminstancesGetallBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v2 pvminstances getall bad request response has a 2xx status code
func (o *PcloudV2PvminstancesGetallBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v2 pvminstances getall bad request response has a 3xx status code
func (o *PcloudV2PvminstancesGetallBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 pvminstances getall bad request response has a 4xx status code
func (o *PcloudV2PvminstancesGetallBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud v2 pvminstances getall bad request response has a 5xx status code
func (o *PcloudV2PvminstancesGetallBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v2 pvminstances getall bad request response a status code equal to that given
func (o *PcloudV2PvminstancesGetallBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud v2 pvminstances getall bad request response
func (o *PcloudV2PvminstancesGetallBadRequest) Code() int {
	return 400
}

func (o *PcloudV2PvminstancesGetallBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudV2PvminstancesGetallBadRequest %s", 400, payload)
}

func (o *PcloudV2PvminstancesGetallBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudV2PvminstancesGetallBadRequest %s", 400, payload)
}

func (o *PcloudV2PvminstancesGetallBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2PvminstancesGetallBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2PvminstancesGetallUnauthorized creates a PcloudV2PvminstancesGetallUnauthorized with default headers values
func NewPcloudV2PvminstancesGetallUnauthorized() *PcloudV2PvminstancesGetallUnauthorized {
	return &PcloudV2PvminstancesGetallUnauthorized{}
}

/*
PcloudV2PvminstancesGetallUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudV2PvminstancesGetallUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v2 pvminstances getall unauthorized response has a 2xx status code
func (o *PcloudV2PvminstancesGetallUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v2 pvminstances getall unauthorized response has a 3xx status code
func (o *PcloudV2PvminstancesGetallUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 pvminstances getall unauthorized response has a 4xx status code
func (o *PcloudV2PvminstancesGetallUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud v2 pvminstances getall unauthorized response has a 5xx status code
func (o *PcloudV2PvminstancesGetallUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v2 pvminstances getall unauthorized response a status code equal to that given
func (o *PcloudV2PvminstancesGetallUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud v2 pvminstances getall unauthorized response
func (o *PcloudV2PvminstancesGetallUnauthorized) Code() int {
	return 401
}

func (o *PcloudV2PvminstancesGetallUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudV2PvminstancesGetallUnauthorized %s", 401, payload)
}

func (o *PcloudV2PvminstancesGetallUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudV2PvminstancesGetallUnauthorized %s", 401, payload)
}

func (o *PcloudV2PvminstancesGetallUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2PvminstancesGetallUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2PvminstancesGetallForbidden creates a PcloudV2PvminstancesGetallForbidden with default headers values
func NewPcloudV2PvminstancesGetallForbidden() *PcloudV2PvminstancesGetallForbidden {
	return &PcloudV2PvminstancesGetallForbidden{}
}

/*
PcloudV2PvminstancesGetallForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudV2PvminstancesGetallForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v2 pvminstances getall forbidden response has a 2xx status code
func (o *PcloudV2PvminstancesGetallForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v2 pvminstances getall forbidden response has a 3xx status code
func (o *PcloudV2PvminstancesGetallForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 pvminstances getall forbidden response has a 4xx status code
func (o *PcloudV2PvminstancesGetallForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud v2 pvminstances getall forbidden response has a 5xx status code
func (o *PcloudV2PvminstancesGetallForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v2 pvminstances getall forbidden response a status code equal to that given
func (o *PcloudV2PvminstancesGetallForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud v2 pvminstances getall forbidden response
func (o *PcloudV2PvminstancesGetallForbidden) Code() int {
	return 403
}

func (o *PcloudV2PvminstancesGetallForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudV2PvminstancesGetallForbidden %s", 403, payload)
}

func (o *PcloudV2PvminstancesGetallForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudV2PvminstancesGetallForbidden %s", 403, payload)
}

func (o *PcloudV2PvminstancesGetallForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2PvminstancesGetallForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2PvminstancesGetallNotFound creates a PcloudV2PvminstancesGetallNotFound with default headers values
func NewPcloudV2PvminstancesGetallNotFound() *PcloudV2PvminstancesGetallNotFound {
	return &PcloudV2PvminstancesGetallNotFound{}
}

/*
PcloudV2PvminstancesGetallNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudV2PvminstancesGetallNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v2 pvminstances getall not found response has a 2xx status code
func (o *PcloudV2PvminstancesGetallNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v2 pvminstances getall not found response has a 3xx status code
func (o *PcloudV2PvminstancesGetallNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 pvminstances getall not found response has a 4xx status code
func (o *PcloudV2PvminstancesGetallNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud v2 pvminstances getall not found response has a 5xx status code
func (o *PcloudV2PvminstancesGetallNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v2 pvminstances getall not found response a status code equal to that given
func (o *PcloudV2PvminstancesGetallNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud v2 pvminstances getall not found response
func (o *PcloudV2PvminstancesGetallNotFound) Code() int {
	return 404
}

func (o *PcloudV2PvminstancesGetallNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudV2PvminstancesGetallNotFound %s", 404, payload)
}

func (o *PcloudV2PvminstancesGetallNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudV2PvminstancesGetallNotFound %s", 404, payload)
}

func (o *PcloudV2PvminstancesGetallNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2PvminstancesGetallNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2PvminstancesGetallRequestTimeout creates a PcloudV2PvminstancesGetallRequestTimeout with default headers values
func NewPcloudV2PvminstancesGetallRequestTimeout() *PcloudV2PvminstancesGetallRequestTimeout {
	return &PcloudV2PvminstancesGetallRequestTimeout{}
}

/*
PcloudV2PvminstancesGetallRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type PcloudV2PvminstancesGetallRequestTimeout struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v2 pvminstances getall request timeout response has a 2xx status code
func (o *PcloudV2PvminstancesGetallRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v2 pvminstances getall request timeout response has a 3xx status code
func (o *PcloudV2PvminstancesGetallRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 pvminstances getall request timeout response has a 4xx status code
func (o *PcloudV2PvminstancesGetallRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud v2 pvminstances getall request timeout response has a 5xx status code
func (o *PcloudV2PvminstancesGetallRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v2 pvminstances getall request timeout response a status code equal to that given
func (o *PcloudV2PvminstancesGetallRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the pcloud v2 pvminstances getall request timeout response
func (o *PcloudV2PvminstancesGetallRequestTimeout) Code() int {
	return 408
}

func (o *PcloudV2PvminstancesGetallRequestTimeout) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudV2PvminstancesGetallRequestTimeout %s", 408, payload)
}

func (o *PcloudV2PvminstancesGetallRequestTimeout) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudV2PvminstancesGetallRequestTimeout %s", 408, payload)
}

func (o *PcloudV2PvminstancesGetallRequestTimeout) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2PvminstancesGetallRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2PvminstancesGetallInternalServerError creates a PcloudV2PvminstancesGetallInternalServerError with default headers values
func NewPcloudV2PvminstancesGetallInternalServerError() *PcloudV2PvminstancesGetallInternalServerError {
	return &PcloudV2PvminstancesGetallInternalServerError{}
}

/*
PcloudV2PvminstancesGetallInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudV2PvminstancesGetallInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v2 pvminstances getall internal server error response has a 2xx status code
func (o *PcloudV2PvminstancesGetallInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v2 pvminstances getall internal server error response has a 3xx status code
func (o *PcloudV2PvminstancesGetallInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 pvminstances getall internal server error response has a 4xx status code
func (o *PcloudV2PvminstancesGetallInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud v2 pvminstances getall internal server error response has a 5xx status code
func (o *PcloudV2PvminstancesGetallInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud v2 pvminstances getall internal server error response a status code equal to that given
func (o *PcloudV2PvminstancesGetallInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud v2 pvminstances getall internal server error response
func (o *PcloudV2PvminstancesGetallInternalServerError) Code() int {
	return 500
}

func (o *PcloudV2PvminstancesGetallInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudV2PvminstancesGetallInternalServerError %s", 500, payload)
}

func (o *PcloudV2PvminstancesGetallInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudV2PvminstancesGetallInternalServerError %s", 500, payload)
}

func (o *PcloudV2PvminstancesGetallInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2PvminstancesGetallInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
