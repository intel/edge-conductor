// Code generated by go-swagger; DO NOT EDIT.

//
//   Copyright (c) 2022 Intel Corporation.
//
//   SPDX-License-Identifier: Apache-2.0
//
//
//

package workflow

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Plugins plugins
//
// swagger:model plugins
type Plugins struct {

	// plugins
	Plugins []*PluginsItems0 `json:"plugins"`
}

// Validate validates this plugins
func (m *Plugins) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePlugins(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Plugins) validatePlugins(formats strfmt.Registry) error {
	if swag.IsZero(m.Plugins) { // not required
		return nil
	}

	for i := 0; i < len(m.Plugins); i++ {
		if swag.IsZero(m.Plugins[i]) { // not required
			continue
		}

		if m.Plugins[i] != nil {
			if err := m.Plugins[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("plugins" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("plugins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this plugins based on the context it is used
func (m *Plugins) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePlugins(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Plugins) contextValidatePlugins(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Plugins); i++ {

		if m.Plugins[i] != nil {
			if err := m.Plugins[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("plugins" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("plugins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Plugins) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Plugins) UnmarshalBinary(b []byte) error {
	var res Plugins
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PluginsItems0 plugins items0
//
// swagger:model PluginsItems0
type PluginsItems0 struct {

	// input
	Input []*PluginsItems0InputItems0 `json:"input"`

	// name
	// Pattern: ^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$
	Name string `json:"name,omitempty"`

	// output
	Output []*PluginsItems0OutputItems0 `json:"output"`
}

// Validate validates this plugins items0
func (m *PluginsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInput(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutput(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PluginsItems0) validateInput(formats strfmt.Registry) error {
	if swag.IsZero(m.Input) { // not required
		return nil
	}

	for i := 0; i < len(m.Input); i++ {
		if swag.IsZero(m.Input[i]) { // not required
			continue
		}

		if m.Input[i] != nil {
			if err := m.Input[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("input" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("input" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PluginsItems0) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", m.Name, `^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *PluginsItems0) validateOutput(formats strfmt.Registry) error {
	if swag.IsZero(m.Output) { // not required
		return nil
	}

	for i := 0; i < len(m.Output); i++ {
		if swag.IsZero(m.Output[i]) { // not required
			continue
		}

		if m.Output[i] != nil {
			if err := m.Output[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("output" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("output" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this plugins items0 based on the context it is used
func (m *PluginsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInput(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOutput(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PluginsItems0) contextValidateInput(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Input); i++ {

		if m.Input[i] != nil {
			if err := m.Input[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("input" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("input" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PluginsItems0) contextValidateOutput(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Output); i++ {

		if m.Output[i] != nil {
			if err := m.Output[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("output" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("output" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PluginsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PluginsItems0) UnmarshalBinary(b []byte) error {
	var res PluginsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PluginsItems0InputItems0 plugins items0 input items0
//
// swagger:model PluginsItems0InputItems0
type PluginsItems0InputItems0 struct {

	// description
	Description string `json:"description,omitempty"`

	// name
	// Pattern: ^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$
	Name string `json:"name,omitempty"`

	// schema
	// Pattern: ^[a-zA-Z.\/][a-zA-Z0-9-_.\/]*$
	Schema string `json:"schema,omitempty"`
}

// Validate validates this plugins items0 input items0
func (m *PluginsItems0InputItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchema(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PluginsItems0InputItems0) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", m.Name, `^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *PluginsItems0InputItems0) validateSchema(formats strfmt.Registry) error {
	if swag.IsZero(m.Schema) { // not required
		return nil
	}

	if err := validate.Pattern("schema", "body", m.Schema, `^[a-zA-Z.\/][a-zA-Z0-9-_.\/]*$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this plugins items0 input items0 based on context it is used
func (m *PluginsItems0InputItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PluginsItems0InputItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PluginsItems0InputItems0) UnmarshalBinary(b []byte) error {
	var res PluginsItems0InputItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PluginsItems0OutputItems0 plugins items0 output items0
//
// swagger:model PluginsItems0OutputItems0
type PluginsItems0OutputItems0 struct {

	// description
	Description string `json:"description,omitempty"`

	// name
	// Pattern: ^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$
	Name string `json:"name,omitempty"`

	// schema
	// Pattern: ^[a-zA-Z.\/][a-zA-Z0-9-_.\/]*$
	Schema string `json:"schema,omitempty"`
}

// Validate validates this plugins items0 output items0
func (m *PluginsItems0OutputItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchema(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PluginsItems0OutputItems0) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", m.Name, `^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *PluginsItems0OutputItems0) validateSchema(formats strfmt.Registry) error {
	if swag.IsZero(m.Schema) { // not required
		return nil
	}

	if err := validate.Pattern("schema", "body", m.Schema, `^[a-zA-Z.\/][a-zA-Z0-9-_.\/]*$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this plugins items0 output items0 based on context it is used
func (m *PluginsItems0OutputItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PluginsItems0OutputItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PluginsItems0OutputItems0) UnmarshalBinary(b []byte) error {
	var res PluginsItems0OutputItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
