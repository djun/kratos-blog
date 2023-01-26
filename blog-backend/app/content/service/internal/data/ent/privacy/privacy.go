// Code generated by ent, DO NOT EDIT.

package privacy

import (
	"context"
	"fmt"
	"kratos-blog/app/content/service/internal/data/ent"

	"entgo.io/ent/entql"
	"entgo.io/ent/privacy"
)

var (
	// Allow may be returned by rules to indicate that the policy
	// evaluation should terminate with allow decision.
	Allow = privacy.Allow

	// Deny may be returned by rules to indicate that the policy
	// evaluation should terminate with deny decision.
	Deny = privacy.Deny

	// Skip may be returned by rules to indicate that the policy
	// evaluation should continue to the next rule.
	Skip = privacy.Skip
)

// Allowf returns an formatted wrapped Allow decision.
func Allowf(format string, a ...any) error {
	return fmt.Errorf(format+": %w", append(a, Allow)...)
}

// Denyf returns an formatted wrapped Deny decision.
func Denyf(format string, a ...any) error {
	return fmt.Errorf(format+": %w", append(a, Deny)...)
}

// Skipf returns an formatted wrapped Skip decision.
func Skipf(format string, a ...any) error {
	return fmt.Errorf(format+": %w", append(a, Skip)...)
}

// DecisionContext creates a new context from the given parent context with
// a policy decision attach to it.
func DecisionContext(parent context.Context, decision error) context.Context {
	return privacy.DecisionContext(parent, decision)
}

// DecisionFromContext retrieves the policy decision from the context.
func DecisionFromContext(ctx context.Context) (error, bool) {
	return privacy.DecisionFromContext(ctx)
}

type (
	// Policy groups query and mutation policies.
	Policy = privacy.Policy

	// QueryRule defines the interface deciding whether a
	// query is allowed and optionally modify it.
	QueryRule = privacy.QueryRule
	// QueryPolicy combines multiple query rules into a single policy.
	QueryPolicy = privacy.QueryPolicy

	// MutationRule defines the interface which decides whether a
	// mutation is allowed and optionally modifies it.
	MutationRule = privacy.MutationRule
	// MutationPolicy combines multiple mutation rules into a single policy.
	MutationPolicy = privacy.MutationPolicy
)

// QueryRuleFunc type is an adapter to allow the use of
// ordinary functions as query rules.
type QueryRuleFunc func(context.Context, ent.Query) error

// Eval returns f(ctx, q).
func (f QueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	return f(ctx, q)
}

// MutationRuleFunc type is an adapter which allows the use of
// ordinary functions as mutation rules.
type MutationRuleFunc func(context.Context, ent.Mutation) error

// EvalMutation returns f(ctx, m).
func (f MutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return f(ctx, m)
}

// QueryMutationRule is an interface which groups query and mutation rules.
type QueryMutationRule interface {
	QueryRule
	MutationRule
}

// AlwaysAllowRule returns a rule that returns an allow decision.
func AlwaysAllowRule() QueryMutationRule {
	return fixedDecision{Allow}
}

// AlwaysDenyRule returns a rule that returns a deny decision.
func AlwaysDenyRule() QueryMutationRule {
	return fixedDecision{Deny}
}

type fixedDecision struct {
	decision error
}

func (f fixedDecision) EvalQuery(context.Context, ent.Query) error {
	return f.decision
}

func (f fixedDecision) EvalMutation(context.Context, ent.Mutation) error {
	return f.decision
}

type contextDecision struct {
	eval func(context.Context) error
}

// ContextQueryMutationRule creates a query/mutation rule from a context eval func.
func ContextQueryMutationRule(eval func(context.Context) error) QueryMutationRule {
	return contextDecision{eval}
}

func (c contextDecision) EvalQuery(ctx context.Context, _ ent.Query) error {
	return c.eval(ctx)
}

func (c contextDecision) EvalMutation(ctx context.Context, _ ent.Mutation) error {
	return c.eval(ctx)
}

// OnMutationOperation evaluates the given rule only on a given mutation operation.
func OnMutationOperation(rule MutationRule, op ent.Op) MutationRule {
	return MutationRuleFunc(func(ctx context.Context, m ent.Mutation) error {
		if m.Op().Is(op) {
			return rule.EvalMutation(ctx, m)
		}
		return Skip
	})
}

// DenyMutationOperationRule returns a rule denying specified mutation operation.
func DenyMutationOperationRule(op ent.Op) MutationRule {
	rule := MutationRuleFunc(func(_ context.Context, m ent.Mutation) error {
		return Denyf("ent/privacy: operation %s is not allowed", m.Op())
	})
	return OnMutationOperation(rule, op)
}

// The CategoryQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type CategoryQueryRuleFunc func(context.Context, *ent.CategoryQuery) error

// EvalQuery return f(ctx, q).
func (f CategoryQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.CategoryQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.CategoryQuery", q)
}

// The CategoryMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type CategoryMutationRuleFunc func(context.Context, *ent.CategoryMutation) error

// EvalMutation calls f(ctx, m).
func (f CategoryMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.CategoryMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.CategoryMutation", m)
}

// The LinkQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type LinkQueryRuleFunc func(context.Context, *ent.LinkQuery) error

// EvalQuery return f(ctx, q).
func (f LinkQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.LinkQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.LinkQuery", q)
}

// The LinkMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type LinkMutationRuleFunc func(context.Context, *ent.LinkMutation) error

// EvalMutation calls f(ctx, m).
func (f LinkMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.LinkMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.LinkMutation", m)
}

// The MenuQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type MenuQueryRuleFunc func(context.Context, *ent.MenuQuery) error

// EvalQuery return f(ctx, q).
func (f MenuQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.MenuQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.MenuQuery", q)
}

// The MenuMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type MenuMutationRuleFunc func(context.Context, *ent.MenuMutation) error

// EvalMutation calls f(ctx, m).
func (f MenuMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.MenuMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.MenuMutation", m)
}

// The PhotoQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type PhotoQueryRuleFunc func(context.Context, *ent.PhotoQuery) error

// EvalQuery return f(ctx, q).
func (f PhotoQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PhotoQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.PhotoQuery", q)
}

// The PhotoMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type PhotoMutationRuleFunc func(context.Context, *ent.PhotoMutation) error

// EvalMutation calls f(ctx, m).
func (f PhotoMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.PhotoMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.PhotoMutation", m)
}

// The PostQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type PostQueryRuleFunc func(context.Context, *ent.PostQuery) error

// EvalQuery return f(ctx, q).
func (f PostQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PostQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.PostQuery", q)
}

// The PostMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type PostMutationRuleFunc func(context.Context, *ent.PostMutation) error

// EvalMutation calls f(ctx, m).
func (f PostMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.PostMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.PostMutation", m)
}

// The TagQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type TagQueryRuleFunc func(context.Context, *ent.TagQuery) error

// EvalQuery return f(ctx, q).
func (f TagQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.TagQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.TagQuery", q)
}

// The TagMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type TagMutationRuleFunc func(context.Context, *ent.TagMutation) error

// EvalMutation calls f(ctx, m).
func (f TagMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.TagMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.TagMutation", m)
}

type (
	// Filter is the interface that wraps the Where function
	// for filtering nodes in queries and mutations.
	Filter interface {
		// Where applies a filter on the executed query/mutation.
		Where(entql.P)
	}

	// The FilterFunc type is an adapter that allows the use of ordinary
	// functions as filters for query and mutation types.
	FilterFunc func(context.Context, Filter) error
)

// EvalQuery calls f(ctx, q) if the query implements the Filter interface, otherwise it is denied.
func (f FilterFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	fr, err := queryFilter(q)
	if err != nil {
		return err
	}
	return f(ctx, fr)
}

// EvalMutation calls f(ctx, q) if the mutation implements the Filter interface, otherwise it is denied.
func (f FilterFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	fr, err := mutationFilter(m)
	if err != nil {
		return err
	}
	return f(ctx, fr)
}

var _ QueryMutationRule = FilterFunc(nil)

func queryFilter(q ent.Query) (Filter, error) {
	switch q := q.(type) {
	case *ent.CategoryQuery:
		return q.Filter(), nil
	case *ent.LinkQuery:
		return q.Filter(), nil
	case *ent.MenuQuery:
		return q.Filter(), nil
	case *ent.PhotoQuery:
		return q.Filter(), nil
	case *ent.PostQuery:
		return q.Filter(), nil
	case *ent.TagQuery:
		return q.Filter(), nil
	default:
		return nil, Denyf("ent/privacy: unexpected query type %T for query filter", q)
	}
}

func mutationFilter(m ent.Mutation) (Filter, error) {
	switch m := m.(type) {
	case *ent.CategoryMutation:
		return m.Filter(), nil
	case *ent.LinkMutation:
		return m.Filter(), nil
	case *ent.MenuMutation:
		return m.Filter(), nil
	case *ent.PhotoMutation:
		return m.Filter(), nil
	case *ent.PostMutation:
		return m.Filter(), nil
	case *ent.TagMutation:
		return m.Filter(), nil
	default:
		return nil, Denyf("ent/privacy: unexpected mutation type %T for mutation filter", m)
	}
}
