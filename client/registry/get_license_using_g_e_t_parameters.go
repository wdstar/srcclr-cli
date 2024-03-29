// Code generated by go-swagger; DO NOT EDIT.

package registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetLicenseUsingGETParams creates a new GetLicenseUsingGETParams object
// with the default values initialized.
func NewGetLicenseUsingGETParams() *GetLicenseUsingGETParams {
	var ()
	return &GetLicenseUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLicenseUsingGETParamsWithTimeout creates a new GetLicenseUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLicenseUsingGETParamsWithTimeout(timeout time.Duration) *GetLicenseUsingGETParams {
	var ()
	return &GetLicenseUsingGETParams{

		timeout: timeout,
	}
}

// NewGetLicenseUsingGETParamsWithContext creates a new GetLicenseUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLicenseUsingGETParamsWithContext(ctx context.Context) *GetLicenseUsingGETParams {
	var ()
	return &GetLicenseUsingGETParams{

		Context: ctx,
	}
}

// NewGetLicenseUsingGETParamsWithHTTPClient creates a new GetLicenseUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLicenseUsingGETParamsWithHTTPClient(client *http.Client) *GetLicenseUsingGETParams {
	var ()
	return &GetLicenseUsingGETParams{
		HTTPClient: client,
	}
}

/*GetLicenseUsingGETParams contains all the parameters to send to the API endpoint
for the get license using g e t operation typically these are written to a http.Request
*/
type GetLicenseUsingGETParams struct {

	/*ID
	  id

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get license using g e t params
func (o *GetLicenseUsingGETParams) WithTimeout(timeout time.Duration) *GetLicenseUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get license using g e t params
func (o *GetLicenseUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get license using g e t params
func (o *GetLicenseUsingGETParams) WithContext(ctx context.Context) *GetLicenseUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get license using g e t params
func (o *GetLicenseUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get license using g e t params
func (o *GetLicenseUsingGETParams) WithHTTPClient(client *http.Client) *GetLicenseUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get license using g e t params
func (o *GetLicenseUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get license using g e t params
func (o *GetLicenseUsingGETParams) WithID(id string) *GetLicenseUsingGETParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get license using g e t params
func (o *GetLicenseUsingGETParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetLicenseUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
