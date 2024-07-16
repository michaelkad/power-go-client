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

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// NewPcloudV2VolumesGetallParams creates a new PcloudV2VolumesGetallParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudV2VolumesGetallParams() *PcloudV2VolumesGetallParams {
	return &PcloudV2VolumesGetallParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudV2VolumesGetallParamsWithTimeout creates a new PcloudV2VolumesGetallParams object
// with the ability to set a timeout on a request.
func NewPcloudV2VolumesGetallParamsWithTimeout(timeout time.Duration) *PcloudV2VolumesGetallParams {
	return &PcloudV2VolumesGetallParams{
		timeout: timeout,
	}
}

// NewPcloudV2VolumesGetallParamsWithContext creates a new PcloudV2VolumesGetallParams object
// with the ability to set a context for a request.
func NewPcloudV2VolumesGetallParamsWithContext(ctx context.Context) *PcloudV2VolumesGetallParams {
	return &PcloudV2VolumesGetallParams{
		Context: ctx,
	}
}

// NewPcloudV2VolumesGetallParamsWithHTTPClient creates a new PcloudV2VolumesGetallParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudV2VolumesGetallParamsWithHTTPClient(client *http.Client) *PcloudV2VolumesGetallParams {
	return &PcloudV2VolumesGetallParams{
		HTTPClient: client,
	}
}

/*
PcloudV2VolumesGetallParams contains all the parameters to send to the API endpoint

	for the pcloud v2 volumes getall operation.

	Typically these are written to a http.Request.
*/
type PcloudV2VolumesGetallParams struct {

	/* Body.

	   Parameters for fetching list of specified volumes
	*/
	Body *models.GetBulkVolume

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud v2 volumes getall params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudV2VolumesGetallParams) WithDefaults() *PcloudV2VolumesGetallParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud v2 volumes getall params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudV2VolumesGetallParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud v2 volumes getall params
func (o *PcloudV2VolumesGetallParams) WithTimeout(timeout time.Duration) *PcloudV2VolumesGetallParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud v2 volumes getall params
func (o *PcloudV2VolumesGetallParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud v2 volumes getall params
func (o *PcloudV2VolumesGetallParams) WithContext(ctx context.Context) *PcloudV2VolumesGetallParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud v2 volumes getall params
func (o *PcloudV2VolumesGetallParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud v2 volumes getall params
func (o *PcloudV2VolumesGetallParams) WithHTTPClient(client *http.Client) *PcloudV2VolumesGetallParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud v2 volumes getall params
func (o *PcloudV2VolumesGetallParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pcloud v2 volumes getall params
func (o *PcloudV2VolumesGetallParams) WithBody(body *models.GetBulkVolume) *PcloudV2VolumesGetallParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pcloud v2 volumes getall params
func (o *PcloudV2VolumesGetallParams) SetBody(body *models.GetBulkVolume) {
	o.Body = body
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud v2 volumes getall params
func (o *PcloudV2VolumesGetallParams) WithCloudInstanceID(cloudInstanceID string) *PcloudV2VolumesGetallParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud v2 volumes getall params
func (o *PcloudV2VolumesGetallParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudV2VolumesGetallParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
