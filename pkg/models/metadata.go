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
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Metadata Metadata of an entity
// swagger:model Metadata
type Metadata struct {

	// Creation time
	// Required: true
	// Format: date-time
	CreatedTime *strfmt.DateTime `json:"created_time"`

	// Modification time
	// Required: true
	// Format: date-time
	ModifiedTime *strfmt.DateTime `json:"modified_time"`

	// Version
	// Required: true
	Version *int32 `json:"version"`
}

// Validate validates this metadata
func (m *Metadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedTime(formats); err != nil {
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

func (m *Metadata) validateCreatedTime(formats strfmt.Registry) error {

	if err := validate.Required("created_time", "body", m.CreatedTime); err != nil {
		return err
	}

	if err := validate.FormatOf("created_time", "body", "date-time", m.CreatedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Metadata) validateModifiedTime(formats strfmt.Registry) error {

	if err := validate.Required("modified_time", "body", m.ModifiedTime); err != nil {
		return err
	}

	if err := validate.FormatOf("modified_time", "body", "date-time", m.ModifiedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Metadata) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Metadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Metadata) UnmarshalBinary(b []byte) error {
	var res Metadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
