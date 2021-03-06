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

package clusters_apm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// SetApmNameReader is a Reader for the SetApmName structure.
type SetApmNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetApmNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetApmNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewSetApmNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewSetApmNameRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetApmNameOK creates a SetApmNameOK with default headers values
func NewSetApmNameOK() *SetApmNameOK {
	return &SetApmNameOK{}
}

/*SetApmNameOK handles this case with default header values.

The cluster name was successfully changed
*/
type SetApmNameOK struct {
	Payload models.EmptyResponse
}

func (o *SetApmNameOK) Error() string {
	return fmt.Sprintf("[PUT /clusters/apm/{cluster_id}/metadata/name/{new_name}][%d] setApmNameOK  %+v", 200, o.Payload)
}

func (o *SetApmNameOK) GetPayload() models.EmptyResponse {
	return o.Payload
}

func (o *SetApmNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetApmNameNotFound creates a SetApmNameNotFound with default headers values
func NewSetApmNameNotFound() *SetApmNameNotFound {
	return &SetApmNameNotFound{}
}

/*SetApmNameNotFound handles this case with default header values.

The cluster specified by {cluster_id} cannot be found (code: 'clusters.cluster_not_found')
*/
type SetApmNameNotFound struct {
	Payload *models.BasicFailedReply
}

func (o *SetApmNameNotFound) Error() string {
	return fmt.Sprintf("[PUT /clusters/apm/{cluster_id}/metadata/name/{new_name}][%d] setApmNameNotFound  %+v", 404, o.Payload)
}

func (o *SetApmNameNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *SetApmNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetApmNameRetryWith creates a SetApmNameRetryWith with default headers values
func NewSetApmNameRetryWith() *SetApmNameRetryWith {
	return &SetApmNameRetryWith{}
}

/*SetApmNameRetryWith handles this case with default header values.

Elevated permissions are required. (code: '"root.unauthorized.rbac.elevated_permissions_required"')
*/
type SetApmNameRetryWith struct {
	Payload *models.BasicFailedReply
}

func (o *SetApmNameRetryWith) Error() string {
	return fmt.Sprintf("[PUT /clusters/apm/{cluster_id}/metadata/name/{new_name}][%d] setApmNameRetryWith  %+v", 449, o.Payload)
}

func (o *SetApmNameRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *SetApmNameRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
