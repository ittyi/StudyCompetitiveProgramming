package tenselectedpastquestions

import "testing"

func TestABC083_B(t *testing.T) {
	// ABC083_B(n: 20, a: 2, b: 5)

	type args struct {
		n int
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "n: 20, a: 2, b: 5 で 84",
			args: args{n: 20, a: 2, b: 5},
			want: 84,
		},
		{
			name: "n: 10, a: 1, b: 21 で 13",
			args: args{n: 10, a: 1, b: 2},
			want: 13,
		},
		{
			name: "n: 100, a: 4, b: 16 で 4554",
			args: args{n: 100, a: 4, b: 16},
			want: 4554,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC083_B(tt.args.n, tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("ABC086_A() = %v, want %v", got, tt.want)
			}
		})
	}
}
