package gen

import (
	"context"

	"github.com/jinzhu/gorm"
)

func (s ExportSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]string, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("exports"), sorts, joins)
}
func (s ExportSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]string, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("id")+" "+s.ID.String())
	}

	if s.Type != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("type")+" "+s.Type.String())
	}

	if s.Metadata != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("metadata")+" "+s.Metadata.String())
	}

	if s.FileID != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("fileId")+" "+s.FileID.String())
	}

	if s.UpdatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("updatedAt")+" "+s.UpdatedAt.String())
	}

	if s.CreatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("createdAt")+" "+s.CreatedAt.String())
	}

	if s.UpdatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("updatedBy")+" "+s.UpdatedBy.String())
	}

	if s.CreatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("createdBy")+" "+s.CreatedBy.String())
	}

	return nil
}
