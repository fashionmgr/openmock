// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ActionPublishKafka publish a message on kafka
// swagger:model ActionPublishKafka
type ActionPublishKafka struct {

	// string payload to send on AMQP
	Payload string `json:"payload,omitempty"`

	// file name (relative to working directory of OM) to load HTTP body from
	PayloadFromFile string `json:"payload_from_file,omitempty"`

	// kafka topic to publish on
	// Required: true
	Topic *string `json:"topic"`
}

// Validate validates this action publish kafka
func (m *ActionPublishKafka) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTopic(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActionPublishKafka) validateTopic(formats strfmt.Registry) error {

	if err := validate.Required("topic", "body", m.Topic); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ActionPublishKafka) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActionPublishKafka) UnmarshalBinary(b []byte) error {
	var res ActionPublishKafka
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
