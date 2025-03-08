package to

import (
	"github.com/viebiz/clgt/types"
	"reflect"
	"testing"
)

func TestCast(t *testing.T) {
	type args struct {
		val interface{}
	}
	type testCase[TType types.Struct] struct {
		name string
		args args
		want TType
	}
	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Cast(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cast() = %v, want %v", got, tt.want)
			}
		})
	}
}
