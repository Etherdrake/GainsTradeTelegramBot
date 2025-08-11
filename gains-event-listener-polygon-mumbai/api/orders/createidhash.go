package orders

import (
	"github.com/segmentio/ksuid"
)

// CreateKsuid generates a ksuid based on https://github.com/segmentio/ksuid
func CreateKsuid() string {
	ksuid := ksuid.New()
	return ksuid.String()
}
