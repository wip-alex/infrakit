package templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewTemplatesHeadByNameParams creates a new TemplatesHeadByNameParams object
// with the default values initialized.
func NewTemplatesHeadByNameParams() *TemplatesHeadByNameParams {
	var ()
	return &TemplatesHeadByNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTemplatesHeadByNameParamsWithTimeout creates a new TemplatesHeadByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTemplatesHeadByNameParamsWithTimeout(timeout time.Duration) *TemplatesHeadByNameParams {
	var ()
	return &TemplatesHeadByNameParams{

		timeout: timeout,
	}
}

// NewTemplatesHeadByNameParamsWithContext creates a new TemplatesHeadByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewTemplatesHeadByNameParamsWithContext(ctx context.Context) *TemplatesHeadByNameParams {
	var ()
	return &TemplatesHeadByNameParams{

		Context: ctx,
	}
}

// NewTemplatesHeadByNameParamsWithHTTPClient creates a new TemplatesHeadByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTemplatesHeadByNameParamsWithHTTPClient(client *http.Client) *TemplatesHeadByNameParams {
	var ()
	return &TemplatesHeadByNameParams{
		HTTPClient: client,
	}
}

/*TemplatesHeadByNameParams contains all the parameters to send to the API endpoint
for the templates head by name operation typically these are written to a http.Request
*/
type TemplatesHeadByNameParams struct {

	/*Macs
	  List of valid MAC addresses to lookup

	*/
	Macs []string
	/*Name
	  The name of the template

	*/
	Name string
	/*NodeID
	  The node identifier

	*/
	NodeID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the templates head by name params
func (o *TemplatesHeadByNameParams) WithTimeout(timeout time.Duration) *TemplatesHeadByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the templates head by name params
func (o *TemplatesHeadByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the templates head by name params
func (o *TemplatesHeadByNameParams) WithContext(ctx context.Context) *TemplatesHeadByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the templates head by name params
func (o *TemplatesHeadByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the templates head by name params
func (o *TemplatesHeadByNameParams) WithHTTPClient(client *http.Client) *TemplatesHeadByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the templates head by name params
func (o *TemplatesHeadByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMacs adds the macs to the templates head by name params
func (o *TemplatesHeadByNameParams) WithMacs(macs []string) *TemplatesHeadByNameParams {
	o.SetMacs(macs)
	return o
}

// SetMacs adds the macs to the templates head by name params
func (o *TemplatesHeadByNameParams) SetMacs(macs []string) {
	o.Macs = macs
}

// WithName adds the name to the templates head by name params
func (o *TemplatesHeadByNameParams) WithName(name string) *TemplatesHeadByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the templates head by name params
func (o *TemplatesHeadByNameParams) SetName(name string) {
	o.Name = name
}

// WithNodeID adds the nodeID to the templates head by name params
func (o *TemplatesHeadByNameParams) WithNodeID(nodeID *string) *TemplatesHeadByNameParams {
	o.SetNodeID(nodeID)
	return o
}

// SetNodeID adds the nodeId to the templates head by name params
func (o *TemplatesHeadByNameParams) SetNodeID(nodeID *string) {
	o.NodeID = nodeID
}

// WriteToRequest writes these params to a swagger request
func (o *TemplatesHeadByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesMacs := o.Macs

	joinedMacs := swag.JoinByFormat(valuesMacs, "multi")
	// query array param macs
	if err := r.SetQueryParam("macs", joinedMacs...); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if o.NodeID != nil {

		// query param nodeId
		var qrNodeID string
		if o.NodeID != nil {
			qrNodeID = *o.NodeID
		}
		qNodeID := qrNodeID
		if qNodeID != "" {
			if err := r.SetQueryParam("nodeId", qNodeID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
