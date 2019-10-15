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

package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/elastic/cloud-sdk-go/pkg/models"
)

// ShutdownDeploymentEsResourceReader is a Reader for the ShutdownDeploymentEsResource structure.
type ShutdownDeploymentEsResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShutdownDeploymentEsResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewShutdownDeploymentEsResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewShutdownDeploymentEsResourceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewShutdownDeploymentEsResourceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewShutdownDeploymentEsResourceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewShutdownDeploymentEsResourceOK creates a ShutdownDeploymentEsResourceOK with default headers values
func NewShutdownDeploymentEsResourceOK() *ShutdownDeploymentEsResourceOK {
	return &ShutdownDeploymentEsResourceOK{}
}

/*ShutdownDeploymentEsResourceOK handles this case with default header values.

Standard response
*/
type ShutdownDeploymentEsResourceOK struct {
	Payload models.DeploymentResourceCommandResponse
}

func (o *ShutdownDeploymentEsResourceOK) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/elasticsearch/{ref_id}/_shutdown][%d] shutdownDeploymentEsResourceOK  %+v", 200, o.Payload)
}

func (o *ShutdownDeploymentEsResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShutdownDeploymentEsResourceUnauthorized creates a ShutdownDeploymentEsResourceUnauthorized with default headers values
func NewShutdownDeploymentEsResourceUnauthorized() *ShutdownDeploymentEsResourceUnauthorized {
	return &ShutdownDeploymentEsResourceUnauthorized{}
}

/*ShutdownDeploymentEsResourceUnauthorized handles this case with default header values.

You are not authorized to perform this action
*/
type ShutdownDeploymentEsResourceUnauthorized struct {
	Payload *models.BasicFailedReply
}

func (o *ShutdownDeploymentEsResourceUnauthorized) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/elasticsearch/{ref_id}/_shutdown][%d] shutdownDeploymentEsResourceUnauthorized  %+v", 401, o.Payload)
}

func (o *ShutdownDeploymentEsResourceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShutdownDeploymentEsResourceNotFound creates a ShutdownDeploymentEsResourceNotFound with default headers values
func NewShutdownDeploymentEsResourceNotFound() *ShutdownDeploymentEsResourceNotFound {
	return &ShutdownDeploymentEsResourceNotFound{}
}

/*ShutdownDeploymentEsResourceNotFound handles this case with default header values.

The Deployment specified by {deployment_id} cannot be found
*/
type ShutdownDeploymentEsResourceNotFound struct {
	Payload *models.BasicFailedReply
}

func (o *ShutdownDeploymentEsResourceNotFound) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/elasticsearch/{ref_id}/_shutdown][%d] shutdownDeploymentEsResourceNotFound  %+v", 404, o.Payload)
}

func (o *ShutdownDeploymentEsResourceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShutdownDeploymentEsResourceInternalServerError creates a ShutdownDeploymentEsResourceInternalServerError with default headers values
func NewShutdownDeploymentEsResourceInternalServerError() *ShutdownDeploymentEsResourceInternalServerError {
	return &ShutdownDeploymentEsResourceInternalServerError{}
}

/*ShutdownDeploymentEsResourceInternalServerError handles this case with default header values.

We have failed you.
*/
type ShutdownDeploymentEsResourceInternalServerError struct {
	Payload *models.BasicFailedReply
}

func (o *ShutdownDeploymentEsResourceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/elasticsearch/{ref_id}/_shutdown][%d] shutdownDeploymentEsResourceInternalServerError  %+v", 500, o.Payload)
}

func (o *ShutdownDeploymentEsResourceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}