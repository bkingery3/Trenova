// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/ent/generalledgeraccount"
	"github.com/emoss08/trenova/ent/predicate"
	"github.com/emoss08/trenova/ent/tag"
	"github.com/google/uuid"
)

// TagUpdate is the builder for updating Tag entities.
type TagUpdate struct {
	config
	hooks     []Hook
	mutation  *TagMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the TagUpdate builder.
func (tu *TagUpdate) Where(ps ...predicate.Tag) *TagUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *TagUpdate) SetUpdatedAt(t time.Time) *TagUpdate {
	tu.mutation.SetUpdatedAt(t)
	return tu
}

// SetName sets the "name" field.
func (tu *TagUpdate) SetName(s string) *TagUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tu *TagUpdate) SetNillableName(s *string) *TagUpdate {
	if s != nil {
		tu.SetName(*s)
	}
	return tu
}

// SetDescription sets the "description" field.
func (tu *TagUpdate) SetDescription(s string) *TagUpdate {
	tu.mutation.SetDescription(s)
	return tu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tu *TagUpdate) SetNillableDescription(s *string) *TagUpdate {
	if s != nil {
		tu.SetDescription(*s)
	}
	return tu
}

// ClearDescription clears the value of the "description" field.
func (tu *TagUpdate) ClearDescription() *TagUpdate {
	tu.mutation.ClearDescription()
	return tu
}

// SetColor sets the "color" field.
func (tu *TagUpdate) SetColor(s string) *TagUpdate {
	tu.mutation.SetColor(s)
	return tu
}

// SetNillableColor sets the "color" field if the given value is not nil.
func (tu *TagUpdate) SetNillableColor(s *string) *TagUpdate {
	if s != nil {
		tu.SetColor(*s)
	}
	return tu
}

// ClearColor clears the value of the "color" field.
func (tu *TagUpdate) ClearColor() *TagUpdate {
	tu.mutation.ClearColor()
	return tu
}

// AddGeneralLedgerAccountIDs adds the "general_ledger_account" edge to the GeneralLedgerAccount entity by IDs.
func (tu *TagUpdate) AddGeneralLedgerAccountIDs(ids ...uuid.UUID) *TagUpdate {
	tu.mutation.AddGeneralLedgerAccountIDs(ids...)
	return tu
}

// AddGeneralLedgerAccount adds the "general_ledger_account" edges to the GeneralLedgerAccount entity.
func (tu *TagUpdate) AddGeneralLedgerAccount(g ...*GeneralLedgerAccount) *TagUpdate {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return tu.AddGeneralLedgerAccountIDs(ids...)
}

// Mutation returns the TagMutation object of the builder.
func (tu *TagUpdate) Mutation() *TagMutation {
	return tu.mutation
}

// ClearGeneralLedgerAccount clears all "general_ledger_account" edges to the GeneralLedgerAccount entity.
func (tu *TagUpdate) ClearGeneralLedgerAccount() *TagUpdate {
	tu.mutation.ClearGeneralLedgerAccount()
	return tu
}

// RemoveGeneralLedgerAccountIDs removes the "general_ledger_account" edge to GeneralLedgerAccount entities by IDs.
func (tu *TagUpdate) RemoveGeneralLedgerAccountIDs(ids ...uuid.UUID) *TagUpdate {
	tu.mutation.RemoveGeneralLedgerAccountIDs(ids...)
	return tu
}

// RemoveGeneralLedgerAccount removes "general_ledger_account" edges to GeneralLedgerAccount entities.
func (tu *TagUpdate) RemoveGeneralLedgerAccount(g ...*GeneralLedgerAccount) *TagUpdate {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return tu.RemoveGeneralLedgerAccountIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TagUpdate) Save(ctx context.Context) (int, error) {
	tu.defaults()
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TagUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TagUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TagUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TagUpdate) defaults() {
	if _, ok := tu.mutation.UpdatedAt(); !ok {
		v := tag.UpdateDefaultUpdatedAt()
		tu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TagUpdate) check() error {
	if v, ok := tu.mutation.Name(); ok {
		if err := tag.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Tag.name": %w`, err)}
		}
	}
	if _, ok := tu.mutation.BusinessUnitID(); tu.mutation.BusinessUnitCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Tag.business_unit"`)
	}
	if _, ok := tu.mutation.OrganizationID(); tu.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Tag.organization"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tu *TagUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TagUpdate {
	tu.modifiers = append(tu.modifiers, modifiers...)
	return tu
}

func (tu *TagUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(tag.Table, tag.Columns, sqlgraph.NewFieldSpec(tag.FieldID, field.TypeUUID))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.SetField(tag.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.SetField(tag.FieldName, field.TypeString, value)
	}
	if value, ok := tu.mutation.Description(); ok {
		_spec.SetField(tag.FieldDescription, field.TypeString, value)
	}
	if tu.mutation.DescriptionCleared() {
		_spec.ClearField(tag.FieldDescription, field.TypeString)
	}
	if value, ok := tu.mutation.Color(); ok {
		_spec.SetField(tag.FieldColor, field.TypeString, value)
	}
	if tu.mutation.ColorCleared() {
		_spec.ClearField(tag.FieldColor, field.TypeString)
	}
	if tu.mutation.GeneralLedgerAccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.GeneralLedgerAccountTable,
			Columns: tag.GeneralLedgerAccountPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(generalledgeraccount.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedGeneralLedgerAccountIDs(); len(nodes) > 0 && !tu.mutation.GeneralLedgerAccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.GeneralLedgerAccountTable,
			Columns: tag.GeneralLedgerAccountPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(generalledgeraccount.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.GeneralLedgerAccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.GeneralLedgerAccountTable,
			Columns: tag.GeneralLedgerAccountPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(generalledgeraccount.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(tu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TagUpdateOne is the builder for updating a single Tag entity.
type TagUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *TagMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *TagUpdateOne) SetUpdatedAt(t time.Time) *TagUpdateOne {
	tuo.mutation.SetUpdatedAt(t)
	return tuo
}

// SetName sets the "name" field.
func (tuo *TagUpdateOne) SetName(s string) *TagUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tuo *TagUpdateOne) SetNillableName(s *string) *TagUpdateOne {
	if s != nil {
		tuo.SetName(*s)
	}
	return tuo
}

// SetDescription sets the "description" field.
func (tuo *TagUpdateOne) SetDescription(s string) *TagUpdateOne {
	tuo.mutation.SetDescription(s)
	return tuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tuo *TagUpdateOne) SetNillableDescription(s *string) *TagUpdateOne {
	if s != nil {
		tuo.SetDescription(*s)
	}
	return tuo
}

// ClearDescription clears the value of the "description" field.
func (tuo *TagUpdateOne) ClearDescription() *TagUpdateOne {
	tuo.mutation.ClearDescription()
	return tuo
}

// SetColor sets the "color" field.
func (tuo *TagUpdateOne) SetColor(s string) *TagUpdateOne {
	tuo.mutation.SetColor(s)
	return tuo
}

// SetNillableColor sets the "color" field if the given value is not nil.
func (tuo *TagUpdateOne) SetNillableColor(s *string) *TagUpdateOne {
	if s != nil {
		tuo.SetColor(*s)
	}
	return tuo
}

// ClearColor clears the value of the "color" field.
func (tuo *TagUpdateOne) ClearColor() *TagUpdateOne {
	tuo.mutation.ClearColor()
	return tuo
}

// AddGeneralLedgerAccountIDs adds the "general_ledger_account" edge to the GeneralLedgerAccount entity by IDs.
func (tuo *TagUpdateOne) AddGeneralLedgerAccountIDs(ids ...uuid.UUID) *TagUpdateOne {
	tuo.mutation.AddGeneralLedgerAccountIDs(ids...)
	return tuo
}

// AddGeneralLedgerAccount adds the "general_ledger_account" edges to the GeneralLedgerAccount entity.
func (tuo *TagUpdateOne) AddGeneralLedgerAccount(g ...*GeneralLedgerAccount) *TagUpdateOne {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return tuo.AddGeneralLedgerAccountIDs(ids...)
}

// Mutation returns the TagMutation object of the builder.
func (tuo *TagUpdateOne) Mutation() *TagMutation {
	return tuo.mutation
}

// ClearGeneralLedgerAccount clears all "general_ledger_account" edges to the GeneralLedgerAccount entity.
func (tuo *TagUpdateOne) ClearGeneralLedgerAccount() *TagUpdateOne {
	tuo.mutation.ClearGeneralLedgerAccount()
	return tuo
}

// RemoveGeneralLedgerAccountIDs removes the "general_ledger_account" edge to GeneralLedgerAccount entities by IDs.
func (tuo *TagUpdateOne) RemoveGeneralLedgerAccountIDs(ids ...uuid.UUID) *TagUpdateOne {
	tuo.mutation.RemoveGeneralLedgerAccountIDs(ids...)
	return tuo
}

// RemoveGeneralLedgerAccount removes "general_ledger_account" edges to GeneralLedgerAccount entities.
func (tuo *TagUpdateOne) RemoveGeneralLedgerAccount(g ...*GeneralLedgerAccount) *TagUpdateOne {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return tuo.RemoveGeneralLedgerAccountIDs(ids...)
}

// Where appends a list predicates to the TagUpdate builder.
func (tuo *TagUpdateOne) Where(ps ...predicate.Tag) *TagUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TagUpdateOne) Select(field string, fields ...string) *TagUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Tag entity.
func (tuo *TagUpdateOne) Save(ctx context.Context) (*Tag, error) {
	tuo.defaults()
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TagUpdateOne) SaveX(ctx context.Context) *Tag {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TagUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TagUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TagUpdateOne) defaults() {
	if _, ok := tuo.mutation.UpdatedAt(); !ok {
		v := tag.UpdateDefaultUpdatedAt()
		tuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TagUpdateOne) check() error {
	if v, ok := tuo.mutation.Name(); ok {
		if err := tag.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Tag.name": %w`, err)}
		}
	}
	if _, ok := tuo.mutation.BusinessUnitID(); tuo.mutation.BusinessUnitCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Tag.business_unit"`)
	}
	if _, ok := tuo.mutation.OrganizationID(); tuo.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Tag.organization"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tuo *TagUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TagUpdateOne {
	tuo.modifiers = append(tuo.modifiers, modifiers...)
	return tuo
}

func (tuo *TagUpdateOne) sqlSave(ctx context.Context) (_node *Tag, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(tag.Table, tag.Columns, sqlgraph.NewFieldSpec(tag.FieldID, field.TypeUUID))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Tag.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tag.FieldID)
		for _, f := range fields {
			if !tag.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.SetField(tag.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.SetField(tag.FieldName, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Description(); ok {
		_spec.SetField(tag.FieldDescription, field.TypeString, value)
	}
	if tuo.mutation.DescriptionCleared() {
		_spec.ClearField(tag.FieldDescription, field.TypeString)
	}
	if value, ok := tuo.mutation.Color(); ok {
		_spec.SetField(tag.FieldColor, field.TypeString, value)
	}
	if tuo.mutation.ColorCleared() {
		_spec.ClearField(tag.FieldColor, field.TypeString)
	}
	if tuo.mutation.GeneralLedgerAccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.GeneralLedgerAccountTable,
			Columns: tag.GeneralLedgerAccountPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(generalledgeraccount.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedGeneralLedgerAccountIDs(); len(nodes) > 0 && !tuo.mutation.GeneralLedgerAccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.GeneralLedgerAccountTable,
			Columns: tag.GeneralLedgerAccountPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(generalledgeraccount.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.GeneralLedgerAccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.GeneralLedgerAccountTable,
			Columns: tag.GeneralLedgerAccountPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(generalledgeraccount.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(tuo.modifiers...)
	_node = &Tag{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
