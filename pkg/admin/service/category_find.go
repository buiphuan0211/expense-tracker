package service

import (
	"context"
	"expense-tracker/internal/util/mgquery"
	"expense-tracker/pkg/admin/dao"
	responsemodel "expense-tracker/pkg/admin/model/response"
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
