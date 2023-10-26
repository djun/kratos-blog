// Code generated by ent, DO NOT EDIT.

package category

import (
	"kratos-cms/app/core/service/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.Category {
	return predicate.Category(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.Category {
	return predicate.Category(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.Category {
	return predicate.Category(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.Category {
	return predicate.Category(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.Category {
	return predicate.Category(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.Category {
	return predicate.Category(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.Category {
	return predicate.Category(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v int64) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v int64) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldUpdateTime, v))
}

// DeleteTime applies equality check predicate on the "delete_time" field. It's identical to DeleteTimeEQ.
func DeleteTime(v int64) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldDeleteTime, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldName, v))
}

// Slug applies equality check predicate on the "slug" field. It's identical to SlugEQ.
func Slug(v string) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldSlug, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldDescription, v))
}

// Thumbnail applies equality check predicate on the "thumbnail" field. It's identical to ThumbnailEQ.
func Thumbnail(v string) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldThumbnail, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldPassword, v))
}

// FullPath applies equality check predicate on the "full_path" field. It's identical to FullPathEQ.
func FullPath(v string) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldFullPath, v))
}

// ParentID applies equality check predicate on the "parent_id" field. It's identical to ParentIDEQ.
func ParentID(v uint32) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldParentID, v))
}

// Priority applies equality check predicate on the "priority" field. It's identical to PriorityEQ.
func Priority(v int32) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldPriority, v))
}

// PostCount applies equality check predicate on the "post_count" field. It's identical to PostCountEQ.
func PostCount(v uint32) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldPostCount, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v int64) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v int64) predicate.Category {
	return predicate.Category(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...int64) predicate.Category {
	return predicate.Category(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...int64) predicate.Category {
	return predicate.Category(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v int64) predicate.Category {
	return predicate.Category(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v int64) predicate.Category {
	return predicate.Category(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v int64) predicate.Category {
	return predicate.Category(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v int64) predicate.Category {
	return predicate.Category(sql.FieldLTE(FieldCreateTime, v))
}

// CreateTimeIsNil applies the IsNil predicate on the "create_time" field.
func CreateTimeIsNil() predicate.Category {
	return predicate.Category(sql.FieldIsNull(FieldCreateTime))
}

// CreateTimeNotNil applies the NotNil predicate on the "create_time" field.
func CreateTimeNotNil() predicate.Category {
	return predicate.Category(sql.FieldNotNull(FieldCreateTime))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v int64) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v int64) predicate.Category {
	return predicate.Category(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...int64) predicate.Category {
	return predicate.Category(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...int64) predicate.Category {
	return predicate.Category(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v int64) predicate.Category {
	return predicate.Category(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v int64) predicate.Category {
	return predicate.Category(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v int64) predicate.Category {
	return predicate.Category(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v int64) predicate.Category {
	return predicate.Category(sql.FieldLTE(FieldUpdateTime, v))
}

// UpdateTimeIsNil applies the IsNil predicate on the "update_time" field.
func UpdateTimeIsNil() predicate.Category {
	return predicate.Category(sql.FieldIsNull(FieldUpdateTime))
}

// UpdateTimeNotNil applies the NotNil predicate on the "update_time" field.
func UpdateTimeNotNil() predicate.Category {
	return predicate.Category(sql.FieldNotNull(FieldUpdateTime))
}

// DeleteTimeEQ applies the EQ predicate on the "delete_time" field.
func DeleteTimeEQ(v int64) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldDeleteTime, v))
}

// DeleteTimeNEQ applies the NEQ predicate on the "delete_time" field.
func DeleteTimeNEQ(v int64) predicate.Category {
	return predicate.Category(sql.FieldNEQ(FieldDeleteTime, v))
}

// DeleteTimeIn applies the In predicate on the "delete_time" field.
func DeleteTimeIn(vs ...int64) predicate.Category {
	return predicate.Category(sql.FieldIn(FieldDeleteTime, vs...))
}

// DeleteTimeNotIn applies the NotIn predicate on the "delete_time" field.
func DeleteTimeNotIn(vs ...int64) predicate.Category {
	return predicate.Category(sql.FieldNotIn(FieldDeleteTime, vs...))
}

// DeleteTimeGT applies the GT predicate on the "delete_time" field.
func DeleteTimeGT(v int64) predicate.Category {
	return predicate.Category(sql.FieldGT(FieldDeleteTime, v))
}

// DeleteTimeGTE applies the GTE predicate on the "delete_time" field.
func DeleteTimeGTE(v int64) predicate.Category {
	return predicate.Category(sql.FieldGTE(FieldDeleteTime, v))
}

// DeleteTimeLT applies the LT predicate on the "delete_time" field.
func DeleteTimeLT(v int64) predicate.Category {
	return predicate.Category(sql.FieldLT(FieldDeleteTime, v))
}

// DeleteTimeLTE applies the LTE predicate on the "delete_time" field.
func DeleteTimeLTE(v int64) predicate.Category {
	return predicate.Category(sql.FieldLTE(FieldDeleteTime, v))
}

// DeleteTimeIsNil applies the IsNil predicate on the "delete_time" field.
func DeleteTimeIsNil() predicate.Category {
	return predicate.Category(sql.FieldIsNull(FieldDeleteTime))
}

// DeleteTimeNotNil applies the NotNil predicate on the "delete_time" field.
func DeleteTimeNotNil() predicate.Category {
	return predicate.Category(sql.FieldNotNull(FieldDeleteTime))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Category {
	return predicate.Category(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Category {
	return predicate.Category(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Category {
	return predicate.Category(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Category {
	return predicate.Category(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Category {
	return predicate.Category(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Category {
	return predicate.Category(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Category {
	return predicate.Category(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Category {
	return predicate.Category(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Category {
	return predicate.Category(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Category {
	return predicate.Category(sql.FieldHasSuffix(FieldName, v))
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.Category {
	return predicate.Category(sql.FieldIsNull(FieldName))
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.Category {
	return predicate.Category(sql.FieldNotNull(FieldName))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Category {
	return predicate.Category(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Category {
	return predicate.Category(sql.FieldContainsFold(FieldName, v))
}

// SlugEQ applies the EQ predicate on the "slug" field.
func SlugEQ(v string) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldSlug, v))
}

// SlugNEQ applies the NEQ predicate on the "slug" field.
func SlugNEQ(v string) predicate.Category {
	return predicate.Category(sql.FieldNEQ(FieldSlug, v))
}

// SlugIn applies the In predicate on the "slug" field.
func SlugIn(vs ...string) predicate.Category {
	return predicate.Category(sql.FieldIn(FieldSlug, vs...))
}

// SlugNotIn applies the NotIn predicate on the "slug" field.
func SlugNotIn(vs ...string) predicate.Category {
	return predicate.Category(sql.FieldNotIn(FieldSlug, vs...))
}

// SlugGT applies the GT predicate on the "slug" field.
func SlugGT(v string) predicate.Category {
	return predicate.Category(sql.FieldGT(FieldSlug, v))
}

// SlugGTE applies the GTE predicate on the "slug" field.
func SlugGTE(v string) predicate.Category {
	return predicate.Category(sql.FieldGTE(FieldSlug, v))
}

// SlugLT applies the LT predicate on the "slug" field.
func SlugLT(v string) predicate.Category {
	return predicate.Category(sql.FieldLT(FieldSlug, v))
}

// SlugLTE applies the LTE predicate on the "slug" field.
func SlugLTE(v string) predicate.Category {
	return predicate.Category(sql.FieldLTE(FieldSlug, v))
}

// SlugContains applies the Contains predicate on the "slug" field.
func SlugContains(v string) predicate.Category {
	return predicate.Category(sql.FieldContains(FieldSlug, v))
}

// SlugHasPrefix applies the HasPrefix predicate on the "slug" field.
func SlugHasPrefix(v string) predicate.Category {
	return predicate.Category(sql.FieldHasPrefix(FieldSlug, v))
}

// SlugHasSuffix applies the HasSuffix predicate on the "slug" field.
func SlugHasSuffix(v string) predicate.Category {
	return predicate.Category(sql.FieldHasSuffix(FieldSlug, v))
}

// SlugIsNil applies the IsNil predicate on the "slug" field.
func SlugIsNil() predicate.Category {
	return predicate.Category(sql.FieldIsNull(FieldSlug))
}

// SlugNotNil applies the NotNil predicate on the "slug" field.
func SlugNotNil() predicate.Category {
	return predicate.Category(sql.FieldNotNull(FieldSlug))
}

// SlugEqualFold applies the EqualFold predicate on the "slug" field.
func SlugEqualFold(v string) predicate.Category {
	return predicate.Category(sql.FieldEqualFold(FieldSlug, v))
}

// SlugContainsFold applies the ContainsFold predicate on the "slug" field.
func SlugContainsFold(v string) predicate.Category {
	return predicate.Category(sql.FieldContainsFold(FieldSlug, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Category {
	return predicate.Category(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Category {
	return predicate.Category(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Category {
	return predicate.Category(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Category {
	return predicate.Category(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Category {
	return predicate.Category(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Category {
	return predicate.Category(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Category {
	return predicate.Category(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Category {
	return predicate.Category(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Category {
	return predicate.Category(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Category {
	return predicate.Category(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Category {
	return predicate.Category(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Category {
	return predicate.Category(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Category {
	return predicate.Category(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Category {
	return predicate.Category(sql.FieldContainsFold(FieldDescription, v))
}

// ThumbnailEQ applies the EQ predicate on the "thumbnail" field.
func ThumbnailEQ(v string) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldThumbnail, v))
}

// ThumbnailNEQ applies the NEQ predicate on the "thumbnail" field.
func ThumbnailNEQ(v string) predicate.Category {
	return predicate.Category(sql.FieldNEQ(FieldThumbnail, v))
}

// ThumbnailIn applies the In predicate on the "thumbnail" field.
func ThumbnailIn(vs ...string) predicate.Category {
	return predicate.Category(sql.FieldIn(FieldThumbnail, vs...))
}

// ThumbnailNotIn applies the NotIn predicate on the "thumbnail" field.
func ThumbnailNotIn(vs ...string) predicate.Category {
	return predicate.Category(sql.FieldNotIn(FieldThumbnail, vs...))
}

// ThumbnailGT applies the GT predicate on the "thumbnail" field.
func ThumbnailGT(v string) predicate.Category {
	return predicate.Category(sql.FieldGT(FieldThumbnail, v))
}

// ThumbnailGTE applies the GTE predicate on the "thumbnail" field.
func ThumbnailGTE(v string) predicate.Category {
	return predicate.Category(sql.FieldGTE(FieldThumbnail, v))
}

// ThumbnailLT applies the LT predicate on the "thumbnail" field.
func ThumbnailLT(v string) predicate.Category {
	return predicate.Category(sql.FieldLT(FieldThumbnail, v))
}

// ThumbnailLTE applies the LTE predicate on the "thumbnail" field.
func ThumbnailLTE(v string) predicate.Category {
	return predicate.Category(sql.FieldLTE(FieldThumbnail, v))
}

// ThumbnailContains applies the Contains predicate on the "thumbnail" field.
func ThumbnailContains(v string) predicate.Category {
	return predicate.Category(sql.FieldContains(FieldThumbnail, v))
}

// ThumbnailHasPrefix applies the HasPrefix predicate on the "thumbnail" field.
func ThumbnailHasPrefix(v string) predicate.Category {
	return predicate.Category(sql.FieldHasPrefix(FieldThumbnail, v))
}

// ThumbnailHasSuffix applies the HasSuffix predicate on the "thumbnail" field.
func ThumbnailHasSuffix(v string) predicate.Category {
	return predicate.Category(sql.FieldHasSuffix(FieldThumbnail, v))
}

// ThumbnailIsNil applies the IsNil predicate on the "thumbnail" field.
func ThumbnailIsNil() predicate.Category {
	return predicate.Category(sql.FieldIsNull(FieldThumbnail))
}

// ThumbnailNotNil applies the NotNil predicate on the "thumbnail" field.
func ThumbnailNotNil() predicate.Category {
	return predicate.Category(sql.FieldNotNull(FieldThumbnail))
}

// ThumbnailEqualFold applies the EqualFold predicate on the "thumbnail" field.
func ThumbnailEqualFold(v string) predicate.Category {
	return predicate.Category(sql.FieldEqualFold(FieldThumbnail, v))
}

// ThumbnailContainsFold applies the ContainsFold predicate on the "thumbnail" field.
func ThumbnailContainsFold(v string) predicate.Category {
	return predicate.Category(sql.FieldContainsFold(FieldThumbnail, v))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.Category {
	return predicate.Category(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.Category {
	return predicate.Category(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.Category {
	return predicate.Category(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.Category {
	return predicate.Category(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.Category {
	return predicate.Category(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.Category {
	return predicate.Category(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.Category {
	return predicate.Category(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.Category {
	return predicate.Category(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.Category {
	return predicate.Category(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.Category {
	return predicate.Category(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordIsNil applies the IsNil predicate on the "password" field.
func PasswordIsNil() predicate.Category {
	return predicate.Category(sql.FieldIsNull(FieldPassword))
}

// PasswordNotNil applies the NotNil predicate on the "password" field.
func PasswordNotNil() predicate.Category {
	return predicate.Category(sql.FieldNotNull(FieldPassword))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.Category {
	return predicate.Category(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.Category {
	return predicate.Category(sql.FieldContainsFold(FieldPassword, v))
}

// FullPathEQ applies the EQ predicate on the "full_path" field.
func FullPathEQ(v string) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldFullPath, v))
}

// FullPathNEQ applies the NEQ predicate on the "full_path" field.
func FullPathNEQ(v string) predicate.Category {
	return predicate.Category(sql.FieldNEQ(FieldFullPath, v))
}

// FullPathIn applies the In predicate on the "full_path" field.
func FullPathIn(vs ...string) predicate.Category {
	return predicate.Category(sql.FieldIn(FieldFullPath, vs...))
}

// FullPathNotIn applies the NotIn predicate on the "full_path" field.
func FullPathNotIn(vs ...string) predicate.Category {
	return predicate.Category(sql.FieldNotIn(FieldFullPath, vs...))
}

// FullPathGT applies the GT predicate on the "full_path" field.
func FullPathGT(v string) predicate.Category {
	return predicate.Category(sql.FieldGT(FieldFullPath, v))
}

// FullPathGTE applies the GTE predicate on the "full_path" field.
func FullPathGTE(v string) predicate.Category {
	return predicate.Category(sql.FieldGTE(FieldFullPath, v))
}

// FullPathLT applies the LT predicate on the "full_path" field.
func FullPathLT(v string) predicate.Category {
	return predicate.Category(sql.FieldLT(FieldFullPath, v))
}

// FullPathLTE applies the LTE predicate on the "full_path" field.
func FullPathLTE(v string) predicate.Category {
	return predicate.Category(sql.FieldLTE(FieldFullPath, v))
}

// FullPathContains applies the Contains predicate on the "full_path" field.
func FullPathContains(v string) predicate.Category {
	return predicate.Category(sql.FieldContains(FieldFullPath, v))
}

// FullPathHasPrefix applies the HasPrefix predicate on the "full_path" field.
func FullPathHasPrefix(v string) predicate.Category {
	return predicate.Category(sql.FieldHasPrefix(FieldFullPath, v))
}

// FullPathHasSuffix applies the HasSuffix predicate on the "full_path" field.
func FullPathHasSuffix(v string) predicate.Category {
	return predicate.Category(sql.FieldHasSuffix(FieldFullPath, v))
}

// FullPathIsNil applies the IsNil predicate on the "full_path" field.
func FullPathIsNil() predicate.Category {
	return predicate.Category(sql.FieldIsNull(FieldFullPath))
}

// FullPathNotNil applies the NotNil predicate on the "full_path" field.
func FullPathNotNil() predicate.Category {
	return predicate.Category(sql.FieldNotNull(FieldFullPath))
}

// FullPathEqualFold applies the EqualFold predicate on the "full_path" field.
func FullPathEqualFold(v string) predicate.Category {
	return predicate.Category(sql.FieldEqualFold(FieldFullPath, v))
}

// FullPathContainsFold applies the ContainsFold predicate on the "full_path" field.
func FullPathContainsFold(v string) predicate.Category {
	return predicate.Category(sql.FieldContainsFold(FieldFullPath, v))
}

// ParentIDEQ applies the EQ predicate on the "parent_id" field.
func ParentIDEQ(v uint32) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldParentID, v))
}

// ParentIDNEQ applies the NEQ predicate on the "parent_id" field.
func ParentIDNEQ(v uint32) predicate.Category {
	return predicate.Category(sql.FieldNEQ(FieldParentID, v))
}

// ParentIDIn applies the In predicate on the "parent_id" field.
func ParentIDIn(vs ...uint32) predicate.Category {
	return predicate.Category(sql.FieldIn(FieldParentID, vs...))
}

// ParentIDNotIn applies the NotIn predicate on the "parent_id" field.
func ParentIDNotIn(vs ...uint32) predicate.Category {
	return predicate.Category(sql.FieldNotIn(FieldParentID, vs...))
}

// ParentIDGT applies the GT predicate on the "parent_id" field.
func ParentIDGT(v uint32) predicate.Category {
	return predicate.Category(sql.FieldGT(FieldParentID, v))
}

// ParentIDGTE applies the GTE predicate on the "parent_id" field.
func ParentIDGTE(v uint32) predicate.Category {
	return predicate.Category(sql.FieldGTE(FieldParentID, v))
}

// ParentIDLT applies the LT predicate on the "parent_id" field.
func ParentIDLT(v uint32) predicate.Category {
	return predicate.Category(sql.FieldLT(FieldParentID, v))
}

// ParentIDLTE applies the LTE predicate on the "parent_id" field.
func ParentIDLTE(v uint32) predicate.Category {
	return predicate.Category(sql.FieldLTE(FieldParentID, v))
}

// ParentIDIsNil applies the IsNil predicate on the "parent_id" field.
func ParentIDIsNil() predicate.Category {
	return predicate.Category(sql.FieldIsNull(FieldParentID))
}

// ParentIDNotNil applies the NotNil predicate on the "parent_id" field.
func ParentIDNotNil() predicate.Category {
	return predicate.Category(sql.FieldNotNull(FieldParentID))
}

// PriorityEQ applies the EQ predicate on the "priority" field.
func PriorityEQ(v int32) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldPriority, v))
}

// PriorityNEQ applies the NEQ predicate on the "priority" field.
func PriorityNEQ(v int32) predicate.Category {
	return predicate.Category(sql.FieldNEQ(FieldPriority, v))
}

// PriorityIn applies the In predicate on the "priority" field.
func PriorityIn(vs ...int32) predicate.Category {
	return predicate.Category(sql.FieldIn(FieldPriority, vs...))
}

// PriorityNotIn applies the NotIn predicate on the "priority" field.
func PriorityNotIn(vs ...int32) predicate.Category {
	return predicate.Category(sql.FieldNotIn(FieldPriority, vs...))
}

// PriorityGT applies the GT predicate on the "priority" field.
func PriorityGT(v int32) predicate.Category {
	return predicate.Category(sql.FieldGT(FieldPriority, v))
}

// PriorityGTE applies the GTE predicate on the "priority" field.
func PriorityGTE(v int32) predicate.Category {
	return predicate.Category(sql.FieldGTE(FieldPriority, v))
}

// PriorityLT applies the LT predicate on the "priority" field.
func PriorityLT(v int32) predicate.Category {
	return predicate.Category(sql.FieldLT(FieldPriority, v))
}

// PriorityLTE applies the LTE predicate on the "priority" field.
func PriorityLTE(v int32) predicate.Category {
	return predicate.Category(sql.FieldLTE(FieldPriority, v))
}

// PriorityIsNil applies the IsNil predicate on the "priority" field.
func PriorityIsNil() predicate.Category {
	return predicate.Category(sql.FieldIsNull(FieldPriority))
}

// PriorityNotNil applies the NotNil predicate on the "priority" field.
func PriorityNotNil() predicate.Category {
	return predicate.Category(sql.FieldNotNull(FieldPriority))
}

// PostCountEQ applies the EQ predicate on the "post_count" field.
func PostCountEQ(v uint32) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldPostCount, v))
}

// PostCountNEQ applies the NEQ predicate on the "post_count" field.
func PostCountNEQ(v uint32) predicate.Category {
	return predicate.Category(sql.FieldNEQ(FieldPostCount, v))
}

// PostCountIn applies the In predicate on the "post_count" field.
func PostCountIn(vs ...uint32) predicate.Category {
	return predicate.Category(sql.FieldIn(FieldPostCount, vs...))
}

// PostCountNotIn applies the NotIn predicate on the "post_count" field.
func PostCountNotIn(vs ...uint32) predicate.Category {
	return predicate.Category(sql.FieldNotIn(FieldPostCount, vs...))
}

// PostCountGT applies the GT predicate on the "post_count" field.
func PostCountGT(v uint32) predicate.Category {
	return predicate.Category(sql.FieldGT(FieldPostCount, v))
}

// PostCountGTE applies the GTE predicate on the "post_count" field.
func PostCountGTE(v uint32) predicate.Category {
	return predicate.Category(sql.FieldGTE(FieldPostCount, v))
}

// PostCountLT applies the LT predicate on the "post_count" field.
func PostCountLT(v uint32) predicate.Category {
	return predicate.Category(sql.FieldLT(FieldPostCount, v))
}

// PostCountLTE applies the LTE predicate on the "post_count" field.
func PostCountLTE(v uint32) predicate.Category {
	return predicate.Category(sql.FieldLTE(FieldPostCount, v))
}

// PostCountIsNil applies the IsNil predicate on the "post_count" field.
func PostCountIsNil() predicate.Category {
	return predicate.Category(sql.FieldIsNull(FieldPostCount))
}

// PostCountNotNil applies the NotNil predicate on the "post_count" field.
func PostCountNotNil() predicate.Category {
	return predicate.Category(sql.FieldNotNull(FieldPostCount))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Category) predicate.Category {
	return predicate.Category(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Category) predicate.Category {
	return predicate.Category(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Category) predicate.Category {
	return predicate.Category(sql.NotPredicates(p))
}
