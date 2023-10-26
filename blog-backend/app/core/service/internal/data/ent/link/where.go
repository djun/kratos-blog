// Code generated by ent, DO NOT EDIT.

package link

import (
	"kratos-cms/app/core/service/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.Link {
	return predicate.Link(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.Link {
	return predicate.Link(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.Link {
	return predicate.Link(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.Link {
	return predicate.Link(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.Link {
	return predicate.Link(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.Link {
	return predicate.Link(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.Link {
	return predicate.Link(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.Link {
	return predicate.Link(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.Link {
	return predicate.Link(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v int64) predicate.Link {
	return predicate.Link(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v int64) predicate.Link {
	return predicate.Link(sql.FieldEQ(FieldUpdateTime, v))
}

// DeleteTime applies equality check predicate on the "delete_time" field. It's identical to DeleteTimeEQ.
func DeleteTime(v int64) predicate.Link {
	return predicate.Link(sql.FieldEQ(FieldDeleteTime, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Link {
	return predicate.Link(sql.FieldEQ(FieldName, v))
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.Link {
	return predicate.Link(sql.FieldEQ(FieldURL, v))
}

// Logo applies equality check predicate on the "logo" field. It's identical to LogoEQ.
func Logo(v string) predicate.Link {
	return predicate.Link(sql.FieldEQ(FieldLogo, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Link {
	return predicate.Link(sql.FieldEQ(FieldDescription, v))
}

// Team applies equality check predicate on the "team" field. It's identical to TeamEQ.
func Team(v string) predicate.Link {
	return predicate.Link(sql.FieldEQ(FieldTeam, v))
}

// Priority applies equality check predicate on the "priority" field. It's identical to PriorityEQ.
func Priority(v int32) predicate.Link {
	return predicate.Link(sql.FieldEQ(FieldPriority, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v int64) predicate.Link {
	return predicate.Link(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v int64) predicate.Link {
	return predicate.Link(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...int64) predicate.Link {
	return predicate.Link(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...int64) predicate.Link {
	return predicate.Link(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v int64) predicate.Link {
	return predicate.Link(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v int64) predicate.Link {
	return predicate.Link(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v int64) predicate.Link {
	return predicate.Link(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v int64) predicate.Link {
	return predicate.Link(sql.FieldLTE(FieldCreateTime, v))
}

// CreateTimeIsNil applies the IsNil predicate on the "create_time" field.
func CreateTimeIsNil() predicate.Link {
	return predicate.Link(sql.FieldIsNull(FieldCreateTime))
}

// CreateTimeNotNil applies the NotNil predicate on the "create_time" field.
func CreateTimeNotNil() predicate.Link {
	return predicate.Link(sql.FieldNotNull(FieldCreateTime))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v int64) predicate.Link {
	return predicate.Link(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v int64) predicate.Link {
	return predicate.Link(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...int64) predicate.Link {
	return predicate.Link(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...int64) predicate.Link {
	return predicate.Link(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v int64) predicate.Link {
	return predicate.Link(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v int64) predicate.Link {
	return predicate.Link(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v int64) predicate.Link {
	return predicate.Link(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v int64) predicate.Link {
	return predicate.Link(sql.FieldLTE(FieldUpdateTime, v))
}

// UpdateTimeIsNil applies the IsNil predicate on the "update_time" field.
func UpdateTimeIsNil() predicate.Link {
	return predicate.Link(sql.FieldIsNull(FieldUpdateTime))
}

// UpdateTimeNotNil applies the NotNil predicate on the "update_time" field.
func UpdateTimeNotNil() predicate.Link {
	return predicate.Link(sql.FieldNotNull(FieldUpdateTime))
}

// DeleteTimeEQ applies the EQ predicate on the "delete_time" field.
func DeleteTimeEQ(v int64) predicate.Link {
	return predicate.Link(sql.FieldEQ(FieldDeleteTime, v))
}

// DeleteTimeNEQ applies the NEQ predicate on the "delete_time" field.
func DeleteTimeNEQ(v int64) predicate.Link {
	return predicate.Link(sql.FieldNEQ(FieldDeleteTime, v))
}

// DeleteTimeIn applies the In predicate on the "delete_time" field.
func DeleteTimeIn(vs ...int64) predicate.Link {
	return predicate.Link(sql.FieldIn(FieldDeleteTime, vs...))
}

// DeleteTimeNotIn applies the NotIn predicate on the "delete_time" field.
func DeleteTimeNotIn(vs ...int64) predicate.Link {
	return predicate.Link(sql.FieldNotIn(FieldDeleteTime, vs...))
}

// DeleteTimeGT applies the GT predicate on the "delete_time" field.
func DeleteTimeGT(v int64) predicate.Link {
	return predicate.Link(sql.FieldGT(FieldDeleteTime, v))
}

// DeleteTimeGTE applies the GTE predicate on the "delete_time" field.
func DeleteTimeGTE(v int64) predicate.Link {
	return predicate.Link(sql.FieldGTE(FieldDeleteTime, v))
}

// DeleteTimeLT applies the LT predicate on the "delete_time" field.
func DeleteTimeLT(v int64) predicate.Link {
	return predicate.Link(sql.FieldLT(FieldDeleteTime, v))
}

// DeleteTimeLTE applies the LTE predicate on the "delete_time" field.
func DeleteTimeLTE(v int64) predicate.Link {
	return predicate.Link(sql.FieldLTE(FieldDeleteTime, v))
}

// DeleteTimeIsNil applies the IsNil predicate on the "delete_time" field.
func DeleteTimeIsNil() predicate.Link {
	return predicate.Link(sql.FieldIsNull(FieldDeleteTime))
}

// DeleteTimeNotNil applies the NotNil predicate on the "delete_time" field.
func DeleteTimeNotNil() predicate.Link {
	return predicate.Link(sql.FieldNotNull(FieldDeleteTime))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Link {
	return predicate.Link(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Link {
	return predicate.Link(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Link {
	return predicate.Link(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Link {
	return predicate.Link(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Link {
	return predicate.Link(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Link {
	return predicate.Link(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Link {
	return predicate.Link(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Link {
	return predicate.Link(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Link {
	return predicate.Link(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Link {
	return predicate.Link(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Link {
	return predicate.Link(sql.FieldHasSuffix(FieldName, v))
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.Link {
	return predicate.Link(sql.FieldIsNull(FieldName))
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.Link {
	return predicate.Link(sql.FieldNotNull(FieldName))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Link {
	return predicate.Link(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Link {
	return predicate.Link(sql.FieldContainsFold(FieldName, v))
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.Link {
	return predicate.Link(sql.FieldEQ(FieldURL, v))
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.Link {
	return predicate.Link(sql.FieldNEQ(FieldURL, v))
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.Link {
	return predicate.Link(sql.FieldIn(FieldURL, vs...))
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.Link {
	return predicate.Link(sql.FieldNotIn(FieldURL, vs...))
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.Link {
	return predicate.Link(sql.FieldGT(FieldURL, v))
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.Link {
	return predicate.Link(sql.FieldGTE(FieldURL, v))
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.Link {
	return predicate.Link(sql.FieldLT(FieldURL, v))
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.Link {
	return predicate.Link(sql.FieldLTE(FieldURL, v))
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.Link {
	return predicate.Link(sql.FieldContains(FieldURL, v))
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.Link {
	return predicate.Link(sql.FieldHasPrefix(FieldURL, v))
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.Link {
	return predicate.Link(sql.FieldHasSuffix(FieldURL, v))
}

// URLIsNil applies the IsNil predicate on the "url" field.
func URLIsNil() predicate.Link {
	return predicate.Link(sql.FieldIsNull(FieldURL))
}

// URLNotNil applies the NotNil predicate on the "url" field.
func URLNotNil() predicate.Link {
	return predicate.Link(sql.FieldNotNull(FieldURL))
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.Link {
	return predicate.Link(sql.FieldEqualFold(FieldURL, v))
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.Link {
	return predicate.Link(sql.FieldContainsFold(FieldURL, v))
}

// LogoEQ applies the EQ predicate on the "logo" field.
func LogoEQ(v string) predicate.Link {
	return predicate.Link(sql.FieldEQ(FieldLogo, v))
}

// LogoNEQ applies the NEQ predicate on the "logo" field.
func LogoNEQ(v string) predicate.Link {
	return predicate.Link(sql.FieldNEQ(FieldLogo, v))
}

// LogoIn applies the In predicate on the "logo" field.
func LogoIn(vs ...string) predicate.Link {
	return predicate.Link(sql.FieldIn(FieldLogo, vs...))
}

// LogoNotIn applies the NotIn predicate on the "logo" field.
func LogoNotIn(vs ...string) predicate.Link {
	return predicate.Link(sql.FieldNotIn(FieldLogo, vs...))
}

// LogoGT applies the GT predicate on the "logo" field.
func LogoGT(v string) predicate.Link {
	return predicate.Link(sql.FieldGT(FieldLogo, v))
}

// LogoGTE applies the GTE predicate on the "logo" field.
func LogoGTE(v string) predicate.Link {
	return predicate.Link(sql.FieldGTE(FieldLogo, v))
}

// LogoLT applies the LT predicate on the "logo" field.
func LogoLT(v string) predicate.Link {
	return predicate.Link(sql.FieldLT(FieldLogo, v))
}

// LogoLTE applies the LTE predicate on the "logo" field.
func LogoLTE(v string) predicate.Link {
	return predicate.Link(sql.FieldLTE(FieldLogo, v))
}

// LogoContains applies the Contains predicate on the "logo" field.
func LogoContains(v string) predicate.Link {
	return predicate.Link(sql.FieldContains(FieldLogo, v))
}

// LogoHasPrefix applies the HasPrefix predicate on the "logo" field.
func LogoHasPrefix(v string) predicate.Link {
	return predicate.Link(sql.FieldHasPrefix(FieldLogo, v))
}

// LogoHasSuffix applies the HasSuffix predicate on the "logo" field.
func LogoHasSuffix(v string) predicate.Link {
	return predicate.Link(sql.FieldHasSuffix(FieldLogo, v))
}

// LogoIsNil applies the IsNil predicate on the "logo" field.
func LogoIsNil() predicate.Link {
	return predicate.Link(sql.FieldIsNull(FieldLogo))
}

// LogoNotNil applies the NotNil predicate on the "logo" field.
func LogoNotNil() predicate.Link {
	return predicate.Link(sql.FieldNotNull(FieldLogo))
}

// LogoEqualFold applies the EqualFold predicate on the "logo" field.
func LogoEqualFold(v string) predicate.Link {
	return predicate.Link(sql.FieldEqualFold(FieldLogo, v))
}

// LogoContainsFold applies the ContainsFold predicate on the "logo" field.
func LogoContainsFold(v string) predicate.Link {
	return predicate.Link(sql.FieldContainsFold(FieldLogo, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Link {
	return predicate.Link(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Link {
	return predicate.Link(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Link {
	return predicate.Link(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Link {
	return predicate.Link(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Link {
	return predicate.Link(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Link {
	return predicate.Link(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Link {
	return predicate.Link(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Link {
	return predicate.Link(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Link {
	return predicate.Link(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Link {
	return predicate.Link(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Link {
	return predicate.Link(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Link {
	return predicate.Link(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Link {
	return predicate.Link(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Link {
	return predicate.Link(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Link {
	return predicate.Link(sql.FieldContainsFold(FieldDescription, v))
}

// TeamEQ applies the EQ predicate on the "team" field.
func TeamEQ(v string) predicate.Link {
	return predicate.Link(sql.FieldEQ(FieldTeam, v))
}

// TeamNEQ applies the NEQ predicate on the "team" field.
func TeamNEQ(v string) predicate.Link {
	return predicate.Link(sql.FieldNEQ(FieldTeam, v))
}

// TeamIn applies the In predicate on the "team" field.
func TeamIn(vs ...string) predicate.Link {
	return predicate.Link(sql.FieldIn(FieldTeam, vs...))
}

// TeamNotIn applies the NotIn predicate on the "team" field.
func TeamNotIn(vs ...string) predicate.Link {
	return predicate.Link(sql.FieldNotIn(FieldTeam, vs...))
}

// TeamGT applies the GT predicate on the "team" field.
func TeamGT(v string) predicate.Link {
	return predicate.Link(sql.FieldGT(FieldTeam, v))
}

// TeamGTE applies the GTE predicate on the "team" field.
func TeamGTE(v string) predicate.Link {
	return predicate.Link(sql.FieldGTE(FieldTeam, v))
}

// TeamLT applies the LT predicate on the "team" field.
func TeamLT(v string) predicate.Link {
	return predicate.Link(sql.FieldLT(FieldTeam, v))
}

// TeamLTE applies the LTE predicate on the "team" field.
func TeamLTE(v string) predicate.Link {
	return predicate.Link(sql.FieldLTE(FieldTeam, v))
}

// TeamContains applies the Contains predicate on the "team" field.
func TeamContains(v string) predicate.Link {
	return predicate.Link(sql.FieldContains(FieldTeam, v))
}

// TeamHasPrefix applies the HasPrefix predicate on the "team" field.
func TeamHasPrefix(v string) predicate.Link {
	return predicate.Link(sql.FieldHasPrefix(FieldTeam, v))
}

// TeamHasSuffix applies the HasSuffix predicate on the "team" field.
func TeamHasSuffix(v string) predicate.Link {
	return predicate.Link(sql.FieldHasSuffix(FieldTeam, v))
}

// TeamIsNil applies the IsNil predicate on the "team" field.
func TeamIsNil() predicate.Link {
	return predicate.Link(sql.FieldIsNull(FieldTeam))
}

// TeamNotNil applies the NotNil predicate on the "team" field.
func TeamNotNil() predicate.Link {
	return predicate.Link(sql.FieldNotNull(FieldTeam))
}

// TeamEqualFold applies the EqualFold predicate on the "team" field.
func TeamEqualFold(v string) predicate.Link {
	return predicate.Link(sql.FieldEqualFold(FieldTeam, v))
}

// TeamContainsFold applies the ContainsFold predicate on the "team" field.
func TeamContainsFold(v string) predicate.Link {
	return predicate.Link(sql.FieldContainsFold(FieldTeam, v))
}

// PriorityEQ applies the EQ predicate on the "priority" field.
func PriorityEQ(v int32) predicate.Link {
	return predicate.Link(sql.FieldEQ(FieldPriority, v))
}

// PriorityNEQ applies the NEQ predicate on the "priority" field.
func PriorityNEQ(v int32) predicate.Link {
	return predicate.Link(sql.FieldNEQ(FieldPriority, v))
}

// PriorityIn applies the In predicate on the "priority" field.
func PriorityIn(vs ...int32) predicate.Link {
	return predicate.Link(sql.FieldIn(FieldPriority, vs...))
}

// PriorityNotIn applies the NotIn predicate on the "priority" field.
func PriorityNotIn(vs ...int32) predicate.Link {
	return predicate.Link(sql.FieldNotIn(FieldPriority, vs...))
}

// PriorityGT applies the GT predicate on the "priority" field.
func PriorityGT(v int32) predicate.Link {
	return predicate.Link(sql.FieldGT(FieldPriority, v))
}

// PriorityGTE applies the GTE predicate on the "priority" field.
func PriorityGTE(v int32) predicate.Link {
	return predicate.Link(sql.FieldGTE(FieldPriority, v))
}

// PriorityLT applies the LT predicate on the "priority" field.
func PriorityLT(v int32) predicate.Link {
	return predicate.Link(sql.FieldLT(FieldPriority, v))
}

// PriorityLTE applies the LTE predicate on the "priority" field.
func PriorityLTE(v int32) predicate.Link {
	return predicate.Link(sql.FieldLTE(FieldPriority, v))
}

// PriorityIsNil applies the IsNil predicate on the "priority" field.
func PriorityIsNil() predicate.Link {
	return predicate.Link(sql.FieldIsNull(FieldPriority))
}

// PriorityNotNil applies the NotNil predicate on the "priority" field.
func PriorityNotNil() predicate.Link {
	return predicate.Link(sql.FieldNotNull(FieldPriority))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Link) predicate.Link {
	return predicate.Link(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Link) predicate.Link {
	return predicate.Link(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Link) predicate.Link {
	return predicate.Link(sql.NotPredicates(p))
}
