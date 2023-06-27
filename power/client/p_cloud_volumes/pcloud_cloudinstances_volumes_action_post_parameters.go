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

// NewPcloudCloudinstancesVolumesActionPostParams creates a new PcloudCloudinstancesVolumesActionPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudCloudinstancesVolumesActionPostParams() *PcloudCloudinstancesVolumesActionPostParams {
	return &PcloudCloudinstancesVolumesActionPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudCloudinstancesVolumesActionPostParamsWithTimeout creates a new PcloudCloudinstancesVolumesActionPostParams object
// with the ability to set a timeout on a request.
func NewPcloudCloudinstancesVolumesActionPostParamsWithTimeout(timeout time.Duration) *PcloudCloudinstancesVolumesActionPostParams {
	return &PcloudCloudinstancesVolumesActionPostParams{
		timeout: timeout,
	}
}

// NewPcloudCloudinstancesVolumesActionPostParamsWithContext creates a new PcloudCloudinstancesVolumesActionPostParams object
// with the ability to set a context for a request.
func NewPcloudCloudinstancesVolumesActionPostParamsWithContext(ctx context.Context) *PcloudCloudinstancesVolumesActionPostParams {
	return &PcloudCloudinstancesVolumesActionPostParams{
		Context: ctx,
	}
}

// NewPcloudCloudinstancesVolumesActionPostParamsWithHTTPClient creates a new PcloudCloudinstancesVolumesActionPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudCloudinstancesVolumesActionPostParamsWithHTTPClient(client *http.Client) *PcloudCloudinstancesVolumesActionPostParams {
	return &PcloudCloudinstancesVolumesActionPostParams{
		HTTPClient: client,
	}
}

/*
PcloudCloudinstancesVolumesActionPostParams contains all the parameters to send to the API endpoint

	for the pcloud cloudinstances volumes action post operation.

	Typically these are written to a http.Request.
*/
type PcloudCloudinstancesVolumesActionPostParams struct {

	/* Body.

	   Parameters for the desired action
	*/
	Body *models.VolumeAction

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	/* VolumeID.

	   Volume ID
	*/
	VolumeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud cloudinstances volumes action post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudCloudinstancesVolumesActionPostParams) WithDefaults() *PcloudCloudinstancesVolumesActionPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud cloudinstances volumes action post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudCloudinstancesVolumesActionPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud cloudinstances volumes action post params
func (o *PcloudCloudinstancesVolumesActionPostParams) WithTimeout(timeout time.Duration) *PcloudCloudinstancesVolumesActionPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud cloudinstances volumes action post params
func (o *PcloudCloudinstancesVolumesActionPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud cloudinstances volumes action post params
func (o *PcloudCloudinstancesVolumesActionPostParams) WithContext(ctx context.Context) *PcloudCloudinstancesVolumesActionPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud cloudinstances volumes action post params
func (o *PcloudCloudinstancesVolumesActionPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud cloudinstances volumes action post params
func (o *PcloudCloudinstancesVolumesActionPostParams) WithHTTPClient(client *http.Client) *PcloudCloudinstancesVolumesActionPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud cloudinstances volumes action post params
func (o *PcloudCloudinstancesVolumesActionPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pcloud cloudinstances volumes action post params
func (o *PcloudCloudinstancesVolumesActionPostParams) WithBody(body *models.VolumeAction) *PcloudCloudinstancesVolumesActionPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pcloud cloudinstances volumes action post params
func (o *PcloudCloudinstancesVolumesActionPostParams) SetBody(body *models.VolumeAction) {
	o.Body = body
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud cloudinstances volumes action post params
func (o *PcloudCloudinstancesVolumesActionPostParams) WithCloudInstanceID(cloudInstanceID string) *PcloudCloudinstancesVolumesActionPostParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud cloudinstances volumes action post params
func (o *PcloudCloudinstancesVolumesActionPostParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithVolumeID adds the volumeID to the pcloud cloudinstances volumes action post params
func (o *PcloudCloudinstancesVolumesActionPostParams) WithVolumeID(volumeID string) *PcloudCloudinstancesVolumesActionPostParams {
	o.SetVolumeID(volumeID)
	return o
}

// SetVolumeID adds the volumeId to the pcloud cloudinstances volumes action post params
func (o *PcloudCloudinstancesVolumesActionPostParams) SetVolumeID(volumeID string) {
	o.VolumeID = volumeID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudCloudinstancesVolumesActionPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param volume_id
	if err := r.SetPathParam("volume_id", o.VolumeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
