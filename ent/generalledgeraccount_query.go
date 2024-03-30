// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/ent/businessunit"
	"github.com/emoss08/trenova/ent/generalledgeraccount"
	"github.com/emoss08/trenova/ent/organization"
	"github.com/emoss08/trenova/ent/predicate"
	"github.com/emoss08/trenova/ent/tag"
	"github.com/google/uuid"
)

// GeneralLedgerAccountQuery is the builder for querying GeneralLedgerAccount entities.
type GeneralLedgerAccountQuery struct {
	config
	ctx              *QueryContext
	order            []generalledgeraccount.OrderOption
	inters           []Interceptor
	predicates       []predicate.GeneralLedgerAccount
	withBusinessUnit *BusinessUnitQuery
	withOrganization *OrganizationQuery
	withTags         *TagQuery
	modifiers        []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GeneralLedgerAccountQuery builder.
func (glaq *GeneralLedgerAccountQuery) Where(ps ...predicate.GeneralLedgerAccount) *GeneralLedgerAccountQuery {
	glaq.predicates = append(glaq.predicates, ps...)
	return glaq
}

// Limit the number of records to be returned by this query.
func (glaq *GeneralLedgerAccountQuery) Limit(limit int) *GeneralLedgerAccountQuery {
	glaq.ctx.Limit = &limit
	return glaq
}

// Offset to start from.
func (glaq *GeneralLedgerAccountQuery) Offset(offset int) *GeneralLedgerAccountQuery {
	glaq.ctx.Offset = &offset
	return glaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (glaq *GeneralLedgerAccountQuery) Unique(unique bool) *GeneralLedgerAccountQuery {
	glaq.ctx.Unique = &unique
	return glaq
}

// Order specifies how the records should be ordered.
func (glaq *GeneralLedgerAccountQuery) Order(o ...generalledgeraccount.OrderOption) *GeneralLedgerAccountQuery {
	glaq.order = append(glaq.order, o...)
	return glaq
}

// QueryBusinessUnit chains the current query on the "business_unit" edge.
func (glaq *GeneralLedgerAccountQuery) QueryBusinessUnit() *BusinessUnitQuery {
	query := (&BusinessUnitClient{config: glaq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := glaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := glaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(generalledgeraccount.Table, generalledgeraccount.FieldID, selector),
			sqlgraph.To(businessunit.Table, businessunit.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, generalledgeraccount.BusinessUnitTable, generalledgeraccount.BusinessUnitColumn),
		)
		fromU = sqlgraph.SetNeighbors(glaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrganization chains the current query on the "organization" edge.
func (glaq *GeneralLedgerAccountQuery) QueryOrganization() *OrganizationQuery {
	query := (&OrganizationClient{config: glaq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := glaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := glaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(generalledgeraccount.Table, generalledgeraccount.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, generalledgeraccount.OrganizationTable, generalledgeraccount.OrganizationColumn),
		)
		fromU = sqlgraph.SetNeighbors(glaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTags chains the current query on the "tags" edge.
func (glaq *GeneralLedgerAccountQuery) QueryTags() *TagQuery {
	query := (&TagClient{config: glaq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := glaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := glaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(generalledgeraccount.Table, generalledgeraccount.FieldID, selector),
			sqlgraph.To(tag.Table, tag.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, generalledgeraccount.TagsTable, generalledgeraccount.TagsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(glaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first GeneralLedgerAccount entity from the query.
// Returns a *NotFoundError when no GeneralLedgerAccount was found.
func (glaq *GeneralLedgerAccountQuery) First(ctx context.Context) (*GeneralLedgerAccount, error) {
	nodes, err := glaq.Limit(1).All(setContextOp(ctx, glaq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{generalledgeraccount.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (glaq *GeneralLedgerAccountQuery) FirstX(ctx context.Context) *GeneralLedgerAccount {
	node, err := glaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GeneralLedgerAccount ID from the query.
// Returns a *NotFoundError when no GeneralLedgerAccount ID was found.
func (glaq *GeneralLedgerAccountQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = glaq.Limit(1).IDs(setContextOp(ctx, glaq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{generalledgeraccount.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (glaq *GeneralLedgerAccountQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := glaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GeneralLedgerAccount entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GeneralLedgerAccount entity is found.
// Returns a *NotFoundError when no GeneralLedgerAccount entities are found.
func (glaq *GeneralLedgerAccountQuery) Only(ctx context.Context) (*GeneralLedgerAccount, error) {
	nodes, err := glaq.Limit(2).All(setContextOp(ctx, glaq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{generalledgeraccount.Label}
	default:
		return nil, &NotSingularError{generalledgeraccount.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (glaq *GeneralLedgerAccountQuery) OnlyX(ctx context.Context) *GeneralLedgerAccount {
	node, err := glaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GeneralLedgerAccount ID in the query.
// Returns a *NotSingularError when more than one GeneralLedgerAccount ID is found.
// Returns a *NotFoundError when no entities are found.
func (glaq *GeneralLedgerAccountQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = glaq.Limit(2).IDs(setContextOp(ctx, glaq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{generalledgeraccount.Label}
	default:
		err = &NotSingularError{generalledgeraccount.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (glaq *GeneralLedgerAccountQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := glaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GeneralLedgerAccounts.
func (glaq *GeneralLedgerAccountQuery) All(ctx context.Context) ([]*GeneralLedgerAccount, error) {
	ctx = setContextOp(ctx, glaq.ctx, "All")
	if err := glaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*GeneralLedgerAccount, *GeneralLedgerAccountQuery]()
	return withInterceptors[[]*GeneralLedgerAccount](ctx, glaq, qr, glaq.inters)
}

// AllX is like All, but panics if an error occurs.
func (glaq *GeneralLedgerAccountQuery) AllX(ctx context.Context) []*GeneralLedgerAccount {
	nodes, err := glaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GeneralLedgerAccount IDs.
func (glaq *GeneralLedgerAccountQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if glaq.ctx.Unique == nil && glaq.path != nil {
		glaq.Unique(true)
	}
	ctx = setContextOp(ctx, glaq.ctx, "IDs")
	if err = glaq.Select(generalledgeraccount.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (glaq *GeneralLedgerAccountQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := glaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (glaq *GeneralLedgerAccountQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, glaq.ctx, "Count")
	if err := glaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, glaq, querierCount[*GeneralLedgerAccountQuery](), glaq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (glaq *GeneralLedgerAccountQuery) CountX(ctx context.Context) int {
	count, err := glaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (glaq *GeneralLedgerAccountQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, glaq.ctx, "Exist")
	switch _, err := glaq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (glaq *GeneralLedgerAccountQuery) ExistX(ctx context.Context) bool {
	exist, err := glaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GeneralLedgerAccountQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (glaq *GeneralLedgerAccountQuery) Clone() *GeneralLedgerAccountQuery {
	if glaq == nil {
		return nil
	}
	return &GeneralLedgerAccountQuery{
		config:           glaq.config,
		ctx:              glaq.ctx.Clone(),
		order:            append([]generalledgeraccount.OrderOption{}, glaq.order...),
		inters:           append([]Interceptor{}, glaq.inters...),
		predicates:       append([]predicate.GeneralLedgerAccount{}, glaq.predicates...),
		withBusinessUnit: glaq.withBusinessUnit.Clone(),
		withOrganization: glaq.withOrganization.Clone(),
		withTags:         glaq.withTags.Clone(),
		// clone intermediate query.
		sql:  glaq.sql.Clone(),
		path: glaq.path,
	}
}

// WithBusinessUnit tells the query-builder to eager-load the nodes that are connected to
// the "business_unit" edge. The optional arguments are used to configure the query builder of the edge.
func (glaq *GeneralLedgerAccountQuery) WithBusinessUnit(opts ...func(*BusinessUnitQuery)) *GeneralLedgerAccountQuery {
	query := (&BusinessUnitClient{config: glaq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	glaq.withBusinessUnit = query
	return glaq
}

// WithOrganization tells the query-builder to eager-load the nodes that are connected to
// the "organization" edge. The optional arguments are used to configure the query builder of the edge.
func (glaq *GeneralLedgerAccountQuery) WithOrganization(opts ...func(*OrganizationQuery)) *GeneralLedgerAccountQuery {
	query := (&OrganizationClient{config: glaq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	glaq.withOrganization = query
	return glaq
}

// WithTags tells the query-builder to eager-load the nodes that are connected to
// the "tags" edge. The optional arguments are used to configure the query builder of the edge.
func (glaq *GeneralLedgerAccountQuery) WithTags(opts ...func(*TagQuery)) *GeneralLedgerAccountQuery {
	query := (&TagClient{config: glaq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	glaq.withTags = query
	return glaq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		BusinessUnitID uuid.UUID `json:"businessUnitId"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GeneralLedgerAccount.Query().
//		GroupBy(generalledgeraccount.FieldBusinessUnitID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (glaq *GeneralLedgerAccountQuery) GroupBy(field string, fields ...string) *GeneralLedgerAccountGroupBy {
	glaq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &GeneralLedgerAccountGroupBy{build: glaq}
	grbuild.flds = &glaq.ctx.Fields
	grbuild.label = generalledgeraccount.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		BusinessUnitID uuid.UUID `json:"businessUnitId"`
//	}
//
//	client.GeneralLedgerAccount.Query().
//		Select(generalledgeraccount.FieldBusinessUnitID).
//		Scan(ctx, &v)
func (glaq *GeneralLedgerAccountQuery) Select(fields ...string) *GeneralLedgerAccountSelect {
	glaq.ctx.Fields = append(glaq.ctx.Fields, fields...)
	sbuild := &GeneralLedgerAccountSelect{GeneralLedgerAccountQuery: glaq}
	sbuild.label = generalledgeraccount.Label
	sbuild.flds, sbuild.scan = &glaq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a GeneralLedgerAccountSelect configured with the given aggregations.
func (glaq *GeneralLedgerAccountQuery) Aggregate(fns ...AggregateFunc) *GeneralLedgerAccountSelect {
	return glaq.Select().Aggregate(fns...)
}

func (glaq *GeneralLedgerAccountQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range glaq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, glaq); err != nil {
				return err
			}
		}
	}
	for _, f := range glaq.ctx.Fields {
		if !generalledgeraccount.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if glaq.path != nil {
		prev, err := glaq.path(ctx)
		if err != nil {
			return err
		}
		glaq.sql = prev
	}
	return nil
}

func (glaq *GeneralLedgerAccountQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*GeneralLedgerAccount, error) {
	var (
		nodes       = []*GeneralLedgerAccount{}
		_spec       = glaq.querySpec()
		loadedTypes = [3]bool{
			glaq.withBusinessUnit != nil,
			glaq.withOrganization != nil,
			glaq.withTags != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*GeneralLedgerAccount).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &GeneralLedgerAccount{config: glaq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(glaq.modifiers) > 0 {
		_spec.Modifiers = glaq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, glaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := glaq.withBusinessUnit; query != nil {
		if err := glaq.loadBusinessUnit(ctx, query, nodes, nil,
			func(n *GeneralLedgerAccount, e *BusinessUnit) { n.Edges.BusinessUnit = e }); err != nil {
			return nil, err
		}
	}
	if query := glaq.withOrganization; query != nil {
		if err := glaq.loadOrganization(ctx, query, nodes, nil,
			func(n *GeneralLedgerAccount, e *Organization) { n.Edges.Organization = e }); err != nil {
			return nil, err
		}
	}
	if query := glaq.withTags; query != nil {
		if err := glaq.loadTags(ctx, query, nodes,
			func(n *GeneralLedgerAccount) { n.Edges.Tags = []*Tag{} },
			func(n *GeneralLedgerAccount, e *Tag) { n.Edges.Tags = append(n.Edges.Tags, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (glaq *GeneralLedgerAccountQuery) loadBusinessUnit(ctx context.Context, query *BusinessUnitQuery, nodes []*GeneralLedgerAccount, init func(*GeneralLedgerAccount), assign func(*GeneralLedgerAccount, *BusinessUnit)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*GeneralLedgerAccount)
	for i := range nodes {
		fk := nodes[i].BusinessUnitID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(businessunit.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "business_unit_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (glaq *GeneralLedgerAccountQuery) loadOrganization(ctx context.Context, query *OrganizationQuery, nodes []*GeneralLedgerAccount, init func(*GeneralLedgerAccount), assign func(*GeneralLedgerAccount, *Organization)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*GeneralLedgerAccount)
	for i := range nodes {
		fk := nodes[i].OrganizationID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(organization.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "organization_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (glaq *GeneralLedgerAccountQuery) loadTags(ctx context.Context, query *TagQuery, nodes []*GeneralLedgerAccount, init func(*GeneralLedgerAccount), assign func(*GeneralLedgerAccount, *Tag)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*GeneralLedgerAccount)
	nids := make(map[uuid.UUID]map[*GeneralLedgerAccount]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(generalledgeraccount.TagsTable)
		s.Join(joinT).On(s.C(tag.FieldID), joinT.C(generalledgeraccount.TagsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(generalledgeraccount.TagsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(generalledgeraccount.TagsPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(uuid.UUID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := *values[0].(*uuid.UUID)
				inValue := *values[1].(*uuid.UUID)
				if nids[inValue] == nil {
					nids[inValue] = map[*GeneralLedgerAccount]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Tag](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "tags" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (glaq *GeneralLedgerAccountQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := glaq.querySpec()
	if len(glaq.modifiers) > 0 {
		_spec.Modifiers = glaq.modifiers
	}
	_spec.Node.Columns = glaq.ctx.Fields
	if len(glaq.ctx.Fields) > 0 {
		_spec.Unique = glaq.ctx.Unique != nil && *glaq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, glaq.driver, _spec)
}

func (glaq *GeneralLedgerAccountQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(generalledgeraccount.Table, generalledgeraccount.Columns, sqlgraph.NewFieldSpec(generalledgeraccount.FieldID, field.TypeUUID))
	_spec.From = glaq.sql
	if unique := glaq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if glaq.path != nil {
		_spec.Unique = true
	}
	if fields := glaq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, generalledgeraccount.FieldID)
		for i := range fields {
			if fields[i] != generalledgeraccount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if glaq.withBusinessUnit != nil {
			_spec.Node.AddColumnOnce(generalledgeraccount.FieldBusinessUnitID)
		}
		if glaq.withOrganization != nil {
			_spec.Node.AddColumnOnce(generalledgeraccount.FieldOrganizationID)
		}
	}
	if ps := glaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := glaq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := glaq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := glaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (glaq *GeneralLedgerAccountQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(glaq.driver.Dialect())
	t1 := builder.Table(generalledgeraccount.Table)
	columns := glaq.ctx.Fields
	if len(columns) == 0 {
		columns = generalledgeraccount.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if glaq.sql != nil {
		selector = glaq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if glaq.ctx.Unique != nil && *glaq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range glaq.modifiers {
		m(selector)
	}
	for _, p := range glaq.predicates {
		p(selector)
	}
	for _, p := range glaq.order {
		p(selector)
	}
	if offset := glaq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := glaq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (glaq *GeneralLedgerAccountQuery) Modify(modifiers ...func(s *sql.Selector)) *GeneralLedgerAccountSelect {
	glaq.modifiers = append(glaq.modifiers, modifiers...)
	return glaq.Select()
}

// GeneralLedgerAccountGroupBy is the group-by builder for GeneralLedgerAccount entities.
type GeneralLedgerAccountGroupBy struct {
	selector
	build *GeneralLedgerAccountQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (glagb *GeneralLedgerAccountGroupBy) Aggregate(fns ...AggregateFunc) *GeneralLedgerAccountGroupBy {
	glagb.fns = append(glagb.fns, fns...)
	return glagb
}

// Scan applies the selector query and scans the result into the given value.
func (glagb *GeneralLedgerAccountGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, glagb.build.ctx, "GroupBy")
	if err := glagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GeneralLedgerAccountQuery, *GeneralLedgerAccountGroupBy](ctx, glagb.build, glagb, glagb.build.inters, v)
}

func (glagb *GeneralLedgerAccountGroupBy) sqlScan(ctx context.Context, root *GeneralLedgerAccountQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(glagb.fns))
	for _, fn := range glagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*glagb.flds)+len(glagb.fns))
		for _, f := range *glagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*glagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := glagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// GeneralLedgerAccountSelect is the builder for selecting fields of GeneralLedgerAccount entities.
type GeneralLedgerAccountSelect struct {
	*GeneralLedgerAccountQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (glas *GeneralLedgerAccountSelect) Aggregate(fns ...AggregateFunc) *GeneralLedgerAccountSelect {
	glas.fns = append(glas.fns, fns...)
	return glas
}

// Scan applies the selector query and scans the result into the given value.
func (glas *GeneralLedgerAccountSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, glas.ctx, "Select")
	if err := glas.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GeneralLedgerAccountQuery, *GeneralLedgerAccountSelect](ctx, glas.GeneralLedgerAccountQuery, glas, glas.inters, v)
}

func (glas *GeneralLedgerAccountSelect) sqlScan(ctx context.Context, root *GeneralLedgerAccountQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(glas.fns))
	for _, fn := range glas.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*glas.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := glas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (glas *GeneralLedgerAccountSelect) Modify(modifiers ...func(s *sql.Selector)) *GeneralLedgerAccountSelect {
	glas.modifiers = append(glas.modifiers, modifiers...)
	return glas
}
