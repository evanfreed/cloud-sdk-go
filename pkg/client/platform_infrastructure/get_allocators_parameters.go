// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by go-swagger; DO NOT EDIT.

package platform_infrastructure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAllocatorsParams creates a new GetAllocatorsParams object
// with the default values initialized.
func NewGetAllocatorsParams() *GetAllocatorsParams {
	var ()
	return &GetAllocatorsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllocatorsParamsWithTimeout creates a new GetAllocatorsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAllocatorsParamsWithTimeout(timeout time.Duration) *GetAllocatorsParams {
	var ()
	return &GetAllocatorsParams{

		timeout: timeout,
	}
}

// NewGetAllocatorsParamsWithContext creates a new GetAllocatorsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAllocatorsParamsWithContext(ctx context.Context) *GetAllocatorsParams {
	var ()
	return &GetAllocatorsParams{

		Context: ctx,
	}
}

// NewGetAllocatorsParamsWithHTTPClient creates a new GetAllocatorsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAllocatorsParamsWithHTTPClient(client *http.Client) *GetAllocatorsParams {
	var ()
	return &GetAllocatorsParams{
		HTTPClient: client,
	}
}

/*GetAllocatorsParams contains all the parameters to send to the API endpoint
for the get allocators operation typically these are written to a http.Request
*/
type GetAllocatorsParams struct {

	/*From
	  (Optional) The offset from the first result you want to fetch. Defaults to 0.

	*/
	From *int64
	/*Q
	  (Optional) The query that filters the allocators. Maps to an Elasticsearch `query_string` query.

	*/
	Q *string
	/*Size
	  (Optional) The maximum number of search results to return. Defaults to 100.

	*/
	Size *int64
	/*Sort
	  (Optional) An comma-separated array of fields to sort the search results by. Defaults to `allocator_id`.

	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get allocators params
func (o *GetAllocatorsParams) WithTimeout(timeout time.Duration) *GetAllocatorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get allocators params
func (o *GetAllocatorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get allocators params
func (o *GetAllocatorsParams) WithContext(ctx context.Context) *GetAllocatorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get allocators params
func (o *GetAllocatorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get allocators params
func (o *GetAllocatorsParams) WithHTTPClient(client *http.Client) *GetAllocatorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get allocators params
func (o *GetAllocatorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFrom adds the from to the get allocators params
func (o *GetAllocatorsParams) WithFrom(from *int64) *GetAllocatorsParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get allocators params
func (o *GetAllocatorsParams) SetFrom(from *int64) {
	o.From = from
}

// WithQ adds the q to the get allocators params
func (o *GetAllocatorsParams) WithQ(q *string) *GetAllocatorsParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the get allocators params
func (o *GetAllocatorsParams) SetQ(q *string) {
	o.Q = q
}

// WithSize adds the size to the get allocators params
func (o *GetAllocatorsParams) WithSize(size *int64) *GetAllocatorsParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get allocators params
func (o *GetAllocatorsParams) SetSize(size *int64) {
	o.Size = size
}

// WithSort adds the sort to the get allocators params
func (o *GetAllocatorsParams) WithSort(sort *string) *GetAllocatorsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get allocators params
func (o *GetAllocatorsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllocatorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.From != nil {

		// query param from
		var qrFrom int64
		if o.From != nil {
			qrFrom = *o.From
		}
		qFrom := swag.FormatInt64(qrFrom)
		if qFrom != "" {
			if err := r.SetQueryParam("from", qFrom); err != nil {
				return err
			}
		}

	}

	if o.Q != nil {

		// query param q
		var qrQ string
		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {
			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}

	}

	if o.Size != nil {

		// query param size
		var qrSize int64
		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := swag.FormatInt64(qrSize)
		if qSize != "" {
			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}

	}

	if o.Sort != nil {

		// query param sort
		var qrSort string
		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {
			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}