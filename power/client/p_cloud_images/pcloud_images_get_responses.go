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

// PcloudImagesGetReader is a Reader for the PcloudImagesGet structure.
type PcloudImagesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudImagesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudImagesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudImagesGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudImagesGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudImagesGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudImagesGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudImagesGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /pcloud/v1/images/{image_id}] pcloud.images.get", response, response.Code())
	}
}

// NewPcloudImagesGetOK creates a PcloudImagesGetOK with default headers values
func NewPcloudImagesGetOK() *PcloudImagesGetOK {
	return &PcloudImagesGetOK{}
}

/*
PcloudImagesGetOK describes a response with status code 200, with default header values.

OK
*/
type PcloudImagesGetOK struct {
	Payload *models.Image
}

// IsSuccess returns true when this pcloud images get o k response has a 2xx status code
func (o *PcloudImagesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud images get o k response has a 3xx status code
func (o *PcloudImagesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud images get o k response has a 4xx status code
func (o *PcloudImagesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud images get o k response has a 5xx status code
func (o *PcloudImagesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud images get o k response a status code equal to that given
func (o *PcloudImagesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud images get o k response
func (o *PcloudImagesGetOK) Code() int {
	return 200
}

func (o *PcloudImagesGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/images/{image_id}][%d] pcloudImagesGetOK  %+v", 200, o.Payload)
}

func (o *PcloudImagesGetOK) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/images/{image_id}][%d] pcloudImagesGetOK  %+v", 200, o.Payload)
}

func (o *PcloudImagesGetOK) GetPayload() *models.Image {
	return o.Payload
}

func (o *PcloudImagesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Image)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudImagesGetBadRequest creates a PcloudImagesGetBadRequest with default headers values
func NewPcloudImagesGetBadRequest() *PcloudImagesGetBadRequest {
	return &PcloudImagesGetBadRequest{}
}

/*
PcloudImagesGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudImagesGetBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud images get bad request response has a 2xx status code
func (o *PcloudImagesGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud images get bad request response has a 3xx status code
func (o *PcloudImagesGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud images get bad request response has a 4xx status code
func (o *PcloudImagesGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud images get bad request response has a 5xx status code
func (o *PcloudImagesGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud images get bad request response a status code equal to that given
func (o *PcloudImagesGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud images get bad request response
func (o *PcloudImagesGetBadRequest) Code() int {
	return 400
}

func (o *PcloudImagesGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/images/{image_id}][%d] pcloudImagesGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudImagesGetBadRequest) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/images/{image_id}][%d] pcloudImagesGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudImagesGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudImagesGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudImagesGetUnauthorized creates a PcloudImagesGetUnauthorized with default headers values
func NewPcloudImagesGetUnauthorized() *PcloudImagesGetUnauthorized {
	return &PcloudImagesGetUnauthorized{}
}

/*
PcloudImagesGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudImagesGetUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud images get unauthorized response has a 2xx status code
func (o *PcloudImagesGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud images get unauthorized response has a 3xx status code
func (o *PcloudImagesGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud images get unauthorized response has a 4xx status code
func (o *PcloudImagesGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud images get unauthorized response has a 5xx status code
func (o *PcloudImagesGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud images get unauthorized response a status code equal to that given
func (o *PcloudImagesGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud images get unauthorized response
func (o *PcloudImagesGetUnauthorized) Code() int {
	return 401
}

func (o *PcloudImagesGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/images/{image_id}][%d] pcloudImagesGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudImagesGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/images/{image_id}][%d] pcloudImagesGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudImagesGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudImagesGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudImagesGetForbidden creates a PcloudImagesGetForbidden with default headers values
func NewPcloudImagesGetForbidden() *PcloudImagesGetForbidden {
	return &PcloudImagesGetForbidden{}
}

/*
PcloudImagesGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudImagesGetForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud images get forbidden response has a 2xx status code
func (o *PcloudImagesGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud images get forbidden response has a 3xx status code
func (o *PcloudImagesGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud images get forbidden response has a 4xx status code
func (o *PcloudImagesGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud images get forbidden response has a 5xx status code
func (o *PcloudImagesGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud images get forbidden response a status code equal to that given
func (o *PcloudImagesGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud images get forbidden response
func (o *PcloudImagesGetForbidden) Code() int {
	return 403
}

func (o *PcloudImagesGetForbidden) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/images/{image_id}][%d] pcloudImagesGetForbidden  %+v", 403, o.Payload)
}

func (o *PcloudImagesGetForbidden) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/images/{image_id}][%d] pcloudImagesGetForbidden  %+v", 403, o.Payload)
}

func (o *PcloudImagesGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudImagesGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudImagesGetNotFound creates a PcloudImagesGetNotFound with default headers values
func NewPcloudImagesGetNotFound() *PcloudImagesGetNotFound {
	return &PcloudImagesGetNotFound{}
}

/*
PcloudImagesGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudImagesGetNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud images get not found response has a 2xx status code
func (o *PcloudImagesGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud images get not found response has a 3xx status code
func (o *PcloudImagesGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud images get not found response has a 4xx status code
func (o *PcloudImagesGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud images get not found response has a 5xx status code
func (o *PcloudImagesGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud images get not found response a status code equal to that given
func (o *PcloudImagesGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud images get not found response
func (o *PcloudImagesGetNotFound) Code() int {
	return 404
}

func (o *PcloudImagesGetNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/images/{image_id}][%d] pcloudImagesGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudImagesGetNotFound) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/images/{image_id}][%d] pcloudImagesGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudImagesGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudImagesGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudImagesGetInternalServerError creates a PcloudImagesGetInternalServerError with default headers values
func NewPcloudImagesGetInternalServerError() *PcloudImagesGetInternalServerError {
	return &PcloudImagesGetInternalServerError{}
}

/*
PcloudImagesGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudImagesGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud images get internal server error response has a 2xx status code
func (o *PcloudImagesGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud images get internal server error response has a 3xx status code
func (o *PcloudImagesGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud images get internal server error response has a 4xx status code
func (o *PcloudImagesGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud images get internal server error response has a 5xx status code
func (o *PcloudImagesGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud images get internal server error response a status code equal to that given
func (o *PcloudImagesGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud images get internal server error response
func (o *PcloudImagesGetInternalServerError) Code() int {
	return 500
}

func (o *PcloudImagesGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/images/{image_id}][%d] pcloudImagesGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudImagesGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/images/{image_id}][%d] pcloudImagesGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudImagesGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudImagesGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
