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

// CreateKibanaRequest The request body for creating one or more Kibana instances.
// swagger:model CreateKibanaRequest
type CreateKibanaRequest struct {

	// The human readable name for the Kibana cluster (default: takes the name of its Elasticsearch cluster)
	ClusterName string `json:"cluster_name,omitempty"`

	// The Id of the Elasticsearch cluster to which this Kibana will be connected
	// Required: true
	ElasticsearchClusterID *string `json:"elasticsearch_cluster_id"`

	// plan
	// Required: true
	Plan *KibanaClusterPlan `json:"plan"`
}

// Validate validates this create kibana request
func (m *CreateKibanaRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateElasticsearchClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlan(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateKibanaRequest) validateElasticsearchClusterID(formats strfmt.Registry) error {

	if err := validate.Required("elasticsearch_cluster_id", "body", m.ElasticsearchClusterID); err != nil {
		return err
	}

	return nil
}

func (m *CreateKibanaRequest) validatePlan(formats strfmt.Registry) error {

	if err := validate.Required("plan", "body", m.Plan); err != nil {
		return err
	}

	if m.Plan != nil {
		if err := m.Plan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("plan")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateKibanaRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateKibanaRequest) UnmarshalBinary(b []byte) error {
	var res CreateKibanaRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}