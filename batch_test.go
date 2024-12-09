package batch_test

import (
	"fmt"
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
			CaseName:      "zero rows count",
			RowsCount:     0,
			ArgPerRow:     2,
			Template:      "INSERT INTO users (id, name) VALUES %s",
			ExpectedQuery: "INSERT INTO users (id, name) VALUES ()",
		},
		&QueryTest{
			CaseName:      "zero args count",
			RowsCount:     2,
			ArgPerRow:     0,
			Template:      "INSERT INTO users (id, name) VALUES %s",
			ExpectedQuery: "INSERT INTO users (id, name) VALUES ()",
		},
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

type QueryMapTest struct {
	CaseName      string
	Template      string
	RowsCount     uint64
	ArgPerRow     uint64
	ExpectedQuery string
	ArgsMap       map[uint64]batch.MapFunc
}

func (q *QueryMapTest) Name() string {
	return q.CaseName
}

func (q *QueryMapTest) Test(t *testing.T) {
	actualQuery := batch.QueryMap(q.Template, q.RowsCount, q.ArgPerRow, q.ArgsMap)

	if q.ExpectedQuery != actualQuery {
		t.Fatalf("query not equal, expected %s, actual %s", q.ExpectedQuery, actualQuery)
	}
}

func Test_QueryMap(t *testing.T) {
	tester.RunNamedTesters(t,
		&QueryMapTest{
			CaseName:      "zero rows count",
			RowsCount:     0,
			ArgPerRow:     2,
			Template:      "INSERT INTO users (id, name) VALUES %s",
			ExpectedQuery: "INSERT INTO users (id, name) VALUES ()",
		},
		&QueryMapTest{
			CaseName:      "zero args count",
			RowsCount:     2,
			ArgPerRow:     0,
			Template:      "INSERT INTO users (id, name) VALUES %s",
			ExpectedQuery: "INSERT INTO users (id, name) VALUES ()",
		},
		&QueryMapTest{
			CaseName:      "single row",
			RowsCount:     1,
			ArgPerRow:     2,
			Template:      "INSERT INTO users (id, name) VALUES %s",
			ExpectedQuery: "INSERT INTO users (id, name) VALUES ($1,$2)",
		},
		&QueryMapTest{
			CaseName:      "many rows",
			RowsCount:     3,
			ArgPerRow:     2,
			Template:      "INSERT INTO users (id, name) VALUES %s",
			ExpectedQuery: "INSERT INTO users (id, name) VALUES ($1,$2),($3,$4),($5,$6)",
		},
		&QueryMapTest{
			CaseName:  "map",
			RowsCount: 3,
			ArgPerRow: 3,
			Template:  "INSERT INTO users (id, name, toy_id) VALUES %s",
			ArgsMap: map[uint64]batch.MapFunc{
				2: func(argNumber uint64) string {
					return fmt.Sprintf("(SELECT id FROM toys WHERE name = $%d)", argNumber)
				},
			},
			ExpectedQuery: "INSERT INTO users (id, name, toy_id) VALUES ($1,$2,(SELECT id FROM toys WHERE name = $3)),($4,$5,(SELECT id FROM toys WHERE name = $6)),($7,$8,(SELECT id FROM toys WHERE name = $9))",
		},
	)
}
