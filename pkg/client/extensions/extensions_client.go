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

package extensions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new extensions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for extensions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateExtension creates an extension

Creates an extension.
*/
func (a *Client) CreateExtension(params *CreateExtensionParams, authInfo runtime.ClientAuthInfoWriter) (*CreateExtensionCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateExtensionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "create-extension",
		Method:             "POST",
		PathPattern:        "/deployments/extensions",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateExtensionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateExtensionCreated), nil

}

/*
DeleteExtension deletes extension

Deletes a Extension.
*/
func (a *Client) DeleteExtension(params *DeleteExtensionParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteExtensionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteExtensionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "delete-extension",
		Method:             "DELETE",
		PathPattern:        "/deployments/extensions/{extension_id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteExtensionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteExtensionOK), nil

}

/*
GetExtension gets extension

Retrieves an Extension.
*/
func (a *Client) GetExtension(params *GetExtensionParams, authInfo runtime.ClientAuthInfoWriter) (*GetExtensionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetExtensionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get-extension",
		Method:             "GET",
		PathPattern:        "/deployments/extensions/{extension_id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetExtensionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetExtensionOK), nil

}

/*
ListExtensions lists extensions

Retrieves all of the available extensions.
*/
func (a *Client) ListExtensions(params *ListExtensionsParams, authInfo runtime.ClientAuthInfoWriter) (*ListExtensionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListExtensionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "list-extensions",
		Method:             "GET",
		PathPattern:        "/deployments/extensions",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListExtensionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListExtensionsOK), nil

}

/*
UpdateExtension updates extension

Updates a Extension.
*/
func (a *Client) UpdateExtension(params *UpdateExtensionParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateExtensionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateExtensionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "update-extension",
		Method:             "POST",
		PathPattern:        "/deployments/extensions/{extension_id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateExtensionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateExtensionOK), nil

}

/*
UploadExtension uploads the extension

Uploads archive for an Extension.
*/
func (a *Client) UploadExtension(params *UploadExtensionParams, authInfo runtime.ClientAuthInfoWriter) (*UploadExtensionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUploadExtensionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "upload-extension",
		Method:             "PUT",
		PathPattern:        "/deployments/extensions/{extension_id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UploadExtensionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UploadExtensionOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}