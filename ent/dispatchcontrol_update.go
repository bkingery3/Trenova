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
	"github.com/emoss08/trenova/ent/businessunit"
	"github.com/emoss08/trenova/ent/dispatchcontrol"
	"github.com/emoss08/trenova/ent/organization"
	"github.com/emoss08/trenova/ent/predicate"
	"github.com/google/uuid"
)

// DispatchControlUpdate is the builder for updating DispatchControl entities.
type DispatchControlUpdate struct {
	config
	hooks    []Hook
	mutation *DispatchControlMutation
}

// Where appends a list predicates to the DispatchControlUpdate builder.
func (dcu *DispatchControlUpdate) Where(ps ...predicate.DispatchControl) *DispatchControlUpdate {
	dcu.mutation.Where(ps...)
	return dcu
}

// SetUpdatedAt sets the "updated_at" field.
func (dcu *DispatchControlUpdate) SetUpdatedAt(t time.Time) *DispatchControlUpdate {
	dcu.mutation.SetUpdatedAt(t)
	return dcu
}

// SetRecordServiceIncident sets the "record_service_incident" field.
func (dcu *DispatchControlUpdate) SetRecordServiceIncident(dsi dispatchcontrol.RecordServiceIncident) *DispatchControlUpdate {
	dcu.mutation.SetRecordServiceIncident(dsi)
	return dcu
}

// SetNillableRecordServiceIncident sets the "record_service_incident" field if the given value is not nil.
func (dcu *DispatchControlUpdate) SetNillableRecordServiceIncident(dsi *dispatchcontrol.RecordServiceIncident) *DispatchControlUpdate {
	if dsi != nil {
		dcu.SetRecordServiceIncident(*dsi)
	}
	return dcu
}

// SetDeadheadTarget sets the "deadhead_target" field.
func (dcu *DispatchControlUpdate) SetDeadheadTarget(f float64) *DispatchControlUpdate {
	dcu.mutation.ResetDeadheadTarget()
	dcu.mutation.SetDeadheadTarget(f)
	return dcu
}

// SetNillableDeadheadTarget sets the "deadhead_target" field if the given value is not nil.
func (dcu *DispatchControlUpdate) SetNillableDeadheadTarget(f *float64) *DispatchControlUpdate {
	if f != nil {
		dcu.SetDeadheadTarget(*f)
	}
	return dcu
}

// AddDeadheadTarget adds f to the "deadhead_target" field.
func (dcu *DispatchControlUpdate) AddDeadheadTarget(f float64) *DispatchControlUpdate {
	dcu.mutation.AddDeadheadTarget(f)
	return dcu
}

// SetMaxShipmentWeightLimit sets the "max_shipment_weight_limit" field.
func (dcu *DispatchControlUpdate) SetMaxShipmentWeightLimit(i int) *DispatchControlUpdate {
	dcu.mutation.ResetMaxShipmentWeightLimit()
	dcu.mutation.SetMaxShipmentWeightLimit(i)
	return dcu
}

// SetNillableMaxShipmentWeightLimit sets the "max_shipment_weight_limit" field if the given value is not nil.
func (dcu *DispatchControlUpdate) SetNillableMaxShipmentWeightLimit(i *int) *DispatchControlUpdate {
	if i != nil {
		dcu.SetMaxShipmentWeightLimit(*i)
	}
	return dcu
}

// AddMaxShipmentWeightLimit adds i to the "max_shipment_weight_limit" field.
func (dcu *DispatchControlUpdate) AddMaxShipmentWeightLimit(i int) *DispatchControlUpdate {
	dcu.mutation.AddMaxShipmentWeightLimit(i)
	return dcu
}

// SetGracePeriod sets the "grace_period" field.
func (dcu *DispatchControlUpdate) SetGracePeriod(u uint8) *DispatchControlUpdate {
	dcu.mutation.ResetGracePeriod()
	dcu.mutation.SetGracePeriod(u)
	return dcu
}

// SetNillableGracePeriod sets the "grace_period" field if the given value is not nil.
func (dcu *DispatchControlUpdate) SetNillableGracePeriod(u *uint8) *DispatchControlUpdate {
	if u != nil {
		dcu.SetGracePeriod(*u)
	}
	return dcu
}

// AddGracePeriod adds u to the "grace_period" field.
func (dcu *DispatchControlUpdate) AddGracePeriod(u int8) *DispatchControlUpdate {
	dcu.mutation.AddGracePeriod(u)
	return dcu
}

// SetEnforceWorkerAssign sets the "enforce_worker_assign" field.
func (dcu *DispatchControlUpdate) SetEnforceWorkerAssign(b bool) *DispatchControlUpdate {
	dcu.mutation.SetEnforceWorkerAssign(b)
	return dcu
}

// SetNillableEnforceWorkerAssign sets the "enforce_worker_assign" field if the given value is not nil.
func (dcu *DispatchControlUpdate) SetNillableEnforceWorkerAssign(b *bool) *DispatchControlUpdate {
	if b != nil {
		dcu.SetEnforceWorkerAssign(*b)
	}
	return dcu
}

// SetTrailerContinuity sets the "trailer_continuity" field.
func (dcu *DispatchControlUpdate) SetTrailerContinuity(b bool) *DispatchControlUpdate {
	dcu.mutation.SetTrailerContinuity(b)
	return dcu
}

// SetNillableTrailerContinuity sets the "trailer_continuity" field if the given value is not nil.
func (dcu *DispatchControlUpdate) SetNillableTrailerContinuity(b *bool) *DispatchControlUpdate {
	if b != nil {
		dcu.SetTrailerContinuity(*b)
	}
	return dcu
}

// SetDupeTrailerCheck sets the "dupe_trailer_check" field.
func (dcu *DispatchControlUpdate) SetDupeTrailerCheck(b bool) *DispatchControlUpdate {
	dcu.mutation.SetDupeTrailerCheck(b)
	return dcu
}

// SetNillableDupeTrailerCheck sets the "dupe_trailer_check" field if the given value is not nil.
func (dcu *DispatchControlUpdate) SetNillableDupeTrailerCheck(b *bool) *DispatchControlUpdate {
	if b != nil {
		dcu.SetDupeTrailerCheck(*b)
	}
	return dcu
}

// SetMaintenanceCompliance sets the "maintenance_compliance" field.
func (dcu *DispatchControlUpdate) SetMaintenanceCompliance(b bool) *DispatchControlUpdate {
	dcu.mutation.SetMaintenanceCompliance(b)
	return dcu
}

// SetNillableMaintenanceCompliance sets the "maintenance_compliance" field if the given value is not nil.
func (dcu *DispatchControlUpdate) SetNillableMaintenanceCompliance(b *bool) *DispatchControlUpdate {
	if b != nil {
		dcu.SetMaintenanceCompliance(*b)
	}
	return dcu
}

// SetRegulatoryCheck sets the "regulatory_check" field.
func (dcu *DispatchControlUpdate) SetRegulatoryCheck(b bool) *DispatchControlUpdate {
	dcu.mutation.SetRegulatoryCheck(b)
	return dcu
}

// SetNillableRegulatoryCheck sets the "regulatory_check" field if the given value is not nil.
func (dcu *DispatchControlUpdate) SetNillableRegulatoryCheck(b *bool) *DispatchControlUpdate {
	if b != nil {
		dcu.SetRegulatoryCheck(*b)
	}
	return dcu
}

// SetPrevShipmentOnHold sets the "prev_shipment_on_hold" field.
func (dcu *DispatchControlUpdate) SetPrevShipmentOnHold(b bool) *DispatchControlUpdate {
	dcu.mutation.SetPrevShipmentOnHold(b)
	return dcu
}

// SetNillablePrevShipmentOnHold sets the "prev_shipment_on_hold" field if the given value is not nil.
func (dcu *DispatchControlUpdate) SetNillablePrevShipmentOnHold(b *bool) *DispatchControlUpdate {
	if b != nil {
		dcu.SetPrevShipmentOnHold(*b)
	}
	return dcu
}

// SetWorkerTimeAwayRestriction sets the "worker_time_away_restriction" field.
func (dcu *DispatchControlUpdate) SetWorkerTimeAwayRestriction(b bool) *DispatchControlUpdate {
	dcu.mutation.SetWorkerTimeAwayRestriction(b)
	return dcu
}

// SetNillableWorkerTimeAwayRestriction sets the "worker_time_away_restriction" field if the given value is not nil.
func (dcu *DispatchControlUpdate) SetNillableWorkerTimeAwayRestriction(b *bool) *DispatchControlUpdate {
	if b != nil {
		dcu.SetWorkerTimeAwayRestriction(*b)
	}
	return dcu
}

// SetTractorWorkerFleetConstraint sets the "tractor_worker_fleet_constraint" field.
func (dcu *DispatchControlUpdate) SetTractorWorkerFleetConstraint(b bool) *DispatchControlUpdate {
	dcu.mutation.SetTractorWorkerFleetConstraint(b)
	return dcu
}

// SetNillableTractorWorkerFleetConstraint sets the "tractor_worker_fleet_constraint" field if the given value is not nil.
func (dcu *DispatchControlUpdate) SetNillableTractorWorkerFleetConstraint(b *bool) *DispatchControlUpdate {
	if b != nil {
		dcu.SetTractorWorkerFleetConstraint(*b)
	}
	return dcu
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (dcu *DispatchControlUpdate) SetOrganizationID(id uuid.UUID) *DispatchControlUpdate {
	dcu.mutation.SetOrganizationID(id)
	return dcu
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (dcu *DispatchControlUpdate) SetOrganization(o *Organization) *DispatchControlUpdate {
	return dcu.SetOrganizationID(o.ID)
}

// SetBusinessUnitID sets the "business_unit" edge to the BusinessUnit entity by ID.
func (dcu *DispatchControlUpdate) SetBusinessUnitID(id uuid.UUID) *DispatchControlUpdate {
	dcu.mutation.SetBusinessUnitID(id)
	return dcu
}

// SetBusinessUnit sets the "business_unit" edge to the BusinessUnit entity.
func (dcu *DispatchControlUpdate) SetBusinessUnit(b *BusinessUnit) *DispatchControlUpdate {
	return dcu.SetBusinessUnitID(b.ID)
}

// Mutation returns the DispatchControlMutation object of the builder.
func (dcu *DispatchControlUpdate) Mutation() *DispatchControlMutation {
	return dcu.mutation
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (dcu *DispatchControlUpdate) ClearOrganization() *DispatchControlUpdate {
	dcu.mutation.ClearOrganization()
	return dcu
}

// ClearBusinessUnit clears the "business_unit" edge to the BusinessUnit entity.
func (dcu *DispatchControlUpdate) ClearBusinessUnit() *DispatchControlUpdate {
	dcu.mutation.ClearBusinessUnit()
	return dcu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (dcu *DispatchControlUpdate) Save(ctx context.Context) (int, error) {
	dcu.defaults()
	return withHooks(ctx, dcu.sqlSave, dcu.mutation, dcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (dcu *DispatchControlUpdate) SaveX(ctx context.Context) int {
	affected, err := dcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (dcu *DispatchControlUpdate) Exec(ctx context.Context) error {
	_, err := dcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcu *DispatchControlUpdate) ExecX(ctx context.Context) {
	if err := dcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dcu *DispatchControlUpdate) defaults() {
	if _, ok := dcu.mutation.UpdatedAt(); !ok {
		v := dispatchcontrol.UpdateDefaultUpdatedAt()
		dcu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dcu *DispatchControlUpdate) check() error {
	if v, ok := dcu.mutation.RecordServiceIncident(); ok {
		if err := dispatchcontrol.RecordServiceIncidentValidator(v); err != nil {
			return &ValidationError{Name: "record_service_incident", err: fmt.Errorf(`ent: validator failed for field "DispatchControl.record_service_incident": %w`, err)}
		}
	}
	if v, ok := dcu.mutation.MaxShipmentWeightLimit(); ok {
		if err := dispatchcontrol.MaxShipmentWeightLimitValidator(v); err != nil {
			return &ValidationError{Name: "max_shipment_weight_limit", err: fmt.Errorf(`ent: validator failed for field "DispatchControl.max_shipment_weight_limit": %w`, err)}
		}
	}
	if _, ok := dcu.mutation.OrganizationID(); dcu.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "DispatchControl.organization"`)
	}
	if _, ok := dcu.mutation.BusinessUnitID(); dcu.mutation.BusinessUnitCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "DispatchControl.business_unit"`)
	}
	return nil
}

func (dcu *DispatchControlUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := dcu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(dispatchcontrol.Table, dispatchcontrol.Columns, sqlgraph.NewFieldSpec(dispatchcontrol.FieldID, field.TypeUUID))
	if ps := dcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dcu.mutation.UpdatedAt(); ok {
		_spec.SetField(dispatchcontrol.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := dcu.mutation.RecordServiceIncident(); ok {
		_spec.SetField(dispatchcontrol.FieldRecordServiceIncident, field.TypeEnum, value)
	}
	if value, ok := dcu.mutation.DeadheadTarget(); ok {
		_spec.SetField(dispatchcontrol.FieldDeadheadTarget, field.TypeFloat64, value)
	}
	if value, ok := dcu.mutation.AddedDeadheadTarget(); ok {
		_spec.AddField(dispatchcontrol.FieldDeadheadTarget, field.TypeFloat64, value)
	}
	if value, ok := dcu.mutation.MaxShipmentWeightLimit(); ok {
		_spec.SetField(dispatchcontrol.FieldMaxShipmentWeightLimit, field.TypeInt, value)
	}
	if value, ok := dcu.mutation.AddedMaxShipmentWeightLimit(); ok {
		_spec.AddField(dispatchcontrol.FieldMaxShipmentWeightLimit, field.TypeInt, value)
	}
	if value, ok := dcu.mutation.GracePeriod(); ok {
		_spec.SetField(dispatchcontrol.FieldGracePeriod, field.TypeUint8, value)
	}
	if value, ok := dcu.mutation.AddedGracePeriod(); ok {
		_spec.AddField(dispatchcontrol.FieldGracePeriod, field.TypeUint8, value)
	}
	if value, ok := dcu.mutation.EnforceWorkerAssign(); ok {
		_spec.SetField(dispatchcontrol.FieldEnforceWorkerAssign, field.TypeBool, value)
	}
	if value, ok := dcu.mutation.TrailerContinuity(); ok {
		_spec.SetField(dispatchcontrol.FieldTrailerContinuity, field.TypeBool, value)
	}
	if value, ok := dcu.mutation.DupeTrailerCheck(); ok {
		_spec.SetField(dispatchcontrol.FieldDupeTrailerCheck, field.TypeBool, value)
	}
	if value, ok := dcu.mutation.MaintenanceCompliance(); ok {
		_spec.SetField(dispatchcontrol.FieldMaintenanceCompliance, field.TypeBool, value)
	}
	if value, ok := dcu.mutation.RegulatoryCheck(); ok {
		_spec.SetField(dispatchcontrol.FieldRegulatoryCheck, field.TypeBool, value)
	}
	if value, ok := dcu.mutation.PrevShipmentOnHold(); ok {
		_spec.SetField(dispatchcontrol.FieldPrevShipmentOnHold, field.TypeBool, value)
	}
	if value, ok := dcu.mutation.WorkerTimeAwayRestriction(); ok {
		_spec.SetField(dispatchcontrol.FieldWorkerTimeAwayRestriction, field.TypeBool, value)
	}
	if value, ok := dcu.mutation.TractorWorkerFleetConstraint(); ok {
		_spec.SetField(dispatchcontrol.FieldTractorWorkerFleetConstraint, field.TypeBool, value)
	}
	if dcu.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   dispatchcontrol.OrganizationTable,
			Columns: []string{dispatchcontrol.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dcu.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   dispatchcontrol.OrganizationTable,
			Columns: []string{dispatchcontrol.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if dcu.mutation.BusinessUnitCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   dispatchcontrol.BusinessUnitTable,
			Columns: []string{dispatchcontrol.BusinessUnitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(businessunit.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dcu.mutation.BusinessUnitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   dispatchcontrol.BusinessUnitTable,
			Columns: []string{dispatchcontrol.BusinessUnitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(businessunit.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, dcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dispatchcontrol.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	dcu.mutation.done = true
	return n, nil
}

// DispatchControlUpdateOne is the builder for updating a single DispatchControl entity.
type DispatchControlUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DispatchControlMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (dcuo *DispatchControlUpdateOne) SetUpdatedAt(t time.Time) *DispatchControlUpdateOne {
	dcuo.mutation.SetUpdatedAt(t)
	return dcuo
}

// SetRecordServiceIncident sets the "record_service_incident" field.
func (dcuo *DispatchControlUpdateOne) SetRecordServiceIncident(dsi dispatchcontrol.RecordServiceIncident) *DispatchControlUpdateOne {
	dcuo.mutation.SetRecordServiceIncident(dsi)
	return dcuo
}

// SetNillableRecordServiceIncident sets the "record_service_incident" field if the given value is not nil.
func (dcuo *DispatchControlUpdateOne) SetNillableRecordServiceIncident(dsi *dispatchcontrol.RecordServiceIncident) *DispatchControlUpdateOne {
	if dsi != nil {
		dcuo.SetRecordServiceIncident(*dsi)
	}
	return dcuo
}

// SetDeadheadTarget sets the "deadhead_target" field.
func (dcuo *DispatchControlUpdateOne) SetDeadheadTarget(f float64) *DispatchControlUpdateOne {
	dcuo.mutation.ResetDeadheadTarget()
	dcuo.mutation.SetDeadheadTarget(f)
	return dcuo
}

// SetNillableDeadheadTarget sets the "deadhead_target" field if the given value is not nil.
func (dcuo *DispatchControlUpdateOne) SetNillableDeadheadTarget(f *float64) *DispatchControlUpdateOne {
	if f != nil {
		dcuo.SetDeadheadTarget(*f)
	}
	return dcuo
}

// AddDeadheadTarget adds f to the "deadhead_target" field.
func (dcuo *DispatchControlUpdateOne) AddDeadheadTarget(f float64) *DispatchControlUpdateOne {
	dcuo.mutation.AddDeadheadTarget(f)
	return dcuo
}

// SetMaxShipmentWeightLimit sets the "max_shipment_weight_limit" field.
func (dcuo *DispatchControlUpdateOne) SetMaxShipmentWeightLimit(i int) *DispatchControlUpdateOne {
	dcuo.mutation.ResetMaxShipmentWeightLimit()
	dcuo.mutation.SetMaxShipmentWeightLimit(i)
	return dcuo
}

// SetNillableMaxShipmentWeightLimit sets the "max_shipment_weight_limit" field if the given value is not nil.
func (dcuo *DispatchControlUpdateOne) SetNillableMaxShipmentWeightLimit(i *int) *DispatchControlUpdateOne {
	if i != nil {
		dcuo.SetMaxShipmentWeightLimit(*i)
	}
	return dcuo
}

// AddMaxShipmentWeightLimit adds i to the "max_shipment_weight_limit" field.
func (dcuo *DispatchControlUpdateOne) AddMaxShipmentWeightLimit(i int) *DispatchControlUpdateOne {
	dcuo.mutation.AddMaxShipmentWeightLimit(i)
	return dcuo
}

// SetGracePeriod sets the "grace_period" field.
func (dcuo *DispatchControlUpdateOne) SetGracePeriod(u uint8) *DispatchControlUpdateOne {
	dcuo.mutation.ResetGracePeriod()
	dcuo.mutation.SetGracePeriod(u)
	return dcuo
}

// SetNillableGracePeriod sets the "grace_period" field if the given value is not nil.
func (dcuo *DispatchControlUpdateOne) SetNillableGracePeriod(u *uint8) *DispatchControlUpdateOne {
	if u != nil {
		dcuo.SetGracePeriod(*u)
	}
	return dcuo
}

// AddGracePeriod adds u to the "grace_period" field.
func (dcuo *DispatchControlUpdateOne) AddGracePeriod(u int8) *DispatchControlUpdateOne {
	dcuo.mutation.AddGracePeriod(u)
	return dcuo
}

// SetEnforceWorkerAssign sets the "enforce_worker_assign" field.
func (dcuo *DispatchControlUpdateOne) SetEnforceWorkerAssign(b bool) *DispatchControlUpdateOne {
	dcuo.mutation.SetEnforceWorkerAssign(b)
	return dcuo
}

// SetNillableEnforceWorkerAssign sets the "enforce_worker_assign" field if the given value is not nil.
func (dcuo *DispatchControlUpdateOne) SetNillableEnforceWorkerAssign(b *bool) *DispatchControlUpdateOne {
	if b != nil {
		dcuo.SetEnforceWorkerAssign(*b)
	}
	return dcuo
}

// SetTrailerContinuity sets the "trailer_continuity" field.
func (dcuo *DispatchControlUpdateOne) SetTrailerContinuity(b bool) *DispatchControlUpdateOne {
	dcuo.mutation.SetTrailerContinuity(b)
	return dcuo
}

// SetNillableTrailerContinuity sets the "trailer_continuity" field if the given value is not nil.
func (dcuo *DispatchControlUpdateOne) SetNillableTrailerContinuity(b *bool) *DispatchControlUpdateOne {
	if b != nil {
		dcuo.SetTrailerContinuity(*b)
	}
	return dcuo
}

// SetDupeTrailerCheck sets the "dupe_trailer_check" field.
func (dcuo *DispatchControlUpdateOne) SetDupeTrailerCheck(b bool) *DispatchControlUpdateOne {
	dcuo.mutation.SetDupeTrailerCheck(b)
	return dcuo
}

// SetNillableDupeTrailerCheck sets the "dupe_trailer_check" field if the given value is not nil.
func (dcuo *DispatchControlUpdateOne) SetNillableDupeTrailerCheck(b *bool) *DispatchControlUpdateOne {
	if b != nil {
		dcuo.SetDupeTrailerCheck(*b)
	}
	return dcuo
}

// SetMaintenanceCompliance sets the "maintenance_compliance" field.
func (dcuo *DispatchControlUpdateOne) SetMaintenanceCompliance(b bool) *DispatchControlUpdateOne {
	dcuo.mutation.SetMaintenanceCompliance(b)
	return dcuo
}

// SetNillableMaintenanceCompliance sets the "maintenance_compliance" field if the given value is not nil.
func (dcuo *DispatchControlUpdateOne) SetNillableMaintenanceCompliance(b *bool) *DispatchControlUpdateOne {
	if b != nil {
		dcuo.SetMaintenanceCompliance(*b)
	}
	return dcuo
}

// SetRegulatoryCheck sets the "regulatory_check" field.
func (dcuo *DispatchControlUpdateOne) SetRegulatoryCheck(b bool) *DispatchControlUpdateOne {
	dcuo.mutation.SetRegulatoryCheck(b)
	return dcuo
}

// SetNillableRegulatoryCheck sets the "regulatory_check" field if the given value is not nil.
func (dcuo *DispatchControlUpdateOne) SetNillableRegulatoryCheck(b *bool) *DispatchControlUpdateOne {
	if b != nil {
		dcuo.SetRegulatoryCheck(*b)
	}
	return dcuo
}

// SetPrevShipmentOnHold sets the "prev_shipment_on_hold" field.
func (dcuo *DispatchControlUpdateOne) SetPrevShipmentOnHold(b bool) *DispatchControlUpdateOne {
	dcuo.mutation.SetPrevShipmentOnHold(b)
	return dcuo
}

// SetNillablePrevShipmentOnHold sets the "prev_shipment_on_hold" field if the given value is not nil.
func (dcuo *DispatchControlUpdateOne) SetNillablePrevShipmentOnHold(b *bool) *DispatchControlUpdateOne {
	if b != nil {
		dcuo.SetPrevShipmentOnHold(*b)
	}
	return dcuo
}

// SetWorkerTimeAwayRestriction sets the "worker_time_away_restriction" field.
func (dcuo *DispatchControlUpdateOne) SetWorkerTimeAwayRestriction(b bool) *DispatchControlUpdateOne {
	dcuo.mutation.SetWorkerTimeAwayRestriction(b)
	return dcuo
}

// SetNillableWorkerTimeAwayRestriction sets the "worker_time_away_restriction" field if the given value is not nil.
func (dcuo *DispatchControlUpdateOne) SetNillableWorkerTimeAwayRestriction(b *bool) *DispatchControlUpdateOne {
	if b != nil {
		dcuo.SetWorkerTimeAwayRestriction(*b)
	}
	return dcuo
}

// SetTractorWorkerFleetConstraint sets the "tractor_worker_fleet_constraint" field.
func (dcuo *DispatchControlUpdateOne) SetTractorWorkerFleetConstraint(b bool) *DispatchControlUpdateOne {
	dcuo.mutation.SetTractorWorkerFleetConstraint(b)
	return dcuo
}

// SetNillableTractorWorkerFleetConstraint sets the "tractor_worker_fleet_constraint" field if the given value is not nil.
func (dcuo *DispatchControlUpdateOne) SetNillableTractorWorkerFleetConstraint(b *bool) *DispatchControlUpdateOne {
	if b != nil {
		dcuo.SetTractorWorkerFleetConstraint(*b)
	}
	return dcuo
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (dcuo *DispatchControlUpdateOne) SetOrganizationID(id uuid.UUID) *DispatchControlUpdateOne {
	dcuo.mutation.SetOrganizationID(id)
	return dcuo
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (dcuo *DispatchControlUpdateOne) SetOrganization(o *Organization) *DispatchControlUpdateOne {
	return dcuo.SetOrganizationID(o.ID)
}

// SetBusinessUnitID sets the "business_unit" edge to the BusinessUnit entity by ID.
func (dcuo *DispatchControlUpdateOne) SetBusinessUnitID(id uuid.UUID) *DispatchControlUpdateOne {
	dcuo.mutation.SetBusinessUnitID(id)
	return dcuo
}

// SetBusinessUnit sets the "business_unit" edge to the BusinessUnit entity.
func (dcuo *DispatchControlUpdateOne) SetBusinessUnit(b *BusinessUnit) *DispatchControlUpdateOne {
	return dcuo.SetBusinessUnitID(b.ID)
}

// Mutation returns the DispatchControlMutation object of the builder.
func (dcuo *DispatchControlUpdateOne) Mutation() *DispatchControlMutation {
	return dcuo.mutation
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (dcuo *DispatchControlUpdateOne) ClearOrganization() *DispatchControlUpdateOne {
	dcuo.mutation.ClearOrganization()
	return dcuo
}

// ClearBusinessUnit clears the "business_unit" edge to the BusinessUnit entity.
func (dcuo *DispatchControlUpdateOne) ClearBusinessUnit() *DispatchControlUpdateOne {
	dcuo.mutation.ClearBusinessUnit()
	return dcuo
}

// Where appends a list predicates to the DispatchControlUpdate builder.
func (dcuo *DispatchControlUpdateOne) Where(ps ...predicate.DispatchControl) *DispatchControlUpdateOne {
	dcuo.mutation.Where(ps...)
	return dcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (dcuo *DispatchControlUpdateOne) Select(field string, fields ...string) *DispatchControlUpdateOne {
	dcuo.fields = append([]string{field}, fields...)
	return dcuo
}

// Save executes the query and returns the updated DispatchControl entity.
func (dcuo *DispatchControlUpdateOne) Save(ctx context.Context) (*DispatchControl, error) {
	dcuo.defaults()
	return withHooks(ctx, dcuo.sqlSave, dcuo.mutation, dcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (dcuo *DispatchControlUpdateOne) SaveX(ctx context.Context) *DispatchControl {
	node, err := dcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (dcuo *DispatchControlUpdateOne) Exec(ctx context.Context) error {
	_, err := dcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcuo *DispatchControlUpdateOne) ExecX(ctx context.Context) {
	if err := dcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dcuo *DispatchControlUpdateOne) defaults() {
	if _, ok := dcuo.mutation.UpdatedAt(); !ok {
		v := dispatchcontrol.UpdateDefaultUpdatedAt()
		dcuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dcuo *DispatchControlUpdateOne) check() error {
	if v, ok := dcuo.mutation.RecordServiceIncident(); ok {
		if err := dispatchcontrol.RecordServiceIncidentValidator(v); err != nil {
			return &ValidationError{Name: "record_service_incident", err: fmt.Errorf(`ent: validator failed for field "DispatchControl.record_service_incident": %w`, err)}
		}
	}
	if v, ok := dcuo.mutation.MaxShipmentWeightLimit(); ok {
		if err := dispatchcontrol.MaxShipmentWeightLimitValidator(v); err != nil {
			return &ValidationError{Name: "max_shipment_weight_limit", err: fmt.Errorf(`ent: validator failed for field "DispatchControl.max_shipment_weight_limit": %w`, err)}
		}
	}
	if _, ok := dcuo.mutation.OrganizationID(); dcuo.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "DispatchControl.organization"`)
	}
	if _, ok := dcuo.mutation.BusinessUnitID(); dcuo.mutation.BusinessUnitCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "DispatchControl.business_unit"`)
	}
	return nil
}

func (dcuo *DispatchControlUpdateOne) sqlSave(ctx context.Context) (_node *DispatchControl, err error) {
	if err := dcuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(dispatchcontrol.Table, dispatchcontrol.Columns, sqlgraph.NewFieldSpec(dispatchcontrol.FieldID, field.TypeUUID))
	id, ok := dcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "DispatchControl.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := dcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dispatchcontrol.FieldID)
		for _, f := range fields {
			if !dispatchcontrol.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != dispatchcontrol.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := dcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dcuo.mutation.UpdatedAt(); ok {
		_spec.SetField(dispatchcontrol.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := dcuo.mutation.RecordServiceIncident(); ok {
		_spec.SetField(dispatchcontrol.FieldRecordServiceIncident, field.TypeEnum, value)
	}
	if value, ok := dcuo.mutation.DeadheadTarget(); ok {
		_spec.SetField(dispatchcontrol.FieldDeadheadTarget, field.TypeFloat64, value)
	}
	if value, ok := dcuo.mutation.AddedDeadheadTarget(); ok {
		_spec.AddField(dispatchcontrol.FieldDeadheadTarget, field.TypeFloat64, value)
	}
	if value, ok := dcuo.mutation.MaxShipmentWeightLimit(); ok {
		_spec.SetField(dispatchcontrol.FieldMaxShipmentWeightLimit, field.TypeInt, value)
	}
	if value, ok := dcuo.mutation.AddedMaxShipmentWeightLimit(); ok {
		_spec.AddField(dispatchcontrol.FieldMaxShipmentWeightLimit, field.TypeInt, value)
	}
	if value, ok := dcuo.mutation.GracePeriod(); ok {
		_spec.SetField(dispatchcontrol.FieldGracePeriod, field.TypeUint8, value)
	}
	if value, ok := dcuo.mutation.AddedGracePeriod(); ok {
		_spec.AddField(dispatchcontrol.FieldGracePeriod, field.TypeUint8, value)
	}
	if value, ok := dcuo.mutation.EnforceWorkerAssign(); ok {
		_spec.SetField(dispatchcontrol.FieldEnforceWorkerAssign, field.TypeBool, value)
	}
	if value, ok := dcuo.mutation.TrailerContinuity(); ok {
		_spec.SetField(dispatchcontrol.FieldTrailerContinuity, field.TypeBool, value)
	}
	if value, ok := dcuo.mutation.DupeTrailerCheck(); ok {
		_spec.SetField(dispatchcontrol.FieldDupeTrailerCheck, field.TypeBool, value)
	}
	if value, ok := dcuo.mutation.MaintenanceCompliance(); ok {
		_spec.SetField(dispatchcontrol.FieldMaintenanceCompliance, field.TypeBool, value)
	}
	if value, ok := dcuo.mutation.RegulatoryCheck(); ok {
		_spec.SetField(dispatchcontrol.FieldRegulatoryCheck, field.TypeBool, value)
	}
	if value, ok := dcuo.mutation.PrevShipmentOnHold(); ok {
		_spec.SetField(dispatchcontrol.FieldPrevShipmentOnHold, field.TypeBool, value)
	}
	if value, ok := dcuo.mutation.WorkerTimeAwayRestriction(); ok {
		_spec.SetField(dispatchcontrol.FieldWorkerTimeAwayRestriction, field.TypeBool, value)
	}
	if value, ok := dcuo.mutation.TractorWorkerFleetConstraint(); ok {
		_spec.SetField(dispatchcontrol.FieldTractorWorkerFleetConstraint, field.TypeBool, value)
	}
	if dcuo.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   dispatchcontrol.OrganizationTable,
			Columns: []string{dispatchcontrol.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dcuo.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   dispatchcontrol.OrganizationTable,
			Columns: []string{dispatchcontrol.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if dcuo.mutation.BusinessUnitCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   dispatchcontrol.BusinessUnitTable,
			Columns: []string{dispatchcontrol.BusinessUnitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(businessunit.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dcuo.mutation.BusinessUnitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   dispatchcontrol.BusinessUnitTable,
			Columns: []string{dispatchcontrol.BusinessUnitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(businessunit.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &DispatchControl{config: dcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, dcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dispatchcontrol.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	dcuo.mutation.done = true
	return _node, nil
}
