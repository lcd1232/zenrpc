package zenrpc

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStorage(t *testing.T) {
	c := NewContext(nil, nil)
	c.Set("name", "John Doe")
	assert.Equal(t, "John Doe", c.Get("name"))
}

func TestRequest(t *testing.T) {
	r := &http.Request{
		Host: "example.com",
	}
	c := NewContext(r, nil)
	assert.Equal(t, r, c.Request())
}

func TestResponse(t *testing.T) {
	r := &httptest.ResponseRecorder{}
	c := NewContext(nil, r)
	assert.Equal(t, r, c.Response())
}

func TestRealIP(t *testing.T) {
	type fields struct {
		request *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "X-Forwarded-For multiple",
			fields: fields{
				request: &http.Request{
					Header: map[string][]string{
						"X-Forwarded-For": {"203.0.113.195, 70.41.3.18, 150.172.238.178"},
						"X-Real-Ip":       {"8.8.8.8"},
					},
					RemoteAddr: "9.9.9.9:41234",
				},
			},
			want: "203.0.113.195",
		},
		{
			name: "X-Forwarded-For one",
			fields: fields{
				request: &http.Request{
					Header: map[string][]string{
						"X-Forwarded-For": {"203.0.113.195"},
						"X-Real-Ip":       {"8.8.8.8"},
					},
					RemoteAddr: "9.9.9.9:41234",
				},
			},
			want: "203.0.113.195",
		},
		{
			name: "X-Real-IP one",
			fields: fields{
				request: &http.Request{
					Header: map[string][]string{
						"X-Real-Ip": {"8.8.8.8"},
					},
					RemoteAddr: "9.9.9.9:41234",
				},
			},
			want: "8.8.8.8",
		},
		{
			name: "No headers",
			fields: fields{
				request: &http.Request{
					RemoteAddr: "9.9.9.9:41234",
				},
			},
			want: "9.9.9.9",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewContext(tt.fields.request, nil)
			assert.Equal(t, tt.want, c.RealIP())
		})
	}
}

func TestCookie(t *testing.T) {
	cookie := &http.Cookie{
		Name:  "foo",
		Value: "bar",
	}
	r := &http.Request{
		Header: make(map[string][]string),
	}
	r.AddCookie(cookie)

	c := NewContext(r, nil)
	got, err := c.Cookie("foo")
	require.NoError(t, err)
	assert.Equal(t, cookie, got)
}

func TestSetCookie(t *testing.T) {
	cookie := &http.Cookie{
		Name:  "foo",
		Value: "bar",
	}
	resp := httptest.NewRecorder()

	c := NewContext(nil, resp)
	c.SetCookie(cookie)
	require.Len(t, c.Response().Header().Values("Set-Cookie"), 1)
}

func TestCookies(t *testing.T) {
	cookie := &http.Cookie{
		Name:  "foo",
		Value: "bar",
	}
	cookie2 := &http.Cookie{
		Name:  "foo2",
		Value: "bar2",
	}

	r := &http.Request{
		Header: make(map[string][]string),
	}
	r.AddCookie(cookie)
	r.AddCookie(cookie2)

	c := NewContext(r, nil)
	require.Len(t, c.Cookies(), 2)
}

func TestID(t *testing.T) {
	c := NewContext(nil, nil)
	assert.Zero(t, c.ID())
	c.SetID(&id{
		int:   10,
		state: IDStateInt,
	})
	assert.Equal(t, &id{
		int:   10,
		state: IDStateInt,
	}, c.ID())
}

func TestNamespace(t *testing.T) {
	c := NewContext(nil, nil)
	assert.Zero(t, c.Namespace())
	c.SetNamespace("auth")
	assert.Equal(t, "auth", c.Namespace())
}
