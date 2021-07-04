package mobile

import (
	"errors"
	"fmt"
	"github.com/gopub/sql"
	"reflect"
	"sync"
	"time"
)

type TableRecord interface {
	RecordKey() interface{}
}

type TableKeyType int

const (
	TableKeyInt64 = iota
	TableKeyString
)

type Table struct {
	name    string
	keyType TableKeyType
	db      *sql.DB
	mu      sync.RWMutex
	stmts   struct {
		insert            *sql.Stmt
		update            *sql.Stmt
		save              *sql.Stmt
		get               *sql.Stmt
		listAll           *sql.Stmt
		listGreaterThan   *sql.Stmt
		listLessThan      *sql.Stmt
		delete            *sql.Stmt
		deleteGreaterThan *sql.Stmt
		deleteLessThan    *sql.Stmt
	}
	Now Now
}

func NewTable(db *sql.DB, name string, keyType TableKeyType) (*Table, error) {
	var typ string
	switch keyType {
	case TableKeyInt64:
		typ = "BIGINT"
	case TableKeyString:
		typ = "VARCHAR(64)"
	default:
		return nil, errors.New("invalid key type")
	}

	_, err := db.Exec(fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s(
id %s PRIMARY KEY,
data BLOB,
updated_at BIGINT
)`, name, typ))
	if err != nil {
		return nil, err
	}

	t := &Table{
		name:    name,
		keyType: keyType,
		db:      db,
	}

	t.stmts.insert = sql.MustPrepare(db, `INSERT INTO %s(id,data,updated_at) VALUES(?,?,?)`, name)
	t.stmts.update = sql.MustPrepare(db, `UPDATE %s SET data=?,updated_at=? WHERE id=?`, name)
	t.stmts.save = sql.MustPrepare(db, `REPLACE INTO %s(id,data,updated_at) VALUES(?,?,?)`, name)
	t.stmts.get = sql.MustPrepare(db, `SELECT data FROM %s WHERE id=?`, name)
	t.stmts.listAll = sql.MustPrepare(db, `SELECT data FROM %s`, name)
	t.stmts.listGreaterThan = sql.MustPrepare(db, `SELECT data FROM %s WHERE id>? ORDER BY id ASC LIMIT ?`, name)
	t.stmts.listLessThan = sql.MustPrepare(db, `SELECT data FROM %s WHERE id<? ORDER BY id DESC LIMIT ?`, name)
	t.stmts.delete = sql.MustPrepare(db, `DELETE FROM %s WHERE id=?`, name)
	t.stmts.deleteGreaterThan = sql.MustPrepare(db, `DELETE FROM %s WHERE id>?`, name)
	t.stmts.deleteLessThan = sql.MustPrepare(db, `DELETE FROM %s WHERE id<?`, name)
	return t, nil
}

func (t *Table) now() int64 {
	if t.Now != nil {
		return t.Now.Now()
	}
	return time.Now().Unix()
}

func (t *Table) Insert(record TableRecord) error {
	t.mu.Lock()
	_, err := t.stmts.insert.Exec(record.RecordKey(), sql.JSON(record), t.now())
	t.mu.Unlock()
	return err
}

func (t *Table) Update(record TableRecord) error {
	t.mu.Lock()
	_, err := t.stmts.update.Exec(record.RecordKey(), sql.JSON(record), t.now())
	t.mu.Unlock()
	return err
}

func (t *Table) Save(record TableRecord) error {
	t.mu.Lock()
	_, err := t.stmts.save.Exec(record.RecordKey(), sql.JSON(record), t.now())
	t.mu.Unlock()
	return err
}

func (t *Table) Get(key interface{}, ptrToRecord interface{}) error {
	t.mu.RLock()
	err := t.stmts.get.QueryRow(key).Scan(sql.JSON(ptrToRecord))
	t.mu.RUnlock()
	return err
}

func (t *Table) ListAll(ptrToSlice interface{}) error {
	t.mu.RLock()
	defer t.mu.RUnlock()
	rows, err := t.stmts.listAll.Query()
	if err != nil {
		return err
	}
	defer rows.Close()
	return t.readList(rows, ptrToSlice)
}

func (t *Table) ListGreaterThan(key interface{}, ptrToSlice interface{}, limit int) error {
	t.mu.RLock()
	defer t.mu.RUnlock()
	rows, err := t.stmts.listGreaterThan.Query(key, limit)
	if err != nil {
		return err
	}
	defer rows.Close()
	return t.readList(rows, ptrToSlice)
}

func (t *Table) ListLessThan(key interface{}, ptrToSlice interface{}, limit int) error {
	t.mu.RLock()
	defer t.mu.RUnlock()
	rows, err := t.stmts.listLessThan.Query(key, limit)
	if err != nil {
		return err
	}
	defer rows.Close()
	err = t.readList(rows, ptrToSlice)
	if err != nil {
		return err
	}
	l := reflect.ValueOf(ptrToSlice).Elem()
	swapF := reflect.Swapper(l.Interface())
	for i, j := 0, l.Len()-1; i < j; i, j = i+1, j-1 {
		swapF(i, j)
	}
	return nil
}

func (t *Table) readList(rows *sql.Rows, ptrToSlice interface{}) error {
	l := reflect.ValueOf(ptrToSlice).Elem()
	if l.Kind() != reflect.Slice {
		return errors.New("not pointer to ptrToSlice")
	}
	e := l.Type().Elem()
	for rows.Next() {
		record := reflect.New(e)
		err := rows.Scan(sql.JSON(record.Interface()))
		if err != nil {
			return err
		}
		l = reflect.Append(l, record.Elem())
	}
	reflect.ValueOf(ptrToSlice).Elem().Set(l)
	return nil
}

func (t *Table) Delete(key interface{}) error {
	_, err := t.stmts.delete.Exec(key)
	return err
}

func (t *Table) DeleteGreaterThan(key interface{}) error {
	_, err := t.stmts.deleteGreaterThan.Exec(key)
	return err
}

func (t *Table) DeleteLessThan(key interface{}) error {
	_, err := t.stmts.deleteLessThan.Exec(key)
	return err
}
