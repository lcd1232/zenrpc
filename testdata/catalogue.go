package testdata

import (
	"github.com/lcd1232/zenrpc/v3"
)

type Group struct {
	Id       int      `json:"id"`
	Title    string   `json:"title"`
	Nodes    []Group  `json:"nodes"`
	Groups   []Group  `json:"group"`
	ChildOpt *Group   `json:"child"`
	Sub      SubGroup `json:"sub"`
}

type SubGroup struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	//Nodes []Group `json:"nodes"` TODO still causes infinite recursion
}

type Campaign struct {
	Id     int     `json:"id"`
	Groups []Group `json:"group"`
}

type CatalogueService struct{ zenrpc.Service }

func (s CatalogueService) First(groups []Group) (bool, error) {
	return true, nil
}

func (s CatalogueService) Second(campaigns []Campaign) (bool, error) {
	return true, nil
}

func (s CatalogueService) Third() (Campaign, error) {
	return Campaign{}, nil
}

//go:generate zenrpc
