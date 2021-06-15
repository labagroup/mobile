package mobile_test

import (
	"database/sql"
	"github.com/labagroup/mobile"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func createTable(t *testing.T, name string, keyType mobile.TableKeyType) *mobile.Table {
	db, err := sql.Open("sqlite3", "file::memory:")
	if err != nil {
		t.Error(err)
	}

	tbl, err := mobile.NewTable(db, name, mobile.TableKeyInt64)
	if err != nil {
		t.Error(err)
	}
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

func TestIntTable(t *testing.T) {
	tbl := createTable(t, "foo", mobile.TableKeyInt64)
	err := tbl.Insert(&testItem{ID: 123, Name: "haha"})
	if err != nil {
		t.Error(err)
	}
	var v *testItem
	err = tbl.Get(123, &v)
	if err != nil {
		t.Error(err)
	}
	t.Log(v.ID, v.Name)

	var l []*testItem
	err = tbl.ListAll(&l)
	if err != nil {
		t.Error(err)
	}
}
