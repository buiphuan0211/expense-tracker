package service

import (
	"context"
	mgmodel "expense-tracker/internal/model/mg"
	"expense-tracker/internal/util/mgquery"
	"expense-tracker/pkg/admin/dao"
	responsemodel "expense-tracker/pkg/admin/model/response"
	"go.mongodb.org/mongo-driver/bson"
	"sync"
)

// All ...
func (s staffImpl) All(ctx context.Context, q mgquery.AppQuery) (res responsemodel.StaffAll) {
	var cond = bson.M{}
	q.AssignKeyword(cond)
	q.AssignStatus(cond)

	var (
		wg = sync.WaitGroup{}
		d  = dao.Staff()
	)

	wg.Add(1)
	func() {
		defer wg.Done()

		res.List = make([]responsemodel.StaffBrief, 0)

		findOpts := q.GetFindOptionsWithPage()

		staffs := d.FindByCondition(ctx, cond, findOpts)

		res.List = s.getListStaffBrief(ctx, staffs)
	}()

	wg.Add(1)
	func() {
		defer wg.Done()
		res.Total = d.CountByCondition(ctx, cond)
	}()

	wg.Wait()

	res.Limit = q.Limit

	return
}

// ExistedByEmail ...
func (staffImpl) ExistedByEmail(ctx context.Context, email string) (res bool) {
	var (
		d    = dao.Staff()
		cond = bson.M{"email": email}
	)

	_, err := d.FindOneByCondition(ctx, cond)
	if err != nil {
		return
	}

	return true
}

// FindRawByEmail ...
func (staffImpl) FindRawByEmail(ctx context.Context, email string) (res mgmodel.Staff, err error) {
	return dao.Staff().FindOneByCondition(ctx, bson.M{"email": email})
}
