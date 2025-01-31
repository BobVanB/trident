// Code generated by go-swagger; DO NOT EDIT.

package storage

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

// NewStorageBridgeGetParams creates a new StorageBridgeGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStorageBridgeGetParams() *StorageBridgeGetParams {
	return &StorageBridgeGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStorageBridgeGetParamsWithTimeout creates a new StorageBridgeGetParams object
// with the ability to set a timeout on a request.
func NewStorageBridgeGetParamsWithTimeout(timeout time.Duration) *StorageBridgeGetParams {
	return &StorageBridgeGetParams{
		timeout: timeout,
	}
}

// NewStorageBridgeGetParamsWithContext creates a new StorageBridgeGetParams object
// with the ability to set a context for a request.
func NewStorageBridgeGetParamsWithContext(ctx context.Context) *StorageBridgeGetParams {
	return &StorageBridgeGetParams{
		Context: ctx,
	}
}

// NewStorageBridgeGetParamsWithHTTPClient creates a new StorageBridgeGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewStorageBridgeGetParamsWithHTTPClient(client *http.Client) *StorageBridgeGetParams {
	return &StorageBridgeGetParams{
		HTTPClient: client,
	}
}

/* StorageBridgeGetParams contains all the parameters to send to the API endpoint
   for the storage bridge get operation.

   Typically these are written to a http.Request.
*/
type StorageBridgeGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	// Wwn.
	WwnPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the storage bridge get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StorageBridgeGetParams) WithDefaults() *StorageBridgeGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the storage bridge get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StorageBridgeGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the storage bridge get params
func (o *StorageBridgeGetParams) WithTimeout(timeout time.Duration) *StorageBridgeGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage bridge get params
func (o *StorageBridgeGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage bridge get params
func (o *StorageBridgeGetParams) WithContext(ctx context.Context) *StorageBridgeGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage bridge get params
func (o *StorageBridgeGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage bridge get params
func (o *StorageBridgeGetParams) WithHTTPClient(client *http.Client) *StorageBridgeGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage bridge get params
func (o *StorageBridgeGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the storage bridge get params
func (o *StorageBridgeGetParams) WithFieldsQueryParameter(fields []string) *StorageBridgeGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the storage bridge get params
func (o *StorageBridgeGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithWwnPathParameter adds the wwn to the storage bridge get params
func (o *StorageBridgeGetParams) WithWwnPathParameter(wwn string) *StorageBridgeGetParams {
	o.SetWwnPathParameter(wwn)
	return o
}

// SetWwnPathParameter adds the wwn to the storage bridge get params
func (o *StorageBridgeGetParams) SetWwnPathParameter(wwn string) {
	o.WwnPathParameter = wwn
}

// WriteToRequest writes these params to a swagger request
func (o *StorageBridgeGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FieldsQueryParameter != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	// path param wwn
	if err := r.SetPathParam("wwn", o.WwnPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamStorageBridgeGet binds the parameter fields
func (o *StorageBridgeGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.FieldsQueryParameter

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}
