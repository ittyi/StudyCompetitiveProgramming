package tenselectedpastquestions

import "testing"

func TestABC087_B(t *testing.T) {
	type args struct {
		a int
		b int
		c int
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{2, 2, 2, 100},
			want: 2,
		},
		{
			name: "2",
			args: args{5, 1, 0, 150},
			want: 0,
		},
		{
			name: "3",
			args: args{30, 40, 50, 6000},
			want: 213,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC087_B(tt.args.a, tt.args.b, tt.args.c, tt.args.x); got != tt.want {
				t.Errorf("ABC087_B() = %v, want %v", got, tt.want)
			}
		})
	}
}