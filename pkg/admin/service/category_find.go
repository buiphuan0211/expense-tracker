package service

import (
	"context"
	"expense-tracker/internal/util/mgquery"
	"expense-tracker/internal/util/pconvert"
	"expense-tracker/pkg/admin/dao"
	"expense-tracker/pkg/admin/errorcode"
	responsemodel "expense-tracker/pkg/admin/model/response"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"sync"
)

// All ...
func (s categoryImpl) All(ctx context.Context, q mgquery.AppQuery) (res responsemodel.CategoryAll) {
	var (
		d  = dao.Category()
		wg = sync.WaitGroup{}
	)

	cond := bson.M{}
	q.AssignKeyword(cond)

	wg.Add(1)
	go func() {
		defer wg.Done()

		res.List = make([]responsemodel.CategoryBrief, 0)

		findOpts := q.GetFindOptionsWithPage()

		docs := d.FindByCondition(ctx, cond, findOpts)

		res.List = s.getListCategoryBrief(ctx, docs)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		res.Total = d.CountByCondition(ctx, cond)
	}()

	wg.Wait()

	res.Limit = q.Limit
	return
}

// Detail ...
func (s categoryImpl) Detail(ctx context.Context, id string) (res responsemodel.CategoryDetail, err error) {
	var (
		d    = dao.Category()
		cond = bson.M{"_id": pconvert.StringToObjectID(id)}
	)

	doc, err := d.FindOnByCondition(ctx, cond)
	if err != nil {
		err = errors.New(errorcode.CategoryNotFound)
		return
	}

	res = s.detail(ctx, doc)
	return
}
