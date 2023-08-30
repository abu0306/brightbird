// Code generated by go-swagger; DO NOT EDIT.

package plugin

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

// NewUploadPluginFilesParamsParams creates a new UploadPluginFilesParamsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUploadPluginFilesParamsParams() *UploadPluginFilesParamsParams {
	return &UploadPluginFilesParamsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUploadPluginFilesParamsParamsWithTimeout creates a new UploadPluginFilesParamsParams object
// with the ability to set a timeout on a request.
func NewUploadPluginFilesParamsParamsWithTimeout(timeout time.Duration) *UploadPluginFilesParamsParams {
	return &UploadPluginFilesParamsParams{
		timeout: timeout,
	}
}

// NewUploadPluginFilesParamsParamsWithContext creates a new UploadPluginFilesParamsParams object
// with the ability to set a context for a request.
func NewUploadPluginFilesParamsParamsWithContext(ctx context.Context) *UploadPluginFilesParamsParams {
	return &UploadPluginFilesParamsParams{
		Context: ctx,
	}
}

// NewUploadPluginFilesParamsParamsWithHTTPClient creates a new UploadPluginFilesParamsParams object
// with the ability to set a custom HTTPClient for a request.
func NewUploadPluginFilesParamsParamsWithHTTPClient(client *http.Client) *UploadPluginFilesParamsParams {
	return &UploadPluginFilesParamsParams{
		HTTPClient: client,
	}
}

/*
UploadPluginFilesParamsParams contains all the parameters to send to the API endpoint

	for the upload plugin files params operation.

	Typically these are written to a http.Request.
*/
type UploadPluginFilesParamsParams struct {

	/* Labels.

	   PluginLabels Plugin Labels
	*/
	Labels []string

	/* Plugin.

	   Plugin file.
	*/
	PluginFile runtime.NamedReadCloser

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the upload plugin files params params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadPluginFilesParamsParams) WithDefaults() *UploadPluginFilesParamsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the upload plugin files params params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadPluginFilesParamsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the upload plugin files params params
func (o *UploadPluginFilesParamsParams) WithTimeout(timeout time.Duration) *UploadPluginFilesParamsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upload plugin files params params
func (o *UploadPluginFilesParamsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upload plugin files params params
func (o *UploadPluginFilesParamsParams) WithContext(ctx context.Context) *UploadPluginFilesParamsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upload plugin files params params
func (o *UploadPluginFilesParamsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upload plugin files params params
func (o *UploadPluginFilesParamsParams) WithHTTPClient(client *http.Client) *UploadPluginFilesParamsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upload plugin files params params
func (o *UploadPluginFilesParamsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLabels adds the labels to the upload plugin files params params
func (o *UploadPluginFilesParamsParams) WithLabels(labels []string) *UploadPluginFilesParamsParams {
	o.SetLabels(labels)
	return o
}

// SetLabels adds the labels to the upload plugin files params params
func (o *UploadPluginFilesParamsParams) SetLabels(labels []string) {
	o.Labels = labels
}

// WithPluginFile adds the plugin to the upload plugin files params params
func (o *UploadPluginFilesParamsParams) WithPluginFile(plugin runtime.NamedReadCloser) *UploadPluginFilesParamsParams {
	o.SetPluginFile(plugin)
	return o
}

// SetPluginFile adds the plugin to the upload plugin files params params
func (o *UploadPluginFilesParamsParams) SetPluginFile(plugin runtime.NamedReadCloser) {
	o.PluginFile = plugin
}

// WriteToRequest writes these params to a swagger request
func (o *UploadPluginFilesParamsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Labels != nil {

		// binding items for labels
		joinedLabels := o.bindParamLabels(reg)

		// form array param labels
		if err := r.SetFormParam("labels", joinedLabels...); err != nil {
			return err
		}
	}

	if o.PluginFile != nil {

		if o.PluginFile != nil {
			// form file param plugin
			if err := r.SetFileParam("plugin", o.PluginFile); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamUploadPluginFilesParams binds the parameter labels
func (o *UploadPluginFilesParamsParams) bindParamLabels(formats strfmt.Registry) []string {
	labelsIR := o.Labels

	var labelsIC []string
	for _, labelsIIR := range labelsIR { // explode []string

		labelsIIV := labelsIIR // string as string
		labelsIC = append(labelsIC, labelsIIV)
	}

	// items.CollectionFormat: ""
	labelsIS := swag.JoinByFormat(labelsIC, "")

	return labelsIS
}
