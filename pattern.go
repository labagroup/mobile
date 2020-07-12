package mobile

import (
	"net/mail"
	"net/url"

	"github.com/gopub/conv"
	"github.com/gopub/types"
)

func IsEmailAddress(s string) bool {
	_, err := mail.ParseAddress(s)
	return err == nil
}

func IsURL(s string) bool {
	u, err := url.Parse(s)
	if err != nil {
		return false
	}
	return u.Scheme != "" && u.Host != ""
}

func IsPhoneNumber(phoneNumber string) bool {
	return types.IsPhoneNumber(phoneNumber)
}

func IsDate(s string) bool {
	if len(s) > 10 {
		return false
	}
	_, err := conv.ToTime(s)
	return err == nil
}
