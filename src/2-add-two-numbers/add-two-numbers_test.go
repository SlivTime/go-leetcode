package addTwoNumbers

import (
	"fmt"
	"testing"
)

func isEqual(l1 *ListNode, l2 *ListNode) bool {
	currentL1 := l1
	currentL2 := l2
	for {
		if currentL1 == nil && currentL2 == nil {
			break
		}
		if currentL1 == nil || currentL2 == nil || currentL1.Val != currentL2.Val {
			return false
		}
		currentL1 = currentL1.Next
		currentL2 = currentL2.Next
	}
	return true
}

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

func Test_addTwoNumbersLoop(t *testing.T) {
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
			"Example",
			args{
				&ListNode{
					2,
					&ListNode{
						4,
						&ListNode{
							3,
							nil,
						},
					},
				},
				&ListNode{
					5,
					&ListNode{
						6,
						&ListNode{
							4,
							nil,
						},
					},
				},
			},
			&ListNode{
				7,
				&ListNode{
					0,
					&ListNode{
						8,
						nil,
					},
				},
			},
		},
		{
			"Carry on end",
			args{
				&ListNode{
					5,
					nil,
				},
				&ListNode{
					5,
					nil,
				},
			},
			&ListNode{
				0,
				&ListNode{
					1,
					nil,
				},
			},
		},
		{
			"Second is zero",
			args{
				&ListNode{
					1,
					&ListNode{
						8,
						nil,
					},
				},
				&ListNode{
					0,
					nil,
				},
			},
			&ListNode{
				1,
				&ListNode{
					8,
					nil,
				},
			},
		},
		{
			"First is zero",
			args{
				&ListNode{
					0,
					nil,
				},
				&ListNode{
					7,
					&ListNode{
						3,
						nil,
					},
				},
			},
			&ListNode{
				7,
				&ListNode{
					3,
					nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbersLoop(tt.args.l1, tt.args.l2); !isEqual(got, tt.want) {
				t.Errorf("addTwoNumbersLoop() = %v, want %v", toString(got), toString(tt.want))
			}
		})
	}
}

func Test_listToInt(t *testing.T) {
	type args struct {
		l *ListNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Explicit zero",
			args{
				&ListNode{
					0,
					nil,
				},
			},
			0,
		},
		{
			"Implicit zero",
			args{
				&ListNode{},
			},
			0,
		},
		{
			"1",
			args{
				&ListNode{
					1,
					nil,
				},
			},
			1,
		},
		{
			"37",
			args{
				&ListNode{
					7,
					&ListNode{
						3,
						nil,
					},
				},
			},
			37,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := listToInt(tt.args.l); got != tt.want {
				t.Errorf("listToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intToList(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"Explicit zero",
			args{
				0,
			},
			&ListNode{
				0,
				nil,
			},
		},
		{
			"1",
			args{
				1,
			},
			&ListNode{
				1,
				nil,
			},
		},
		{
			"37",
			args{
				37,
			},
			&ListNode{
				7,
				&ListNode{
					3,
					nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToList(tt.args.n); !isEqual(got, tt.want) {
				t.Errorf("intToList() = %v, want %v", toString(got), toString(tt.want))
			}
		})
	}
}

func Test_addTwoNumbers(t *testing.T) {
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
			"Example",
			args{
				&ListNode{
					2,
					&ListNode{
						4,
						&ListNode{
							3,
							nil,
						},
					},
				},
				&ListNode{
					5,
					&ListNode{
						6,
						&ListNode{
							4,
							nil,
						},
					},
				},
			},
			&ListNode{
				7,
				&ListNode{
					0,
					&ListNode{
						8,
						nil,
					},
				},
			},
		},
		{
			"Carry on end",
			args{
				&ListNode{
					5,
					nil,
				},
				&ListNode{
					5,
					nil,
				},
			},
			&ListNode{
				0,
				&ListNode{
					1,
					nil,
				},
			},
		},
		{
			"Second is zero",
			args{
				&ListNode{
					1,
					&ListNode{
						8,
						nil,
					},
				},
				&ListNode{
					0,
					nil,
				},
			},
			&ListNode{
				1,
				&ListNode{
					8,
					nil,
				},
			},
		},
		{
			"First is zero",
			args{
				&ListNode{
					0,
					nil,
				},
				&ListNode{
					7,
					&ListNode{
						3,
						nil,
					},
				},
			},
			&ListNode{
				7,
				&ListNode{
					3,
					nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !isEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", toString(got), toString(tt.want))
			}
		})
	}
}
