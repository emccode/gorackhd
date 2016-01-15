package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
)

/*catalog Catalog catalog

swagger:model catalog
*/
type Catalog struct {

	/* arbitrary json object
	 */
	Data interface{} `json:"data,omitempty"`

	/* ID id

	Required: true
	*/
	ID string `json:"id,omitempty"`

	/* Node node

	Required: true
	*/
	Node *Node `json:"node,omitempty"`

	/* Source source

	Required: true
	*/
	Source string `json:"source,omitempty"`
}

// Validate validates this catalog
func (m *Catalog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Catalog) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Catalog) validateNode(formats strfmt.Registry) error {

	if m.Node != nil {

		if err := m.Node.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *Catalog) validateSource(formats strfmt.Registry) error {

	if err := validate.Required("source", "body", string(m.Source)); err != nil {
		return err
	}

	return nil
}
