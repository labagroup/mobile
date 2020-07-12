package mobile

import (
	"github.com/gopub/types"
)

type PhoneNumber types.PhoneNumber

func (pn *PhoneNumber) String() string {
	return (*types.PhoneNumber)(pn).String()
}

func (pn *PhoneNumber) InternationalFormat() string {
	return (*types.PhoneNumber)(pn).InternationalFormat()
}

type PhoneNumberList struct {
	List []*types.PhoneNumber
}

func (l *PhoneNumberList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.List)
}

func (l *PhoneNumberList) Get(index int) *PhoneNumber {
	return (*PhoneNumber)(l.List[index])
}

func NewPhoneNumberList() *PhoneNumberList {
	return &PhoneNumberList{}
}

func (l *PhoneNumberList) Add(phoneNumber *PhoneNumber) {
	l.List = append(l.List, (*types.PhoneNumber)(phoneNumber))
}

func (l *PhoneNumberList) Contains(phoneNumber *PhoneNumber) bool {
	for _, pn := range l.List {
		if pn.String() == (*types.PhoneNumber)(phoneNumber).String() {
			return true
		}
	}
	return false
}

func (l *PhoneNumberList) ContainsString(phoneNumber string) bool {
	for _, pn := range l.List {
		if pn.String() == phoneNumber {
			return true
		}
	}
	return false
}

func NewPhoneNumber(callingCode int, number int64) *PhoneNumber {
	return &PhoneNumber{
		Code:   callingCode,
		Number: number,
	}
}

func ParsePhoneNumber(s string, code int) *PhoneNumber {
	pn, _ := types.NewPhoneNumberV2(s, code)
	return (*PhoneNumber)(pn)
}
