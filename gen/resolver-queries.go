package gen

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/ast"
)

type GeneratedQueryResolver struct{ *GeneratedResolver }

type QueryExportHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *ExportFilterType
}

func (r *GeneratedQueryResolver) Export(ctx context.Context, id *string, q *string, filter *ExportFilterType) (*Export, error) {
	opts := QueryExportHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryExport(ctx, r.GeneratedResolver, opts)
}
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
	qb := r.DB.Query()
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
		return nil, &NotFoundError{Entity: "Export"}
	}
	return items[0], err
}

type QueryExportsHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*ExportSortType
	Filter *ExportFilterType
}

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

type GeneratedExportResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedExportResultTypeResolver) Items(ctx context.Context, obj *ExportResultType) (items []*Export, err error) {
	giOpts := GetItemsOptions{
		Alias:      TableName("exports"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.DB.db, giOpts, &items)

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

func (r *GeneratedExportResultTypeResolver) Count(ctx context.Context, obj *ExportResultType) (count int, err error) {
	return obj.GetCount(ctx, r.DB.db, &Export{})
}

type GeneratedExportResolver struct{ *GeneratedResolver }

func (r *GeneratedExportResolver) File(ctx context.Context, obj *Export) (res *File, err error) {
	return r.Handlers.ExportFile(ctx, r.GeneratedResolver, obj)
}
func ExportFileHandler(ctx context.Context, r *GeneratedResolver, obj *Export) (res *File, err error) {

	if obj.FileID != nil {
		res = &File{ID: *obj.FileID}
	}

	return
}
