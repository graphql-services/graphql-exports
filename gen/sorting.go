package gen

import (
	"context"

	"github.com/jinzhu/gorm"
)

// Apply ...
func (s ExportSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]SortInfo, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("exports"), sorts, joins)
}

// ApplyWithAlias ...
func (s ExportSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]SortInfo, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("id"), Direction: s.ID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.IDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.IDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Type != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("type"), Direction: s.Type.String()}
		*sorts = append(*sorts, sort)
	}

	if s.TypeMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("type") + ")", Direction: s.TypeMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.TypeMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("type") + ")", Direction: s.TypeMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Metadata != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("metadata"), Direction: s.Metadata.String()}
		*sorts = append(*sorts, sort)
	}

	if s.MetadataMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("metadata") + ")", Direction: s.MetadataMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.MetadataMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("metadata") + ")", Direction: s.MetadataMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Progress != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("progress"), Direction: s.Progress.String()}
		*sorts = append(*sorts, sort)
	}

	if s.ProgressMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("progress") + ")", Direction: s.ProgressMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.ProgressMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("progress") + ")", Direction: s.ProgressMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.ProgressAvg != nil {
		sort := SortInfo{Field: "Avg(" + aliasPrefix + dialect.Quote("progress") + ")", Direction: s.ProgressAvg.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.ProgressSum != nil {
		sort := SortInfo{Field: "Sum(" + aliasPrefix + dialect.Quote("progress") + ")", Direction: s.ProgressSum.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.ErrorDescription != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("errorDescription"), Direction: s.ErrorDescription.String()}
		*sorts = append(*sorts, sort)
	}

	if s.ErrorDescriptionMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("errorDescription") + ")", Direction: s.ErrorDescriptionMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.ErrorDescriptionMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("errorDescription") + ")", Direction: s.ErrorDescriptionMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.FileID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("fileId"), Direction: s.FileID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.FileIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("fileId") + ")", Direction: s.FileIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.FileIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("fileId") + ")", Direction: s.FileIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedAt"), Direction: s.UpdatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdAt"), Direction: s.CreatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedBy"), Direction: s.UpdatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdBy"), Direction: s.CreatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	return nil
}
