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

package platform_configuration_security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/elastic/cloud-sdk-go/pkg/models"
)

// SetTLSCertificateReader is a Reader for the SetTLSCertificate structure.
type SetTLSCertificateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetTLSCertificateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewSetTLSCertificateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 449:
		result := NewSetTLSCertificateRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetTLSCertificateAccepted creates a SetTLSCertificateAccepted with default headers values
func NewSetTLSCertificateAccepted() *SetTLSCertificateAccepted {
	return &SetTLSCertificateAccepted{}
}

/*SetTLSCertificateAccepted handles this case with default header values.

The TLS update has been accepted for the given service and will take effect throughout the system
*/
type SetTLSCertificateAccepted struct {
	Payload models.EmptyResponse
}

func (o *SetTLSCertificateAccepted) Error() string {
	return fmt.Sprintf("[POST /platform/configuration/security/tls/{service_name}][%d] setTlsCertificateAccepted  %+v", 202, o.Payload)
}

func (o *SetTLSCertificateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetTLSCertificateRetryWith creates a SetTLSCertificateRetryWith with default headers values
func NewSetTLSCertificateRetryWith() *SetTLSCertificateRetryWith {
	return &SetTLSCertificateRetryWith{}
}

/*SetTLSCertificateRetryWith handles this case with default header values.

elevated permissions are required. (code: '"root.unauthorized.rbac.elevated_permissions_required"')
*/
type SetTLSCertificateRetryWith struct {
	Payload *models.BasicFailedReply
}

func (o *SetTLSCertificateRetryWith) Error() string {
	return fmt.Sprintf("[POST /platform/configuration/security/tls/{service_name}][%d] setTlsCertificateRetryWith  %+v", 449, o.Payload)
}

func (o *SetTLSCertificateRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}