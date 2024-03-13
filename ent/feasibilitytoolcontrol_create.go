// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/ent/businessunit"
	"github.com/emoss08/trenova/ent/feasibilitytoolcontrol"
	"github.com/emoss08/trenova/ent/organization"
	"github.com/google/uuid"
)

// FeasibilityToolControlCreate is the builder for creating a FeasibilityToolControl entity.
type FeasibilityToolControlCreate struct {
	config
	mutation *FeasibilityToolControlMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ftcc *FeasibilityToolControlCreate) SetCreatedAt(t time.Time) *FeasibilityToolControlCreate {
	ftcc.mutation.SetCreatedAt(t)
	return ftcc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ftcc *FeasibilityToolControlCreate) SetNillableCreatedAt(t *time.Time) *FeasibilityToolControlCreate {
	if t != nil {
		ftcc.SetCreatedAt(*t)
	}
	return ftcc
}

// SetUpdatedAt sets the "updated_at" field.
func (ftcc *FeasibilityToolControlCreate) SetUpdatedAt(t time.Time) *FeasibilityToolControlCreate {
	ftcc.mutation.SetUpdatedAt(t)
	return ftcc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ftcc *FeasibilityToolControlCreate) SetNillableUpdatedAt(t *time.Time) *FeasibilityToolControlCreate {
	if t != nil {
		ftcc.SetUpdatedAt(*t)
	}
	return ftcc
}

// SetOtpOperator sets the "otp_operator" field.
func (ftcc *FeasibilityToolControlCreate) SetOtpOperator(fo feasibilitytoolcontrol.OtpOperator) *FeasibilityToolControlCreate {
	ftcc.mutation.SetOtpOperator(fo)
	return ftcc
}

// SetNillableOtpOperator sets the "otp_operator" field if the given value is not nil.
func (ftcc *FeasibilityToolControlCreate) SetNillableOtpOperator(fo *feasibilitytoolcontrol.OtpOperator) *FeasibilityToolControlCreate {
	if fo != nil {
		ftcc.SetOtpOperator(*fo)
	}
	return ftcc
}

// SetOtpValue sets the "otp_value" field.
func (ftcc *FeasibilityToolControlCreate) SetOtpValue(f float64) *FeasibilityToolControlCreate {
	ftcc.mutation.SetOtpValue(f)
	return ftcc
}

// SetNillableOtpValue sets the "otp_value" field if the given value is not nil.
func (ftcc *FeasibilityToolControlCreate) SetNillableOtpValue(f *float64) *FeasibilityToolControlCreate {
	if f != nil {
		ftcc.SetOtpValue(*f)
	}
	return ftcc
}

// SetMpwOperator sets the "mpw_operator" field.
func (ftcc *FeasibilityToolControlCreate) SetMpwOperator(fo feasibilitytoolcontrol.MpwOperator) *FeasibilityToolControlCreate {
	ftcc.mutation.SetMpwOperator(fo)
	return ftcc
}

// SetNillableMpwOperator sets the "mpw_operator" field if the given value is not nil.
func (ftcc *FeasibilityToolControlCreate) SetNillableMpwOperator(fo *feasibilitytoolcontrol.MpwOperator) *FeasibilityToolControlCreate {
	if fo != nil {
		ftcc.SetMpwOperator(*fo)
	}
	return ftcc
}

// SetMpwValue sets the "mpw_value" field.
func (ftcc *FeasibilityToolControlCreate) SetMpwValue(f float64) *FeasibilityToolControlCreate {
	ftcc.mutation.SetMpwValue(f)
	return ftcc
}

// SetNillableMpwValue sets the "mpw_value" field if the given value is not nil.
func (ftcc *FeasibilityToolControlCreate) SetNillableMpwValue(f *float64) *FeasibilityToolControlCreate {
	if f != nil {
		ftcc.SetMpwValue(*f)
	}
	return ftcc
}

// SetMpdOperator sets the "mpd_operator" field.
func (ftcc *FeasibilityToolControlCreate) SetMpdOperator(fo feasibilitytoolcontrol.MpdOperator) *FeasibilityToolControlCreate {
	ftcc.mutation.SetMpdOperator(fo)
	return ftcc
}

// SetNillableMpdOperator sets the "mpd_operator" field if the given value is not nil.
func (ftcc *FeasibilityToolControlCreate) SetNillableMpdOperator(fo *feasibilitytoolcontrol.MpdOperator) *FeasibilityToolControlCreate {
	if fo != nil {
		ftcc.SetMpdOperator(*fo)
	}
	return ftcc
}

// SetMpdValue sets the "mpd_value" field.
func (ftcc *FeasibilityToolControlCreate) SetMpdValue(f float64) *FeasibilityToolControlCreate {
	ftcc.mutation.SetMpdValue(f)
	return ftcc
}

// SetNillableMpdValue sets the "mpd_value" field if the given value is not nil.
func (ftcc *FeasibilityToolControlCreate) SetNillableMpdValue(f *float64) *FeasibilityToolControlCreate {
	if f != nil {
		ftcc.SetMpdValue(*f)
	}
	return ftcc
}

// SetMpgOperator sets the "mpg_operator" field.
func (ftcc *FeasibilityToolControlCreate) SetMpgOperator(fo feasibilitytoolcontrol.MpgOperator) *FeasibilityToolControlCreate {
	ftcc.mutation.SetMpgOperator(fo)
	return ftcc
}

// SetNillableMpgOperator sets the "mpg_operator" field if the given value is not nil.
func (ftcc *FeasibilityToolControlCreate) SetNillableMpgOperator(fo *feasibilitytoolcontrol.MpgOperator) *FeasibilityToolControlCreate {
	if fo != nil {
		ftcc.SetMpgOperator(*fo)
	}
	return ftcc
}

// SetMpgValue sets the "mpg_value" field.
func (ftcc *FeasibilityToolControlCreate) SetMpgValue(f float64) *FeasibilityToolControlCreate {
	ftcc.mutation.SetMpgValue(f)
	return ftcc
}

// SetNillableMpgValue sets the "mpg_value" field if the given value is not nil.
func (ftcc *FeasibilityToolControlCreate) SetNillableMpgValue(f *float64) *FeasibilityToolControlCreate {
	if f != nil {
		ftcc.SetMpgValue(*f)
	}
	return ftcc
}

// SetID sets the "id" field.
func (ftcc *FeasibilityToolControlCreate) SetID(u uuid.UUID) *FeasibilityToolControlCreate {
	ftcc.mutation.SetID(u)
	return ftcc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ftcc *FeasibilityToolControlCreate) SetNillableID(u *uuid.UUID) *FeasibilityToolControlCreate {
	if u != nil {
		ftcc.SetID(*u)
	}
	return ftcc
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (ftcc *FeasibilityToolControlCreate) SetOrganizationID(id uuid.UUID) *FeasibilityToolControlCreate {
	ftcc.mutation.SetOrganizationID(id)
	return ftcc
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (ftcc *FeasibilityToolControlCreate) SetOrganization(o *Organization) *FeasibilityToolControlCreate {
	return ftcc.SetOrganizationID(o.ID)
}

// SetBusinessUnitID sets the "business_unit" edge to the BusinessUnit entity by ID.
func (ftcc *FeasibilityToolControlCreate) SetBusinessUnitID(id uuid.UUID) *FeasibilityToolControlCreate {
	ftcc.mutation.SetBusinessUnitID(id)
	return ftcc
}

// SetBusinessUnit sets the "business_unit" edge to the BusinessUnit entity.
func (ftcc *FeasibilityToolControlCreate) SetBusinessUnit(b *BusinessUnit) *FeasibilityToolControlCreate {
	return ftcc.SetBusinessUnitID(b.ID)
}

// Mutation returns the FeasibilityToolControlMutation object of the builder.
func (ftcc *FeasibilityToolControlCreate) Mutation() *FeasibilityToolControlMutation {
	return ftcc.mutation
}

// Save creates the FeasibilityToolControl in the database.
func (ftcc *FeasibilityToolControlCreate) Save(ctx context.Context) (*FeasibilityToolControl, error) {
	ftcc.defaults()
	return withHooks(ctx, ftcc.sqlSave, ftcc.mutation, ftcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ftcc *FeasibilityToolControlCreate) SaveX(ctx context.Context) *FeasibilityToolControl {
	v, err := ftcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ftcc *FeasibilityToolControlCreate) Exec(ctx context.Context) error {
	_, err := ftcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftcc *FeasibilityToolControlCreate) ExecX(ctx context.Context) {
	if err := ftcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ftcc *FeasibilityToolControlCreate) defaults() {
	if _, ok := ftcc.mutation.CreatedAt(); !ok {
		v := feasibilitytoolcontrol.DefaultCreatedAt
		ftcc.mutation.SetCreatedAt(v)
	}
	if _, ok := ftcc.mutation.UpdatedAt(); !ok {
		v := feasibilitytoolcontrol.DefaultUpdatedAt
		ftcc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ftcc.mutation.OtpOperator(); !ok {
		v := feasibilitytoolcontrol.DefaultOtpOperator
		ftcc.mutation.SetOtpOperator(v)
	}
	if _, ok := ftcc.mutation.OtpValue(); !ok {
		v := feasibilitytoolcontrol.DefaultOtpValue
		ftcc.mutation.SetOtpValue(v)
	}
	if _, ok := ftcc.mutation.MpwOperator(); !ok {
		v := feasibilitytoolcontrol.DefaultMpwOperator
		ftcc.mutation.SetMpwOperator(v)
	}
	if _, ok := ftcc.mutation.MpwValue(); !ok {
		v := feasibilitytoolcontrol.DefaultMpwValue
		ftcc.mutation.SetMpwValue(v)
	}
	if _, ok := ftcc.mutation.MpdOperator(); !ok {
		v := feasibilitytoolcontrol.DefaultMpdOperator
		ftcc.mutation.SetMpdOperator(v)
	}
	if _, ok := ftcc.mutation.MpdValue(); !ok {
		v := feasibilitytoolcontrol.DefaultMpdValue
		ftcc.mutation.SetMpdValue(v)
	}
	if _, ok := ftcc.mutation.MpgOperator(); !ok {
		v := feasibilitytoolcontrol.DefaultMpgOperator
		ftcc.mutation.SetMpgOperator(v)
	}
	if _, ok := ftcc.mutation.MpgValue(); !ok {
		v := feasibilitytoolcontrol.DefaultMpgValue
		ftcc.mutation.SetMpgValue(v)
	}
	if _, ok := ftcc.mutation.ID(); !ok {
		v := feasibilitytoolcontrol.DefaultID()
		ftcc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ftcc *FeasibilityToolControlCreate) check() error {
	if _, ok := ftcc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "FeasibilityToolControl.created_at"`)}
	}
	if _, ok := ftcc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "FeasibilityToolControl.updated_at"`)}
	}
	if _, ok := ftcc.mutation.OtpOperator(); !ok {
		return &ValidationError{Name: "otp_operator", err: errors.New(`ent: missing required field "FeasibilityToolControl.otp_operator"`)}
	}
	if v, ok := ftcc.mutation.OtpOperator(); ok {
		if err := feasibilitytoolcontrol.OtpOperatorValidator(v); err != nil {
			return &ValidationError{Name: "otp_operator", err: fmt.Errorf(`ent: validator failed for field "FeasibilityToolControl.otp_operator": %w`, err)}
		}
	}
	if _, ok := ftcc.mutation.OtpValue(); !ok {
		return &ValidationError{Name: "otp_value", err: errors.New(`ent: missing required field "FeasibilityToolControl.otp_value"`)}
	}
	if _, ok := ftcc.mutation.MpwOperator(); !ok {
		return &ValidationError{Name: "mpw_operator", err: errors.New(`ent: missing required field "FeasibilityToolControl.mpw_operator"`)}
	}
	if v, ok := ftcc.mutation.MpwOperator(); ok {
		if err := feasibilitytoolcontrol.MpwOperatorValidator(v); err != nil {
			return &ValidationError{Name: "mpw_operator", err: fmt.Errorf(`ent: validator failed for field "FeasibilityToolControl.mpw_operator": %w`, err)}
		}
	}
	if _, ok := ftcc.mutation.MpwValue(); !ok {
		return &ValidationError{Name: "mpw_value", err: errors.New(`ent: missing required field "FeasibilityToolControl.mpw_value"`)}
	}
	if _, ok := ftcc.mutation.MpdOperator(); !ok {
		return &ValidationError{Name: "mpd_operator", err: errors.New(`ent: missing required field "FeasibilityToolControl.mpd_operator"`)}
	}
	if v, ok := ftcc.mutation.MpdOperator(); ok {
		if err := feasibilitytoolcontrol.MpdOperatorValidator(v); err != nil {
			return &ValidationError{Name: "mpd_operator", err: fmt.Errorf(`ent: validator failed for field "FeasibilityToolControl.mpd_operator": %w`, err)}
		}
	}
	if _, ok := ftcc.mutation.MpdValue(); !ok {
		return &ValidationError{Name: "mpd_value", err: errors.New(`ent: missing required field "FeasibilityToolControl.mpd_value"`)}
	}
	if _, ok := ftcc.mutation.MpgOperator(); !ok {
		return &ValidationError{Name: "mpg_operator", err: errors.New(`ent: missing required field "FeasibilityToolControl.mpg_operator"`)}
	}
	if v, ok := ftcc.mutation.MpgOperator(); ok {
		if err := feasibilitytoolcontrol.MpgOperatorValidator(v); err != nil {
			return &ValidationError{Name: "mpg_operator", err: fmt.Errorf(`ent: validator failed for field "FeasibilityToolControl.mpg_operator": %w`, err)}
		}
	}
	if _, ok := ftcc.mutation.MpgValue(); !ok {
		return &ValidationError{Name: "mpg_value", err: errors.New(`ent: missing required field "FeasibilityToolControl.mpg_value"`)}
	}
	if _, ok := ftcc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization", err: errors.New(`ent: missing required edge "FeasibilityToolControl.organization"`)}
	}
	if _, ok := ftcc.mutation.BusinessUnitID(); !ok {
		return &ValidationError{Name: "business_unit", err: errors.New(`ent: missing required edge "FeasibilityToolControl.business_unit"`)}
	}
	return nil
}

func (ftcc *FeasibilityToolControlCreate) sqlSave(ctx context.Context) (*FeasibilityToolControl, error) {
	if err := ftcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ftcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ftcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	ftcc.mutation.id = &_node.ID
	ftcc.mutation.done = true
	return _node, nil
}

func (ftcc *FeasibilityToolControlCreate) createSpec() (*FeasibilityToolControl, *sqlgraph.CreateSpec) {
	var (
		_node = &FeasibilityToolControl{config: ftcc.config}
		_spec = sqlgraph.NewCreateSpec(feasibilitytoolcontrol.Table, sqlgraph.NewFieldSpec(feasibilitytoolcontrol.FieldID, field.TypeUUID))
	)
	if id, ok := ftcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ftcc.mutation.CreatedAt(); ok {
		_spec.SetField(feasibilitytoolcontrol.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ftcc.mutation.UpdatedAt(); ok {
		_spec.SetField(feasibilitytoolcontrol.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ftcc.mutation.OtpOperator(); ok {
		_spec.SetField(feasibilitytoolcontrol.FieldOtpOperator, field.TypeEnum, value)
		_node.OtpOperator = value
	}
	if value, ok := ftcc.mutation.OtpValue(); ok {
		_spec.SetField(feasibilitytoolcontrol.FieldOtpValue, field.TypeFloat64, value)
		_node.OtpValue = value
	}
	if value, ok := ftcc.mutation.MpwOperator(); ok {
		_spec.SetField(feasibilitytoolcontrol.FieldMpwOperator, field.TypeEnum, value)
		_node.MpwOperator = value
	}
	if value, ok := ftcc.mutation.MpwValue(); ok {
		_spec.SetField(feasibilitytoolcontrol.FieldMpwValue, field.TypeFloat64, value)
		_node.MpwValue = value
	}
	if value, ok := ftcc.mutation.MpdOperator(); ok {
		_spec.SetField(feasibilitytoolcontrol.FieldMpdOperator, field.TypeEnum, value)
		_node.MpdOperator = value
	}
	if value, ok := ftcc.mutation.MpdValue(); ok {
		_spec.SetField(feasibilitytoolcontrol.FieldMpdValue, field.TypeFloat64, value)
		_node.MpdValue = value
	}
	if value, ok := ftcc.mutation.MpgOperator(); ok {
		_spec.SetField(feasibilitytoolcontrol.FieldMpgOperator, field.TypeEnum, value)
		_node.MpgOperator = value
	}
	if value, ok := ftcc.mutation.MpgValue(); ok {
		_spec.SetField(feasibilitytoolcontrol.FieldMpgValue, field.TypeFloat64, value)
		_node.MpgValue = value
	}
	if nodes := ftcc.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   feasibilitytoolcontrol.OrganizationTable,
			Columns: []string{feasibilitytoolcontrol.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.organization_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ftcc.mutation.BusinessUnitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   feasibilitytoolcontrol.BusinessUnitTable,
			Columns: []string{feasibilitytoolcontrol.BusinessUnitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(businessunit.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.business_unit_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// FeasibilityToolControlCreateBulk is the builder for creating many FeasibilityToolControl entities in bulk.
type FeasibilityToolControlCreateBulk struct {
	config
	err      error
	builders []*FeasibilityToolControlCreate
}

// Save creates the FeasibilityToolControl entities in the database.
func (ftccb *FeasibilityToolControlCreateBulk) Save(ctx context.Context) ([]*FeasibilityToolControl, error) {
	if ftccb.err != nil {
		return nil, ftccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ftccb.builders))
	nodes := make([]*FeasibilityToolControl, len(ftccb.builders))
	mutators := make([]Mutator, len(ftccb.builders))
	for i := range ftccb.builders {
		func(i int, root context.Context) {
			builder := ftccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FeasibilityToolControlMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ftccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ftccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ftccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ftccb *FeasibilityToolControlCreateBulk) SaveX(ctx context.Context) []*FeasibilityToolControl {
	v, err := ftccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ftccb *FeasibilityToolControlCreateBulk) Exec(ctx context.Context) error {
	_, err := ftccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftccb *FeasibilityToolControlCreateBulk) ExecX(ctx context.Context) {
	if err := ftccb.Exec(ctx); err != nil {
		panic(err)
	}
}
