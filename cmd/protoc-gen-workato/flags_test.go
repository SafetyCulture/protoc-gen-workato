package main_test

import (
	"bytes"
	"fmt"
	"testing"

	. "github.com/SafetyCulture/protoc-gen-workato/cmd/protoc-gen-workato"
	"github.com/stretchr/testify/require"
)

func TestCode(t *testing.T) {
	f := ParseFlags(nil, []string{"app", "-help"})
	require.Zero(t, f.Code())

	f = ParseFlags(nil, []string{"app", "-whoawhoawhoa"})
	require.Equal(t, 1, f.Code())
}

func TestHasMatch(t *testing.T) {
	f := ParseFlags(nil, []string{"app", "-help"})
	require.True(t, f.HasMatch())

	f = ParseFlags(nil, []string{"app", "-version"})
	require.True(t, f.HasMatch())

	f = ParseFlags(nil, []string{"app", "-watthewhat"})
	require.True(t, f.HasMatch())

	f = ParseFlags(nil, []string{"app"})
	require.False(t, f.HasMatch())
}

func TestShowHelp(t *testing.T) {
	f := ParseFlags(nil, []string{"app", "-help"})
	require.True(t, f.ShowHelp())

	f = ParseFlags(nil, []string{"app", "-version"})
	require.False(t, f.ShowHelp())
}

func TestShowVersion(t *testing.T) {
	f := ParseFlags(nil, []string{"app", "-version"})
	require.True(t, f.ShowVersion())

	f = ParseFlags(nil, []string{"app", "-help"})
	require.False(t, f.ShowVersion())
}

func TestPrintHelp(t *testing.T) {
	buf := new(bytes.Buffer)

	f := ParseFlags(buf, []string{"app"})
	f.PrintHelp()

	result := buf.String()
	require.Contains(t, result, "Usage of app:\n\n")
	require.Contains(t, result, "FLAGS\n")
	require.Contains(t, result, "-help")
	require.Contains(t, result, "-version")
}

func TestPrintVersion(t *testing.T) {
	buf := new(bytes.Buffer)

	f := ParseFlags(buf, []string{"app"})
	f.PrintVersion()

	require.Equal(t, fmt.Sprintf("app version %s\n", Version()), buf.String())
}

func TestInvalidFlags(t *testing.T) {
	buf := new(bytes.Buffer)

	f := ParseFlags(buf, []string{"app", "-wat"})
	require.Contains(t, buf.String(), "flag provided but not defined: -wat\n")
	require.True(t, f.HasMatch())
	require.True(t, f.ShowHelp())
}
