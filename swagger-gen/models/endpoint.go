// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Endpoint Endpoint
//
// The network context of a node in the service graph
// swagger:model Endpoint

type Endpoint struct {

	// The text representation of the primary IPv4 address associated with this
	// a connection. Ex. 192.168.99.100 Absent if unknown.
	//
	IPV4 strfmt.IPv4 `json:"ipv4,omitempty"`

	// The text representation of the primary IPv6 address associated with this
	// a connection. Ex. 2001:db8::c001 Absent if unknown.
	//
	// Prefer using the ipv4 field for mapped addresses.
	//
	IPV6 strfmt.IPv6 `json:"ipv6,omitempty"`

	// Depending on context, this could be a listen port or the client-side of a
	// socket. Absent if unknown
	//
	Port int64 `json:"port,omitempty"`

	// Lower-case label of this node in the service graph, such as "favstar". Leave
	// absent if unknown.
	//
	// This is a primary label for trace lookup and aggregation, so it should be
	// intuitive and consistent. Many use a name from service discovery.
	//
	ServiceName string `json:"serviceName,omitempty"`
}

/* polymorph Endpoint ipv4 false */

/* polymorph Endpoint ipv6 false */

/* polymorph Endpoint port false */

/* polymorph Endpoint serviceName false */

// Validate validates this endpoint
func (m *Endpoint) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Endpoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Endpoint) UnmarshalBinary(b []byte) error {
	var res Endpoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}