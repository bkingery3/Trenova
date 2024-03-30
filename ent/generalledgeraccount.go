// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/emoss08/trenova/ent/businessunit"
	"github.com/emoss08/trenova/ent/generalledgeraccount"
	"github.com/emoss08/trenova/ent/organization"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// GeneralLedgerAccount is the model entity for the GeneralLedgerAccount schema.
type GeneralLedgerAccount struct {
	config `json:"-" validate:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// BusinessUnitID holds the value of the "business_unit_id" field.
	BusinessUnitID uuid.UUID `json:"businessUnitId"`
	// OrganizationID holds the value of the "organization_id" field.
	OrganizationID uuid.UUID `json:"organizationId"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updatedAt"`
	// Status holds the value of the "status" field.
	Status generalledgeraccount.Status `json:"status" validate:"required,oneof=A I"`
	// AccountNumber holds the value of the "account_number" field.
	AccountNumber string `json:"accountNumber" validate:"required,max=7"`
	// AccountType holds the value of the "account_type" field.
	AccountType generalledgeraccount.AccountType `json:"accountType" validate:"required"`
	// CashFlowType holds the value of the "cash_flow_type" field.
	CashFlowType string `json:"cashFlowType" validate:"omitempty"`
	// AccountSubType holds the value of the "account_sub_type" field.
	AccountSubType string `json:"accountSubType" validate:"omitempty"`
	// AccountClass holds the value of the "account_class" field.
	AccountClass string `json:"accountClass" validate:"omitempty"`
	// Balance holds the value of the "balance" field.
	Balance float64 `json:"balance" validate:"omitempty"`
	// InterestRate holds the value of the "interest_rate" field.
	InterestRate float64 `json:"interestRate" validate:"omitempty"`
	// DateOpened holds the value of the "date_opened" field.
	DateOpened *pgtype.Date `json:"dateOpened" validate:"omitempty"`
	// DateClosed holds the value of the "date_closed" field.
	DateClosed *pgtype.Date `json:"dateClosed" validate:"omitempty"`
	// Notes holds the value of the "notes" field.
	Notes string `json:"notes,omitempty"`
	// IsTaxRelevant holds the value of the "is_tax_relevant" field.
	IsTaxRelevant bool `json:"isTaxRelevant" validate:"omitempty"`
	// IsReconciled holds the value of the "is_reconciled" field.
	IsReconciled bool `json:"isReconciled" validate:"omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GeneralLedgerAccountQuery when eager-loading is set.
	Edges        GeneralLedgerAccountEdges `json:"edges"`
	selectValues sql.SelectValues
}

// GeneralLedgerAccountEdges holds the relations/edges for other nodes in the graph.
type GeneralLedgerAccountEdges struct {
	// BusinessUnit holds the value of the business_unit edge.
	BusinessUnit *BusinessUnit `json:"business_unit,omitempty"`
	// Organization holds the value of the organization edge.
	Organization *Organization `json:"organization,omitempty"`
	// Tags holds the value of the tags edge.
	Tags []*Tag `json:"tags,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// BusinessUnitOrErr returns the BusinessUnit value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GeneralLedgerAccountEdges) BusinessUnitOrErr() (*BusinessUnit, error) {
	if e.BusinessUnit != nil {
		return e.BusinessUnit, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: businessunit.Label}
	}
	return nil, &NotLoadedError{edge: "business_unit"}
}

// OrganizationOrErr returns the Organization value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GeneralLedgerAccountEdges) OrganizationOrErr() (*Organization, error) {
	if e.Organization != nil {
		return e.Organization, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: organization.Label}
	}
	return nil, &NotLoadedError{edge: "organization"}
}

// TagsOrErr returns the Tags value or an error if the edge
// was not loaded in eager-loading.
func (e GeneralLedgerAccountEdges) TagsOrErr() ([]*Tag, error) {
	if e.loadedTypes[2] {
		return e.Tags, nil
	}
	return nil, &NotLoadedError{edge: "tags"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GeneralLedgerAccount) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case generalledgeraccount.FieldDateClosed:
			values[i] = &sql.NullScanner{S: new(pgtype.Date)}
		case generalledgeraccount.FieldDateOpened:
			values[i] = new(pgtype.Date)
		case generalledgeraccount.FieldIsTaxRelevant, generalledgeraccount.FieldIsReconciled:
			values[i] = new(sql.NullBool)
		case generalledgeraccount.FieldBalance, generalledgeraccount.FieldInterestRate:
			values[i] = new(sql.NullFloat64)
		case generalledgeraccount.FieldStatus, generalledgeraccount.FieldAccountNumber, generalledgeraccount.FieldAccountType, generalledgeraccount.FieldCashFlowType, generalledgeraccount.FieldAccountSubType, generalledgeraccount.FieldAccountClass, generalledgeraccount.FieldNotes:
			values[i] = new(sql.NullString)
		case generalledgeraccount.FieldCreatedAt, generalledgeraccount.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case generalledgeraccount.FieldID, generalledgeraccount.FieldBusinessUnitID, generalledgeraccount.FieldOrganizationID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GeneralLedgerAccount fields.
func (gla *GeneralLedgerAccount) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case generalledgeraccount.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				gla.ID = *value
			}
		case generalledgeraccount.FieldBusinessUnitID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field business_unit_id", values[i])
			} else if value != nil {
				gla.BusinessUnitID = *value
			}
		case generalledgeraccount.FieldOrganizationID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field organization_id", values[i])
			} else if value != nil {
				gla.OrganizationID = *value
			}
		case generalledgeraccount.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				gla.CreatedAt = value.Time
			}
		case generalledgeraccount.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				gla.UpdatedAt = value.Time
			}
		case generalledgeraccount.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				gla.Status = generalledgeraccount.Status(value.String)
			}
		case generalledgeraccount.FieldAccountNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field account_number", values[i])
			} else if value.Valid {
				gla.AccountNumber = value.String
			}
		case generalledgeraccount.FieldAccountType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field account_type", values[i])
			} else if value.Valid {
				gla.AccountType = generalledgeraccount.AccountType(value.String)
			}
		case generalledgeraccount.FieldCashFlowType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cash_flow_type", values[i])
			} else if value.Valid {
				gla.CashFlowType = value.String
			}
		case generalledgeraccount.FieldAccountSubType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field account_sub_type", values[i])
			} else if value.Valid {
				gla.AccountSubType = value.String
			}
		case generalledgeraccount.FieldAccountClass:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field account_class", values[i])
			} else if value.Valid {
				gla.AccountClass = value.String
			}
		case generalledgeraccount.FieldBalance:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field balance", values[i])
			} else if value.Valid {
				gla.Balance = value.Float64
			}
		case generalledgeraccount.FieldInterestRate:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field interest_rate", values[i])
			} else if value.Valid {
				gla.InterestRate = value.Float64
			}
		case generalledgeraccount.FieldDateOpened:
			if value, ok := values[i].(*pgtype.Date); !ok {
				return fmt.Errorf("unexpected type %T for field date_opened", values[i])
			} else if value != nil {
				gla.DateOpened = value
			}
		case generalledgeraccount.FieldDateClosed:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field date_closed", values[i])
			} else if value.Valid {
				gla.DateClosed = value.S.(*pgtype.Date)
			}
		case generalledgeraccount.FieldNotes:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field notes", values[i])
			} else if value.Valid {
				gla.Notes = value.String
			}
		case generalledgeraccount.FieldIsTaxRelevant:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_tax_relevant", values[i])
			} else if value.Valid {
				gla.IsTaxRelevant = value.Bool
			}
		case generalledgeraccount.FieldIsReconciled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_reconciled", values[i])
			} else if value.Valid {
				gla.IsReconciled = value.Bool
			}
		default:
			gla.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the GeneralLedgerAccount.
// This includes values selected through modifiers, order, etc.
func (gla *GeneralLedgerAccount) Value(name string) (ent.Value, error) {
	return gla.selectValues.Get(name)
}

// QueryBusinessUnit queries the "business_unit" edge of the GeneralLedgerAccount entity.
func (gla *GeneralLedgerAccount) QueryBusinessUnit() *BusinessUnitQuery {
	return NewGeneralLedgerAccountClient(gla.config).QueryBusinessUnit(gla)
}

// QueryOrganization queries the "organization" edge of the GeneralLedgerAccount entity.
func (gla *GeneralLedgerAccount) QueryOrganization() *OrganizationQuery {
	return NewGeneralLedgerAccountClient(gla.config).QueryOrganization(gla)
}

// QueryTags queries the "tags" edge of the GeneralLedgerAccount entity.
func (gla *GeneralLedgerAccount) QueryTags() *TagQuery {
	return NewGeneralLedgerAccountClient(gla.config).QueryTags(gla)
}

// Update returns a builder for updating this GeneralLedgerAccount.
// Note that you need to call GeneralLedgerAccount.Unwrap() before calling this method if this GeneralLedgerAccount
// was returned from a transaction, and the transaction was committed or rolled back.
func (gla *GeneralLedgerAccount) Update() *GeneralLedgerAccountUpdateOne {
	return NewGeneralLedgerAccountClient(gla.config).UpdateOne(gla)
}

// Unwrap unwraps the GeneralLedgerAccount entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gla *GeneralLedgerAccount) Unwrap() *GeneralLedgerAccount {
	_tx, ok := gla.config.driver.(*txDriver)
	if !ok {
		panic("ent: GeneralLedgerAccount is not a transactional entity")
	}
	gla.config.driver = _tx.drv
	return gla
}

// String implements the fmt.Stringer.
func (gla *GeneralLedgerAccount) String() string {
	var builder strings.Builder
	builder.WriteString("GeneralLedgerAccount(")
	builder.WriteString(fmt.Sprintf("id=%v, ", gla.ID))
	builder.WriteString("business_unit_id=")
	builder.WriteString(fmt.Sprintf("%v", gla.BusinessUnitID))
	builder.WriteString(", ")
	builder.WriteString("organization_id=")
	builder.WriteString(fmt.Sprintf("%v", gla.OrganizationID))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(gla.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(gla.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", gla.Status))
	builder.WriteString(", ")
	builder.WriteString("account_number=")
	builder.WriteString(gla.AccountNumber)
	builder.WriteString(", ")
	builder.WriteString("account_type=")
	builder.WriteString(fmt.Sprintf("%v", gla.AccountType))
	builder.WriteString(", ")
	builder.WriteString("cash_flow_type=")
	builder.WriteString(gla.CashFlowType)
	builder.WriteString(", ")
	builder.WriteString("account_sub_type=")
	builder.WriteString(gla.AccountSubType)
	builder.WriteString(", ")
	builder.WriteString("account_class=")
	builder.WriteString(gla.AccountClass)
	builder.WriteString(", ")
	builder.WriteString("balance=")
	builder.WriteString(fmt.Sprintf("%v", gla.Balance))
	builder.WriteString(", ")
	builder.WriteString("interest_rate=")
	builder.WriteString(fmt.Sprintf("%v", gla.InterestRate))
	builder.WriteString(", ")
	builder.WriteString("date_opened=")
	builder.WriteString(fmt.Sprintf("%v", gla.DateOpened))
	builder.WriteString(", ")
	if v := gla.DateClosed; v != nil {
		builder.WriteString("date_closed=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("notes=")
	builder.WriteString(gla.Notes)
	builder.WriteString(", ")
	builder.WriteString("is_tax_relevant=")
	builder.WriteString(fmt.Sprintf("%v", gla.IsTaxRelevant))
	builder.WriteString(", ")
	builder.WriteString("is_reconciled=")
	builder.WriteString(fmt.Sprintf("%v", gla.IsReconciled))
	builder.WriteByte(')')
	return builder.String()
}

// GeneralLedgerAccounts is a parsable slice of GeneralLedgerAccount.
type GeneralLedgerAccounts []*GeneralLedgerAccount
