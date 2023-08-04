package template

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTitle(t *testing.T) {
	tests := map[string]struct {
		have string
		want string
	}{
		"should replace space #1": {"GetStatus", "Get status"},
		"should replace space #2": {"GetTaskIDsForUser", "Get task ids for user"},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			res := GetTitle(tt.have)
			assert.Equal(t, tt.want, res)
		})
	}
}
