package mobile

import (
	"encoding/json"

	"github.com/gopub/conv"

	"github.com/gopub/types"
)

type Int64List struct {
	List []int64
}

func NewInt64List() *Int64List {
	l := &Int64List{}
	return l
}

func (l *Int64List) Len() int {
	if l == nil {
		return 0
	}
	return len(l.List)
}

func (l *Int64List) Get(index int) int64 {
	return l.List[index]
}

func (l *Int64List) Add(val int64) {
	l.List = append(l.List, val)
}

func (l *Int64List) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &l.List)
}

func (l *Int64List) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.List)
}

type IntList struct {
	List []int
}

func NewIntList() *IntList {
	l := &IntList{}
	return l
}

func (l *IntList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.List)
}

func (l *IntList) Get(index int) int {
	return l.List[index]
}

func (l *IntList) Add(val int) {
	l.List = append(l.List, val)
}

func (l *IntList) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &l.List)
}

func (l *IntList) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.List)
}

type Int64Set types.Int64Set

func (s *Int64Set) Add(i int64) {
	((*types.Int64Set)(s)).Add(i)
}

func (s *Int64Set) Contains(i int64) bool {
	return ((*types.Int64Set)(s)).Contains(i)
}

func (s *Int64Set) Remove(i int64) {
	((*types.Int64Set)(s)).Remove(i)
}

func (s *Int64Set) String() string {
	return conv.MustJSONString((*types.Int64Set)(s).Slice())
}
