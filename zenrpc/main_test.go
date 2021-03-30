package main

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/lcd1232/zenrpc/v3/parser"
)

var testCases = []struct {
	name         string
	filename     string
	wantFilename string
	wantEmpty    bool
}{
	{
		name:         "1",
		filename:     "../testdata/content1/content.go",
		wantFilename: "../testdata/want1.txt",
	},
}

func TestGenerator(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			pi, err := parser.NewPackageInfo(tc.filename)
			require.NoError(t, err)

			require.NoError(t, pi.Parse(tc.filename))

			if tc.wantEmpty {
				require.Len(t, pi.Services, 0)
				return
			}
			require.True(t, len(pi.Services) > 0)

			got, err := generateSourceCode(pi)
			require.NoError(t, err)
			want, err := ioutil.ReadFile(tc.wantFilename)
			require.NoError(t, err)
			assert.Equal(t, string(want), string(got))
		})
	}
}
