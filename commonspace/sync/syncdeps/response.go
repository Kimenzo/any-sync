package syncdeps

import (
	"github.com/Kimenzo/any-sync/util/multiqueue"
)

type Response interface {
	multiqueue.Sizeable
}
