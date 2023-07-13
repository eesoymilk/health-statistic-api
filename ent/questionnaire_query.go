// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/eesoymilk/health-statistic-api/ent/predicate"
	"github.com/eesoymilk/health-statistic-api/ent/questionnaire"
)

// QuestionnaireQuery is the builder for querying Questionnaire entities.
type QuestionnaireQuery struct {
	config
	ctx        *QueryContext
	order      []questionnaire.OrderOption
	inters     []Interceptor
	predicates []predicate.Questionnaire
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the QuestionnaireQuery builder.
func (qq *QuestionnaireQuery) Where(ps ...predicate.Questionnaire) *QuestionnaireQuery {
	qq.predicates = append(qq.predicates, ps...)
	return qq
}

// Limit the number of records to be returned by this query.
func (qq *QuestionnaireQuery) Limit(limit int) *QuestionnaireQuery {
	qq.ctx.Limit = &limit
	return qq
}

// Offset to start from.
func (qq *QuestionnaireQuery) Offset(offset int) *QuestionnaireQuery {
	qq.ctx.Offset = &offset
	return qq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (qq *QuestionnaireQuery) Unique(unique bool) *QuestionnaireQuery {
	qq.ctx.Unique = &unique
	return qq
}

// Order specifies how the records should be ordered.
func (qq *QuestionnaireQuery) Order(o ...questionnaire.OrderOption) *QuestionnaireQuery {
	qq.order = append(qq.order, o...)
	return qq
}

// First returns the first Questionnaire entity from the query.
// Returns a *NotFoundError when no Questionnaire was found.
func (qq *QuestionnaireQuery) First(ctx context.Context) (*Questionnaire, error) {
	nodes, err := qq.Limit(1).All(setContextOp(ctx, qq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{questionnaire.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (qq *QuestionnaireQuery) FirstX(ctx context.Context) *Questionnaire {
	node, err := qq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Questionnaire ID from the query.
// Returns a *NotFoundError when no Questionnaire ID was found.
func (qq *QuestionnaireQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = qq.Limit(1).IDs(setContextOp(ctx, qq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{questionnaire.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (qq *QuestionnaireQuery) FirstIDX(ctx context.Context) int {
	id, err := qq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Questionnaire entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Questionnaire entity is found.
// Returns a *NotFoundError when no Questionnaire entities are found.
func (qq *QuestionnaireQuery) Only(ctx context.Context) (*Questionnaire, error) {
	nodes, err := qq.Limit(2).All(setContextOp(ctx, qq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{questionnaire.Label}
	default:
		return nil, &NotSingularError{questionnaire.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (qq *QuestionnaireQuery) OnlyX(ctx context.Context) *Questionnaire {
	node, err := qq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Questionnaire ID in the query.
// Returns a *NotSingularError when more than one Questionnaire ID is found.
// Returns a *NotFoundError when no entities are found.
func (qq *QuestionnaireQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = qq.Limit(2).IDs(setContextOp(ctx, qq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{questionnaire.Label}
	default:
		err = &NotSingularError{questionnaire.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (qq *QuestionnaireQuery) OnlyIDX(ctx context.Context) int {
	id, err := qq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Questionnaires.
func (qq *QuestionnaireQuery) All(ctx context.Context) ([]*Questionnaire, error) {
	ctx = setContextOp(ctx, qq.ctx, "All")
	if err := qq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Questionnaire, *QuestionnaireQuery]()
	return withInterceptors[[]*Questionnaire](ctx, qq, qr, qq.inters)
}

// AllX is like All, but panics if an error occurs.
func (qq *QuestionnaireQuery) AllX(ctx context.Context) []*Questionnaire {
	nodes, err := qq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Questionnaire IDs.
func (qq *QuestionnaireQuery) IDs(ctx context.Context) (ids []int, err error) {
	if qq.ctx.Unique == nil && qq.path != nil {
		qq.Unique(true)
	}
	ctx = setContextOp(ctx, qq.ctx, "IDs")
	if err = qq.Select(questionnaire.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (qq *QuestionnaireQuery) IDsX(ctx context.Context) []int {
	ids, err := qq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (qq *QuestionnaireQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, qq.ctx, "Count")
	if err := qq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, qq, querierCount[*QuestionnaireQuery](), qq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (qq *QuestionnaireQuery) CountX(ctx context.Context) int {
	count, err := qq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (qq *QuestionnaireQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, qq.ctx, "Exist")
	switch _, err := qq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (qq *QuestionnaireQuery) ExistX(ctx context.Context) bool {
	exist, err := qq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the QuestionnaireQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (qq *QuestionnaireQuery) Clone() *QuestionnaireQuery {
	if qq == nil {
		return nil
	}
	return &QuestionnaireQuery{
		config:     qq.config,
		ctx:        qq.ctx.Clone(),
		order:      append([]questionnaire.OrderOption{}, qq.order...),
		inters:     append([]Interceptor{}, qq.inters...),
		predicates: append([]predicate.Questionnaire{}, qq.predicates...),
		// clone intermediate query.
		sql:  qq.sql.Clone(),
		path: qq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (qq *QuestionnaireQuery) GroupBy(field string, fields ...string) *QuestionnaireGroupBy {
	qq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &QuestionnaireGroupBy{build: qq}
	grbuild.flds = &qq.ctx.Fields
	grbuild.label = questionnaire.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (qq *QuestionnaireQuery) Select(fields ...string) *QuestionnaireSelect {
	qq.ctx.Fields = append(qq.ctx.Fields, fields...)
	sbuild := &QuestionnaireSelect{QuestionnaireQuery: qq}
	sbuild.label = questionnaire.Label
	sbuild.flds, sbuild.scan = &qq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a QuestionnaireSelect configured with the given aggregations.
func (qq *QuestionnaireQuery) Aggregate(fns ...AggregateFunc) *QuestionnaireSelect {
	return qq.Select().Aggregate(fns...)
}

func (qq *QuestionnaireQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range qq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, qq); err != nil {
				return err
			}
		}
	}
	for _, f := range qq.ctx.Fields {
		if !questionnaire.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if qq.path != nil {
		prev, err := qq.path(ctx)
		if err != nil {
			return err
		}
		qq.sql = prev
	}
	return nil
}

func (qq *QuestionnaireQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Questionnaire, error) {
	var (
		nodes = []*Questionnaire{}
		_spec = qq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Questionnaire).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Questionnaire{config: qq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, qq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (qq *QuestionnaireQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := qq.querySpec()
	_spec.Node.Columns = qq.ctx.Fields
	if len(qq.ctx.Fields) > 0 {
		_spec.Unique = qq.ctx.Unique != nil && *qq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, qq.driver, _spec)
}

func (qq *QuestionnaireQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(questionnaire.Table, questionnaire.Columns, sqlgraph.NewFieldSpec(questionnaire.FieldID, field.TypeInt))
	_spec.From = qq.sql
	if unique := qq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if qq.path != nil {
		_spec.Unique = true
	}
	if fields := qq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, questionnaire.FieldID)
		for i := range fields {
			if fields[i] != questionnaire.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := qq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := qq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := qq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := qq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (qq *QuestionnaireQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(qq.driver.Dialect())
	t1 := builder.Table(questionnaire.Table)
	columns := qq.ctx.Fields
	if len(columns) == 0 {
		columns = questionnaire.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if qq.sql != nil {
		selector = qq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if qq.ctx.Unique != nil && *qq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range qq.predicates {
		p(selector)
	}
	for _, p := range qq.order {
		p(selector)
	}
	if offset := qq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := qq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// QuestionnaireGroupBy is the group-by builder for Questionnaire entities.
type QuestionnaireGroupBy struct {
	selector
	build *QuestionnaireQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (qgb *QuestionnaireGroupBy) Aggregate(fns ...AggregateFunc) *QuestionnaireGroupBy {
	qgb.fns = append(qgb.fns, fns...)
	return qgb
}

// Scan applies the selector query and scans the result into the given value.
func (qgb *QuestionnaireGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, qgb.build.ctx, "GroupBy")
	if err := qgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*QuestionnaireQuery, *QuestionnaireGroupBy](ctx, qgb.build, qgb, qgb.build.inters, v)
}

func (qgb *QuestionnaireGroupBy) sqlScan(ctx context.Context, root *QuestionnaireQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(qgb.fns))
	for _, fn := range qgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*qgb.flds)+len(qgb.fns))
		for _, f := range *qgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*qgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := qgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// QuestionnaireSelect is the builder for selecting fields of Questionnaire entities.
type QuestionnaireSelect struct {
	*QuestionnaireQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (qs *QuestionnaireSelect) Aggregate(fns ...AggregateFunc) *QuestionnaireSelect {
	qs.fns = append(qs.fns, fns...)
	return qs
}

// Scan applies the selector query and scans the result into the given value.
func (qs *QuestionnaireSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, qs.ctx, "Select")
	if err := qs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*QuestionnaireQuery, *QuestionnaireSelect](ctx, qs.QuestionnaireQuery, qs, qs.inters, v)
}

func (qs *QuestionnaireSelect) sqlScan(ctx context.Context, root *QuestionnaireQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(qs.fns))
	for _, fn := range qs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*qs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := qs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
