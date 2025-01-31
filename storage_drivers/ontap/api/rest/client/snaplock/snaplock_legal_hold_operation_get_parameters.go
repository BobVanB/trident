// Code generated by go-swagger; DO NOT EDIT.

package snaplock

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

// NewSnaplockLegalHoldOperationGetParams creates a new SnaplockLegalHoldOperationGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSnaplockLegalHoldOperationGetParams() *SnaplockLegalHoldOperationGetParams {
	return &SnaplockLegalHoldOperationGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSnaplockLegalHoldOperationGetParamsWithTimeout creates a new SnaplockLegalHoldOperationGetParams object
// with the ability to set a timeout on a request.
func NewSnaplockLegalHoldOperationGetParamsWithTimeout(timeout time.Duration) *SnaplockLegalHoldOperationGetParams {
	return &SnaplockLegalHoldOperationGetParams{
		timeout: timeout,
	}
}

// NewSnaplockLegalHoldOperationGetParamsWithContext creates a new SnaplockLegalHoldOperationGetParams object
// with the ability to set a context for a request.
func NewSnaplockLegalHoldOperationGetParamsWithContext(ctx context.Context) *SnaplockLegalHoldOperationGetParams {
	return &SnaplockLegalHoldOperationGetParams{
		Context: ctx,
	}
}

// NewSnaplockLegalHoldOperationGetParamsWithHTTPClient creates a new SnaplockLegalHoldOperationGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSnaplockLegalHoldOperationGetParamsWithHTTPClient(client *http.Client) *SnaplockLegalHoldOperationGetParams {
	return &SnaplockLegalHoldOperationGetParams{
		HTTPClient: client,
	}
}

/* SnaplockLegalHoldOperationGetParams contains all the parameters to send to the API endpoint
   for the snaplock legal hold operation get operation.

   Typically these are written to a http.Request.
*/
type SnaplockLegalHoldOperationGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* ID.

	   Operation ID.
	*/
	IDPathParameter string

	/* LitigationID.

	   Litigation ID
	*/
	LitigationIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the snaplock legal hold operation get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnaplockLegalHoldOperationGetParams) WithDefaults() *SnaplockLegalHoldOperationGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the snaplock legal hold operation get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnaplockLegalHoldOperationGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the snaplock legal hold operation get params
func (o *SnaplockLegalHoldOperationGetParams) WithTimeout(timeout time.Duration) *SnaplockLegalHoldOperationGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snaplock legal hold operation get params
func (o *SnaplockLegalHoldOperationGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snaplock legal hold operation get params
func (o *SnaplockLegalHoldOperationGetParams) WithContext(ctx context.Context) *SnaplockLegalHoldOperationGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snaplock legal hold operation get params
func (o *SnaplockLegalHoldOperationGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snaplock legal hold operation get params
func (o *SnaplockLegalHoldOperationGetParams) WithHTTPClient(client *http.Client) *SnaplockLegalHoldOperationGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snaplock legal hold operation get params
func (o *SnaplockLegalHoldOperationGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the snaplock legal hold operation get params
func (o *SnaplockLegalHoldOperationGetParams) WithFieldsQueryParameter(fields []string) *SnaplockLegalHoldOperationGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the snaplock legal hold operation get params
func (o *SnaplockLegalHoldOperationGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithIDPathParameter adds the id to the snaplock legal hold operation get params
func (o *SnaplockLegalHoldOperationGetParams) WithIDPathParameter(id string) *SnaplockLegalHoldOperationGetParams {
	o.SetIDPathParameter(id)
	return o
}

// SetIDPathParameter adds the id to the snaplock legal hold operation get params
func (o *SnaplockLegalHoldOperationGetParams) SetIDPathParameter(id string) {
	o.IDPathParameter = id
}

// WithLitigationIDPathParameter adds the litigationID to the snaplock legal hold operation get params
func (o *SnaplockLegalHoldOperationGetParams) WithLitigationIDPathParameter(litigationID string) *SnaplockLegalHoldOperationGetParams {
	o.SetLitigationIDPathParameter(litigationID)
	return o
}

// SetLitigationIDPathParameter adds the litigationId to the snaplock legal hold operation get params
func (o *SnaplockLegalHoldOperationGetParams) SetLitigationIDPathParameter(litigationID string) {
	o.LitigationIDPathParameter = litigationID
}

// WriteToRequest writes these params to a swagger request
func (o *SnaplockLegalHoldOperationGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.IDPathParameter); err != nil {
		return err
	}

	// path param litigation.id
	if err := r.SetPathParam("litigation.id", o.LitigationIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSnaplockLegalHoldOperationGet binds the parameter fields
func (o *SnaplockLegalHoldOperationGetParams) bindParamFields(formats strfmt.Registry) []string {
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
