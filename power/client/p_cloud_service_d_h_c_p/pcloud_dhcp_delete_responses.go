// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_service_d_h_c_p

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudDhcpDeleteReader is a Reader for the PcloudDhcpDelete structure.
type PcloudDhcpDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudDhcpDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPcloudDhcpDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudDhcpDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudDhcpDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudDhcpDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudDhcpDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/services/dhcp/{dhcp_id}] pcloud.dhcp.delete", response, response.Code())
	}
}

// NewPcloudDhcpDeleteAccepted creates a PcloudDhcpDeleteAccepted with default headers values
func NewPcloudDhcpDeleteAccepted() *PcloudDhcpDeleteAccepted {
	return &PcloudDhcpDeleteAccepted{}
}

/*
PcloudDhcpDeleteAccepted describes a response with status code 202, with default header values.

OK
*/
type PcloudDhcpDeleteAccepted struct {
	Payload models.Object
}

// IsSuccess returns true when this pcloud dhcp delete accepted response has a 2xx status code
func (o *PcloudDhcpDeleteAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud dhcp delete accepted response has a 3xx status code
func (o *PcloudDhcpDeleteAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud dhcp delete accepted response has a 4xx status code
func (o *PcloudDhcpDeleteAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud dhcp delete accepted response has a 5xx status code
func (o *PcloudDhcpDeleteAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud dhcp delete accepted response a status code equal to that given
func (o *PcloudDhcpDeleteAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the pcloud dhcp delete accepted response
func (o *PcloudDhcpDeleteAccepted) Code() int {
	return 202
}

func (o *PcloudDhcpDeleteAccepted) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/services/dhcp/{dhcp_id}][%d] pcloudDhcpDeleteAccepted  %+v", 202, o.Payload)
}

func (o *PcloudDhcpDeleteAccepted) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/services/dhcp/{dhcp_id}][%d] pcloudDhcpDeleteAccepted  %+v", 202, o.Payload)
}

func (o *PcloudDhcpDeleteAccepted) GetPayload() models.Object {
	return o.Payload
}

func (o *PcloudDhcpDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudDhcpDeleteBadRequest creates a PcloudDhcpDeleteBadRequest with default headers values
func NewPcloudDhcpDeleteBadRequest() *PcloudDhcpDeleteBadRequest {
	return &PcloudDhcpDeleteBadRequest{}
}

/*
PcloudDhcpDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudDhcpDeleteBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud dhcp delete bad request response has a 2xx status code
func (o *PcloudDhcpDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud dhcp delete bad request response has a 3xx status code
func (o *PcloudDhcpDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud dhcp delete bad request response has a 4xx status code
func (o *PcloudDhcpDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud dhcp delete bad request response has a 5xx status code
func (o *PcloudDhcpDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud dhcp delete bad request response a status code equal to that given
func (o *PcloudDhcpDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud dhcp delete bad request response
func (o *PcloudDhcpDeleteBadRequest) Code() int {
	return 400
}

func (o *PcloudDhcpDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/services/dhcp/{dhcp_id}][%d] pcloudDhcpDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudDhcpDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/services/dhcp/{dhcp_id}][%d] pcloudDhcpDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudDhcpDeleteBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudDhcpDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudDhcpDeleteForbidden creates a PcloudDhcpDeleteForbidden with default headers values
func NewPcloudDhcpDeleteForbidden() *PcloudDhcpDeleteForbidden {
	return &PcloudDhcpDeleteForbidden{}
}

/*
PcloudDhcpDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudDhcpDeleteForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud dhcp delete forbidden response has a 2xx status code
func (o *PcloudDhcpDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud dhcp delete forbidden response has a 3xx status code
func (o *PcloudDhcpDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud dhcp delete forbidden response has a 4xx status code
func (o *PcloudDhcpDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud dhcp delete forbidden response has a 5xx status code
func (o *PcloudDhcpDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud dhcp delete forbidden response a status code equal to that given
func (o *PcloudDhcpDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud dhcp delete forbidden response
func (o *PcloudDhcpDeleteForbidden) Code() int {
	return 403
}

func (o *PcloudDhcpDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/services/dhcp/{dhcp_id}][%d] pcloudDhcpDeleteForbidden  %+v", 403, o.Payload)
}

func (o *PcloudDhcpDeleteForbidden) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/services/dhcp/{dhcp_id}][%d] pcloudDhcpDeleteForbidden  %+v", 403, o.Payload)
}

func (o *PcloudDhcpDeleteForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudDhcpDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudDhcpDeleteNotFound creates a PcloudDhcpDeleteNotFound with default headers values
func NewPcloudDhcpDeleteNotFound() *PcloudDhcpDeleteNotFound {
	return &PcloudDhcpDeleteNotFound{}
}

/*
PcloudDhcpDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudDhcpDeleteNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud dhcp delete not found response has a 2xx status code
func (o *PcloudDhcpDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud dhcp delete not found response has a 3xx status code
func (o *PcloudDhcpDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud dhcp delete not found response has a 4xx status code
func (o *PcloudDhcpDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud dhcp delete not found response has a 5xx status code
func (o *PcloudDhcpDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud dhcp delete not found response a status code equal to that given
func (o *PcloudDhcpDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud dhcp delete not found response
func (o *PcloudDhcpDeleteNotFound) Code() int {
	return 404
}

func (o *PcloudDhcpDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/services/dhcp/{dhcp_id}][%d] pcloudDhcpDeleteNotFound  %+v", 404, o.Payload)
}

func (o *PcloudDhcpDeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/services/dhcp/{dhcp_id}][%d] pcloudDhcpDeleteNotFound  %+v", 404, o.Payload)
}

func (o *PcloudDhcpDeleteNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudDhcpDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudDhcpDeleteInternalServerError creates a PcloudDhcpDeleteInternalServerError with default headers values
func NewPcloudDhcpDeleteInternalServerError() *PcloudDhcpDeleteInternalServerError {
	return &PcloudDhcpDeleteInternalServerError{}
}

/*
PcloudDhcpDeleteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudDhcpDeleteInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud dhcp delete internal server error response has a 2xx status code
func (o *PcloudDhcpDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud dhcp delete internal server error response has a 3xx status code
func (o *PcloudDhcpDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud dhcp delete internal server error response has a 4xx status code
func (o *PcloudDhcpDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud dhcp delete internal server error response has a 5xx status code
func (o *PcloudDhcpDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud dhcp delete internal server error response a status code equal to that given
func (o *PcloudDhcpDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud dhcp delete internal server error response
func (o *PcloudDhcpDeleteInternalServerError) Code() int {
	return 500
}

func (o *PcloudDhcpDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/services/dhcp/{dhcp_id}][%d] pcloudDhcpDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudDhcpDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/services/dhcp/{dhcp_id}][%d] pcloudDhcpDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudDhcpDeleteInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudDhcpDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
