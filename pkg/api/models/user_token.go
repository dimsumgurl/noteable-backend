// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserToken user token
//
// swagger:model userToken
type UserToken struct {

	// token
	Token string `json:"token,omitempty"`
}

// Validate validates this user token
func (m *UserToken) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user token based on context it is used
func (m *UserToken) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserToken) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserToken) UnmarshalBinary(b []byte) error {
	var res UserToken
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
