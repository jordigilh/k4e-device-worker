// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MetricsConfiguration Defines metrics configuration for the device
//
// swagger:model metrics-configuration
type MetricsConfiguration struct {

	// Defines metrics data retention limits
	Retention *MetricsRetention `json:"retention,omitempty"`
}

// Validate validates this metrics configuration
func (m *MetricsConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRetention(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetricsConfiguration) validateRetention(formats strfmt.Registry) error {

	if swag.IsZero(m.Retention) { // not required
		return nil
	}

	if m.Retention != nil {
		if err := m.Retention.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("retention")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetricsConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetricsConfiguration) UnmarshalBinary(b []byte) error {
	var res MetricsConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
