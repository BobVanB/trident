// Code generated by go-swagger; DO NOT EDIT.

package cluster

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
	"github.com/go-openapi/swag"
)

// NewNodeDeleteParams creates a new NodeDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNodeDeleteParams() *NodeDeleteParams {
	return &NodeDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNodeDeleteParamsWithTimeout creates a new NodeDeleteParams object
// with the ability to set a timeout on a request.
func NewNodeDeleteParamsWithTimeout(timeout time.Duration) *NodeDeleteParams {
	return &NodeDeleteParams{
		timeout: timeout,
	}
}

// NewNodeDeleteParamsWithContext creates a new NodeDeleteParams object
// with the ability to set a context for a request.
func NewNodeDeleteParamsWithContext(ctx context.Context) *NodeDeleteParams {
	return &NodeDeleteParams{
		Context: ctx,
	}
}

// NewNodeDeleteParamsWithHTTPClient creates a new NodeDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewNodeDeleteParamsWithHTTPClient(client *http.Client) *NodeDeleteParams {
	return &NodeDeleteParams{
		HTTPClient: client,
	}
}

/* NodeDeleteParams contains all the parameters to send to the API endpoint
   for the node delete operation.

   Typically these are written to a http.Request.
*/
type NodeDeleteParams struct {

	/* Force.

	   Set the force flag to "true" to forcibly remove a node that is down and cannot be brought online to remove its shared resources.

	*/
	ForceQueryParameter *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeoutQueryParameter *int64

	// UUID.
	UUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the node delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NodeDeleteParams) WithDefaults() *NodeDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the node delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NodeDeleteParams) SetDefaults() {
	var (
		forceQueryParameterDefault = bool(false)

		returnTimeoutQueryParameterDefault = int64(0)
	)

	val := NodeDeleteParams{
		ForceQueryParameter:         &forceQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the node delete params
func (o *NodeDeleteParams) WithTimeout(timeout time.Duration) *NodeDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the node delete params
func (o *NodeDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the node delete params
func (o *NodeDeleteParams) WithContext(ctx context.Context) *NodeDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the node delete params
func (o *NodeDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the node delete params
func (o *NodeDeleteParams) WithHTTPClient(client *http.Client) *NodeDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the node delete params
func (o *NodeDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForceQueryParameter adds the force to the node delete params
func (o *NodeDeleteParams) WithForceQueryParameter(force *bool) *NodeDeleteParams {
	o.SetForceQueryParameter(force)
	return o
}

// SetForceQueryParameter adds the force to the node delete params
func (o *NodeDeleteParams) SetForceQueryParameter(force *bool) {
	o.ForceQueryParameter = force
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the node delete params
func (o *NodeDeleteParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *NodeDeleteParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the node delete params
func (o *NodeDeleteParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithUUIDPathParameter adds the uuid to the node delete params
func (o *NodeDeleteParams) WithUUIDPathParameter(uuid string) *NodeDeleteParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the node delete params
func (o *NodeDeleteParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *NodeDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ForceQueryParameter != nil {

		// query param force
		var qrForce bool

		if o.ForceQueryParameter != nil {
			qrForce = *o.ForceQueryParameter
		}
		qForce := swag.FormatBool(qrForce)
		if qForce != "" {

			if err := r.SetQueryParam("force", qForce); err != nil {
				return err
			}
		}
	}

	if o.ReturnTimeoutQueryParameter != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeoutQueryParameter != nil {
			qrReturnTimeout = *o.ReturnTimeoutQueryParameter
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
		}
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
