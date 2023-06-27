// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.ibm.com/power-iaas/service-broker/power/models"
)

// NewPcloudV2VolumesPostParams creates a new PcloudV2VolumesPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudV2VolumesPostParams() *PcloudV2VolumesPostParams {
	return &PcloudV2VolumesPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudV2VolumesPostParamsWithTimeout creates a new PcloudV2VolumesPostParams object
// with the ability to set a timeout on a request.
func NewPcloudV2VolumesPostParamsWithTimeout(timeout time.Duration) *PcloudV2VolumesPostParams {
	return &PcloudV2VolumesPostParams{
		timeout: timeout,
	}
}

// NewPcloudV2VolumesPostParamsWithContext creates a new PcloudV2VolumesPostParams object
// with the ability to set a context for a request.
func NewPcloudV2VolumesPostParamsWithContext(ctx context.Context) *PcloudV2VolumesPostParams {
	return &PcloudV2VolumesPostParams{
		Context: ctx,
	}
}

// NewPcloudV2VolumesPostParamsWithHTTPClient creates a new PcloudV2VolumesPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudV2VolumesPostParamsWithHTTPClient(client *http.Client) *PcloudV2VolumesPostParams {
	return &PcloudV2VolumesPostParams{
		HTTPClient: client,
	}
}

/*
PcloudV2VolumesPostParams contains all the parameters to send to the API endpoint

	for the pcloud v2 volumes post operation.

	Typically these are written to a http.Request.
*/
type PcloudV2VolumesPostParams struct {

	/* Body.

	   Parameters for creating multiple volumes
	*/
	Body *models.MultiVolumesCreate

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud v2 volumes post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudV2VolumesPostParams) WithDefaults() *PcloudV2VolumesPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud v2 volumes post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudV2VolumesPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud v2 volumes post params
func (o *PcloudV2VolumesPostParams) WithTimeout(timeout time.Duration) *PcloudV2VolumesPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud v2 volumes post params
func (o *PcloudV2VolumesPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud v2 volumes post params
func (o *PcloudV2VolumesPostParams) WithContext(ctx context.Context) *PcloudV2VolumesPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud v2 volumes post params
func (o *PcloudV2VolumesPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud v2 volumes post params
func (o *PcloudV2VolumesPostParams) WithHTTPClient(client *http.Client) *PcloudV2VolumesPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud v2 volumes post params
func (o *PcloudV2VolumesPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pcloud v2 volumes post params
func (o *PcloudV2VolumesPostParams) WithBody(body *models.MultiVolumesCreate) *PcloudV2VolumesPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pcloud v2 volumes post params
func (o *PcloudV2VolumesPostParams) SetBody(body *models.MultiVolumesCreate) {
	o.Body = body
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud v2 volumes post params
func (o *PcloudV2VolumesPostParams) WithCloudInstanceID(cloudInstanceID string) *PcloudV2VolumesPostParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud v2 volumes post params
func (o *PcloudV2VolumesPostParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudV2VolumesPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
