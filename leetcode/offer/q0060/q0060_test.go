package q0060

import (
	"testing"
)

func Test_dicesProbability(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "官方用例1",
			args: args{
				n: 1,
			},
			want: []float64{0.16667, 0.16667, 0.16667, 0.16667, 0.16667, 0.16667},
		},
		{
			name: "官方用例2",
			args: args{
				n: 2,
			},
			want: []float64{0.02778, 0.05556, 0.08333, 0.11111, 0.13889, 0.16667, 0.13889, 0.11111, 0.08333, 0.05556, 0.02778},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dicesProbability(tt.args.n); !equal(got, tt.want) {
				t.Errorf("dicesProbability() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equal(a, b []float64) bool {
	const tolerate float64 = 1e-4
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if c := a[i] - b[i]; c > tolerate || c < -tolerate {
			return false
		}
	}
	return true
}
