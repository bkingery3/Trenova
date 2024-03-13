// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/emoss08/trenova/ent/businessunit"
	"github.com/emoss08/trenova/ent/feasibilitytoolcontrol"
	"github.com/emoss08/trenova/ent/organization"
	"github.com/google/uuid"
)

// FeasibilityToolControl is the model entity for the FeasibilityToolControl schema.
type FeasibilityToolControl struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// OtpOperator holds the value of the "otp_operator" field.
	OtpOperator feasibilitytoolcontrol.OtpOperator `json:"otpOperator"`
	// OtpValue holds the value of the "otp_value" field.
	OtpValue float64 `json:"otpValue"`
	// MpwOperator holds the value of the "mpw_operator" field.
	MpwOperator feasibilitytoolcontrol.MpwOperator `json:"mpwOperator"`
	// MpwValue holds the value of the "mpw_value" field.
	MpwValue float64 `json:"mpwValue"`
	// MpdOperator holds the value of the "mpd_operator" field.
	MpdOperator feasibilitytoolcontrol.MpdOperator `json:"mpdOperator"`
	// MpdValue holds the value of the "mpd_value" field.
	MpdValue float64 `json:"mpdValue"`
	// MpgOperator holds the value of the "mpg_operator" field.
	MpgOperator feasibilitytoolcontrol.MpgOperator `json:"mpgOperator"`
	// MpgValue holds the value of the "mpg_value" field.
	MpgValue float64 `json:"mpgValue"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FeasibilityToolControlQuery when eager-loading is set.
	Edges            FeasibilityToolControlEdges `json:"edges"`
	business_unit_id *uuid.UUID
	organization_id  *uuid.UUID
	selectValues     sql.SelectValues
}

// FeasibilityToolControlEdges holds the relations/edges for other nodes in the graph.
type FeasibilityToolControlEdges struct {
	// Organization holds the value of the organization edge.
	Organization *Organization `json:"organization,omitempty"`
	// BusinessUnit holds the value of the business_unit edge.
	BusinessUnit *BusinessUnit `json:"business_unit,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// OrganizationOrErr returns the Organization value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FeasibilityToolControlEdges) OrganizationOrErr() (*Organization, error) {
	if e.Organization != nil {
		return e.Organization, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: organization.Label}
	}
	return nil, &NotLoadedError{edge: "organization"}
}

// BusinessUnitOrErr returns the BusinessUnit value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FeasibilityToolControlEdges) BusinessUnitOrErr() (*BusinessUnit, error) {
	if e.BusinessUnit != nil {
		return e.BusinessUnit, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: businessunit.Label}
	}
	return nil, &NotLoadedError{edge: "business_unit"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FeasibilityToolControl) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case feasibilitytoolcontrol.FieldOtpValue, feasibilitytoolcontrol.FieldMpwValue, feasibilitytoolcontrol.FieldMpdValue, feasibilitytoolcontrol.FieldMpgValue:
			values[i] = new(sql.NullFloat64)
		case feasibilitytoolcontrol.FieldOtpOperator, feasibilitytoolcontrol.FieldMpwOperator, feasibilitytoolcontrol.FieldMpdOperator, feasibilitytoolcontrol.FieldMpgOperator:
			values[i] = new(sql.NullString)
		case feasibilitytoolcontrol.FieldCreatedAt, feasibilitytoolcontrol.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case feasibilitytoolcontrol.FieldID:
			values[i] = new(uuid.UUID)
		case feasibilitytoolcontrol.ForeignKeys[0]: // business_unit_id
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case feasibilitytoolcontrol.ForeignKeys[1]: // organization_id
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FeasibilityToolControl fields.
func (ftc *FeasibilityToolControl) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case feasibilitytoolcontrol.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ftc.ID = *value
			}
		case feasibilitytoolcontrol.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ftc.CreatedAt = value.Time
			}
		case feasibilitytoolcontrol.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ftc.UpdatedAt = value.Time
			}
		case feasibilitytoolcontrol.FieldOtpOperator:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field otp_operator", values[i])
			} else if value.Valid {
				ftc.OtpOperator = feasibilitytoolcontrol.OtpOperator(value.String)
			}
		case feasibilitytoolcontrol.FieldOtpValue:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field otp_value", values[i])
			} else if value.Valid {
				ftc.OtpValue = value.Float64
			}
		case feasibilitytoolcontrol.FieldMpwOperator:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mpw_operator", values[i])
			} else if value.Valid {
				ftc.MpwOperator = feasibilitytoolcontrol.MpwOperator(value.String)
			}
		case feasibilitytoolcontrol.FieldMpwValue:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field mpw_value", values[i])
			} else if value.Valid {
				ftc.MpwValue = value.Float64
			}
		case feasibilitytoolcontrol.FieldMpdOperator:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mpd_operator", values[i])
			} else if value.Valid {
				ftc.MpdOperator = feasibilitytoolcontrol.MpdOperator(value.String)
			}
		case feasibilitytoolcontrol.FieldMpdValue:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field mpd_value", values[i])
			} else if value.Valid {
				ftc.MpdValue = value.Float64
			}
		case feasibilitytoolcontrol.FieldMpgOperator:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mpg_operator", values[i])
			} else if value.Valid {
				ftc.MpgOperator = feasibilitytoolcontrol.MpgOperator(value.String)
			}
		case feasibilitytoolcontrol.FieldMpgValue:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field mpg_value", values[i])
			} else if value.Valid {
				ftc.MpgValue = value.Float64
			}
		case feasibilitytoolcontrol.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field business_unit_id", values[i])
			} else if value.Valid {
				ftc.business_unit_id = new(uuid.UUID)
				*ftc.business_unit_id = *value.S.(*uuid.UUID)
			}
		case feasibilitytoolcontrol.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field organization_id", values[i])
			} else if value.Valid {
				ftc.organization_id = new(uuid.UUID)
				*ftc.organization_id = *value.S.(*uuid.UUID)
			}
		default:
			ftc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the FeasibilityToolControl.
// This includes values selected through modifiers, order, etc.
func (ftc *FeasibilityToolControl) Value(name string) (ent.Value, error) {
	return ftc.selectValues.Get(name)
}

// QueryOrganization queries the "organization" edge of the FeasibilityToolControl entity.
func (ftc *FeasibilityToolControl) QueryOrganization() *OrganizationQuery {
	return NewFeasibilityToolControlClient(ftc.config).QueryOrganization(ftc)
}

// QueryBusinessUnit queries the "business_unit" edge of the FeasibilityToolControl entity.
func (ftc *FeasibilityToolControl) QueryBusinessUnit() *BusinessUnitQuery {
	return NewFeasibilityToolControlClient(ftc.config).QueryBusinessUnit(ftc)
}

// Update returns a builder for updating this FeasibilityToolControl.
// Note that you need to call FeasibilityToolControl.Unwrap() before calling this method if this FeasibilityToolControl
// was returned from a transaction, and the transaction was committed or rolled back.
func (ftc *FeasibilityToolControl) Update() *FeasibilityToolControlUpdateOne {
	return NewFeasibilityToolControlClient(ftc.config).UpdateOne(ftc)
}

// Unwrap unwraps the FeasibilityToolControl entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ftc *FeasibilityToolControl) Unwrap() *FeasibilityToolControl {
	_tx, ok := ftc.config.driver.(*txDriver)
	if !ok {
		panic("ent: FeasibilityToolControl is not a transactional entity")
	}
	ftc.config.driver = _tx.drv
	return ftc
}

// String implements the fmt.Stringer.
func (ftc *FeasibilityToolControl) String() string {
	var builder strings.Builder
	builder.WriteString("FeasibilityToolControl(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ftc.ID))
	builder.WriteString("created_at=")
	builder.WriteString(ftc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ftc.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("otp_operator=")
	builder.WriteString(fmt.Sprintf("%v", ftc.OtpOperator))
	builder.WriteString(", ")
	builder.WriteString("otp_value=")
	builder.WriteString(fmt.Sprintf("%v", ftc.OtpValue))
	builder.WriteString(", ")
	builder.WriteString("mpw_operator=")
	builder.WriteString(fmt.Sprintf("%v", ftc.MpwOperator))
	builder.WriteString(", ")
	builder.WriteString("mpw_value=")
	builder.WriteString(fmt.Sprintf("%v", ftc.MpwValue))
	builder.WriteString(", ")
	builder.WriteString("mpd_operator=")
	builder.WriteString(fmt.Sprintf("%v", ftc.MpdOperator))
	builder.WriteString(", ")
	builder.WriteString("mpd_value=")
	builder.WriteString(fmt.Sprintf("%v", ftc.MpdValue))
	builder.WriteString(", ")
	builder.WriteString("mpg_operator=")
	builder.WriteString(fmt.Sprintf("%v", ftc.MpgOperator))
	builder.WriteString(", ")
	builder.WriteString("mpg_value=")
	builder.WriteString(fmt.Sprintf("%v", ftc.MpgValue))
	builder.WriteByte(')')
	return builder.String()
}

// FeasibilityToolControls is a parsable slice of FeasibilityToolControl.
type FeasibilityToolControls []*FeasibilityToolControl
