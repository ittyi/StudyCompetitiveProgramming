package tenselectedpastquestions

import "testing"

func TestABC088_B(t *testing.T) {
	type args struct {
		n int
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[]int{3, 1}で 2",
			args: args{n: 2, a: []int{3, 1}},
			want: 2,
		},
		{
			name: "[]int{2, 7, 4}} で 5",
			args: args{n: 3, a: []int{2, 7, 4}},
			want: 5,
		},
		{
			name: "[]int{20, 18, 2, 18} で 18",
			args: args{n: 4, a: []int{20, 18, 2, 18}},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC088_B(tt.args.n, tt.args.a); got != tt.want {
				t.Errorf("ABC086_A() = %v, want %v", got, tt.want)
			}
		})
	}
}
