package gen

import (
	"context"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

func (f *ExportFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	values := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &values, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0
}
func (f *ExportFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, values *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("exports"), wheres, values, joins)
}
func (f *ExportFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _values := f.WhereContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*values = append(*values, _values...)

	if f.Or != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_cs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_cs, &vs, &js)
			if err != nil {
				return err
			}
			cs = append(cs, strings.Join(_cs, " AND "))
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, "("+strings.Join(cs, " OR ")+")")
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &cs, &vs, &js)
			if err != nil {
				return err
			}
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, strings.Join(cs, " AND "))
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}

	return nil
}

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

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NOT NULL")
		}
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
