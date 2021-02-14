package q0093

import (
	"testing"
)

func Test_restoreIpAddresses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "官方用例1",
			args: args{"25525511135"},
			want: []string{"255.255.11.135", "255.255.111.35"},
		},
		{
			name: "官方用例2",
			args: args{"0000"},
			want: []string{"0.0.0.0"},
		},
		{
			name: "官方用例3",
			args: args{"1111"},
			want: []string{"1.1.1.1"},
		},
		{
			name: "官方用例4",
			args: args{"010010"},
			want: []string{"0.10.0.10", "0.100.1.0"},
		},
		{
			name: "官方用例5",
			args: args{"101023"},
			want: []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := restoreIpAddresses(tt.args.s); !isEqual(got, tt.want) {
				t.Errorf("restoreIpAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}

func isEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	am, bm := make(map[string]int), make(map[string]int)
	for _, v := range a {
		am[v] += 1
	}
	for _, v := range b {
		bm[v] += 1
	}
	for k, v := range am {
		if bm[k] != v {
			return false
		}
	}
	return true
}
