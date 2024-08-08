// Code generated by go-swagger; DO NOT EDIT.

package network_security_groups

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
)

// NewV1NetworkSecurityGroupsListParams creates a new V1NetworkSecurityGroupsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewV1NetworkSecurityGroupsListParams() *V1NetworkSecurityGroupsListParams {
	return &V1NetworkSecurityGroupsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewV1NetworkSecurityGroupsListParamsWithTimeout creates a new V1NetworkSecurityGroupsListParams object
// with the ability to set a timeout on a request.
func NewV1NetworkSecurityGroupsListParamsWithTimeout(timeout time.Duration) *V1NetworkSecurityGroupsListParams {
	return &V1NetworkSecurityGroupsListParams{
		timeout: timeout,
	}
}

// NewV1NetworkSecurityGroupsListParamsWithContext creates a new V1NetworkSecurityGroupsListParams object
// with the ability to set a context for a request.
func NewV1NetworkSecurityGroupsListParamsWithContext(ctx context.Context) *V1NetworkSecurityGroupsListParams {
	return &V1NetworkSecurityGroupsListParams{
		Context: ctx,
	}
}

// NewV1NetworkSecurityGroupsListParamsWithHTTPClient creates a new V1NetworkSecurityGroupsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewV1NetworkSecurityGroupsListParamsWithHTTPClient(client *http.Client) *V1NetworkSecurityGroupsListParams {
	return &V1NetworkSecurityGroupsListParams{
		HTTPClient: client,
	}
}

/*
V1NetworkSecurityGroupsListParams contains all the parameters to send to the API endpoint

	for the v1 network security groups list operation.

	Typically these are written to a http.Request.
*/
type V1NetworkSecurityGroupsListParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the v1 network security groups list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V1NetworkSecurityGroupsListParams) WithDefaults() *V1NetworkSecurityGroupsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the v1 network security groups list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V1NetworkSecurityGroupsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the v1 network security groups list params
func (o *V1NetworkSecurityGroupsListParams) WithTimeout(timeout time.Duration) *V1NetworkSecurityGroupsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 network security groups list params
func (o *V1NetworkSecurityGroupsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 network security groups list params
func (o *V1NetworkSecurityGroupsListParams) WithContext(ctx context.Context) *V1NetworkSecurityGroupsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 network security groups list params
func (o *V1NetworkSecurityGroupsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 network security groups list params
func (o *V1NetworkSecurityGroupsListParams) WithHTTPClient(client *http.Client) *V1NetworkSecurityGroupsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 network security groups list params
func (o *V1NetworkSecurityGroupsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *V1NetworkSecurityGroupsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}