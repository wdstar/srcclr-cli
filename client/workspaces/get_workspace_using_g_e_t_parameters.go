// Code generated by go-swagger; DO NOT EDIT.

package workspaces

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

// NewGetWorkspaceUsingGETParams creates a new GetWorkspaceUsingGETParams object
// with the default values initialized.
func NewGetWorkspaceUsingGETParams() *GetWorkspaceUsingGETParams {
	var ()
	return &GetWorkspaceUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkspaceUsingGETParamsWithTimeout creates a new GetWorkspaceUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWorkspaceUsingGETParamsWithTimeout(timeout time.Duration) *GetWorkspaceUsingGETParams {
	var ()
	return &GetWorkspaceUsingGETParams{

		timeout: timeout,
	}
}

// NewGetWorkspaceUsingGETParamsWithContext creates a new GetWorkspaceUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWorkspaceUsingGETParamsWithContext(ctx context.Context) *GetWorkspaceUsingGETParams {
	var ()
	return &GetWorkspaceUsingGETParams{

		Context: ctx,
	}
}

// NewGetWorkspaceUsingGETParamsWithHTTPClient creates a new GetWorkspaceUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWorkspaceUsingGETParamsWithHTTPClient(client *http.Client) *GetWorkspaceUsingGETParams {
	var ()
	return &GetWorkspaceUsingGETParams{
		HTTPClient: client,
	}
}

/*GetWorkspaceUsingGETParams contains all the parameters to send to the API endpoint
for the get workspace using g e t operation typically these are written to a http.Request
*/
type GetWorkspaceUsingGETParams struct {

	/*ID
	  The workspace ID.

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get workspace using g e t params
func (o *GetWorkspaceUsingGETParams) WithTimeout(timeout time.Duration) *GetWorkspaceUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workspace using g e t params
func (o *GetWorkspaceUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workspace using g e t params
func (o *GetWorkspaceUsingGETParams) WithContext(ctx context.Context) *GetWorkspaceUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workspace using g e t params
func (o *GetWorkspaceUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workspace using g e t params
func (o *GetWorkspaceUsingGETParams) WithHTTPClient(client *http.Client) *GetWorkspaceUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workspace using g e t params
func (o *GetWorkspaceUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get workspace using g e t params
func (o *GetWorkspaceUsingGETParams) WithID(id strfmt.UUID) *GetWorkspaceUsingGETParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get workspace using g e t params
func (o *GetWorkspaceUsingGETParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkspaceUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
