package split

import (
	"reflect"
	"testing"
)

// 测试函数
func TestSplit(t *testing.T) {
	// 定义测试用例类型
	type test struct {
		input string
		sep   string
		want  []string
	}

	// 定义存储测试用例的切片
	tests := map[string]test{
		"simple": {
			input: "a:b:c",
			sep:   ":",
			want:  []string{"a", "b", "c"},
		},
		"wrong sep": {
			input: "a:b:c",
			sep:   ",",
			want:  []string{"a:b:c"},
		},
		"more sep": {
			input: "apple",
			sep:   "pl",
			want:  []string{"ap", "e"},
		},
		"leading sep": {
			input: "我爱她",
			sep:   "爱",
			want:  []string{"我", "她"},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("expected:%#v, got:%#v", tc.want, got)
			}
		})
	}
}

// 基准测试
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("我爱你你爱我", "爱")
	}
}

// 并行测试
func BenchmarkSplitParallel(b *testing.B) {
	// b.SetParallelism(1) // 设置使用的CPU数
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Split("我爱你你爱我", "爱")
		}
	})
}
