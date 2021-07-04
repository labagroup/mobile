package mobile_test

import (
	"database/sql"
	"github.com/gopub/errors"
	"github.com/gopub/types"
	"github.com/labagroup/mobile"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
	"testing"
)

func createTable(t *testing.T, name string, keyType mobile.TableKeyType) *mobile.Table {
	db, err := sql.Open("sqlite3", "file::memory:")
	if err != nil {
		t.Error(err)
	}

	tbl, err := mobile.NewTable(db, name, mobile.TableKeyInt64)
	require.NoError(t, err)
	return tbl
}

type testItem struct {
	ID    int64
	Name  string
	Score float64
}

func (i *testItem) RecordKey() interface{} {
	return i.ID
}

func newTestItem() *testItem {
	return &testItem{
		ID:    types.RandomID().Int(),
		Name:  types.RandomID().Pretty(),
		Score: float64(types.RandomID()) / float64(3),
	}
}

func TestIntTable(t *testing.T) {
	tbl := createTable(t, "tbl"+types.RandomID().Pretty(), mobile.TableKeyInt64)
	var items []*testItem
	item := newTestItem()
	items = append(items, item)
	err := tbl.Insert(item)
	require.NoError(t, err)
	var v *testItem
	err = tbl.Get(item.ID, &v)
	require.NoError(t, err)
	require.Equal(t, item, v)

	item = newTestItem()
	item.ID = items[0].ID + 1
	err = tbl.Insert(item)
	require.NoError(t, err)
	items = append(items, item)

	var l []*testItem
	err = tbl.ListAll(&l)
	require.NoError(t, err)
	require.NotEmpty(t, l)
	require.Equal(t, items, l)

	l = l[0:0]
	err = tbl.ListGreaterThan(item.ID, &l, 10)
	require.NoError(t, err)
	require.Empty(t, l)

	l = l[0:0]
	err = tbl.ListLessThan(item.ID+1, &l, 10)
	require.NoError(t, err)
	require.Equal(t, 2, len(l))
	//t.Log(l[0].ID, l[1].ID)
	require.True(t, l[0].ID < l[1].ID)

	err = tbl.Delete(item.ID)
	require.NoError(t, err)

	err = tbl.Get(item.ID, &v)
	require.NotEmpty(t, err)
	require.Equal(t, true, errors.Is(err, sql.ErrNoRows))
}
