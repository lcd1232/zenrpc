package zenrpc

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getRawMessage(b []byte) *json.RawMessage {
	if len(b) == 0 {
		return nil
	}
	rm := json.RawMessage(b)
	return &rm
}

func Test_newID(t *testing.T) {
	type args struct {
		rawID *json.RawMessage
	}
	tests := []struct {
		name    string
		args    args
		want    id
		wantErr bool
	}{
		{
			name: "int",
			args: args{
				rawID: getRawMessage([]byte(`25`)),
			},
			want: id{
				int:   25,
				state: IDStateInt,
			},
		},
		{
			name: "string",
			args: args{
				rawID: getRawMessage([]byte(`"25"`)),
			},
			want: id{
				string: "25",
				state:  IDStateString,
			},
		},
		{
			name: "float",
			args: args{
				rawID: getRawMessage([]byte(`25.25`)),
			},
			want: id{
				float: 25.25,
				state: IDStateFloat,
			},
		},
		{
			name: "null",
			args: args{
				rawID: nil,
			},
			want: id{
				state: IDStateNull,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewID(tt.args.rawID)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want.state, got.State())
			assert.Equal(t, tt.want.int, got.Int())
			assert.Equal(t, tt.want.float, got.Float())
			assert.Equal(t, tt.want.string, got.String())
		})
	}
}

func TestID_RawMessage(t *testing.T) {
	type fields struct {
		rawID  *json.RawMessage
		Int    int64
		Float  float64
		String string
		State  IDState
	}
	tests := []struct {
		name   string
		fields fields
		want   *json.RawMessage
	}{
		{
			name: "string",
			fields: fields{
				rawID:  getRawMessage([]byte(`"john"`)),
				String: "john",
				State:  IDStateString,
			},
			want: getRawMessage([]byte(`"john"`)),
		},
		{
			name: "int",
			fields: fields{
				rawID: getRawMessage([]byte(`15`)),
				Int:   15,
				State: IDStateInt,
			},
			want: getRawMessage([]byte(`15`)),
		},
		{
			name: "float",
			fields: fields{
				rawID: getRawMessage([]byte(`26.9`)),
				Float: 26.9,
				State: IDStateFloat,
			},
			want: getRawMessage([]byte(`26.9`)),
		},
		{
			name: "null",
			fields: fields{
				State: IDStateNull,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &id{
				rawID:  tt.fields.rawID,
				int:    tt.fields.Int,
				float:  tt.fields.Float,
				string: tt.fields.String,
				state:  tt.fields.State,
			}
			got := i.RawMessage()
			assert.Equal(t, tt.want, got)
		})
	}
}
