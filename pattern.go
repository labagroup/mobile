package mobile

import (
	"github.com/gopub/conv"
	"github.com/gopub/types"
	"github.com/gopub/wine/urlutil"
)

func IsEmailAddress(s string) bool {
	return conv.IsEmailAddress(s)
}

func IsURL(s string) bool {
	return urlutil.IsURL(s)
}

func IsPhoneNumber(phoneNumber string) bool {
	return types.IsPhoneNumber(phoneNumber)
}

func IsDate(s string) bool {
	return conv.IsDate(s)
}
