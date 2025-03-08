package types

import "testing"

func TestIsUnsignedNumber(t *testing.T) {
	type args struct {
		val interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUnsignedNumber(tt.args.val); got != tt.want {
				t.Errorf("IsUnsignedNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
