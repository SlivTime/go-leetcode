package remove_duplicates

import (
	"fmt"
	"reflect"
	"testing"
)

func toString(l *ListNode) string {
	result := ""
	if l != nil {
		result += "["
		cur := l
		for {
			if cur == nil {
				break
			}
			result += fmt.Sprintf("%d", cur.Val)
			cur = cur.Next
			if cur != nil {
				result += ","
			}
		}
		result += "]"
	}
	return result
}

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"example1",
			args{
				&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, &ListNode{4, &ListNode{4, &ListNode{5, nil}}}}}}},
			},
			&ListNode{1, &ListNode{2, &ListNode{5, nil}}},
		},
		{
			"example2",
			args{
				&ListNode{1, &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, nil}}}}},
			},
			&ListNode{2, &ListNode{3, nil}},
		},
		{
			"empty",
			args{
				nil,
			},
			nil,
		},
		{
			"example1",
			args{
				&ListNode{1, nil},
			},
			&ListNode{1, nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", toString(got), toString(tt.want))
			}
		})
	}
}
