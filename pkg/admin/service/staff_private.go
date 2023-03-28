package service

import (
	"context"
	mgmodel "expense-tracker/internal/model/mg"
	"expense-tracker/internal/util/ptime"
	responsemodel "expense-tracker/pkg/admin/model/response"
)

// getListStaffBrief ...
func (s staffImpl) getListStaffBrief(ctx context.Context, docs []mgmodel.Staff) (res []responsemodel.StaffBrief) {
	var total = len(docs)

	res = make([]responsemodel.StaffBrief, total)
	for i, doc := range docs {
		func(index int, staff mgmodel.Staff) {
			res[i] = s.brief(ctx, staff)
		}(i, doc)
	}

	return
}

// brief ...
func (staffImpl) brief(ctx context.Context, doc mgmodel.Staff) responsemodel.StaffBrief {
	return responsemodel.StaffBrief{
		ID:         doc.ID.Hex(),
		Name:       doc.Name,
		Phone:      doc.Phone,
		Email:      doc.Email,
		Permission: doc.Permissions,
		IsRoot:     doc.IsRoot,
		CreatedAt:  ptime.TimeResponseInit(doc.CreatedAt),
	}
}
