package src

import (
	"context"

	"github.com/graphql-services/graphql-exports/gen"
	"github.com/novacloudcz/graphql-orm/events"
)

func NewResolver(db *gen.DB, ec *events.EventController) *Resolver {
	handlers := gen.DefaultResolutionHandlers()
	return &Resolver{&gen.GeneratedResolver{Handlers: handlers, DB: db, EventController: ec}}
}

type Resolver struct {
	*gen.GeneratedResolver
}

type MutationResolver struct {
	*gen.GeneratedMutationResolver
}

func (r *MutationResolver) BeginTransaction(ctx context.Context, fn func(context.Context) error) error {
	ctx = gen.EnrichContextWithMutations(ctx, r.GeneratedResolver)
	err := fn(ctx)
	if err != nil {
		tx := gen.GetTransaction(ctx)
		tx.Rollback()
		return err
	}
	return gen.FinishMutationContext(ctx, r.GeneratedResolver)
}

type QueryResolver struct {
	*gen.GeneratedQueryResolver
}

func (r *Resolver) Mutation() gen.MutationResolver {
	return &MutationResolver{&gen.GeneratedMutationResolver{r.GeneratedResolver}}
}
func (r *Resolver) Query() gen.QueryResolver {
	return &QueryResolver{&gen.GeneratedQueryResolver{r.GeneratedResolver}}
}

type ExportResultTypeResolver struct {
	*gen.GeneratedExportResultTypeResolver
}

func (r *Resolver) ExportResultType() gen.ExportResultTypeResolver {
	return &ExportResultTypeResolver{&gen.GeneratedExportResultTypeResolver{r.GeneratedResolver}}
}

type ExportResolver struct {
	*gen.GeneratedExportResolver
}

func (r *Resolver) Export() gen.ExportResolver {
	return &ExportResolver{&gen.GeneratedExportResolver{r.GeneratedResolver}}
}

type FileResolver struct {
	*gen.GeneratedFileResolver
}
