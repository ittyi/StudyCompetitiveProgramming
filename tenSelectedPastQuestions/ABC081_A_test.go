package tenselectedpastquestions

import "testing"

func TestABC081_A(t *testing.T) {
	type args struct {
		str string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "偶数",
			args: args{str:"101"},
			want: 2,
		},
		{
			name: "奇数",
			args: args{str: "000"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ABC081_A(tt.args.str); got != tt.want {
				t.Errorf("ABC081_A() = %v, want %v", got, tt.want)
			}
		})
	}
}