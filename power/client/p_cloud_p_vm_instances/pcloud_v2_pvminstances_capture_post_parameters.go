// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_p_vm_instances

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

// NewPcloudV2PvminstancesCapturePostParams creates a new PcloudV2PvminstancesCapturePostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudV2PvminstancesCapturePostParams() *PcloudV2PvminstancesCapturePostParams {
	return &PcloudV2PvminstancesCapturePostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudV2PvminstancesCapturePostParamsWithTimeout creates a new PcloudV2PvminstancesCapturePostParams object
// with the ability to set a timeout on a request.
func NewPcloudV2PvminstancesCapturePostParamsWithTimeout(timeout time.Duration) *PcloudV2PvminstancesCapturePostParams {
	return &PcloudV2PvminstancesCapturePostParams{
		timeout: timeout,
	}
}

// NewPcloudV2PvminstancesCapturePostParamsWithContext creates a new PcloudV2PvminstancesCapturePostParams object
// with the ability to set a context for a request.
func NewPcloudV2PvminstancesCapturePostParamsWithContext(ctx context.Context) *PcloudV2PvminstancesCapturePostParams {
	return &PcloudV2PvminstancesCapturePostParams{
		Context: ctx,
	}
}

// NewPcloudV2PvminstancesCapturePostParamsWithHTTPClient creates a new PcloudV2PvminstancesCapturePostParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudV2PvminstancesCapturePostParamsWithHTTPClient(client *http.Client) *PcloudV2PvminstancesCapturePostParams {
	return &PcloudV2PvminstancesCapturePostParams{
		HTTPClient: client,
	}
}

/*
PcloudV2PvminstancesCapturePostParams contains all the parameters to send to the API endpoint

	for the pcloud v2 pvminstances capture post operation.

	Typically these are written to a http.Request.
*/
type PcloudV2PvminstancesCapturePostParams struct {

	/* Body.

	   Parameters for the capture
	*/
	Body *models.PVMInstanceCapture

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	/* PvmInstanceID.

	   PCloud PVM Instance ID
	*/
	PvmInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud v2 pvminstances capture post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudV2PvminstancesCapturePostParams) WithDefaults() *PcloudV2PvminstancesCapturePostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud v2 pvminstances capture post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudV2PvminstancesCapturePostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud v2 pvminstances capture post params
func (o *PcloudV2PvminstancesCapturePostParams) WithTimeout(timeout time.Duration) *PcloudV2PvminstancesCapturePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud v2 pvminstances capture post params
func (o *PcloudV2PvminstancesCapturePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud v2 pvminstances capture post params
func (o *PcloudV2PvminstancesCapturePostParams) WithContext(ctx context.Context) *PcloudV2PvminstancesCapturePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud v2 pvminstances capture post params
func (o *PcloudV2PvminstancesCapturePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud v2 pvminstances capture post params
func (o *PcloudV2PvminstancesCapturePostParams) WithHTTPClient(client *http.Client) *PcloudV2PvminstancesCapturePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud v2 pvminstances capture post params
func (o *PcloudV2PvminstancesCapturePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pcloud v2 pvminstances capture post params
func (o *PcloudV2PvminstancesCapturePostParams) WithBody(body *models.PVMInstanceCapture) *PcloudV2PvminstancesCapturePostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pcloud v2 pvminstances capture post params
func (o *PcloudV2PvminstancesCapturePostParams) SetBody(body *models.PVMInstanceCapture) {
	o.Body = body
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud v2 pvminstances capture post params
func (o *PcloudV2PvminstancesCapturePostParams) WithCloudInstanceID(cloudInstanceID string) *PcloudV2PvminstancesCapturePostParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud v2 pvminstances capture post params
func (o *PcloudV2PvminstancesCapturePostParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithPvmInstanceID adds the pvmInstanceID to the pcloud v2 pvminstances capture post params
func (o *PcloudV2PvminstancesCapturePostParams) WithPvmInstanceID(pvmInstanceID string) *PcloudV2PvminstancesCapturePostParams {
	o.SetPvmInstanceID(pvmInstanceID)
	return o
}

// SetPvmInstanceID adds the pvmInstanceId to the pcloud v2 pvminstances capture post params
func (o *PcloudV2PvminstancesCapturePostParams) SetPvmInstanceID(pvmInstanceID string) {
	o.PvmInstanceID = pvmInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudV2PvminstancesCapturePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param pvm_instance_id
	if err := r.SetPathParam("pvm_instance_id", o.PvmInstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
