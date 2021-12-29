package gen

import (
	"context"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

// IsEmpty ...
func (f *ExportFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	havings := []string{}
	whereValues := []interface{}{}
	havingValues := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &whereValues, &havings, &havingValues, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0 && len(havings) == 0
}

// Apply ...
func (f *ExportFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("exports"), wheres, whereValues, havings, havingValues, joins)
}

// ApplyWithAlias ...
func (f *ExportFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _whereValues := f.WhereContent(dialect, aliasPrefix)
	_having, _havingValues := f.HavingContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*havings = append(*havings, _having...)
	*whereValues = append(*whereValues, _whereValues...)
	*havingValues = append(*havingValues, _havingValues...)

	if f.Or != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_ws := []string{}
			_hs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_ws, &wvs, &_hs, &hvs, &js)
			if err != nil {
				return err
			}
			if len(_ws) > 0 {
				ws = append(ws, strings.Join(_ws, " AND "))
			}
			if len(_hs) > 0 {
				hs = append(hs, strings.Join(_hs, " AND "))
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, "("+strings.Join(ws, " OR ")+")")
		}
		if len(hs) > 0 {
			*havings = append(*havings, "("+strings.Join(hs, " OR ")+")")
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &ws, &wvs, &hs, &hvs, &js)
			if err != nil {
				return err
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, strings.Join(ws, " AND "))
		}
		if len(hs) > 0 {
			*havings = append(*havings, strings.Join(hs, " AND "))
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}

	return nil
}

// WhereContent ...
func (f *ExportFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}

	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}

	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}

	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}

	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}

	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}

	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.IDNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" NOT IN (?)")
		values = append(values, f.IDNotIn)
	}

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NOT NULL")
		}
	}

	if f.Type != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" = ?")
		values = append(values, f.Type)
	}

	if f.TypeNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" != ?")
		values = append(values, f.TypeNe)
	}

	if f.TypeGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" > ?")
		values = append(values, f.TypeGt)
	}

	if f.TypeLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" < ?")
		values = append(values, f.TypeLt)
	}

	if f.TypeGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" >= ?")
		values = append(values, f.TypeGte)
	}

	if f.TypeLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" <= ?")
		values = append(values, f.TypeLte)
	}

	if f.TypeIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" IN (?)")
		values = append(values, f.TypeIn)
	}

	if f.TypeNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" NOT IN (?)")
		values = append(values, f.TypeNotIn)
	}

	if f.TypeLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.TypeLike, "?", "_", -1), "*", "%", -1))
	}

	if f.TypePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.TypePrefix))
	}

	if f.TypeSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.TypeSuffix))
	}

	if f.TypeNull != nil {
		if *f.TypeNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" IS NOT NULL")
		}
	}

	if f.Metadata != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("metadata")+" = ?")
		values = append(values, f.Metadata)
	}

	if f.MetadataNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("metadata")+" != ?")
		values = append(values, f.MetadataNe)
	}

	if f.MetadataGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("metadata")+" > ?")
		values = append(values, f.MetadataGt)
	}

	if f.MetadataLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("metadata")+" < ?")
		values = append(values, f.MetadataLt)
	}

	if f.MetadataGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("metadata")+" >= ?")
		values = append(values, f.MetadataGte)
	}

	if f.MetadataLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("metadata")+" <= ?")
		values = append(values, f.MetadataLte)
	}

	if f.MetadataIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("metadata")+" IN (?)")
		values = append(values, f.MetadataIn)
	}

	if f.MetadataNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("metadata")+" NOT IN (?)")
		values = append(values, f.MetadataNotIn)
	}

	if f.MetadataLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("metadata")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.MetadataLike, "?", "_", -1), "*", "%", -1))
	}

	if f.MetadataPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("metadata")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.MetadataPrefix))
	}

	if f.MetadataSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("metadata")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.MetadataSuffix))
	}

	if f.MetadataNull != nil {
		if *f.MetadataNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("metadata")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("metadata")+" IS NOT NULL")
		}
	}

	if f.State != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" = ?")
		values = append(values, f.State)
	}

	if f.StateNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" != ?")
		values = append(values, f.StateNe)
	}

	if f.StateGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" > ?")
		values = append(values, f.StateGt)
	}

	if f.StateLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" < ?")
		values = append(values, f.StateLt)
	}

	if f.StateGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" >= ?")
		values = append(values, f.StateGte)
	}

	if f.StateLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" <= ?")
		values = append(values, f.StateLte)
	}

	if f.StateIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" IN (?)")
		values = append(values, f.StateIn)
	}

	if f.StateNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" NOT IN (?)")
		values = append(values, f.StateNotIn)
	}

	if f.StateNull != nil {
		if *f.StateNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" IS NOT NULL")
		}
	}

	if f.Progress != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("progress")+" = ?")
		values = append(values, f.Progress)
	}

	if f.ProgressNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("progress")+" != ?")
		values = append(values, f.ProgressNe)
	}

	if f.ProgressGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("progress")+" > ?")
		values = append(values, f.ProgressGt)
	}

	if f.ProgressLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("progress")+" < ?")
		values = append(values, f.ProgressLt)
	}

	if f.ProgressGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("progress")+" >= ?")
		values = append(values, f.ProgressGte)
	}

	if f.ProgressLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("progress")+" <= ?")
		values = append(values, f.ProgressLte)
	}

	if f.ProgressIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("progress")+" IN (?)")
		values = append(values, f.ProgressIn)
	}

	if f.ProgressNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("progress")+" NOT IN (?)")
		values = append(values, f.ProgressNotIn)
	}

	if f.ProgressNull != nil {
		if *f.ProgressNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("progress")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("progress")+" IS NOT NULL")
		}
	}

	if f.ErrorDescription != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("errorDescription")+" = ?")
		values = append(values, f.ErrorDescription)
	}

	if f.ErrorDescriptionNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("errorDescription")+" != ?")
		values = append(values, f.ErrorDescriptionNe)
	}

	if f.ErrorDescriptionGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("errorDescription")+" > ?")
		values = append(values, f.ErrorDescriptionGt)
	}

	if f.ErrorDescriptionLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("errorDescription")+" < ?")
		values = append(values, f.ErrorDescriptionLt)
	}

	if f.ErrorDescriptionGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("errorDescription")+" >= ?")
		values = append(values, f.ErrorDescriptionGte)
	}

	if f.ErrorDescriptionLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("errorDescription")+" <= ?")
		values = append(values, f.ErrorDescriptionLte)
	}

	if f.ErrorDescriptionIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("errorDescription")+" IN (?)")
		values = append(values, f.ErrorDescriptionIn)
	}

	if f.ErrorDescriptionNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("errorDescription")+" NOT IN (?)")
		values = append(values, f.ErrorDescriptionNotIn)
	}

	if f.ErrorDescriptionLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("errorDescription")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.ErrorDescriptionLike, "?", "_", -1), "*", "%", -1))
	}

	if f.ErrorDescriptionPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("errorDescription")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.ErrorDescriptionPrefix))
	}

	if f.ErrorDescriptionSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("errorDescription")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.ErrorDescriptionSuffix))
	}

	if f.ErrorDescriptionNull != nil {
		if *f.ErrorDescriptionNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("errorDescription")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("errorDescription")+" IS NOT NULL")
		}
	}

	if f.FileID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("fileId")+" = ?")
		values = append(values, f.FileID)
	}

	if f.FileIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("fileId")+" != ?")
		values = append(values, f.FileIDNe)
	}

	if f.FileIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("fileId")+" > ?")
		values = append(values, f.FileIDGt)
	}

	if f.FileIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("fileId")+" < ?")
		values = append(values, f.FileIDLt)
	}

	if f.FileIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("fileId")+" >= ?")
		values = append(values, f.FileIDGte)
	}

	if f.FileIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("fileId")+" <= ?")
		values = append(values, f.FileIDLte)
	}

	if f.FileIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("fileId")+" IN (?)")
		values = append(values, f.FileIDIn)
	}

	if f.FileIDNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("fileId")+" NOT IN (?)")
		values = append(values, f.FileIDNotIn)
	}

	if f.FileIDLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("fileId")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.FileIDLike, "?", "_", -1), "*", "%", -1))
	}

	if f.FileIDPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("fileId")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.FileIDPrefix))
	}

	if f.FileIDSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("fileId")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.FileIDSuffix))
	}

	if f.FileIDNull != nil {
		if *f.FileIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("fileId")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("fileId")+" IS NOT NULL")
		}
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}

	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}

	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}

	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}

	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}

	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}

	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.UpdatedAtNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" NOT IN (?)")
		values = append(values, f.UpdatedAtNotIn)
	}

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NOT NULL")
		}
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}

	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}

	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}

	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}

	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}

	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}

	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.CreatedAtNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" NOT IN (?)")
		values = append(values, f.CreatedAtNotIn)
	}

	if f.CreatedAtNull != nil {
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NOT NULL")
		}
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}

	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}

	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}

	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}

	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}

	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}

	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.UpdatedByNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" NOT IN (?)")
		values = append(values, f.UpdatedByNotIn)
	}

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NOT NULL")
		}
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}

	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}

	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}

	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}

	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}

	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}

	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	if f.CreatedByNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" NOT IN (?)")
		values = append(values, f.CreatedByNotIn)
	}

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NOT NULL")
		}
	}

	return
}

// HavingContent ...
func (f *ExportFilterType) HavingContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.IDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMin)
	}

	if f.IDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMax)
	}

	if f.IDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMinNe)
	}

	if f.IDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMaxNe)
	}

	if f.IDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMinGt)
	}

	if f.IDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMaxGt)
	}

	if f.IDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMinLt)
	}

	if f.IDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMaxLt)
	}

	if f.IDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMinGte)
	}

	if f.IDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMaxGte)
	}

	if f.IDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMinLte)
	}

	if f.IDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMaxLte)
	}

	if f.IDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMinIn)
	}

	if f.IDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMaxIn)
	}

	if f.IDMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") NOT IN (?)")
		values = append(values, f.IDMinNotIn)
	}

	if f.IDMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") NOT IN (?)")
		values = append(values, f.IDMaxNotIn)
	}

	if f.TypeMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("type")+") = ?")
		values = append(values, f.TypeMin)
	}

	if f.TypeMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("type")+") = ?")
		values = append(values, f.TypeMax)
	}

	if f.TypeMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("type")+") != ?")
		values = append(values, f.TypeMinNe)
	}

	if f.TypeMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("type")+") != ?")
		values = append(values, f.TypeMaxNe)
	}

	if f.TypeMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("type")+") > ?")
		values = append(values, f.TypeMinGt)
	}

	if f.TypeMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("type")+") > ?")
		values = append(values, f.TypeMaxGt)
	}

	if f.TypeMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("type")+") < ?")
		values = append(values, f.TypeMinLt)
	}

	if f.TypeMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("type")+") < ?")
		values = append(values, f.TypeMaxLt)
	}

	if f.TypeMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("type")+") >= ?")
		values = append(values, f.TypeMinGte)
	}

	if f.TypeMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("type")+") >= ?")
		values = append(values, f.TypeMaxGte)
	}

	if f.TypeMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("type")+") <= ?")
		values = append(values, f.TypeMinLte)
	}

	if f.TypeMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("type")+") <= ?")
		values = append(values, f.TypeMaxLte)
	}

	if f.TypeMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("type")+") IN (?)")
		values = append(values, f.TypeMinIn)
	}

	if f.TypeMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("type")+") IN (?)")
		values = append(values, f.TypeMaxIn)
	}

	if f.TypeMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("type")+") NOT IN (?)")
		values = append(values, f.TypeMinNotIn)
	}

	if f.TypeMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("type")+") NOT IN (?)")
		values = append(values, f.TypeMaxNotIn)
	}

	if f.TypeMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("type")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.TypeMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.TypeMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("type")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.TypeMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.TypeMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("type")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.TypeMinPrefix))
	}

	if f.TypeMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("type")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.TypeMaxPrefix))
	}

	if f.TypeMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("type")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.TypeMinSuffix))
	}

	if f.TypeMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("type")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.TypeMaxSuffix))
	}

	if f.MetadataMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("metadata")+") = ?")
		values = append(values, f.MetadataMin)
	}

	if f.MetadataMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("metadata")+") = ?")
		values = append(values, f.MetadataMax)
	}

	if f.MetadataMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("metadata")+") != ?")
		values = append(values, f.MetadataMinNe)
	}

	if f.MetadataMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("metadata")+") != ?")
		values = append(values, f.MetadataMaxNe)
	}

	if f.MetadataMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("metadata")+") > ?")
		values = append(values, f.MetadataMinGt)
	}

	if f.MetadataMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("metadata")+") > ?")
		values = append(values, f.MetadataMaxGt)
	}

	if f.MetadataMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("metadata")+") < ?")
		values = append(values, f.MetadataMinLt)
	}

	if f.MetadataMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("metadata")+") < ?")
		values = append(values, f.MetadataMaxLt)
	}

	if f.MetadataMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("metadata")+") >= ?")
		values = append(values, f.MetadataMinGte)
	}

	if f.MetadataMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("metadata")+") >= ?")
		values = append(values, f.MetadataMaxGte)
	}

	if f.MetadataMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("metadata")+") <= ?")
		values = append(values, f.MetadataMinLte)
	}

	if f.MetadataMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("metadata")+") <= ?")
		values = append(values, f.MetadataMaxLte)
	}

	if f.MetadataMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("metadata")+") IN (?)")
		values = append(values, f.MetadataMinIn)
	}

	if f.MetadataMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("metadata")+") IN (?)")
		values = append(values, f.MetadataMaxIn)
	}

	if f.MetadataMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("metadata")+") NOT IN (?)")
		values = append(values, f.MetadataMinNotIn)
	}

	if f.MetadataMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("metadata")+") NOT IN (?)")
		values = append(values, f.MetadataMaxNotIn)
	}

	if f.MetadataMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("metadata")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.MetadataMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.MetadataMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("metadata")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.MetadataMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.MetadataMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("metadata")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.MetadataMinPrefix))
	}

	if f.MetadataMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("metadata")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.MetadataMaxPrefix))
	}

	if f.MetadataMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("metadata")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.MetadataMinSuffix))
	}

	if f.MetadataMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("metadata")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.MetadataMaxSuffix))
	}

	if f.StateMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("state")+") = ?")
		values = append(values, f.StateMin)
	}

	if f.StateMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("state")+") = ?")
		values = append(values, f.StateMax)
	}

	if f.StateMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("state")+") != ?")
		values = append(values, f.StateMinNe)
	}

	if f.StateMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("state")+") != ?")
		values = append(values, f.StateMaxNe)
	}

	if f.StateMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("state")+") > ?")
		values = append(values, f.StateMinGt)
	}

	if f.StateMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("state")+") > ?")
		values = append(values, f.StateMaxGt)
	}

	if f.StateMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("state")+") < ?")
		values = append(values, f.StateMinLt)
	}

	if f.StateMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("state")+") < ?")
		values = append(values, f.StateMaxLt)
	}

	if f.StateMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("state")+") >= ?")
		values = append(values, f.StateMinGte)
	}

	if f.StateMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("state")+") >= ?")
		values = append(values, f.StateMaxGte)
	}

	if f.StateMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("state")+") <= ?")
		values = append(values, f.StateMinLte)
	}

	if f.StateMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("state")+") <= ?")
		values = append(values, f.StateMaxLte)
	}

	if f.StateMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("state")+") IN (?)")
		values = append(values, f.StateMinIn)
	}

	if f.StateMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("state")+") IN (?)")
		values = append(values, f.StateMaxIn)
	}

	if f.StateMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("state")+") NOT IN (?)")
		values = append(values, f.StateMinNotIn)
	}

	if f.StateMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("state")+") NOT IN (?)")
		values = append(values, f.StateMaxNotIn)
	}

	if f.ProgressMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("progress")+") = ?")
		values = append(values, f.ProgressMin)
	}

	if f.ProgressMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("progress")+") = ?")
		values = append(values, f.ProgressMax)
	}

	if f.ProgressAvg != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("progress")+") = ?")
		values = append(values, f.ProgressAvg)
	}

	if f.ProgressSum != nil {
		conditions = append(conditions, "Sum("+aliasPrefix+dialect.Quote("progress")+") = ?")
		values = append(values, f.ProgressSum)
	}

	if f.ProgressMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("progress")+") != ?")
		values = append(values, f.ProgressMinNe)
	}

	if f.ProgressMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("progress")+") != ?")
		values = append(values, f.ProgressMaxNe)
	}

	if f.ProgressAvgNe != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("progress")+") != ?")
		values = append(values, f.ProgressAvgNe)
	}

	if f.ProgressSumNe != nil {
		conditions = append(conditions, "Sum("+aliasPrefix+dialect.Quote("progress")+") != ?")
		values = append(values, f.ProgressSumNe)
	}

	if f.ProgressMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("progress")+") > ?")
		values = append(values, f.ProgressMinGt)
	}

	if f.ProgressMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("progress")+") > ?")
		values = append(values, f.ProgressMaxGt)
	}

	if f.ProgressAvgGt != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("progress")+") > ?")
		values = append(values, f.ProgressAvgGt)
	}

	if f.ProgressSumGt != nil {
		conditions = append(conditions, "Sum("+aliasPrefix+dialect.Quote("progress")+") > ?")
		values = append(values, f.ProgressSumGt)
	}

	if f.ProgressMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("progress")+") < ?")
		values = append(values, f.ProgressMinLt)
	}

	if f.ProgressMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("progress")+") < ?")
		values = append(values, f.ProgressMaxLt)
	}

	if f.ProgressAvgLt != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("progress")+") < ?")
		values = append(values, f.ProgressAvgLt)
	}

	if f.ProgressSumLt != nil {
		conditions = append(conditions, "Sum("+aliasPrefix+dialect.Quote("progress")+") < ?")
		values = append(values, f.ProgressSumLt)
	}

	if f.ProgressMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("progress")+") >= ?")
		values = append(values, f.ProgressMinGte)
	}

	if f.ProgressMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("progress")+") >= ?")
		values = append(values, f.ProgressMaxGte)
	}

	if f.ProgressAvgGte != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("progress")+") >= ?")
		values = append(values, f.ProgressAvgGte)
	}

	if f.ProgressSumGte != nil {
		conditions = append(conditions, "Sum("+aliasPrefix+dialect.Quote("progress")+") >= ?")
		values = append(values, f.ProgressSumGte)
	}

	if f.ProgressMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("progress")+") <= ?")
		values = append(values, f.ProgressMinLte)
	}

	if f.ProgressMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("progress")+") <= ?")
		values = append(values, f.ProgressMaxLte)
	}

	if f.ProgressAvgLte != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("progress")+") <= ?")
		values = append(values, f.ProgressAvgLte)
	}

	if f.ProgressSumLte != nil {
		conditions = append(conditions, "Sum("+aliasPrefix+dialect.Quote("progress")+") <= ?")
		values = append(values, f.ProgressSumLte)
	}

	if f.ProgressMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("progress")+") IN (?)")
		values = append(values, f.ProgressMinIn)
	}

	if f.ProgressMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("progress")+") IN (?)")
		values = append(values, f.ProgressMaxIn)
	}

	if f.ProgressAvgIn != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("progress")+") IN (?)")
		values = append(values, f.ProgressAvgIn)
	}

	if f.ProgressSumIn != nil {
		conditions = append(conditions, "Sum("+aliasPrefix+dialect.Quote("progress")+") IN (?)")
		values = append(values, f.ProgressSumIn)
	}

	if f.ProgressMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("progress")+") NOT IN (?)")
		values = append(values, f.ProgressMinNotIn)
	}

	if f.ProgressMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("progress")+") NOT IN (?)")
		values = append(values, f.ProgressMaxNotIn)
	}

	if f.ProgressAvgNotIn != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("progress")+") NOT IN (?)")
		values = append(values, f.ProgressAvgNotIn)
	}

	if f.ProgressSumNotIn != nil {
		conditions = append(conditions, "Sum("+aliasPrefix+dialect.Quote("progress")+") NOT IN (?)")
		values = append(values, f.ProgressSumNotIn)
	}

	if f.ErrorDescriptionMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("errorDescription")+") = ?")
		values = append(values, f.ErrorDescriptionMin)
	}

	if f.ErrorDescriptionMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("errorDescription")+") = ?")
		values = append(values, f.ErrorDescriptionMax)
	}

	if f.ErrorDescriptionMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("errorDescription")+") != ?")
		values = append(values, f.ErrorDescriptionMinNe)
	}

	if f.ErrorDescriptionMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("errorDescription")+") != ?")
		values = append(values, f.ErrorDescriptionMaxNe)
	}

	if f.ErrorDescriptionMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("errorDescription")+") > ?")
		values = append(values, f.ErrorDescriptionMinGt)
	}

	if f.ErrorDescriptionMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("errorDescription")+") > ?")
		values = append(values, f.ErrorDescriptionMaxGt)
	}

	if f.ErrorDescriptionMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("errorDescription")+") < ?")
		values = append(values, f.ErrorDescriptionMinLt)
	}

	if f.ErrorDescriptionMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("errorDescription")+") < ?")
		values = append(values, f.ErrorDescriptionMaxLt)
	}

	if f.ErrorDescriptionMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("errorDescription")+") >= ?")
		values = append(values, f.ErrorDescriptionMinGte)
	}

	if f.ErrorDescriptionMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("errorDescription")+") >= ?")
		values = append(values, f.ErrorDescriptionMaxGte)
	}

	if f.ErrorDescriptionMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("errorDescription")+") <= ?")
		values = append(values, f.ErrorDescriptionMinLte)
	}

	if f.ErrorDescriptionMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("errorDescription")+") <= ?")
		values = append(values, f.ErrorDescriptionMaxLte)
	}

	if f.ErrorDescriptionMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("errorDescription")+") IN (?)")
		values = append(values, f.ErrorDescriptionMinIn)
	}

	if f.ErrorDescriptionMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("errorDescription")+") IN (?)")
		values = append(values, f.ErrorDescriptionMaxIn)
	}

	if f.ErrorDescriptionMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("errorDescription")+") NOT IN (?)")
		values = append(values, f.ErrorDescriptionMinNotIn)
	}

	if f.ErrorDescriptionMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("errorDescription")+") NOT IN (?)")
		values = append(values, f.ErrorDescriptionMaxNotIn)
	}

	if f.ErrorDescriptionMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("errorDescription")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.ErrorDescriptionMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.ErrorDescriptionMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("errorDescription")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.ErrorDescriptionMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.ErrorDescriptionMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("errorDescription")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.ErrorDescriptionMinPrefix))
	}

	if f.ErrorDescriptionMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("errorDescription")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.ErrorDescriptionMaxPrefix))
	}

	if f.ErrorDescriptionMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("errorDescription")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.ErrorDescriptionMinSuffix))
	}

	if f.ErrorDescriptionMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("errorDescription")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.ErrorDescriptionMaxSuffix))
	}

	if f.FileIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("fileId")+") = ?")
		values = append(values, f.FileIDMin)
	}

	if f.FileIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("fileId")+") = ?")
		values = append(values, f.FileIDMax)
	}

	if f.FileIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("fileId")+") != ?")
		values = append(values, f.FileIDMinNe)
	}

	if f.FileIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("fileId")+") != ?")
		values = append(values, f.FileIDMaxNe)
	}

	if f.FileIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("fileId")+") > ?")
		values = append(values, f.FileIDMinGt)
	}

	if f.FileIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("fileId")+") > ?")
		values = append(values, f.FileIDMaxGt)
	}

	if f.FileIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("fileId")+") < ?")
		values = append(values, f.FileIDMinLt)
	}

	if f.FileIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("fileId")+") < ?")
		values = append(values, f.FileIDMaxLt)
	}

	if f.FileIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("fileId")+") >= ?")
		values = append(values, f.FileIDMinGte)
	}

	if f.FileIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("fileId")+") >= ?")
		values = append(values, f.FileIDMaxGte)
	}

	if f.FileIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("fileId")+") <= ?")
		values = append(values, f.FileIDMinLte)
	}

	if f.FileIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("fileId")+") <= ?")
		values = append(values, f.FileIDMaxLte)
	}

	if f.FileIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("fileId")+") IN (?)")
		values = append(values, f.FileIDMinIn)
	}

	if f.FileIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("fileId")+") IN (?)")
		values = append(values, f.FileIDMaxIn)
	}

	if f.FileIDMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("fileId")+") NOT IN (?)")
		values = append(values, f.FileIDMinNotIn)
	}

	if f.FileIDMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("fileId")+") NOT IN (?)")
		values = append(values, f.FileIDMaxNotIn)
	}

	if f.FileIDMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("fileId")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.FileIDMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.FileIDMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("fileId")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.FileIDMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.FileIDMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("fileId")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.FileIDMinPrefix))
	}

	if f.FileIDMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("fileId")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.FileIDMaxPrefix))
	}

	if f.FileIDMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("fileId")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.FileIDMinSuffix))
	}

	if f.FileIDMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("fileId")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.FileIDMaxSuffix))
	}

	if f.UpdatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMin)
	}

	if f.UpdatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMax)
	}

	if f.UpdatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMinNe)
	}

	if f.UpdatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMaxNe)
	}

	if f.UpdatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMinGt)
	}

	if f.UpdatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMaxGt)
	}

	if f.UpdatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMinLt)
	}

	if f.UpdatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMaxLt)
	}

	if f.UpdatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMinGte)
	}

	if f.UpdatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMaxGte)
	}

	if f.UpdatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMinLte)
	}

	if f.UpdatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMaxLte)
	}

	if f.UpdatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMinIn)
	}

	if f.UpdatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMaxIn)
	}

	if f.UpdatedAtMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") NOT IN (?)")
		values = append(values, f.UpdatedAtMinNotIn)
	}

	if f.UpdatedAtMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") NOT IN (?)")
		values = append(values, f.UpdatedAtMaxNotIn)
	}

	if f.CreatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMin)
	}

	if f.CreatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMax)
	}

	if f.CreatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMinNe)
	}

	if f.CreatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMaxNe)
	}

	if f.CreatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMinGt)
	}

	if f.CreatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMaxGt)
	}

	if f.CreatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMinLt)
	}

	if f.CreatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMaxLt)
	}

	if f.CreatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMinGte)
	}

	if f.CreatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMaxGte)
	}

	if f.CreatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMinLte)
	}

	if f.CreatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMaxLte)
	}

	if f.CreatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMinIn)
	}

	if f.CreatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMaxIn)
	}

	if f.CreatedAtMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") NOT IN (?)")
		values = append(values, f.CreatedAtMinNotIn)
	}

	if f.CreatedAtMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") NOT IN (?)")
		values = append(values, f.CreatedAtMaxNotIn)
	}

	if f.UpdatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMin)
	}

	if f.UpdatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMax)
	}

	if f.UpdatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMinNe)
	}

	if f.UpdatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMaxNe)
	}

	if f.UpdatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMinGt)
	}

	if f.UpdatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMaxGt)
	}

	if f.UpdatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMinLt)
	}

	if f.UpdatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMaxLt)
	}

	if f.UpdatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMinGte)
	}

	if f.UpdatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMaxGte)
	}

	if f.UpdatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMinLte)
	}

	if f.UpdatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMaxLte)
	}

	if f.UpdatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMinIn)
	}

	if f.UpdatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMaxIn)
	}

	if f.UpdatedByMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") NOT IN (?)")
		values = append(values, f.UpdatedByMinNotIn)
	}

	if f.UpdatedByMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") NOT IN (?)")
		values = append(values, f.UpdatedByMaxNotIn)
	}

	if f.CreatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMin)
	}

	if f.CreatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMax)
	}

	if f.CreatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMinNe)
	}

	if f.CreatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMaxNe)
	}

	if f.CreatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMinGt)
	}

	if f.CreatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMaxGt)
	}

	if f.CreatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMinLt)
	}

	if f.CreatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMaxLt)
	}

	if f.CreatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMinGte)
	}

	if f.CreatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMaxGte)
	}

	if f.CreatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMinLte)
	}

	if f.CreatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMaxLte)
	}

	if f.CreatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMinIn)
	}

	if f.CreatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMaxIn)
	}

	if f.CreatedByMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") NOT IN (?)")
		values = append(values, f.CreatedByMinNotIn)
	}

	if f.CreatedByMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") NOT IN (?)")
		values = append(values, f.CreatedByMaxNotIn)
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *ExportFilterType) AndWith(f2 ...*ExportFilterType) *ExportFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &ExportFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *ExportFilterType) OrWith(f2 ...*ExportFilterType) *ExportFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &ExportFilterType{
		Or: append(_f2, f),
	}
}
