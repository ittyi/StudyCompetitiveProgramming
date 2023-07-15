package tenselectedpastquestions

import "testing"

func TestABC085_B(t *testing.T) {
	type args struct {
		n int
		d []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[]int{10, 8, 8, 6} で 3",
			args: args{n: 4, d: []int{10, 8, 8, 6}},
			want: 3,
		},
		{
			name: "[]int{15, 15, 15} で 1",
			args: args{n: 3, d: []int{15, 15, 15}},
			want: 1,
		},
		{
			name: "[]int{50, 30, 50, 100, 50, 80, 30} で 4",
			args: args{n: 7, d: []int{50, 30, 50, 100, 50, 80, 30}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC085_B(tt.args.n, tt.args.d); got != tt.want {
				t.Errorf("ABC086_A() = %v, want %v", got, tt.want)
			}
		})
	}
}
