// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Order order
// swagger:model Order
type Order struct {

	// order ID
	// Required: true
	OrderID *string `json:"orderID"`

	// order lines
	OrderLines []*OrderLine `json:"orderLines"`
}

// Validate validates this order
func (m *Order) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderLines(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Order) validateOrderID(formats strfmt.Registry) error {

	if err := validate.Required("orderID", "body", m.OrderID); err != nil {
		return err
	}

	return nil
}

func (m *Order) validateOrderLines(formats strfmt.Registry) error {

	if swag.IsZero(m.OrderLines) { // not required
		return nil
	}

	for i := 0; i < len(m.OrderLines); i++ {
		if swag.IsZero(m.OrderLines[i]) { // not required
			continue
		}

		if m.OrderLines[i] != nil {
			if err := m.OrderLines[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("orderLines" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Order) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Order) UnmarshalBinary(b []byte) error {
	var res Order
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OrderLine order line
// swagger:model orderLine
type OrderLine struct {

	// purchased item
	// Required: true
	PurchasedItem Item `json:"purchasedItem"`

	// quantity
	// Required: true
	// Minimum: 1
	Quantity *uint32 `json:"quantity"`
}

// Validate validates this order line
func (m *OrderLine) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePurchasedItem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuantity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderLine) validatePurchasedItem(formats strfmt.Registry) error {

	if err := m.PurchasedItem.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("purchasedItem")
		}
		return err
	}

	return nil
}

func (m *OrderLine) validateQuantity(formats strfmt.Registry) error {

	if err := validate.Required("quantity", "body", m.Quantity); err != nil {
		return err
	}

	if err := validate.Minimum("quantity", "body", float64(*m.Quantity), 1, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderLine) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderLine) UnmarshalBinary(b []byte) error {
	var res OrderLine
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
