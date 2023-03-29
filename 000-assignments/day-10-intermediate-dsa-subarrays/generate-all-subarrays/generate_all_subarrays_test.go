package generateallsubarrays

import (
	"reflect"
	"testing"
)

func TestGenerateAllSubarrays(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"test1", args{[]int{1, 2, 3}}, [][]int{{1}, {1, 2}, {1, 2, 3}, {2}, {2, 3}, {3}}},
		{"test2", args{[]int{1, 2, 3, 4}}, [][]int{{1}, {1, 2}, {1, 2, 3}, {1, 2, 3, 4}, {2}, {2, 3}, {2, 3, 4}, {3}, {3, 4}, {4}}},
		{"test3", args{[]int{5, 2, 1, 4}}, [][]int{{5}, {5, 2}, {5, 2, 1}, {5, 2, 1, 4}, {2}, {2, 1}, {2, 1, 4}, {1}, {1, 4}, {4}}},
		{"test4", args{[]int{9, 36, 63, 63, 26, 87, 28, 77, 93, 7}}, [][]int{{9}, {9, 36}, {9, 36, 63}, {9, 36, 63, 63}, {9, 36, 63, 63, 26}, {9, 36, 63, 63, 26, 87}, {9, 36, 63, 63, 26, 87, 28}, {9, 36, 63, 63, 26, 87, 28, 77}, {9, 36, 63, 63, 26, 87, 28, 77, 93}, {9, 36, 63, 63, 26, 87, 28, 77, 93, 7}, {36}, {36, 63}, {36, 63, 63}, {36, 63, 63, 26}, {36, 63, 63, 26, 87}, {36, 63, 63, 26, 87, 28}, {36, 63, 63, 26, 87, 28, 77}, {36, 63, 63, 26, 87, 28, 77, 93}, {36, 63, 63, 26, 87, 28, 77, 93, 7}, {63}, {63, 63}, {63, 63, 26}, {63, 63, 26, 87}, {63, 63, 26, 87, 28}, {63, 63, 26, 87, 28, 77}, {63, 63, 26, 87, 28, 77, 93}, {63, 63, 26, 87, 28, 77, 93, 7}, {63}, {63, 26}, {63, 26, 87}, {63, 26, 87, 28}, {63, 26, 87, 28, 77}, {63, 26, 87, 28, 77, 93}, {63, 26, 87, 28, 77, 93, 7}, {26}, {26, 87}, {26, 87, 28}, {26, 87, 28, 77}, {26, 87, 28, 77, 93}, {26, 87, 28, 77, 93, 7}, {87}, {87, 28}, {87, 28, 77}, {87, 28, 77, 93}, {87, 28, 77, 93, 7}, {28}, {28, 77}, {28, 77, 93}, {28, 77, 93, 7}, {77}, {77, 93}, {77, 93, 7}, {93}, {93, 7}, {7}}},
		{"test_large", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}, [][]int{{1}, {1, 2}, {1, 2, 3}, {1, 2, 3, 4}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5, 6}, {1, 2, 3, 4, 5, 6, 7}, {1, 2, 3, 4, 5, 6, 7, 8}, {1, 2, 3, 4, 5, 6, 7, 8, 9}, {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, {2}, {2, 3}, {2, 3, 4}, {2, 3, 4, 5}, {2, 3, 4, 5, 6}, {2, 3, 4, 5, 6, 7}, {2, 3, 4, 5, 6, 7, 8}, {2, 3, 4, 5, 6, 7, 8, 9}, {2, 3, 4, 5, 6, 7, 8, 9, 10}, {3}, {3, 4}, {3, 4, 5}, {3, 4, 5, 6}, {3, 4, 5, 6, 7}, {3, 4, 5, 6, 7, 8}, {3, 4, 5, 6, 7, 8, 9}, {3, 4, 5, 6, 7, 8, 9, 10}, {4}, {4, 5}, {4, 5, 6}, {4, 5, 6, 7}, {4, 5, 6, 7, 8}, {4, 5, 6, 7, 8, 9}, {4, 5, 6, 7, 8, 9, 10}, {5}, {5, 6}, {5, 6, 7}, {5, 6, 7, 8}, {5, 6, 7, 8, 9}, {5, 6, 7, 8, 9, 10}, {6}, {6, 7}, {6, 7, 8}, {6, 7, 8, 9}, {6, 7, 8, 9, 10}, {7}, {7, 8}, {7, 8, 9}, {7, 8, 9, 10}, {8}, {8, 9}, {8, 9, 10}, {9}, {9, 10}, {10}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateAllSubarrays(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateAllSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkGenerateAllSubarrays-8   	  137402	      7913 ns/op	    7088 B/op	     167 allocs/op
func BenchmarkGenerateAllSubarrays(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateAllSubarrays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	}
}
