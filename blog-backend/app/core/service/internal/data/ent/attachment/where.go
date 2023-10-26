// Code generated by ent, DO NOT EDIT.

package attachment

import (
	"kratos-cms/app/core/service/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.Attachment {
	return predicate.Attachment(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.Attachment {
	return predicate.Attachment(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.Attachment {
	return predicate.Attachment(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.Attachment {
	return predicate.Attachment(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.Attachment {
	return predicate.Attachment(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.Attachment {
	return predicate.Attachment(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.Attachment {
	return predicate.Attachment(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.Attachment {
	return predicate.Attachment(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.Attachment {
	return predicate.Attachment(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v int64) predicate.Attachment {
	return predicate.Attachment(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v int64) predicate.Attachment {
	return predicate.Attachment(sql.FieldEQ(FieldUpdateTime, v))
}

// DeleteTime applies equality check predicate on the "delete_time" field. It's identical to DeleteTimeEQ.
func DeleteTime(v int64) predicate.Attachment {
	return predicate.Attachment(sql.FieldEQ(FieldDeleteTime, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldEQ(FieldName, v))
}

// Path applies equality check predicate on the "path" field. It's identical to PathEQ.
func Path(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldEQ(FieldPath, v))
}

// FileKey applies equality check predicate on the "file_key" field. It's identical to FileKeyEQ.
func FileKey(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldEQ(FieldFileKey, v))
}

// Thumbnail applies equality check predicate on the "thumbnail" field. It's identical to ThumbnailEQ.
func Thumbnail(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldEQ(FieldThumbnail, v))
}

// MediaType applies equality check predicate on the "media_type" field. It's identical to MediaTypeEQ.
func MediaType(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldEQ(FieldMediaType, v))
}

// Suffix applies equality check predicate on the "suffix" field. It's identical to SuffixEQ.
func Suffix(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldEQ(FieldSuffix, v))
}

// Width applies equality check predicate on the "width" field. It's identical to WidthEQ.
func Width(v int32) predicate.Attachment {
	return predicate.Attachment(sql.FieldEQ(FieldWidth, v))
}

// Height applies equality check predicate on the "height" field. It's identical to HeightEQ.
func Height(v int32) predicate.Attachment {
	return predicate.Attachment(sql.FieldEQ(FieldHeight, v))
}

// Size applies equality check predicate on the "Size" field. It's identical to SizeEQ.
func Size(v uint64) predicate.Attachment {
	return predicate.Attachment(sql.FieldEQ(FieldSize, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v int32) predicate.Attachment {
	return predicate.Attachment(sql.FieldEQ(FieldType, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v int64) predicate.Attachment {
	return predicate.Attachment(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v int64) predicate.Attachment {
	return predicate.Attachment(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...int64) predicate.Attachment {
	return predicate.Attachment(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...int64) predicate.Attachment {
	return predicate.Attachment(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v int64) predicate.Attachment {
	return predicate.Attachment(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v int64) predicate.Attachment {
	return predicate.Attachment(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v int64) predicate.Attachment {
	return predicate.Attachment(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v int64) predicate.Attachment {
	return predicate.Attachment(sql.FieldLTE(FieldCreateTime, v))
}

// CreateTimeIsNil applies the IsNil predicate on the "create_time" field.
func CreateTimeIsNil() predicate.Attachment {
	return predicate.Attachment(sql.FieldIsNull(FieldCreateTime))
}

// CreateTimeNotNil applies the NotNil predicate on the "create_time" field.
func CreateTimeNotNil() predicate.Attachment {
	return predicate.Attachment(sql.FieldNotNull(FieldCreateTime))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v int64) predicate.Attachment {
	return predicate.Attachment(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v int64) predicate.Attachment {
	return predicate.Attachment(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...int64) predicate.Attachment {
	return predicate.Attachment(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...int64) predicate.Attachment {
	return predicate.Attachment(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v int64) predicate.Attachment {
	return predicate.Attachment(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v int64) predicate.Attachment {
	return predicate.Attachment(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v int64) predicate.Attachment {
	return predicate.Attachment(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v int64) predicate.Attachment {
	return predicate.Attachment(sql.FieldLTE(FieldUpdateTime, v))
}

// UpdateTimeIsNil applies the IsNil predicate on the "update_time" field.
func UpdateTimeIsNil() predicate.Attachment {
	return predicate.Attachment(sql.FieldIsNull(FieldUpdateTime))
}

// UpdateTimeNotNil applies the NotNil predicate on the "update_time" field.
func UpdateTimeNotNil() predicate.Attachment {
	return predicate.Attachment(sql.FieldNotNull(FieldUpdateTime))
}

// DeleteTimeEQ applies the EQ predicate on the "delete_time" field.
func DeleteTimeEQ(v int64) predicate.Attachment {
	return predicate.Attachment(sql.FieldEQ(FieldDeleteTime, v))
}

// DeleteTimeNEQ applies the NEQ predicate on the "delete_time" field.
func DeleteTimeNEQ(v int64) predicate.Attachment {
	return predicate.Attachment(sql.FieldNEQ(FieldDeleteTime, v))
}

// DeleteTimeIn applies the In predicate on the "delete_time" field.
func DeleteTimeIn(vs ...int64) predicate.Attachment {
	return predicate.Attachment(sql.FieldIn(FieldDeleteTime, vs...))
}

// DeleteTimeNotIn applies the NotIn predicate on the "delete_time" field.
func DeleteTimeNotIn(vs ...int64) predicate.Attachment {
	return predicate.Attachment(sql.FieldNotIn(FieldDeleteTime, vs...))
}

// DeleteTimeGT applies the GT predicate on the "delete_time" field.
func DeleteTimeGT(v int64) predicate.Attachment {
	return predicate.Attachment(sql.FieldGT(FieldDeleteTime, v))
}

// DeleteTimeGTE applies the GTE predicate on the "delete_time" field.
func DeleteTimeGTE(v int64) predicate.Attachment {
	return predicate.Attachment(sql.FieldGTE(FieldDeleteTime, v))
}

// DeleteTimeLT applies the LT predicate on the "delete_time" field.
func DeleteTimeLT(v int64) predicate.Attachment {
	return predicate.Attachment(sql.FieldLT(FieldDeleteTime, v))
}

// DeleteTimeLTE applies the LTE predicate on the "delete_time" field.
func DeleteTimeLTE(v int64) predicate.Attachment {
	return predicate.Attachment(sql.FieldLTE(FieldDeleteTime, v))
}

// DeleteTimeIsNil applies the IsNil predicate on the "delete_time" field.
func DeleteTimeIsNil() predicate.Attachment {
	return predicate.Attachment(sql.FieldIsNull(FieldDeleteTime))
}

// DeleteTimeNotNil applies the NotNil predicate on the "delete_time" field.
func DeleteTimeNotNil() predicate.Attachment {
	return predicate.Attachment(sql.FieldNotNull(FieldDeleteTime))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Attachment {
	return predicate.Attachment(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Attachment {
	return predicate.Attachment(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldHasSuffix(FieldName, v))
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.Attachment {
	return predicate.Attachment(sql.FieldIsNull(FieldName))
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.Attachment {
	return predicate.Attachment(sql.FieldNotNull(FieldName))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldContainsFold(FieldName, v))
}

// PathEQ applies the EQ predicate on the "path" field.
func PathEQ(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldEQ(FieldPath, v))
}

// PathNEQ applies the NEQ predicate on the "path" field.
func PathNEQ(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldNEQ(FieldPath, v))
}

// PathIn applies the In predicate on the "path" field.
func PathIn(vs ...string) predicate.Attachment {
	return predicate.Attachment(sql.FieldIn(FieldPath, vs...))
}

// PathNotIn applies the NotIn predicate on the "path" field.
func PathNotIn(vs ...string) predicate.Attachment {
	return predicate.Attachment(sql.FieldNotIn(FieldPath, vs...))
}

// PathGT applies the GT predicate on the "path" field.
func PathGT(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldGT(FieldPath, v))
}

// PathGTE applies the GTE predicate on the "path" field.
func PathGTE(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldGTE(FieldPath, v))
}

// PathLT applies the LT predicate on the "path" field.
func PathLT(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldLT(FieldPath, v))
}

// PathLTE applies the LTE predicate on the "path" field.
func PathLTE(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldLTE(FieldPath, v))
}

// PathContains applies the Contains predicate on the "path" field.
func PathContains(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldContains(FieldPath, v))
}

// PathHasPrefix applies the HasPrefix predicate on the "path" field.
func PathHasPrefix(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldHasPrefix(FieldPath, v))
}

// PathHasSuffix applies the HasSuffix predicate on the "path" field.
func PathHasSuffix(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldHasSuffix(FieldPath, v))
}

// PathIsNil applies the IsNil predicate on the "path" field.
func PathIsNil() predicate.Attachment {
	return predicate.Attachment(sql.FieldIsNull(FieldPath))
}

// PathNotNil applies the NotNil predicate on the "path" field.
func PathNotNil() predicate.Attachment {
	return predicate.Attachment(sql.FieldNotNull(FieldPath))
}

// PathEqualFold applies the EqualFold predicate on the "path" field.
func PathEqualFold(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldEqualFold(FieldPath, v))
}

// PathContainsFold applies the ContainsFold predicate on the "path" field.
func PathContainsFold(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldContainsFold(FieldPath, v))
}

// FileKeyEQ applies the EQ predicate on the "file_key" field.
func FileKeyEQ(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldEQ(FieldFileKey, v))
}

// FileKeyNEQ applies the NEQ predicate on the "file_key" field.
func FileKeyNEQ(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldNEQ(FieldFileKey, v))
}

// FileKeyIn applies the In predicate on the "file_key" field.
func FileKeyIn(vs ...string) predicate.Attachment {
	return predicate.Attachment(sql.FieldIn(FieldFileKey, vs...))
}

// FileKeyNotIn applies the NotIn predicate on the "file_key" field.
func FileKeyNotIn(vs ...string) predicate.Attachment {
	return predicate.Attachment(sql.FieldNotIn(FieldFileKey, vs...))
}

// FileKeyGT applies the GT predicate on the "file_key" field.
func FileKeyGT(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldGT(FieldFileKey, v))
}

// FileKeyGTE applies the GTE predicate on the "file_key" field.
func FileKeyGTE(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldGTE(FieldFileKey, v))
}

// FileKeyLT applies the LT predicate on the "file_key" field.
func FileKeyLT(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldLT(FieldFileKey, v))
}

// FileKeyLTE applies the LTE predicate on the "file_key" field.
func FileKeyLTE(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldLTE(FieldFileKey, v))
}

// FileKeyContains applies the Contains predicate on the "file_key" field.
func FileKeyContains(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldContains(FieldFileKey, v))
}

// FileKeyHasPrefix applies the HasPrefix predicate on the "file_key" field.
func FileKeyHasPrefix(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldHasPrefix(FieldFileKey, v))
}

// FileKeyHasSuffix applies the HasSuffix predicate on the "file_key" field.
func FileKeyHasSuffix(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldHasSuffix(FieldFileKey, v))
}

// FileKeyIsNil applies the IsNil predicate on the "file_key" field.
func FileKeyIsNil() predicate.Attachment {
	return predicate.Attachment(sql.FieldIsNull(FieldFileKey))
}

// FileKeyNotNil applies the NotNil predicate on the "file_key" field.
func FileKeyNotNil() predicate.Attachment {
	return predicate.Attachment(sql.FieldNotNull(FieldFileKey))
}

// FileKeyEqualFold applies the EqualFold predicate on the "file_key" field.
func FileKeyEqualFold(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldEqualFold(FieldFileKey, v))
}

// FileKeyContainsFold applies the ContainsFold predicate on the "file_key" field.
func FileKeyContainsFold(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldContainsFold(FieldFileKey, v))
}

// ThumbnailEQ applies the EQ predicate on the "thumbnail" field.
func ThumbnailEQ(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldEQ(FieldThumbnail, v))
}

// ThumbnailNEQ applies the NEQ predicate on the "thumbnail" field.
func ThumbnailNEQ(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldNEQ(FieldThumbnail, v))
}

// ThumbnailIn applies the In predicate on the "thumbnail" field.
func ThumbnailIn(vs ...string) predicate.Attachment {
	return predicate.Attachment(sql.FieldIn(FieldThumbnail, vs...))
}

// ThumbnailNotIn applies the NotIn predicate on the "thumbnail" field.
func ThumbnailNotIn(vs ...string) predicate.Attachment {
	return predicate.Attachment(sql.FieldNotIn(FieldThumbnail, vs...))
}

// ThumbnailGT applies the GT predicate on the "thumbnail" field.
func ThumbnailGT(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldGT(FieldThumbnail, v))
}

// ThumbnailGTE applies the GTE predicate on the "thumbnail" field.
func ThumbnailGTE(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldGTE(FieldThumbnail, v))
}

// ThumbnailLT applies the LT predicate on the "thumbnail" field.
func ThumbnailLT(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldLT(FieldThumbnail, v))
}

// ThumbnailLTE applies the LTE predicate on the "thumbnail" field.
func ThumbnailLTE(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldLTE(FieldThumbnail, v))
}

// ThumbnailContains applies the Contains predicate on the "thumbnail" field.
func ThumbnailContains(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldContains(FieldThumbnail, v))
}

// ThumbnailHasPrefix applies the HasPrefix predicate on the "thumbnail" field.
func ThumbnailHasPrefix(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldHasPrefix(FieldThumbnail, v))
}

// ThumbnailHasSuffix applies the HasSuffix predicate on the "thumbnail" field.
func ThumbnailHasSuffix(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldHasSuffix(FieldThumbnail, v))
}

// ThumbnailIsNil applies the IsNil predicate on the "thumbnail" field.
func ThumbnailIsNil() predicate.Attachment {
	return predicate.Attachment(sql.FieldIsNull(FieldThumbnail))
}

// ThumbnailNotNil applies the NotNil predicate on the "thumbnail" field.
func ThumbnailNotNil() predicate.Attachment {
	return predicate.Attachment(sql.FieldNotNull(FieldThumbnail))
}

// ThumbnailEqualFold applies the EqualFold predicate on the "thumbnail" field.
func ThumbnailEqualFold(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldEqualFold(FieldThumbnail, v))
}

// ThumbnailContainsFold applies the ContainsFold predicate on the "thumbnail" field.
func ThumbnailContainsFold(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldContainsFold(FieldThumbnail, v))
}

// MediaTypeEQ applies the EQ predicate on the "media_type" field.
func MediaTypeEQ(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldEQ(FieldMediaType, v))
}

// MediaTypeNEQ applies the NEQ predicate on the "media_type" field.
func MediaTypeNEQ(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldNEQ(FieldMediaType, v))
}

// MediaTypeIn applies the In predicate on the "media_type" field.
func MediaTypeIn(vs ...string) predicate.Attachment {
	return predicate.Attachment(sql.FieldIn(FieldMediaType, vs...))
}

// MediaTypeNotIn applies the NotIn predicate on the "media_type" field.
func MediaTypeNotIn(vs ...string) predicate.Attachment {
	return predicate.Attachment(sql.FieldNotIn(FieldMediaType, vs...))
}

// MediaTypeGT applies the GT predicate on the "media_type" field.
func MediaTypeGT(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldGT(FieldMediaType, v))
}

// MediaTypeGTE applies the GTE predicate on the "media_type" field.
func MediaTypeGTE(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldGTE(FieldMediaType, v))
}

// MediaTypeLT applies the LT predicate on the "media_type" field.
func MediaTypeLT(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldLT(FieldMediaType, v))
}

// MediaTypeLTE applies the LTE predicate on the "media_type" field.
func MediaTypeLTE(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldLTE(FieldMediaType, v))
}

// MediaTypeContains applies the Contains predicate on the "media_type" field.
func MediaTypeContains(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldContains(FieldMediaType, v))
}

// MediaTypeHasPrefix applies the HasPrefix predicate on the "media_type" field.
func MediaTypeHasPrefix(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldHasPrefix(FieldMediaType, v))
}

// MediaTypeHasSuffix applies the HasSuffix predicate on the "media_type" field.
func MediaTypeHasSuffix(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldHasSuffix(FieldMediaType, v))
}

// MediaTypeIsNil applies the IsNil predicate on the "media_type" field.
func MediaTypeIsNil() predicate.Attachment {
	return predicate.Attachment(sql.FieldIsNull(FieldMediaType))
}

// MediaTypeNotNil applies the NotNil predicate on the "media_type" field.
func MediaTypeNotNil() predicate.Attachment {
	return predicate.Attachment(sql.FieldNotNull(FieldMediaType))
}

// MediaTypeEqualFold applies the EqualFold predicate on the "media_type" field.
func MediaTypeEqualFold(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldEqualFold(FieldMediaType, v))
}

// MediaTypeContainsFold applies the ContainsFold predicate on the "media_type" field.
func MediaTypeContainsFold(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldContainsFold(FieldMediaType, v))
}

// SuffixEQ applies the EQ predicate on the "suffix" field.
func SuffixEQ(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldEQ(FieldSuffix, v))
}

// SuffixNEQ applies the NEQ predicate on the "suffix" field.
func SuffixNEQ(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldNEQ(FieldSuffix, v))
}

// SuffixIn applies the In predicate on the "suffix" field.
func SuffixIn(vs ...string) predicate.Attachment {
	return predicate.Attachment(sql.FieldIn(FieldSuffix, vs...))
}

// SuffixNotIn applies the NotIn predicate on the "suffix" field.
func SuffixNotIn(vs ...string) predicate.Attachment {
	return predicate.Attachment(sql.FieldNotIn(FieldSuffix, vs...))
}

// SuffixGT applies the GT predicate on the "suffix" field.
func SuffixGT(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldGT(FieldSuffix, v))
}

// SuffixGTE applies the GTE predicate on the "suffix" field.
func SuffixGTE(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldGTE(FieldSuffix, v))
}

// SuffixLT applies the LT predicate on the "suffix" field.
func SuffixLT(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldLT(FieldSuffix, v))
}

// SuffixLTE applies the LTE predicate on the "suffix" field.
func SuffixLTE(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldLTE(FieldSuffix, v))
}

// SuffixContains applies the Contains predicate on the "suffix" field.
func SuffixContains(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldContains(FieldSuffix, v))
}

// SuffixHasPrefix applies the HasPrefix predicate on the "suffix" field.
func SuffixHasPrefix(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldHasPrefix(FieldSuffix, v))
}

// SuffixHasSuffix applies the HasSuffix predicate on the "suffix" field.
func SuffixHasSuffix(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldHasSuffix(FieldSuffix, v))
}

// SuffixIsNil applies the IsNil predicate on the "suffix" field.
func SuffixIsNil() predicate.Attachment {
	return predicate.Attachment(sql.FieldIsNull(FieldSuffix))
}

// SuffixNotNil applies the NotNil predicate on the "suffix" field.
func SuffixNotNil() predicate.Attachment {
	return predicate.Attachment(sql.FieldNotNull(FieldSuffix))
}

// SuffixEqualFold applies the EqualFold predicate on the "suffix" field.
func SuffixEqualFold(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldEqualFold(FieldSuffix, v))
}

// SuffixContainsFold applies the ContainsFold predicate on the "suffix" field.
func SuffixContainsFold(v string) predicate.Attachment {
	return predicate.Attachment(sql.FieldContainsFold(FieldSuffix, v))
}

// WidthEQ applies the EQ predicate on the "width" field.
func WidthEQ(v int32) predicate.Attachment {
	return predicate.Attachment(sql.FieldEQ(FieldWidth, v))
}

// WidthNEQ applies the NEQ predicate on the "width" field.
func WidthNEQ(v int32) predicate.Attachment {
	return predicate.Attachment(sql.FieldNEQ(FieldWidth, v))
}

// WidthIn applies the In predicate on the "width" field.
func WidthIn(vs ...int32) predicate.Attachment {
	return predicate.Attachment(sql.FieldIn(FieldWidth, vs...))
}

// WidthNotIn applies the NotIn predicate on the "width" field.
func WidthNotIn(vs ...int32) predicate.Attachment {
	return predicate.Attachment(sql.FieldNotIn(FieldWidth, vs...))
}

// WidthGT applies the GT predicate on the "width" field.
func WidthGT(v int32) predicate.Attachment {
	return predicate.Attachment(sql.FieldGT(FieldWidth, v))
}

// WidthGTE applies the GTE predicate on the "width" field.
func WidthGTE(v int32) predicate.Attachment {
	return predicate.Attachment(sql.FieldGTE(FieldWidth, v))
}

// WidthLT applies the LT predicate on the "width" field.
func WidthLT(v int32) predicate.Attachment {
	return predicate.Attachment(sql.FieldLT(FieldWidth, v))
}

// WidthLTE applies the LTE predicate on the "width" field.
func WidthLTE(v int32) predicate.Attachment {
	return predicate.Attachment(sql.FieldLTE(FieldWidth, v))
}

// WidthIsNil applies the IsNil predicate on the "width" field.
func WidthIsNil() predicate.Attachment {
	return predicate.Attachment(sql.FieldIsNull(FieldWidth))
}

// WidthNotNil applies the NotNil predicate on the "width" field.
func WidthNotNil() predicate.Attachment {
	return predicate.Attachment(sql.FieldNotNull(FieldWidth))
}

// HeightEQ applies the EQ predicate on the "height" field.
func HeightEQ(v int32) predicate.Attachment {
	return predicate.Attachment(sql.FieldEQ(FieldHeight, v))
}

// HeightNEQ applies the NEQ predicate on the "height" field.
func HeightNEQ(v int32) predicate.Attachment {
	return predicate.Attachment(sql.FieldNEQ(FieldHeight, v))
}

// HeightIn applies the In predicate on the "height" field.
func HeightIn(vs ...int32) predicate.Attachment {
	return predicate.Attachment(sql.FieldIn(FieldHeight, vs...))
}

// HeightNotIn applies the NotIn predicate on the "height" field.
func HeightNotIn(vs ...int32) predicate.Attachment {
	return predicate.Attachment(sql.FieldNotIn(FieldHeight, vs...))
}

// HeightGT applies the GT predicate on the "height" field.
func HeightGT(v int32) predicate.Attachment {
	return predicate.Attachment(sql.FieldGT(FieldHeight, v))
}

// HeightGTE applies the GTE predicate on the "height" field.
func HeightGTE(v int32) predicate.Attachment {
	return predicate.Attachment(sql.FieldGTE(FieldHeight, v))
}

// HeightLT applies the LT predicate on the "height" field.
func HeightLT(v int32) predicate.Attachment {
	return predicate.Attachment(sql.FieldLT(FieldHeight, v))
}

// HeightLTE applies the LTE predicate on the "height" field.
func HeightLTE(v int32) predicate.Attachment {
	return predicate.Attachment(sql.FieldLTE(FieldHeight, v))
}

// HeightIsNil applies the IsNil predicate on the "height" field.
func HeightIsNil() predicate.Attachment {
	return predicate.Attachment(sql.FieldIsNull(FieldHeight))
}

// HeightNotNil applies the NotNil predicate on the "height" field.
func HeightNotNil() predicate.Attachment {
	return predicate.Attachment(sql.FieldNotNull(FieldHeight))
}

// SizeEQ applies the EQ predicate on the "Size" field.
func SizeEQ(v uint64) predicate.Attachment {
	return predicate.Attachment(sql.FieldEQ(FieldSize, v))
}

// SizeNEQ applies the NEQ predicate on the "Size" field.
func SizeNEQ(v uint64) predicate.Attachment {
	return predicate.Attachment(sql.FieldNEQ(FieldSize, v))
}

// SizeIn applies the In predicate on the "Size" field.
func SizeIn(vs ...uint64) predicate.Attachment {
	return predicate.Attachment(sql.FieldIn(FieldSize, vs...))
}

// SizeNotIn applies the NotIn predicate on the "Size" field.
func SizeNotIn(vs ...uint64) predicate.Attachment {
	return predicate.Attachment(sql.FieldNotIn(FieldSize, vs...))
}

// SizeGT applies the GT predicate on the "Size" field.
func SizeGT(v uint64) predicate.Attachment {
	return predicate.Attachment(sql.FieldGT(FieldSize, v))
}

// SizeGTE applies the GTE predicate on the "Size" field.
func SizeGTE(v uint64) predicate.Attachment {
	return predicate.Attachment(sql.FieldGTE(FieldSize, v))
}

// SizeLT applies the LT predicate on the "Size" field.
func SizeLT(v uint64) predicate.Attachment {
	return predicate.Attachment(sql.FieldLT(FieldSize, v))
}

// SizeLTE applies the LTE predicate on the "Size" field.
func SizeLTE(v uint64) predicate.Attachment {
	return predicate.Attachment(sql.FieldLTE(FieldSize, v))
}

// SizeIsNil applies the IsNil predicate on the "Size" field.
func SizeIsNil() predicate.Attachment {
	return predicate.Attachment(sql.FieldIsNull(FieldSize))
}

// SizeNotNil applies the NotNil predicate on the "Size" field.
func SizeNotNil() predicate.Attachment {
	return predicate.Attachment(sql.FieldNotNull(FieldSize))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v int32) predicate.Attachment {
	return predicate.Attachment(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v int32) predicate.Attachment {
	return predicate.Attachment(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...int32) predicate.Attachment {
	return predicate.Attachment(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...int32) predicate.Attachment {
	return predicate.Attachment(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v int32) predicate.Attachment {
	return predicate.Attachment(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v int32) predicate.Attachment {
	return predicate.Attachment(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v int32) predicate.Attachment {
	return predicate.Attachment(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v int32) predicate.Attachment {
	return predicate.Attachment(sql.FieldLTE(FieldType, v))
}

// TypeIsNil applies the IsNil predicate on the "type" field.
func TypeIsNil() predicate.Attachment {
	return predicate.Attachment(sql.FieldIsNull(FieldType))
}

// TypeNotNil applies the NotNil predicate on the "type" field.
func TypeNotNil() predicate.Attachment {
	return predicate.Attachment(sql.FieldNotNull(FieldType))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Attachment) predicate.Attachment {
	return predicate.Attachment(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Attachment) predicate.Attachment {
	return predicate.Attachment(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Attachment) predicate.Attachment {
	return predicate.Attachment(sql.NotPredicates(p))
}
