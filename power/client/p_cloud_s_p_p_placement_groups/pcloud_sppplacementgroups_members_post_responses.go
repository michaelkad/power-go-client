// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_s_p_p_placement_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.ibm.com/power-iaas/service-broker/power/models"
)

// PcloudSppplacementgroupsMembersPostReader is a Reader for the PcloudSppplacementgroupsMembersPost structure.
type PcloudSppplacementgroupsMembersPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudSppplacementgroupsMembersPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudSppplacementgroupsMembersPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudSppplacementgroupsMembersPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudSppplacementgroupsMembersPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudSppplacementgroupsMembersPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudSppplacementgroupsMembersPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPcloudSppplacementgroupsMembersPostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudSppplacementgroupsMembersPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudSppplacementgroupsMembersPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups/{spp_placement_group_id}/members/{shared_processor_pool_id}] pcloud.sppplacementgroups.members.post", response, response.Code())
	}
}

// NewPcloudSppplacementgroupsMembersPostOK creates a PcloudSppplacementgroupsMembersPostOK with default headers values
func NewPcloudSppplacementgroupsMembersPostOK() *PcloudSppplacementgroupsMembersPostOK {
	return &PcloudSppplacementgroupsMembersPostOK{}
}

/*
PcloudSppplacementgroupsMembersPostOK describes a response with status code 200, with default header values.

OK
*/
type PcloudSppplacementgroupsMembersPostOK struct {
	Payload *models.SPPPlacementGroup
}

// IsSuccess returns true when this pcloud sppplacementgroups members post o k response has a 2xx status code
func (o *PcloudSppplacementgroupsMembersPostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud sppplacementgroups members post o k response has a 3xx status code
func (o *PcloudSppplacementgroupsMembersPostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sppplacementgroups members post o k response has a 4xx status code
func (o *PcloudSppplacementgroupsMembersPostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud sppplacementgroups members post o k response has a 5xx status code
func (o *PcloudSppplacementgroupsMembersPostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sppplacementgroups members post o k response a status code equal to that given
func (o *PcloudSppplacementgroupsMembersPostOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud sppplacementgroups members post o k response
func (o *PcloudSppplacementgroupsMembersPostOK) Code() int {
	return 200
}

func (o *PcloudSppplacementgroupsMembersPostOK) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups/{spp_placement_group_id}/members/{shared_processor_pool_id}][%d] pcloudSppplacementgroupsMembersPostOK  %+v", 200, o.Payload)
}

func (o *PcloudSppplacementgroupsMembersPostOK) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups/{spp_placement_group_id}/members/{shared_processor_pool_id}][%d] pcloudSppplacementgroupsMembersPostOK  %+v", 200, o.Payload)
}

func (o *PcloudSppplacementgroupsMembersPostOK) GetPayload() *models.SPPPlacementGroup {
	return o.Payload
}

func (o *PcloudSppplacementgroupsMembersPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SPPPlacementGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSppplacementgroupsMembersPostBadRequest creates a PcloudSppplacementgroupsMembersPostBadRequest with default headers values
func NewPcloudSppplacementgroupsMembersPostBadRequest() *PcloudSppplacementgroupsMembersPostBadRequest {
	return &PcloudSppplacementgroupsMembersPostBadRequest{}
}

/*
PcloudSppplacementgroupsMembersPostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudSppplacementgroupsMembersPostBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud sppplacementgroups members post bad request response has a 2xx status code
func (o *PcloudSppplacementgroupsMembersPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud sppplacementgroups members post bad request response has a 3xx status code
func (o *PcloudSppplacementgroupsMembersPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sppplacementgroups members post bad request response has a 4xx status code
func (o *PcloudSppplacementgroupsMembersPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud sppplacementgroups members post bad request response has a 5xx status code
func (o *PcloudSppplacementgroupsMembersPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sppplacementgroups members post bad request response a status code equal to that given
func (o *PcloudSppplacementgroupsMembersPostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud sppplacementgroups members post bad request response
func (o *PcloudSppplacementgroupsMembersPostBadRequest) Code() int {
	return 400
}

func (o *PcloudSppplacementgroupsMembersPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups/{spp_placement_group_id}/members/{shared_processor_pool_id}][%d] pcloudSppplacementgroupsMembersPostBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudSppplacementgroupsMembersPostBadRequest) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups/{spp_placement_group_id}/members/{shared_processor_pool_id}][%d] pcloudSppplacementgroupsMembersPostBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudSppplacementgroupsMembersPostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSppplacementgroupsMembersPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSppplacementgroupsMembersPostUnauthorized creates a PcloudSppplacementgroupsMembersPostUnauthorized with default headers values
func NewPcloudSppplacementgroupsMembersPostUnauthorized() *PcloudSppplacementgroupsMembersPostUnauthorized {
	return &PcloudSppplacementgroupsMembersPostUnauthorized{}
}

/*
PcloudSppplacementgroupsMembersPostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudSppplacementgroupsMembersPostUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud sppplacementgroups members post unauthorized response has a 2xx status code
func (o *PcloudSppplacementgroupsMembersPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud sppplacementgroups members post unauthorized response has a 3xx status code
func (o *PcloudSppplacementgroupsMembersPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sppplacementgroups members post unauthorized response has a 4xx status code
func (o *PcloudSppplacementgroupsMembersPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud sppplacementgroups members post unauthorized response has a 5xx status code
func (o *PcloudSppplacementgroupsMembersPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sppplacementgroups members post unauthorized response a status code equal to that given
func (o *PcloudSppplacementgroupsMembersPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud sppplacementgroups members post unauthorized response
func (o *PcloudSppplacementgroupsMembersPostUnauthorized) Code() int {
	return 401
}

func (o *PcloudSppplacementgroupsMembersPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups/{spp_placement_group_id}/members/{shared_processor_pool_id}][%d] pcloudSppplacementgroupsMembersPostUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudSppplacementgroupsMembersPostUnauthorized) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups/{spp_placement_group_id}/members/{shared_processor_pool_id}][%d] pcloudSppplacementgroupsMembersPostUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudSppplacementgroupsMembersPostUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSppplacementgroupsMembersPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSppplacementgroupsMembersPostForbidden creates a PcloudSppplacementgroupsMembersPostForbidden with default headers values
func NewPcloudSppplacementgroupsMembersPostForbidden() *PcloudSppplacementgroupsMembersPostForbidden {
	return &PcloudSppplacementgroupsMembersPostForbidden{}
}

/*
PcloudSppplacementgroupsMembersPostForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudSppplacementgroupsMembersPostForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud sppplacementgroups members post forbidden response has a 2xx status code
func (o *PcloudSppplacementgroupsMembersPostForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud sppplacementgroups members post forbidden response has a 3xx status code
func (o *PcloudSppplacementgroupsMembersPostForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sppplacementgroups members post forbidden response has a 4xx status code
func (o *PcloudSppplacementgroupsMembersPostForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud sppplacementgroups members post forbidden response has a 5xx status code
func (o *PcloudSppplacementgroupsMembersPostForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sppplacementgroups members post forbidden response a status code equal to that given
func (o *PcloudSppplacementgroupsMembersPostForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud sppplacementgroups members post forbidden response
func (o *PcloudSppplacementgroupsMembersPostForbidden) Code() int {
	return 403
}

func (o *PcloudSppplacementgroupsMembersPostForbidden) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups/{spp_placement_group_id}/members/{shared_processor_pool_id}][%d] pcloudSppplacementgroupsMembersPostForbidden  %+v", 403, o.Payload)
}

func (o *PcloudSppplacementgroupsMembersPostForbidden) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups/{spp_placement_group_id}/members/{shared_processor_pool_id}][%d] pcloudSppplacementgroupsMembersPostForbidden  %+v", 403, o.Payload)
}

func (o *PcloudSppplacementgroupsMembersPostForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSppplacementgroupsMembersPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSppplacementgroupsMembersPostNotFound creates a PcloudSppplacementgroupsMembersPostNotFound with default headers values
func NewPcloudSppplacementgroupsMembersPostNotFound() *PcloudSppplacementgroupsMembersPostNotFound {
	return &PcloudSppplacementgroupsMembersPostNotFound{}
}

/*
PcloudSppplacementgroupsMembersPostNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudSppplacementgroupsMembersPostNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud sppplacementgroups members post not found response has a 2xx status code
func (o *PcloudSppplacementgroupsMembersPostNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud sppplacementgroups members post not found response has a 3xx status code
func (o *PcloudSppplacementgroupsMembersPostNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sppplacementgroups members post not found response has a 4xx status code
func (o *PcloudSppplacementgroupsMembersPostNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud sppplacementgroups members post not found response has a 5xx status code
func (o *PcloudSppplacementgroupsMembersPostNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sppplacementgroups members post not found response a status code equal to that given
func (o *PcloudSppplacementgroupsMembersPostNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud sppplacementgroups members post not found response
func (o *PcloudSppplacementgroupsMembersPostNotFound) Code() int {
	return 404
}

func (o *PcloudSppplacementgroupsMembersPostNotFound) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups/{spp_placement_group_id}/members/{shared_processor_pool_id}][%d] pcloudSppplacementgroupsMembersPostNotFound  %+v", 404, o.Payload)
}

func (o *PcloudSppplacementgroupsMembersPostNotFound) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups/{spp_placement_group_id}/members/{shared_processor_pool_id}][%d] pcloudSppplacementgroupsMembersPostNotFound  %+v", 404, o.Payload)
}

func (o *PcloudSppplacementgroupsMembersPostNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSppplacementgroupsMembersPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSppplacementgroupsMembersPostConflict creates a PcloudSppplacementgroupsMembersPostConflict with default headers values
func NewPcloudSppplacementgroupsMembersPostConflict() *PcloudSppplacementgroupsMembersPostConflict {
	return &PcloudSppplacementgroupsMembersPostConflict{}
}

/*
PcloudSppplacementgroupsMembersPostConflict describes a response with status code 409, with default header values.

Conflict
*/
type PcloudSppplacementgroupsMembersPostConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud sppplacementgroups members post conflict response has a 2xx status code
func (o *PcloudSppplacementgroupsMembersPostConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud sppplacementgroups members post conflict response has a 3xx status code
func (o *PcloudSppplacementgroupsMembersPostConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sppplacementgroups members post conflict response has a 4xx status code
func (o *PcloudSppplacementgroupsMembersPostConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud sppplacementgroups members post conflict response has a 5xx status code
func (o *PcloudSppplacementgroupsMembersPostConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sppplacementgroups members post conflict response a status code equal to that given
func (o *PcloudSppplacementgroupsMembersPostConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the pcloud sppplacementgroups members post conflict response
func (o *PcloudSppplacementgroupsMembersPostConflict) Code() int {
	return 409
}

func (o *PcloudSppplacementgroupsMembersPostConflict) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups/{spp_placement_group_id}/members/{shared_processor_pool_id}][%d] pcloudSppplacementgroupsMembersPostConflict  %+v", 409, o.Payload)
}

func (o *PcloudSppplacementgroupsMembersPostConflict) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups/{spp_placement_group_id}/members/{shared_processor_pool_id}][%d] pcloudSppplacementgroupsMembersPostConflict  %+v", 409, o.Payload)
}

func (o *PcloudSppplacementgroupsMembersPostConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSppplacementgroupsMembersPostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSppplacementgroupsMembersPostUnprocessableEntity creates a PcloudSppplacementgroupsMembersPostUnprocessableEntity with default headers values
func NewPcloudSppplacementgroupsMembersPostUnprocessableEntity() *PcloudSppplacementgroupsMembersPostUnprocessableEntity {
	return &PcloudSppplacementgroupsMembersPostUnprocessableEntity{}
}

/*
PcloudSppplacementgroupsMembersPostUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudSppplacementgroupsMembersPostUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud sppplacementgroups members post unprocessable entity response has a 2xx status code
func (o *PcloudSppplacementgroupsMembersPostUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud sppplacementgroups members post unprocessable entity response has a 3xx status code
func (o *PcloudSppplacementgroupsMembersPostUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sppplacementgroups members post unprocessable entity response has a 4xx status code
func (o *PcloudSppplacementgroupsMembersPostUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud sppplacementgroups members post unprocessable entity response has a 5xx status code
func (o *PcloudSppplacementgroupsMembersPostUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sppplacementgroups members post unprocessable entity response a status code equal to that given
func (o *PcloudSppplacementgroupsMembersPostUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the pcloud sppplacementgroups members post unprocessable entity response
func (o *PcloudSppplacementgroupsMembersPostUnprocessableEntity) Code() int {
	return 422
}

func (o *PcloudSppplacementgroupsMembersPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups/{spp_placement_group_id}/members/{shared_processor_pool_id}][%d] pcloudSppplacementgroupsMembersPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudSppplacementgroupsMembersPostUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups/{spp_placement_group_id}/members/{shared_processor_pool_id}][%d] pcloudSppplacementgroupsMembersPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudSppplacementgroupsMembersPostUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSppplacementgroupsMembersPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSppplacementgroupsMembersPostInternalServerError creates a PcloudSppplacementgroupsMembersPostInternalServerError with default headers values
func NewPcloudSppplacementgroupsMembersPostInternalServerError() *PcloudSppplacementgroupsMembersPostInternalServerError {
	return &PcloudSppplacementgroupsMembersPostInternalServerError{}
}

/*
PcloudSppplacementgroupsMembersPostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudSppplacementgroupsMembersPostInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud sppplacementgroups members post internal server error response has a 2xx status code
func (o *PcloudSppplacementgroupsMembersPostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud sppplacementgroups members post internal server error response has a 3xx status code
func (o *PcloudSppplacementgroupsMembersPostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sppplacementgroups members post internal server error response has a 4xx status code
func (o *PcloudSppplacementgroupsMembersPostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud sppplacementgroups members post internal server error response has a 5xx status code
func (o *PcloudSppplacementgroupsMembersPostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud sppplacementgroups members post internal server error response a status code equal to that given
func (o *PcloudSppplacementgroupsMembersPostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud sppplacementgroups members post internal server error response
func (o *PcloudSppplacementgroupsMembersPostInternalServerError) Code() int {
	return 500
}

func (o *PcloudSppplacementgroupsMembersPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups/{spp_placement_group_id}/members/{shared_processor_pool_id}][%d] pcloudSppplacementgroupsMembersPostInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudSppplacementgroupsMembersPostInternalServerError) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups/{spp_placement_group_id}/members/{shared_processor_pool_id}][%d] pcloudSppplacementgroupsMembersPostInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudSppplacementgroupsMembersPostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSppplacementgroupsMembersPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
