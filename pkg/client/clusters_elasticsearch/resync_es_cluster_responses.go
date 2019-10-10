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

package clusters_elasticsearch

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/elastic/cloud-sdk-go/pkg/models"
)

// ResyncEsClusterReader is a Reader for the ResyncEsCluster structure.
type ResyncEsClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResyncEsClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewResyncEsClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 449:
		result := NewResyncEsClusterRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewResyncEsClusterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewResyncEsClusterOK creates a ResyncEsClusterOK with default headers values
func NewResyncEsClusterOK() *ResyncEsClusterOK {
	return &ResyncEsClusterOK{}
}

/*ResyncEsClusterOK handles this case with default header values.

The cluster resync operation executed successfully
*/
type ResyncEsClusterOK struct {
	Payload models.EmptyResponse
}

func (o *ResyncEsClusterOK) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/_resync][%d] resyncEsClusterOK  %+v", 200, o.Payload)
}

func (o *ResyncEsClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResyncEsClusterRetryWith creates a ResyncEsClusterRetryWith with default headers values
func NewResyncEsClusterRetryWith() *ResyncEsClusterRetryWith {
	return &ResyncEsClusterRetryWith{}
}

/*ResyncEsClusterRetryWith handles this case with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type ResyncEsClusterRetryWith struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *ResyncEsClusterRetryWith) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/_resync][%d] resyncEsClusterRetryWith  %+v", 449, o.Payload)
}

func (o *ResyncEsClusterRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResyncEsClusterInternalServerError creates a ResyncEsClusterInternalServerError with default headers values
func NewResyncEsClusterInternalServerError() *ResyncEsClusterInternalServerError {
	return &ResyncEsClusterInternalServerError{}
}

/*ResyncEsClusterInternalServerError handles this case with default header values.

The cluster resync operation failed for cluster {cluster_id}. (code: `clusters.resync_failed`)
*/
type ResyncEsClusterInternalServerError struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *ResyncEsClusterInternalServerError) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/_resync][%d] resyncEsClusterInternalServerError  %+v", 500, o.Payload)
}

func (o *ResyncEsClusterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}