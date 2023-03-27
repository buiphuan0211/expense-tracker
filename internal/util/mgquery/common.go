package mgquery

import (
	"expense-tracker/internal/util/pstring"
	"go.mongodb.org/mongo-driver/bson"
)

// AssignKeyword ...
func (q *AppQuery) AssignKeyword(cond bson.M) {
	if q.Keyword != "" {
		cond["searchString"] = pstring.GenerateQuerySearchString(q.Keyword)
	}
}

func (q *AppQuery) AssignStatus(cond bson.M) {
	if q.Status != "" {
		cond["status"] = q.Status
	}
}
