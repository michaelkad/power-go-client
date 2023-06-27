// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_images

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.ibm.com/power-iaas/service-broker/power/models"
)

// PcloudCloudinstancesStockimagesGetReader is a Reader for the PcloudCloudinstancesStockimagesGet structure.
type PcloudCloudinstancesStockimagesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudCloudinstancesStockimagesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudCloudinstancesStockimagesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudCloudinstancesStockimagesGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudCloudinstancesStockimagesGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudCloudinstancesStockimagesGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudCloudinstancesStockimagesGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/stock-images/{image_id}] pcloud.cloudinstances.stockimages.get", response, response.Code())
	}
}

// NewPcloudCloudinstancesStockimagesGetOK creates a PcloudCloudinstancesStockimagesGetOK with default headers values
func NewPcloudCloudinstancesStockimagesGetOK() *PcloudCloudinstancesStockimagesGetOK {
	return &PcloudCloudinstancesStockimagesGetOK{}
}

/*
PcloudCloudinstancesStockimagesGetOK describes a response with status code 200, with default header values.

OK
*/
type PcloudCloudinstancesStockimagesGetOK struct {
	Payload *models.Image
}

// IsSuccess returns true when this pcloud cloudinstances stockimages get o k response has a 2xx status code
func (o *PcloudCloudinstancesStockimagesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud cloudinstances stockimages get o k response has a 3xx status code
func (o *PcloudCloudinstancesStockimagesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances stockimages get o k response has a 4xx status code
func (o *PcloudCloudinstancesStockimagesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudinstances stockimages get o k response has a 5xx status code
func (o *PcloudCloudinstancesStockimagesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances stockimages get o k response a status code equal to that given
func (o *PcloudCloudinstancesStockimagesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud cloudinstances stockimages get o k response
func (o *PcloudCloudinstancesStockimagesGetOK) Code() int {
	return 200
}

func (o *PcloudCloudinstancesStockimagesGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/stock-images/{image_id}][%d] pcloudCloudinstancesStockimagesGetOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudinstancesStockimagesGetOK) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/stock-images/{image_id}][%d] pcloudCloudinstancesStockimagesGetOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudinstancesStockimagesGetOK) GetPayload() *models.Image {
	return o.Payload
}

func (o *PcloudCloudinstancesStockimagesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Image)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesStockimagesGetBadRequest creates a PcloudCloudinstancesStockimagesGetBadRequest with default headers values
func NewPcloudCloudinstancesStockimagesGetBadRequest() *PcloudCloudinstancesStockimagesGetBadRequest {
	return &PcloudCloudinstancesStockimagesGetBadRequest{}
}

/*
PcloudCloudinstancesStockimagesGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudCloudinstancesStockimagesGetBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances stockimages get bad request response has a 2xx status code
func (o *PcloudCloudinstancesStockimagesGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances stockimages get bad request response has a 3xx status code
func (o *PcloudCloudinstancesStockimagesGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances stockimages get bad request response has a 4xx status code
func (o *PcloudCloudinstancesStockimagesGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances stockimages get bad request response has a 5xx status code
func (o *PcloudCloudinstancesStockimagesGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances stockimages get bad request response a status code equal to that given
func (o *PcloudCloudinstancesStockimagesGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud cloudinstances stockimages get bad request response
func (o *PcloudCloudinstancesStockimagesGetBadRequest) Code() int {
	return 400
}

func (o *PcloudCloudinstancesStockimagesGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/stock-images/{image_id}][%d] pcloudCloudinstancesStockimagesGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudinstancesStockimagesGetBadRequest) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/stock-images/{image_id}][%d] pcloudCloudinstancesStockimagesGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudinstancesStockimagesGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesStockimagesGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesStockimagesGetUnauthorized creates a PcloudCloudinstancesStockimagesGetUnauthorized with default headers values
func NewPcloudCloudinstancesStockimagesGetUnauthorized() *PcloudCloudinstancesStockimagesGetUnauthorized {
	return &PcloudCloudinstancesStockimagesGetUnauthorized{}
}

/*
PcloudCloudinstancesStockimagesGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudCloudinstancesStockimagesGetUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances stockimages get unauthorized response has a 2xx status code
func (o *PcloudCloudinstancesStockimagesGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances stockimages get unauthorized response has a 3xx status code
func (o *PcloudCloudinstancesStockimagesGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances stockimages get unauthorized response has a 4xx status code
func (o *PcloudCloudinstancesStockimagesGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances stockimages get unauthorized response has a 5xx status code
func (o *PcloudCloudinstancesStockimagesGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances stockimages get unauthorized response a status code equal to that given
func (o *PcloudCloudinstancesStockimagesGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud cloudinstances stockimages get unauthorized response
func (o *PcloudCloudinstancesStockimagesGetUnauthorized) Code() int {
	return 401
}

func (o *PcloudCloudinstancesStockimagesGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/stock-images/{image_id}][%d] pcloudCloudinstancesStockimagesGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudinstancesStockimagesGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/stock-images/{image_id}][%d] pcloudCloudinstancesStockimagesGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudinstancesStockimagesGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesStockimagesGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesStockimagesGetNotFound creates a PcloudCloudinstancesStockimagesGetNotFound with default headers values
func NewPcloudCloudinstancesStockimagesGetNotFound() *PcloudCloudinstancesStockimagesGetNotFound {
	return &PcloudCloudinstancesStockimagesGetNotFound{}
}

/*
PcloudCloudinstancesStockimagesGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudCloudinstancesStockimagesGetNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances stockimages get not found response has a 2xx status code
func (o *PcloudCloudinstancesStockimagesGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances stockimages get not found response has a 3xx status code
func (o *PcloudCloudinstancesStockimagesGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances stockimages get not found response has a 4xx status code
func (o *PcloudCloudinstancesStockimagesGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances stockimages get not found response has a 5xx status code
func (o *PcloudCloudinstancesStockimagesGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances stockimages get not found response a status code equal to that given
func (o *PcloudCloudinstancesStockimagesGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud cloudinstances stockimages get not found response
func (o *PcloudCloudinstancesStockimagesGetNotFound) Code() int {
	return 404
}

func (o *PcloudCloudinstancesStockimagesGetNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/stock-images/{image_id}][%d] pcloudCloudinstancesStockimagesGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudCloudinstancesStockimagesGetNotFound) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/stock-images/{image_id}][%d] pcloudCloudinstancesStockimagesGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudCloudinstancesStockimagesGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesStockimagesGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesStockimagesGetInternalServerError creates a PcloudCloudinstancesStockimagesGetInternalServerError with default headers values
func NewPcloudCloudinstancesStockimagesGetInternalServerError() *PcloudCloudinstancesStockimagesGetInternalServerError {
	return &PcloudCloudinstancesStockimagesGetInternalServerError{}
}

/*
PcloudCloudinstancesStockimagesGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudCloudinstancesStockimagesGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances stockimages get internal server error response has a 2xx status code
func (o *PcloudCloudinstancesStockimagesGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances stockimages get internal server error response has a 3xx status code
func (o *PcloudCloudinstancesStockimagesGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances stockimages get internal server error response has a 4xx status code
func (o *PcloudCloudinstancesStockimagesGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudinstances stockimages get internal server error response has a 5xx status code
func (o *PcloudCloudinstancesStockimagesGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud cloudinstances stockimages get internal server error response a status code equal to that given
func (o *PcloudCloudinstancesStockimagesGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud cloudinstances stockimages get internal server error response
func (o *PcloudCloudinstancesStockimagesGetInternalServerError) Code() int {
	return 500
}

func (o *PcloudCloudinstancesStockimagesGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/stock-images/{image_id}][%d] pcloudCloudinstancesStockimagesGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudinstancesStockimagesGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/stock-images/{image_id}][%d] pcloudCloudinstancesStockimagesGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudinstancesStockimagesGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesStockimagesGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
