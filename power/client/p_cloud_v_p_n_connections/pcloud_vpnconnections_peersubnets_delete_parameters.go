// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_v_p_n_connections

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

// NewPcloudVpnconnectionsPeersubnetsDeleteParams creates a new PcloudVpnconnectionsPeersubnetsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudVpnconnectionsPeersubnetsDeleteParams() *PcloudVpnconnectionsPeersubnetsDeleteParams {
	return &PcloudVpnconnectionsPeersubnetsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudVpnconnectionsPeersubnetsDeleteParamsWithTimeout creates a new PcloudVpnconnectionsPeersubnetsDeleteParams object
// with the ability to set a timeout on a request.
func NewPcloudVpnconnectionsPeersubnetsDeleteParamsWithTimeout(timeout time.Duration) *PcloudVpnconnectionsPeersubnetsDeleteParams {
	return &PcloudVpnconnectionsPeersubnetsDeleteParams{
		timeout: timeout,
	}
}

// NewPcloudVpnconnectionsPeersubnetsDeleteParamsWithContext creates a new PcloudVpnconnectionsPeersubnetsDeleteParams object
// with the ability to set a context for a request.
func NewPcloudVpnconnectionsPeersubnetsDeleteParamsWithContext(ctx context.Context) *PcloudVpnconnectionsPeersubnetsDeleteParams {
	return &PcloudVpnconnectionsPeersubnetsDeleteParams{
		Context: ctx,
	}
}

// NewPcloudVpnconnectionsPeersubnetsDeleteParamsWithHTTPClient creates a new PcloudVpnconnectionsPeersubnetsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudVpnconnectionsPeersubnetsDeleteParamsWithHTTPClient(client *http.Client) *PcloudVpnconnectionsPeersubnetsDeleteParams {
	return &PcloudVpnconnectionsPeersubnetsDeleteParams{
		HTTPClient: client,
	}
}

/*
PcloudVpnconnectionsPeersubnetsDeleteParams contains all the parameters to send to the API endpoint

	for the pcloud vpnconnections peersubnets delete operation.

	Typically these are written to a http.Request.
*/
type PcloudVpnconnectionsPeersubnetsDeleteParams struct {

	/* Body.

	   Peer subnet to detach
	*/
	Body *models.PeerSubnetUpdate

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	/* VpnConnectionID.

	   ID of a VPN connection
	*/
	VpnConnectionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud vpnconnections peersubnets delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudVpnconnectionsPeersubnetsDeleteParams) WithDefaults() *PcloudVpnconnectionsPeersubnetsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud vpnconnections peersubnets delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudVpnconnectionsPeersubnetsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud vpnconnections peersubnets delete params
func (o *PcloudVpnconnectionsPeersubnetsDeleteParams) WithTimeout(timeout time.Duration) *PcloudVpnconnectionsPeersubnetsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud vpnconnections peersubnets delete params
func (o *PcloudVpnconnectionsPeersubnetsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud vpnconnections peersubnets delete params
func (o *PcloudVpnconnectionsPeersubnetsDeleteParams) WithContext(ctx context.Context) *PcloudVpnconnectionsPeersubnetsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud vpnconnections peersubnets delete params
func (o *PcloudVpnconnectionsPeersubnetsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud vpnconnections peersubnets delete params
func (o *PcloudVpnconnectionsPeersubnetsDeleteParams) WithHTTPClient(client *http.Client) *PcloudVpnconnectionsPeersubnetsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud vpnconnections peersubnets delete params
func (o *PcloudVpnconnectionsPeersubnetsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pcloud vpnconnections peersubnets delete params
func (o *PcloudVpnconnectionsPeersubnetsDeleteParams) WithBody(body *models.PeerSubnetUpdate) *PcloudVpnconnectionsPeersubnetsDeleteParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pcloud vpnconnections peersubnets delete params
func (o *PcloudVpnconnectionsPeersubnetsDeleteParams) SetBody(body *models.PeerSubnetUpdate) {
	o.Body = body
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud vpnconnections peersubnets delete params
func (o *PcloudVpnconnectionsPeersubnetsDeleteParams) WithCloudInstanceID(cloudInstanceID string) *PcloudVpnconnectionsPeersubnetsDeleteParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud vpnconnections peersubnets delete params
func (o *PcloudVpnconnectionsPeersubnetsDeleteParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithVpnConnectionID adds the vpnConnectionID to the pcloud vpnconnections peersubnets delete params
func (o *PcloudVpnconnectionsPeersubnetsDeleteParams) WithVpnConnectionID(vpnConnectionID string) *PcloudVpnconnectionsPeersubnetsDeleteParams {
	o.SetVpnConnectionID(vpnConnectionID)
	return o
}

// SetVpnConnectionID adds the vpnConnectionId to the pcloud vpnconnections peersubnets delete params
func (o *PcloudVpnconnectionsPeersubnetsDeleteParams) SetVpnConnectionID(vpnConnectionID string) {
	o.VpnConnectionID = vpnConnectionID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudVpnconnectionsPeersubnetsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param vpn_connection_id
	if err := r.SetPathParam("vpn_connection_id", o.VpnConnectionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
