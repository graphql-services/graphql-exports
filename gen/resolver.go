package gen

import (
	"context"

	"github.com/novacloudcz/graphql-orm/events"
)

type ResolutionHandlers struct {
	OnEvent func(ctx context.Context, r *GeneratedResolver, e *events.Event) error

	CreateExport     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Export, err error)
	UpdateExport     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Export, err error)
	DeleteExport     func(ctx context.Context, r *GeneratedResolver, id string) (item *Export, err error)
	DeleteAllExports func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryExport      func(ctx context.Context, r *GeneratedResolver, opts QueryExportHandlerOptions) (*Export, error)
	QueryExports     func(ctx context.Context, r *GeneratedResolver, opts QueryExportsHandlerOptions) (*ExportResultType, error)

	ExportFile func(ctx context.Context, r *GeneratedResolver, obj *Export) (res *File, err error)
}

func DefaultResolutionHandlers() ResolutionHandlers {
	handlers := ResolutionHandlers{
		OnEvent: func(ctx context.Context, r *GeneratedResolver, e *events.Event) error { return nil },

		CreateExport:     CreateExportHandler,
		UpdateExport:     UpdateExportHandler,
		DeleteExport:     DeleteExportHandler,
		DeleteAllExports: DeleteAllExportsHandler,
		QueryExport:      QueryExportHandler,
		QueryExports:     QueryExportsHandler,

		ExportFile: ExportFileHandler,
	}
	return handlers
}

type GeneratedResolver struct {
	Handlers        ResolutionHandlers
	DB              *DB
	EventController *events.EventController
}
