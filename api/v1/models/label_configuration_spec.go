// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2021 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LabelConfigurationSpec User desired Label configuration of an endpoint
//
// swagger:model LabelConfigurationSpec
type LabelConfigurationSpec struct {

	// Custom labels in addition to orchestration system labels.
	User Labels `json:"user,omitempty"`
}

// Validate validates this label configuration spec
func (m *LabelConfigurationSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LabelConfigurationSpec) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(m.User) { // not required
		return nil
	}

	if err := m.User.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("user")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LabelConfigurationSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LabelConfigurationSpec) UnmarshalBinary(b []byte) error {
	var res LabelConfigurationSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
