package gen

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/ast"
)

// GeneratedQueryResolver ...
type GeneratedQueryResolver struct{ *GeneratedResolver }

// QueryExportHandlerOptions ...
type QueryExportHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *ExportFilterType
}

// Export ...
func (r *GeneratedQueryResolver) Export(ctx context.Context, id *string, q *string, filter *ExportFilterType) (*Export, error) {
	opts := QueryExportHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryExport(ctx, r.GeneratedResolver, opts)
}

// QueryExportHandler ...
func QueryExportHandler(ctx context.Context, r *GeneratedResolver, opts QueryExportHandlerOptions) (*Export, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := ExportQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &ExportResultType{
		EntityResultType: EntityResultType{
			Offset:       &offset,
			Limit:        &limit,
			Query:        &query,
			Filter:       opts.Filter,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.GetDB(ctx)
	if opts.ID != nil {
		qb = qb.Where(TableName("exports")+".id = ?", *opts.ID)
	}

	var items []*Export
	giOpts := GetItemsOptions{
		Alias:      TableName("exports"),
		Preloaders: []string{},
	}
	err := rt.GetItems(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, nil
	}
	return items[0], err
}

// QueryExportsHandlerOptions ...
type QueryExportsHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*ExportSortType
	Filter *ExportFilterType
}

// Exports ...
func (r *GeneratedQueryResolver) Exports(ctx context.Context, offset *int, limit *int, q *string, sort []*ExportSortType, filter *ExportFilterType) (*ExportResultType, error) {
	opts := QueryExportsHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryExports(ctx, r.GeneratedResolver, opts)
}

// ExportsItems ...
func (r *GeneratedResolver) ExportsItems(ctx context.Context, opts QueryExportsHandlerOptions) (res []*Export, err error) {
	resultType, err := r.Handlers.QueryExports(ctx, r, opts)
	if err != nil {
		return
	}
	err = resultType.GetItems(ctx, r.GetDB(ctx), GetItemsOptions{
		Alias: TableName("exports"),
	}, &res)
	if err != nil {
		return
	}
	return
}

// ExportsCount ...
func (r *GeneratedResolver) ExportsCount(ctx context.Context, opts QueryExportsHandlerOptions) (count int, err error) {
	resultType, err := r.Handlers.QueryExports(ctx, r, opts)
	if err != nil {
		return
	}
	return resultType.GetCount(ctx, r.GetDB(ctx), GetItemsOptions{
		Alias: TableName("exports"),
	}, &Export{})
}

// QueryExportsHandler ...
func QueryExportsHandler(ctx context.Context, r *GeneratedResolver, opts QueryExportsHandlerOptions) (*ExportResultType, error) {
	query := ExportQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	_sort := []EntitySort{}
	for _, sort := range opts.Sort {
		_sort = append(_sort, sort)
	}

	return &ExportResultType{
		EntityResultType: EntityResultType{
			Offset:       opts.Offset,
			Limit:        opts.Limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

// GeneratedExportResultTypeResolver ...
type GeneratedExportResultTypeResolver struct{ *GeneratedResolver }

// Items ...
func (r *GeneratedExportResultTypeResolver) Items(ctx context.Context, obj *ExportResultType) (items []*Export, err error) {
	opts := GetItemsOptions{
		Alias:      TableName("exports"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.GetDB(ctx), opts, &items)

	uniqueItems := []*Export{}
	idMap := map[string]bool{}
	for _, item := range items {
		if _, ok := idMap[item.ID]; !ok {
			idMap[item.ID] = true
			uniqueItems = append(uniqueItems, item)
		}
	}
	items = uniqueItems
	return
}

// Count ...
func (r *GeneratedExportResultTypeResolver) Count(ctx context.Context, obj *ExportResultType) (count int, err error) {
	opts := GetItemsOptions{
		Alias:      TableName("exports"),
		Preloaders: []string{},
	}
	return obj.GetCount(ctx, r.GetDB(ctx), opts, &Export{})
}

// Aggregations ...
func (r *GeneratedExportResultTypeResolver) Aggregations(ctx context.Context, obj *ExportResultType) (res *ExportResultAggregations, err error) {
	aggregationsMap := map[string]GetAggregationsAggregationField{

		"errorDescriptionMax": GetAggregationsAggregationField{"errorDescription", "Max"},
		"errorDescriptionMin": GetAggregationsAggregationField{"errorDescription", "Min"},
		"fileIdMax":           GetAggregationsAggregationField{"fileId", "Max"},
		"fileIdMin":           GetAggregationsAggregationField{"fileId", "Min"},
		"metadataMax":         GetAggregationsAggregationField{"metadata", "Max"},
		"metadataMin":         GetAggregationsAggregationField{"metadata", "Min"},
		"progressAvg":         GetAggregationsAggregationField{"progress", "Avg"},
		"progressMax":         GetAggregationsAggregationField{"progress", "Max"},
		"progressMin":         GetAggregationsAggregationField{"progress", "Min"},
		"progressSum":         GetAggregationsAggregationField{"progress", "Sum"},
		"typeMax":             GetAggregationsAggregationField{"type", "Max"},
		"typeMin":             GetAggregationsAggregationField{"type", "Min"},
	}
	aggregationFieldsMap := map[string]string{

		"errorDescriptionMax": "errorDescription",
		"errorDescriptionMin": "errorDescription",
		"fileIdMax":           "fileId",
		"fileIdMin":           "fileId",
		"metadataMax":         "metadata",
		"metadataMin":         "metadata",
		"progressAvg":         "progress",
		"progressMax":         "progress",
		"progressMin":         "progress",
		"progressSum":         "progress",
		"typeMax":             "type",
		"typeMin":             "type",
	}
	fieldsMap := map[string]struct{}{}
	fields := []string{}
	aggregationFields := []GetAggregationsAggregationField{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		aggregationFields = append(aggregationFields, aggregationsMap[f.Field.Name])
		fieldsMap[aggregationFieldsMap[f.Field.Name]] = struct{}{}
	}
	for key, _ := range fieldsMap {
		fields = append(fields, key)
	}

	opts := GetAggregationsOptions{
		Alias:             TableName("exports"),
		Fields:            fields,
		AggregationFields: aggregationFields,
	}
	res = &ExportResultAggregations{}
	err = obj.GetAggregations(ctx, r.GetDB(ctx), opts, &Export{}, res)
	return
}

type GeneratedExportResolver struct{ *GeneratedResolver }

// File ...
func (r *GeneratedExportResolver) File(ctx context.Context, obj *Export) (res *File, err error) {
	return r.Handlers.ExportFile(ctx, r.GeneratedResolver, obj)
}

// FileHandler ...
func ExportFileHandler(ctx context.Context, r *GeneratedResolver, obj *Export) (res *File, err error) {

	if obj.FileID != nil {
		res = &File{ID: *obj.FileID}
	}

	return
}
