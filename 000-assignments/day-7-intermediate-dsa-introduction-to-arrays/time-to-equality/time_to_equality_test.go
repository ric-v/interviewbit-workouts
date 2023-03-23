package timetoequality

import "testing"

func TestTimeToEquality(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{[]int{3, 3, 3, 3}}, 0},
		{"Test 2", args{[]int{1, 2, 3, 4}}, 6},
		{"Test 3", args{[]int{1, 1, 1, 1}}, 0},
		{"Test 4", args{[]int{1, 1, 1, 2}}, 3},
		{"Test 5", args{[]int{1, 1, 2, 2}}, 2},
		{"Test 6", args{[]int{1, 2, 2, 2}}, 1},
		{"Test 7", args{[]int{2, 4, 1, 3, 2}}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeToEquality(tt.args.A); got != tt.want {
				t.Errorf("TimeToEquality() = %v, want %v", got, tt.want)
			}
		})
	}
}
