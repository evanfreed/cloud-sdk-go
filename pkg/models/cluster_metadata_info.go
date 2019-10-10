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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ClusterMetadataInfo Information about the public and internal state, and the configuration settings of an Elasticsearch cluster.
// swagger:model ClusterMetadataInfo
type ClusterMetadataInfo struct {

	// The cloud ID, an encoded string that provides other Elastic services with the necessary information to connect to this Elasticsearch and Kibana (only present if both exist)
	CloudID string `json:"cloud_id,omitempty"`

	// The DNS name of the cluster endpoint, if available
	Endpoint string `json:"endpoint,omitempty"`

	// The most recent time the cluster metadata was changed (ISO format in UTC)
	// Required: true
	// Format: date-time
	LastModified *strfmt.DateTime `json:"last_modified"`

	// The ports that allow communication with the cluster using various protocols.
	Ports *ClusterMetadataPortInfo `json:"ports,omitempty"`

	// An unstructured JSON representation of the public and internal state (can be filtered out via URL parameter). The contents and structure of the `raw` field can change at any time.
	Raw interface{} `json:"raw,omitempty"`

	// The resource version number of the cluster metadata
	// Required: true
	Version *int32 `json:"version"`
}

// Validate validates this cluster metadata info
func (m *ClusterMetadataInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePorts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterMetadataInfo) validateLastModified(formats strfmt.Registry) error {

	if err := validate.Required("last_modified", "body", m.LastModified); err != nil {
		return err
	}

	if err := validate.FormatOf("last_modified", "body", "date-time", m.LastModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ClusterMetadataInfo) validatePorts(formats strfmt.Registry) error {

	if swag.IsZero(m.Ports) { // not required
		return nil
	}

	if m.Ports != nil {
		if err := m.Ports.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ports")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterMetadataInfo) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterMetadataInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterMetadataInfo) UnmarshalBinary(b []byte) error {
	var res ClusterMetadataInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}