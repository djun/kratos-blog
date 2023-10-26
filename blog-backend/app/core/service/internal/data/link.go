package data

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	entgo "github.com/tx7do/go-utils/entgo/query"
	util "github.com/tx7do/go-utils/time"

	"kratos-cms/gen/api/go/common/pagination"
	"kratos-cms/gen/api/go/content/service/v1"

	"kratos-cms/app/core/service/internal/biz"
	"kratos-cms/app/core/service/internal/data/ent"
	"kratos-cms/app/core/service/internal/data/ent/link"
)

var _ biz.LinkRepo = (*LinkRepo)(nil)

type LinkRepo struct {
	data *Data
	log  *log.Helper
}

func NewLinkRepo(data *Data, logger log.Logger) biz.LinkRepo {
	l := log.NewHelper(log.With(logger, "module", "link/repo/core-service"))
	return &LinkRepo{
		data: data,
		log:  l,
	}
}

func (r *LinkRepo) convertEntToProto(in *ent.Link) *v1.Link {
	if in == nil {
		return nil
	}
	return &v1.Link{
		Id:          in.ID,
		Name:        in.Name,
		Url:         in.URL,
		Logo:        in.Logo,
		Description: in.Description,
		Team:        in.Team,
		Priority:    in.Priority,
		CreateTime:  util.UnixMilliToStringPtr(in.CreateTime),
		UpdateTime:  util.UnixMilliToStringPtr(in.UpdateTime),
	}
}

func (r *LinkRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().Link.Query()
	if len(whereCond) != 0 {
		builder.Modify(whereCond...)
	}
	return builder.Count(ctx)
}

func (r *LinkRepo) List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListLinkResponse, error) {
	builder := r.data.db.Client().Link.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(r.data.db.Driver().Dialect(),
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), link.FieldCreateTime)
	if err != nil {
		r.log.Errorf("解析条件发生错误[%s]", err.Error())
		return nil, err
	}

	if querySelectors != nil {
		builder.Modify(querySelectors...)
	}

	results, err := builder.All(ctx)
	if err != nil {
		return nil, err
	}

	items := make([]*v1.Link, 0, len(results))
	for _, res := range results {
		item := r.convertEntToProto(res)
		items = append(items, item)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &v1.ListLinkResponse{
		Total: int32(count),
		Items: items,
	}, nil
}

func (r *LinkRepo) Get(ctx context.Context, req *v1.GetLinkRequest) (*v1.Link, error) {
	res, err := r.data.db.Client().Link.Get(ctx, req.GetId())
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	return r.convertEntToProto(res), err
}

func (r *LinkRepo) Create(ctx context.Context, req *v1.CreateLinkRequest) (*v1.Link, error) {
	res, err := r.data.db.Client().Link.Create().
		SetNillableName(req.Link.Name).
		SetNillableURL(req.Link.Url).
		SetNillableLogo(req.Link.Logo).
		SetNillableDescription(req.Link.Description).
		SetNillableTeam(req.Link.Team).
		SetNillablePriority(req.Link.Priority).
		SetCreateTime(time.Now().UnixMilli()).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.convertEntToProto(res), err
}

func (r *LinkRepo) Update(ctx context.Context, req *v1.UpdateLinkRequest) (*v1.Link, error) {
	builder := r.data.db.Client().Link.UpdateOneID(req.Id).
		SetNillableName(req.Link.Name).
		SetNillableURL(req.Link.Url).
		SetNillableLogo(req.Link.Logo).
		SetNillableDescription(req.Link.Description).
		SetNillableTeam(req.Link.Team).
		SetNillablePriority(req.Link.Priority).
		SetUpdateTime(time.Now().UnixMilli())

	res, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.convertEntToProto(res), err
}

func (r *LinkRepo) Delete(ctx context.Context, req *v1.DeleteLinkRequest) (bool, error) {
	err := r.data.db.Client().Link.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	return err != nil, err
}
