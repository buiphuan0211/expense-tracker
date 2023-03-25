package service

import (
	"context"
	mgmodel "expense-tracker/internal/model/mg"
	"expense-tracker/internal/util/pconvert"
	"expense-tracker/internal/util/pstring"
	"expense-tracker/internal/util/ptime"
	"expense-tracker/pkg/admin/dao"
	responsemodel "expense-tracker/pkg/admin/model/response"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
	"sync"
)

// IsExistedByName ...
func (categoryImpl) IsExistedByName(ctx context.Context, name string) bool {
	var (
		d    = dao.Category()
		cond = bson.M{"slug": pstring.ToSlug(strings.ToLower(name))}
	)
	doc, _ := d.FindOnByCondition(ctx, cond)
	if !doc.ID.IsZero() {
		return true
	}
	return false
}

// IsExistedByID ...
func (categoryImpl) IsExistedByID(ctx context.Context, id string) bool {
	var (
		d    = dao.Category()
		cond = bson.M{"_id": pconvert.StringToObjectID(id)}
	)

	doc, _ := d.FindOnByCondition(ctx, cond)
	if !doc.ID.IsZero() {
		return true
	}
	return false
}

// getListCategoryBrief ...
func (s categoryImpl) getListCategoryBrief(ctx context.Context, docs []mgmodel.Category) (res []responsemodel.CategoryBrief) {
	var (
		wg    = sync.WaitGroup{}
		total = len(docs)
	)
	res = make([]responsemodel.CategoryBrief, total)

	wg.Add(total)
	for i, doc := range docs {
		go func(index int, category mgmodel.Category) {
			defer wg.Done()
			res[index] = s.brief(ctx, category)
		}(i, doc)
	}

	wg.Wait()
	return
}

// brief ...
func (categoryImpl) brief(ctx context.Context, category mgmodel.Category) responsemodel.CategoryBrief {
	return responsemodel.CategoryBrief{
		ID:        category.ID.Hex(),
		Name:      category.Name,
		CreatedAt: ptime.TimeResponseInit(category.CreatedAt),
	}
}

// detail ...
func (categoryImpl) detail(ctx context.Context, category mgmodel.Category) responsemodel.CategoryDetail {
	return responsemodel.CategoryDetail{
		ID:        category.ID.Hex(),
		Name:      category.Name,
		CreatedAt: ptime.TimeResponseInit(category.CreatedAt),
	}
}
