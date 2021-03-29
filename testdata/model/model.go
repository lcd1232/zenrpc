package model

import "github.com/lcd1232/zenrpc/v3/testdata/objects"

type Point struct {
	objects.AbstractObject     // embeded object
	X, Y                   int // coordinate
	Z                      int `json:"-"`
	ConnectedObject        objects.AbstractObject
}
