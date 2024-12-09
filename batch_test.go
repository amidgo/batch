package batch_test

import (
	"testing"

	"github.com/amidgo/batch"
	"github.com/amidgo/tester"
)

type QueryTest struct {
	CaseName      string
	Template      string
	RowsCount     uint64
	ArgPerRow     uint64
	ExpectedQuery string
}

func (q *QueryTest) Name() string {
	return q.CaseName
}

func (q *QueryTest) Test(t *testing.T) {
	actualQuery := batch.Query(q.Template, q.RowsCount, q.ArgPerRow)

	if q.ExpectedQuery != actualQuery {
		t.Fatalf("query not equal, expected %s, actual %s", q.ExpectedQuery, actualQuery)
	}
}

func Test_Query(t *testing.T) {
	tester.RunNamedTesters(t,
		&QueryTest{
			CaseName:      "single row",
			RowsCount:     1,
			ArgPerRow:     2,
			Template:      "INSERT INTO users (id, name) VALUES %s",
			ExpectedQuery: "INSERT INTO users (id, name) VALUES ($1,$2)",
		},
		&QueryTest{
			CaseName:      "many rows",
			RowsCount:     3,
			ArgPerRow:     2,
			Template:      "INSERT INTO users (id, name) VALUES %s",
			ExpectedQuery: "INSERT INTO users (id, name) VALUES ($1,$2),($3,$4),($5,$6)",
		},
	)
}
