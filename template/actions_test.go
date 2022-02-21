package template

import (
	"fmt"
	"testing"
)

func TestGetTitle(t *testing.T) {
	tests := []struct {
		have string
		want string
	}{
		{"GetStatus", "Get status"},
		{"GetTaskIDsForUser", "Get task ids for user"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Testing %s", tt.have), func(t *testing.T) {
			if got := GetTitle(tt.have); got != tt.want {
				t.Errorf("GetTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
