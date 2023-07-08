package tenselectedpastquestions

import "testing"

func TestABC086_A(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "偶数",
			args: args{a: 3, b: 4},
			want: "Even",
		},
		{
			name: "奇数",
			args: args{a: 1, b: 21},
			want: "Odd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC086_A(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}