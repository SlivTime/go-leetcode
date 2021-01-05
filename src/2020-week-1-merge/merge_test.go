package merge

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

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"example1",
			args{
				&ListNode{1, &ListNode{2, &ListNode{4, nil}}},
				&ListNode{1, &ListNode{3, &ListNode{4, nil}}},
			},
			&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}},
		},
		{
			"example2",
			args{
				nil,
				nil,
			},
			nil,
		},
		{
			"example3",
			args{
				nil,
				&ListNode{0, nil},
			},
			&ListNode{0, nil},
		},
		{
			"second empty",
			args{
				&ListNode{0, nil},
				nil,
			},
			&ListNode{0, nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", toString(got), toString(tt.want))
			}
		})
	}
}
