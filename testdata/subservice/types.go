package subarithservice

import (
	"github.com/lcd1232/zenrpc/v3/testdata/objects"
	"time"
)

type Point struct {
	objects.AbstractObject
	A, B int        // coordinate
	C    int        `json:"-"`
	When *time.Time `json:"when"` // when it happened
}
