package zenrpc

import (
	"encoding/json"
)

type IDState int

const (
	IDStateInt IDState = iota
	IDStateFloat
	IDStateString
	IDStateNull
)

type ID interface {
	RawMessage() *json.RawMessage
	Int() int64
	Float() float64
	String() string
	State() IDState
}

type id struct {
	rawID  *json.RawMessage
	int    int64
	float  float64
	string string
	state  IDState
}

func (i *id) RawMessage() *json.RawMessage {
	return i.rawID
}

func (i *id) Int() int64 {
	return i.int
}

func (i *id) Float() float64 {
	return i.float
}

func (i *id) String() string {
	return i.string
}

func (i *id) State() IDState {
	return i.state
}

func newID(rawID *json.RawMessage) (ID, error) {
	if rawID == nil {
		return &id{
			rawID: rawID,
			state: IDStateNull,
		}, nil
	}

	if len(*rawID) > 2 && (*rawID)[0] == '"' && (*rawID)[len(*rawID)-1] == '"' {
		return &id{
			rawID:  rawID,
			state:  IDStateString,
			string: string((*rawID)[1 : len(*rawID)-1]),
		}, nil
	}

	var number json.Number
	if err := json.Unmarshal(*rawID, &number); err != nil {
		return &id{}, err
	}

	a, err := number.Int64()
	if err == nil {
		return &id{
			rawID: rawID,
			state: IDStateInt,
			int:   a,
		}, nil
	}

	f, err := number.Float64()
	if err == nil {
		return &id{
			rawID: rawID,
			state: IDStateFloat,
			float: f,
		}, nil
	}

	return &id{
		rawID: rawID,
	}, err
}
