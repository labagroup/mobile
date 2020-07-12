package mobile

import (
	"time"

	"github.com/gopub/log"
	"github.com/gopub/types"
)

type IntE struct {
	Val int
	Err *Error
}

type BoolE struct {
	Val bool
	Err *Error
}

type StringE struct {
	Val string
	Err *Error
}

type Int64E struct {
	Val int64
	Err *Error
}

type ImageE struct {
	Val *types.Image
	Err *Error
}

type Float64E struct {
	Val float64
	Err *Error
}

type SecretManager interface {
	Get(key string) string
	Set(key, data string) bool
	Del(key string) bool
}

type Uptimer interface {
	Uptime() int64
}

const (
	KeyDeviceID = "device_id"
)

func GetDeviceID(m SecretManager) string {
	id := m.Get(KeyDeviceID)
	if id == "" {
		id = types.NewUUID()
		m.Set(KeyDeviceID, id)
	}
	return id
}

func SetTimeZone(name string, offset int) {
	// LoadLocation get failed on iOS
	loc, err := time.LoadLocation(name)
	if err != nil {
		log.Warnf("Cannot load location %s: %v. Set fixed zone", name, err)
		loc = time.FixedZone(name, offset)
	}
	time.Local = loc
}

func NewUUID() string {
	return types.NewUUID()
}
