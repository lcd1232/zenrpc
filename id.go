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

type ID struct {
	rawID  *json.RawMessage
	Int    int64
	Float  float64
	String string
	State  IDState
}

func (i *ID) RawMessage() *json.RawMessage {
	return i.rawID
}

func newID(rawID *json.RawMessage) (ID, error) {
	if rawID == nil {
		return ID{
			rawID: rawID,
			State: IDStateNull,
		}, nil
	}

	if len(*rawID) > 2 && (*rawID)[0] == '"' && (*rawID)[len(*rawID)-1] == '"' {
		return ID{
			rawID:  rawID,
			State:  IDStateString,
			String: string((*rawID)[1 : len(*rawID)-1]),
		}, nil
	}

	var id json.Number
	if err := json.Unmarshal(*rawID, &id); err != nil {
		return ID{}, err
	}

	a, err := id.Int64()
	if err == nil {
		return ID{
			rawID: rawID,
			State: IDStateInt,
			Int:   a,
		}, nil
	}

	f, err := id.Float64()
	if err == nil {
		return ID{
			rawID: rawID,
			State: IDStateFloat,
			Float: f,
		}, nil
	}

	return ID{
		rawID: rawID,
	}, err
}
