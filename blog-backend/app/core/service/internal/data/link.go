package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	"kratos-cms/gen/api/go/common/pagination"
	"kratos-cms/gen/api/go/content/service/v1"

	"kratos-cms/app/core/service/internal/biz"
	"kratos-cms/app/core/service/internal/data/ent"
	"kratos-cms/app/core/service/internal/data/ent/link"

	"github.com/tx7do/kratos-utils/entgo"
	paging "github.com/tx7do/kratos-utils/pagination"
	util "github.com/tx7do/kratos-utils/time"
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

func (r *LinkRepo) Count(ctx context.Context, whereCond entgo.WhereConditions) (int, error) {
	builder := r.data.db.Link.Query()
	if len(whereCond) != 0 {
		for _, cond := range whereCond {
			builder = builder.Where(cond)
		}
	}
	return builder.Count(ctx)
}

func (r *LinkRepo) List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListLinkResponse, error) {
	whereCond, orderCond := entgo.QueryCommandToSelector(req.GetQuery(), req.GetOrderBy())

	builder := r.data.db.Link.Query()
	if len(whereCond) != 0 {
		for _, v := range whereCond {
			builder.Where(v)
		}
	}
	if len(orderCond) != 0 {
		for _, v := range orderCond {
			builder.Order(v)
		}
	} else {
		builder.Order(ent.Desc(link.FieldCreateTime))
	}
	if req.GetPage() > 0 && req.GetPageSize() > 0 && !req.GetNopaging() {
		builder.
			Offset(paging.GetPageOffset(req.GetPage(), req.GetPageSize())).
			Limit(int(req.GetPageSize()))
	}
	links, err := builder.All(ctx)
	if err != nil {
		return nil, err
	}

	items := make([]*v1.Link, 0, len(links))
	for _, po := range links {
		item := r.convertEntToProto(po)
		items = append(items, item)
	}

	count, err := r.Count(ctx, whereCond)
	if err != nil {
		return nil, err
	}

	ret := v1.ListLinkResponse{
		Total: int32(count),
		Items: items,
	}

	return &ret, err
}

func (r *LinkRepo) Get(ctx context.Context, req *v1.GetLinkRequest) (*v1.Link, error) {
	po, err := r.data.db.Link.Get(ctx, req.GetId())
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	return r.convertEntToProto(po), err
}

func (r *LinkRepo) Create(ctx context.Context, req *v1.CreateLinkRequest) (*v1.Link, error) {
	po, err := r.data.db.Link.Create().
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

	return r.convertEntToProto(po), err
}

func (r *LinkRepo) Update(ctx context.Context, req *v1.UpdateLinkRequest) (*v1.Link, error) {
	builder := r.data.db.Link.UpdateOneID(req.Id).
		SetNillableName(req.Link.Name).
		SetNillableURL(req.Link.Url).
		SetNillableLogo(req.Link.Logo).
		SetNillableDescription(req.Link.Description).
		SetNillableTeam(req.Link.Team).
		SetNillablePriority(req.Link.Priority).
		SetUpdateTime(time.Now().UnixMilli())

	po, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.convertEntToProto(po), err
}

func (r *LinkRepo) Delete(ctx context.Context, req *v1.DeleteLinkRequest) (bool, error) {
	err := r.data.db.Link.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	return err != nil, err
}
