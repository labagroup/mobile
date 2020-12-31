package mobile

import (
	"encoding/json"
	"regexp"
	"strings"

	"github.com/gopub/types"
)

type StringList struct {
	List []string
}

func NewStringList() *StringList {
	l := &StringList{}
	return l
}

func (l *StringList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.List)
}

func (l *StringList) Get(index int) string {
	return l.List[index]
}

func (l *StringList) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &l.List)
}

func (l *StringList) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.List)
}

func SmartLen(s string) int {
	n := 0
	for _, c := range s {
		if c <= 255 {
			n++
		} else {
			n += 2
		}
	}

	return n
}

var whitespaceRegexp = regexp.MustCompile(`\\s`)

// CombineSpaces combines multiple spaces to a single space
func CombineSpaces(s string) string {
	s = strings.TrimSpace(s)
	s = whitespaceRegexp.ReplaceAllString(s, " ")
	return s
}

type StringSet types.StringSet

func (s *StringSet) Add(str string) {
	((*types.StringSet)(s)).Add(str)
}

func (s *StringSet) Contains(str string) bool {
	return ((*types.StringSet)(s)).Contains(str)
}

func (s *StringSet) Remove(str string) {
	((*types.StringSet)(s)).Remove(str)
}

type StringMap struct {
	m map[string]string
}

func NewStringMap() *StringMap {
	return &StringMap{
		m: make(map[string]string),
	}
}

func (m *StringMap) Len() int {
	return len(m.m)
}

func (m *StringMap) Get(key string) string {
	return m.m[key]
}

func (m *StringMap) Set(key, val string) {
	m.m[key] = val
}

func (m *StringMap) Keys() *StringList {
	l := new(StringList)
	for k := range m.m {
		l.List = append(l.List, k)
	}
	return l
}
