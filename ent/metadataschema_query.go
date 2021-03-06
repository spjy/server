// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/responserms/server/ent/metadata"
	"github.com/responserms/server/ent/metadataschema"
	"github.com/responserms/server/ent/predicate"
)

// MetadataSchemaQuery is the builder for querying MetadataSchema entities.
type MetadataSchemaQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.MetadataSchema
	// eager-loading edges.
	withMetadata *MetadataQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (msq *MetadataSchemaQuery) Where(ps ...predicate.MetadataSchema) *MetadataSchemaQuery {
	msq.predicates = append(msq.predicates, ps...)
	return msq
}

// Limit adds a limit step to the query.
func (msq *MetadataSchemaQuery) Limit(limit int) *MetadataSchemaQuery {
	msq.limit = &limit
	return msq
}

// Offset adds an offset step to the query.
func (msq *MetadataSchemaQuery) Offset(offset int) *MetadataSchemaQuery {
	msq.offset = &offset
	return msq
}

// Order adds an order step to the query.
func (msq *MetadataSchemaQuery) Order(o ...OrderFunc) *MetadataSchemaQuery {
	msq.order = append(msq.order, o...)
	return msq
}

// QueryMetadata chains the current query on the metadata edge.
func (msq *MetadataSchemaQuery) QueryMetadata() *MetadataQuery {
	query := &MetadataQuery{config: msq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := msq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := msq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(metadataschema.Table, metadataschema.FieldID, selector),
			sqlgraph.To(metadata.Table, metadata.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, metadataschema.MetadataTable, metadataschema.MetadataColumn),
		)
		fromU = sqlgraph.SetNeighbors(msq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MetadataSchema entity in the query. Returns *NotFoundError when no metadataschema was found.
func (msq *MetadataSchemaQuery) First(ctx context.Context) (*MetadataSchema, error) {
	nodes, err := msq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{metadataschema.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (msq *MetadataSchemaQuery) FirstX(ctx context.Context) *MetadataSchema {
	node, err := msq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MetadataSchema id in the query. Returns *NotFoundError when no id was found.
func (msq *MetadataSchemaQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = msq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{metadataschema.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (msq *MetadataSchemaQuery) FirstIDX(ctx context.Context) int {
	id, err := msq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only MetadataSchema entity in the query, returns an error if not exactly one entity was returned.
func (msq *MetadataSchemaQuery) Only(ctx context.Context) (*MetadataSchema, error) {
	nodes, err := msq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{metadataschema.Label}
	default:
		return nil, &NotSingularError{metadataschema.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (msq *MetadataSchemaQuery) OnlyX(ctx context.Context) *MetadataSchema {
	node, err := msq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID returns the only MetadataSchema id in the query, returns an error if not exactly one id was returned.
func (msq *MetadataSchemaQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = msq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{metadataschema.Label}
	default:
		err = &NotSingularError{metadataschema.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (msq *MetadataSchemaQuery) OnlyIDX(ctx context.Context) int {
	id, err := msq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MetadataSchemas.
func (msq *MetadataSchemaQuery) All(ctx context.Context) ([]*MetadataSchema, error) {
	if err := msq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return msq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (msq *MetadataSchemaQuery) AllX(ctx context.Context) []*MetadataSchema {
	nodes, err := msq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MetadataSchema ids.
func (msq *MetadataSchemaQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := msq.Select(metadataschema.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (msq *MetadataSchemaQuery) IDsX(ctx context.Context) []int {
	ids, err := msq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (msq *MetadataSchemaQuery) Count(ctx context.Context) (int, error) {
	if err := msq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return msq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (msq *MetadataSchemaQuery) CountX(ctx context.Context) int {
	count, err := msq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (msq *MetadataSchemaQuery) Exist(ctx context.Context) (bool, error) {
	if err := msq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return msq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (msq *MetadataSchemaQuery) ExistX(ctx context.Context) bool {
	exist, err := msq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (msq *MetadataSchemaQuery) Clone() *MetadataSchemaQuery {
	if msq == nil {
		return nil
	}
	return &MetadataSchemaQuery{
		config:       msq.config,
		limit:        msq.limit,
		offset:       msq.offset,
		order:        append([]OrderFunc{}, msq.order...),
		predicates:   append([]predicate.MetadataSchema{}, msq.predicates...),
		withMetadata: msq.withMetadata.Clone(),
		// clone intermediate query.
		sql:  msq.sql.Clone(),
		path: msq.path,
	}
}

//  WithMetadata tells the query-builder to eager-loads the nodes that are connected to
// the "metadata" edge. The optional arguments used to configure the query builder of the edge.
func (msq *MetadataSchemaQuery) WithMetadata(opts ...func(*MetadataQuery)) *MetadataSchemaQuery {
	query := &MetadataQuery{config: msq.config}
	for _, opt := range opts {
		opt(query)
	}
	msq.withMetadata = query
	return msq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.MetadataSchema.Query().
//		GroupBy(metadataschema.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (msq *MetadataSchemaQuery) GroupBy(field string, fields ...string) *MetadataSchemaGroupBy {
	group := &MetadataSchemaGroupBy{config: msq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := msq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return msq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.MetadataSchema.Query().
//		Select(metadataschema.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (msq *MetadataSchemaQuery) Select(field string, fields ...string) *MetadataSchemaSelect {
	msq.fields = append([]string{field}, fields...)
	return &MetadataSchemaSelect{MetadataSchemaQuery: msq}
}

func (msq *MetadataSchemaQuery) prepareQuery(ctx context.Context) error {
	for _, f := range msq.fields {
		if !metadataschema.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if msq.path != nil {
		prev, err := msq.path(ctx)
		if err != nil {
			return err
		}
		msq.sql = prev
	}
	return nil
}

func (msq *MetadataSchemaQuery) sqlAll(ctx context.Context) ([]*MetadataSchema, error) {
	var (
		nodes       = []*MetadataSchema{}
		_spec       = msq.querySpec()
		loadedTypes = [1]bool{
			msq.withMetadata != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &MetadataSchema{config: msq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, msq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := msq.withMetadata; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*MetadataSchema)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Metadata = []*Metadata{}
		}
		query.withFKs = true
		query.Where(predicate.Metadata(func(s *sql.Selector) {
			s.Where(sql.InValues(metadataschema.MetadataColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.metadata_schema
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "metadata_schema" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "metadata_schema" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Metadata = append(node.Edges.Metadata, n)
		}
	}

	return nodes, nil
}

func (msq *MetadataSchemaQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := msq.querySpec()
	return sqlgraph.CountNodes(ctx, msq.driver, _spec)
}

func (msq *MetadataSchemaQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := msq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (msq *MetadataSchemaQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   metadataschema.Table,
			Columns: metadataschema.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: metadataschema.FieldID,
			},
		},
		From:   msq.sql,
		Unique: true,
	}
	if fields := msq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, metadataschema.FieldID)
		for i := range fields {
			if fields[i] != metadataschema.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := msq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := msq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := msq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := msq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, metadataschema.ValidColumn)
			}
		}
	}
	return _spec
}

func (msq *MetadataSchemaQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(msq.driver.Dialect())
	t1 := builder.Table(metadataschema.Table)
	selector := builder.Select(t1.Columns(metadataschema.Columns...)...).From(t1)
	if msq.sql != nil {
		selector = msq.sql
		selector.Select(selector.Columns(metadataschema.Columns...)...)
	}
	for _, p := range msq.predicates {
		p(selector)
	}
	for _, p := range msq.order {
		p(selector, metadataschema.ValidColumn)
	}
	if offset := msq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := msq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MetadataSchemaGroupBy is the builder for group-by MetadataSchema entities.
type MetadataSchemaGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (msgb *MetadataSchemaGroupBy) Aggregate(fns ...AggregateFunc) *MetadataSchemaGroupBy {
	msgb.fns = append(msgb.fns, fns...)
	return msgb
}

// Scan applies the group-by query and scan the result into the given value.
func (msgb *MetadataSchemaGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := msgb.path(ctx)
	if err != nil {
		return err
	}
	msgb.sql = query
	return msgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (msgb *MetadataSchemaGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := msgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (msgb *MetadataSchemaGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(msgb.fields) > 1 {
		return nil, errors.New("ent: MetadataSchemaGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := msgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (msgb *MetadataSchemaGroupBy) StringsX(ctx context.Context) []string {
	v, err := msgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (msgb *MetadataSchemaGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = msgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{metadataschema.Label}
	default:
		err = fmt.Errorf("ent: MetadataSchemaGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (msgb *MetadataSchemaGroupBy) StringX(ctx context.Context) string {
	v, err := msgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (msgb *MetadataSchemaGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(msgb.fields) > 1 {
		return nil, errors.New("ent: MetadataSchemaGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := msgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (msgb *MetadataSchemaGroupBy) IntsX(ctx context.Context) []int {
	v, err := msgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (msgb *MetadataSchemaGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = msgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{metadataschema.Label}
	default:
		err = fmt.Errorf("ent: MetadataSchemaGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (msgb *MetadataSchemaGroupBy) IntX(ctx context.Context) int {
	v, err := msgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (msgb *MetadataSchemaGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(msgb.fields) > 1 {
		return nil, errors.New("ent: MetadataSchemaGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := msgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (msgb *MetadataSchemaGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := msgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (msgb *MetadataSchemaGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = msgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{metadataschema.Label}
	default:
		err = fmt.Errorf("ent: MetadataSchemaGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (msgb *MetadataSchemaGroupBy) Float64X(ctx context.Context) float64 {
	v, err := msgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (msgb *MetadataSchemaGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(msgb.fields) > 1 {
		return nil, errors.New("ent: MetadataSchemaGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := msgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (msgb *MetadataSchemaGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := msgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (msgb *MetadataSchemaGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = msgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{metadataschema.Label}
	default:
		err = fmt.Errorf("ent: MetadataSchemaGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (msgb *MetadataSchemaGroupBy) BoolX(ctx context.Context) bool {
	v, err := msgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (msgb *MetadataSchemaGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range msgb.fields {
		if !metadataschema.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := msgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := msgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (msgb *MetadataSchemaGroupBy) sqlQuery() *sql.Selector {
	selector := msgb.sql
	columns := make([]string, 0, len(msgb.fields)+len(msgb.fns))
	columns = append(columns, msgb.fields...)
	for _, fn := range msgb.fns {
		columns = append(columns, fn(selector, metadataschema.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(msgb.fields...)
}

// MetadataSchemaSelect is the builder for select fields of MetadataSchema entities.
type MetadataSchemaSelect struct {
	*MetadataSchemaQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (mss *MetadataSchemaSelect) Scan(ctx context.Context, v interface{}) error {
	if err := mss.prepareQuery(ctx); err != nil {
		return err
	}
	mss.sql = mss.MetadataSchemaQuery.sqlQuery()
	return mss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mss *MetadataSchemaSelect) ScanX(ctx context.Context, v interface{}) {
	if err := mss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (mss *MetadataSchemaSelect) Strings(ctx context.Context) ([]string, error) {
	if len(mss.fields) > 1 {
		return nil, errors.New("ent: MetadataSchemaSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := mss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mss *MetadataSchemaSelect) StringsX(ctx context.Context) []string {
	v, err := mss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (mss *MetadataSchemaSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = mss.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{metadataschema.Label}
	default:
		err = fmt.Errorf("ent: MetadataSchemaSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (mss *MetadataSchemaSelect) StringX(ctx context.Context) string {
	v, err := mss.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (mss *MetadataSchemaSelect) Ints(ctx context.Context) ([]int, error) {
	if len(mss.fields) > 1 {
		return nil, errors.New("ent: MetadataSchemaSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := mss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mss *MetadataSchemaSelect) IntsX(ctx context.Context) []int {
	v, err := mss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (mss *MetadataSchemaSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = mss.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{metadataschema.Label}
	default:
		err = fmt.Errorf("ent: MetadataSchemaSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (mss *MetadataSchemaSelect) IntX(ctx context.Context) int {
	v, err := mss.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (mss *MetadataSchemaSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(mss.fields) > 1 {
		return nil, errors.New("ent: MetadataSchemaSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := mss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mss *MetadataSchemaSelect) Float64sX(ctx context.Context) []float64 {
	v, err := mss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (mss *MetadataSchemaSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = mss.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{metadataschema.Label}
	default:
		err = fmt.Errorf("ent: MetadataSchemaSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (mss *MetadataSchemaSelect) Float64X(ctx context.Context) float64 {
	v, err := mss.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (mss *MetadataSchemaSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(mss.fields) > 1 {
		return nil, errors.New("ent: MetadataSchemaSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := mss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mss *MetadataSchemaSelect) BoolsX(ctx context.Context) []bool {
	v, err := mss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (mss *MetadataSchemaSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = mss.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{metadataschema.Label}
	default:
		err = fmt.Errorf("ent: MetadataSchemaSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (mss *MetadataSchemaSelect) BoolX(ctx context.Context) bool {
	v, err := mss.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mss *MetadataSchemaSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := mss.sqlQuery().Query()
	if err := mss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mss *MetadataSchemaSelect) sqlQuery() sql.Querier {
	selector := mss.sql
	selector.Select(selector.Columns(mss.fields...)...)
	return selector
}
