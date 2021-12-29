package src

import (
	"context"

	"github.com/graphql-services/graphql-exports/gen"
)

// NewResolver ...
func NewResolver(db *gen.DB, ec *gen.EventController) *Resolver {
	handlers := gen.DefaultResolutionHandlers()
	return &Resolver{gen.NewGeneratedResolver(handlers, db, ec)}
}

// Resolver ...
type Resolver struct {
	*gen.GeneratedResolver
}

// MutationResolver ...
type MutationResolver struct {
	*gen.GeneratedMutationResolver
}

// BeginTransaction ...
func (r *MutationResolver) BeginTransaction(ctx context.Context, fn func(context.Context) error) error {
	ctx = gen.EnrichContextWithMutations(ctx, r.GeneratedResolver)
	err := fn(ctx)
	if err != nil {
		tx := r.GeneratedResolver.GetDB(ctx)
		tx.Rollback()
		return err
	}
	return gen.FinishMutationContext(ctx, r.GeneratedResolver)
}

// QueryResolver ...
type QueryResolver struct {
	*gen.GeneratedQueryResolver
}

// Mutation ...
func (r *Resolver) Mutation() gen.MutationResolver {
	return &MutationResolver{&gen.GeneratedMutationResolver{r.GeneratedResolver}}
}

// Query ...
func (r *Resolver) Query() gen.QueryResolver {
	return &QueryResolver{&gen.GeneratedQueryResolver{r.GeneratedResolver}}
}

// ExportResultTypeResolver ...
type ExportResultTypeResolver struct {
	*gen.GeneratedExportResultTypeResolver
}

// ExportResultType ...
func (r *Resolver) ExportResultType() gen.ExportResultTypeResolver {
	return &ExportResultTypeResolver{&gen.GeneratedExportResultTypeResolver{r.GeneratedResolver}}
}

// ExportResolver ...
type ExportResolver struct {
	*gen.GeneratedExportResolver
}

// Export ...
func (r *Resolver) Export() gen.ExportResolver {
	return &ExportResolver{&gen.GeneratedExportResolver{r.GeneratedResolver}}
}

// FileResolver ...
type FileResolver struct {
	*gen.GeneratedFileResolver
}
