// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PageMetadata PageMetadata
// swagger:model PageMetadata
type PageMetadata struct {

	// number
	Number int64 `json:"number,omitempty" xml:"number,attr"`

	// size
	Size int64 `json:"size,omitempty" xml:"size,attr"`

	// total elements
	TotalElements int64 `json:"total_elements,omitempty" xml:"total_elements,attr"`

	// total pages
	TotalPages int64 `json:"total_pages,omitempty" xml:"total_pages,attr"`
}

// Validate validates this page metadata
func (m *PageMetadata) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PageMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PageMetadata) UnmarshalBinary(b []byte) error {
	var res PageMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
