package functions

import "testing"

func TestGraphic(t *testing.T) {
	type args struct {
		splitArgs []string
		banner    []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Graphic(tt.args.splitArgs, tt.args.banner); got != tt.want {
				t.Errorf("Graphic() = %v, want %v", got, tt.want)
			}
		})
	}
}
