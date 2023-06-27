// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_v_p_n_connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.ibm.com/power-iaas/service-broker/power/models"
)

// PcloudVpnconnectionsNetworksDeleteReader is a Reader for the PcloudVpnconnectionsNetworksDelete structure.
type PcloudVpnconnectionsNetworksDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudVpnconnectionsNetworksDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPcloudVpnconnectionsNetworksDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudVpnconnectionsNetworksDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudVpnconnectionsNetworksDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudVpnconnectionsNetworksDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudVpnconnectionsNetworksDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudVpnconnectionsNetworksDeleteUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudVpnconnectionsNetworksDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/networks] pcloud.vpnconnections.networks.delete", response, response.Code())
	}
}

// NewPcloudVpnconnectionsNetworksDeleteAccepted creates a PcloudVpnconnectionsNetworksDeleteAccepted with default headers values
func NewPcloudVpnconnectionsNetworksDeleteAccepted() *PcloudVpnconnectionsNetworksDeleteAccepted {
	return &PcloudVpnconnectionsNetworksDeleteAccepted{}
}

/*
PcloudVpnconnectionsNetworksDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type PcloudVpnconnectionsNetworksDeleteAccepted struct {
	Payload *models.JobReference
}

// IsSuccess returns true when this pcloud vpnconnections networks delete accepted response has a 2xx status code
func (o *PcloudVpnconnectionsNetworksDeleteAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud vpnconnections networks delete accepted response has a 3xx status code
func (o *PcloudVpnconnectionsNetworksDeleteAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud vpnconnections networks delete accepted response has a 4xx status code
func (o *PcloudVpnconnectionsNetworksDeleteAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud vpnconnections networks delete accepted response has a 5xx status code
func (o *PcloudVpnconnectionsNetworksDeleteAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud vpnconnections networks delete accepted response a status code equal to that given
func (o *PcloudVpnconnectionsNetworksDeleteAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the pcloud vpnconnections networks delete accepted response
func (o *PcloudVpnconnectionsNetworksDeleteAccepted) Code() int {
	return 202
}

func (o *PcloudVpnconnectionsNetworksDeleteAccepted) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/networks][%d] pcloudVpnconnectionsNetworksDeleteAccepted  %+v", 202, o.Payload)
}

func (o *PcloudVpnconnectionsNetworksDeleteAccepted) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/networks][%d] pcloudVpnconnectionsNetworksDeleteAccepted  %+v", 202, o.Payload)
}

func (o *PcloudVpnconnectionsNetworksDeleteAccepted) GetPayload() *models.JobReference {
	return o.Payload
}

func (o *PcloudVpnconnectionsNetworksDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobReference)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsNetworksDeleteBadRequest creates a PcloudVpnconnectionsNetworksDeleteBadRequest with default headers values
func NewPcloudVpnconnectionsNetworksDeleteBadRequest() *PcloudVpnconnectionsNetworksDeleteBadRequest {
	return &PcloudVpnconnectionsNetworksDeleteBadRequest{}
}

/*
PcloudVpnconnectionsNetworksDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudVpnconnectionsNetworksDeleteBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud vpnconnections networks delete bad request response has a 2xx status code
func (o *PcloudVpnconnectionsNetworksDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud vpnconnections networks delete bad request response has a 3xx status code
func (o *PcloudVpnconnectionsNetworksDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud vpnconnections networks delete bad request response has a 4xx status code
func (o *PcloudVpnconnectionsNetworksDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud vpnconnections networks delete bad request response has a 5xx status code
func (o *PcloudVpnconnectionsNetworksDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud vpnconnections networks delete bad request response a status code equal to that given
func (o *PcloudVpnconnectionsNetworksDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud vpnconnections networks delete bad request response
func (o *PcloudVpnconnectionsNetworksDeleteBadRequest) Code() int {
	return 400
}

func (o *PcloudVpnconnectionsNetworksDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/networks][%d] pcloudVpnconnectionsNetworksDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudVpnconnectionsNetworksDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/networks][%d] pcloudVpnconnectionsNetworksDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudVpnconnectionsNetworksDeleteBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsNetworksDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsNetworksDeleteUnauthorized creates a PcloudVpnconnectionsNetworksDeleteUnauthorized with default headers values
func NewPcloudVpnconnectionsNetworksDeleteUnauthorized() *PcloudVpnconnectionsNetworksDeleteUnauthorized {
	return &PcloudVpnconnectionsNetworksDeleteUnauthorized{}
}

/*
PcloudVpnconnectionsNetworksDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudVpnconnectionsNetworksDeleteUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud vpnconnections networks delete unauthorized response has a 2xx status code
func (o *PcloudVpnconnectionsNetworksDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud vpnconnections networks delete unauthorized response has a 3xx status code
func (o *PcloudVpnconnectionsNetworksDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud vpnconnections networks delete unauthorized response has a 4xx status code
func (o *PcloudVpnconnectionsNetworksDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud vpnconnections networks delete unauthorized response has a 5xx status code
func (o *PcloudVpnconnectionsNetworksDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud vpnconnections networks delete unauthorized response a status code equal to that given
func (o *PcloudVpnconnectionsNetworksDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud vpnconnections networks delete unauthorized response
func (o *PcloudVpnconnectionsNetworksDeleteUnauthorized) Code() int {
	return 401
}

func (o *PcloudVpnconnectionsNetworksDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/networks][%d] pcloudVpnconnectionsNetworksDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudVpnconnectionsNetworksDeleteUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/networks][%d] pcloudVpnconnectionsNetworksDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudVpnconnectionsNetworksDeleteUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsNetworksDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsNetworksDeleteForbidden creates a PcloudVpnconnectionsNetworksDeleteForbidden with default headers values
func NewPcloudVpnconnectionsNetworksDeleteForbidden() *PcloudVpnconnectionsNetworksDeleteForbidden {
	return &PcloudVpnconnectionsNetworksDeleteForbidden{}
}

/*
PcloudVpnconnectionsNetworksDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudVpnconnectionsNetworksDeleteForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud vpnconnections networks delete forbidden response has a 2xx status code
func (o *PcloudVpnconnectionsNetworksDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud vpnconnections networks delete forbidden response has a 3xx status code
func (o *PcloudVpnconnectionsNetworksDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud vpnconnections networks delete forbidden response has a 4xx status code
func (o *PcloudVpnconnectionsNetworksDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud vpnconnections networks delete forbidden response has a 5xx status code
func (o *PcloudVpnconnectionsNetworksDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud vpnconnections networks delete forbidden response a status code equal to that given
func (o *PcloudVpnconnectionsNetworksDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud vpnconnections networks delete forbidden response
func (o *PcloudVpnconnectionsNetworksDeleteForbidden) Code() int {
	return 403
}

func (o *PcloudVpnconnectionsNetworksDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/networks][%d] pcloudVpnconnectionsNetworksDeleteForbidden  %+v", 403, o.Payload)
}

func (o *PcloudVpnconnectionsNetworksDeleteForbidden) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/networks][%d] pcloudVpnconnectionsNetworksDeleteForbidden  %+v", 403, o.Payload)
}

func (o *PcloudVpnconnectionsNetworksDeleteForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsNetworksDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsNetworksDeleteNotFound creates a PcloudVpnconnectionsNetworksDeleteNotFound with default headers values
func NewPcloudVpnconnectionsNetworksDeleteNotFound() *PcloudVpnconnectionsNetworksDeleteNotFound {
	return &PcloudVpnconnectionsNetworksDeleteNotFound{}
}

/*
PcloudVpnconnectionsNetworksDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudVpnconnectionsNetworksDeleteNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud vpnconnections networks delete not found response has a 2xx status code
func (o *PcloudVpnconnectionsNetworksDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud vpnconnections networks delete not found response has a 3xx status code
func (o *PcloudVpnconnectionsNetworksDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud vpnconnections networks delete not found response has a 4xx status code
func (o *PcloudVpnconnectionsNetworksDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud vpnconnections networks delete not found response has a 5xx status code
func (o *PcloudVpnconnectionsNetworksDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud vpnconnections networks delete not found response a status code equal to that given
func (o *PcloudVpnconnectionsNetworksDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud vpnconnections networks delete not found response
func (o *PcloudVpnconnectionsNetworksDeleteNotFound) Code() int {
	return 404
}

func (o *PcloudVpnconnectionsNetworksDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/networks][%d] pcloudVpnconnectionsNetworksDeleteNotFound  %+v", 404, o.Payload)
}

func (o *PcloudVpnconnectionsNetworksDeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/networks][%d] pcloudVpnconnectionsNetworksDeleteNotFound  %+v", 404, o.Payload)
}

func (o *PcloudVpnconnectionsNetworksDeleteNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsNetworksDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsNetworksDeleteUnprocessableEntity creates a PcloudVpnconnectionsNetworksDeleteUnprocessableEntity with default headers values
func NewPcloudVpnconnectionsNetworksDeleteUnprocessableEntity() *PcloudVpnconnectionsNetworksDeleteUnprocessableEntity {
	return &PcloudVpnconnectionsNetworksDeleteUnprocessableEntity{}
}

/*
PcloudVpnconnectionsNetworksDeleteUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudVpnconnectionsNetworksDeleteUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud vpnconnections networks delete unprocessable entity response has a 2xx status code
func (o *PcloudVpnconnectionsNetworksDeleteUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud vpnconnections networks delete unprocessable entity response has a 3xx status code
func (o *PcloudVpnconnectionsNetworksDeleteUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud vpnconnections networks delete unprocessable entity response has a 4xx status code
func (o *PcloudVpnconnectionsNetworksDeleteUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud vpnconnections networks delete unprocessable entity response has a 5xx status code
func (o *PcloudVpnconnectionsNetworksDeleteUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud vpnconnections networks delete unprocessable entity response a status code equal to that given
func (o *PcloudVpnconnectionsNetworksDeleteUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the pcloud vpnconnections networks delete unprocessable entity response
func (o *PcloudVpnconnectionsNetworksDeleteUnprocessableEntity) Code() int {
	return 422
}

func (o *PcloudVpnconnectionsNetworksDeleteUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/networks][%d] pcloudVpnconnectionsNetworksDeleteUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudVpnconnectionsNetworksDeleteUnprocessableEntity) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/networks][%d] pcloudVpnconnectionsNetworksDeleteUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudVpnconnectionsNetworksDeleteUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsNetworksDeleteUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsNetworksDeleteInternalServerError creates a PcloudVpnconnectionsNetworksDeleteInternalServerError with default headers values
func NewPcloudVpnconnectionsNetworksDeleteInternalServerError() *PcloudVpnconnectionsNetworksDeleteInternalServerError {
	return &PcloudVpnconnectionsNetworksDeleteInternalServerError{}
}

/*
PcloudVpnconnectionsNetworksDeleteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudVpnconnectionsNetworksDeleteInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud vpnconnections networks delete internal server error response has a 2xx status code
func (o *PcloudVpnconnectionsNetworksDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud vpnconnections networks delete internal server error response has a 3xx status code
func (o *PcloudVpnconnectionsNetworksDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud vpnconnections networks delete internal server error response has a 4xx status code
func (o *PcloudVpnconnectionsNetworksDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud vpnconnections networks delete internal server error response has a 5xx status code
func (o *PcloudVpnconnectionsNetworksDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud vpnconnections networks delete internal server error response a status code equal to that given
func (o *PcloudVpnconnectionsNetworksDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud vpnconnections networks delete internal server error response
func (o *PcloudVpnconnectionsNetworksDeleteInternalServerError) Code() int {
	return 500
}

func (o *PcloudVpnconnectionsNetworksDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/networks][%d] pcloudVpnconnectionsNetworksDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudVpnconnectionsNetworksDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/networks][%d] pcloudVpnconnectionsNetworksDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudVpnconnectionsNetworksDeleteInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsNetworksDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
