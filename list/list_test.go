package list

import "testing"

func Test_isPalindrome(t *testing.T) {
	head1 := &ListNode{Val: 1}
	head1.Next = &ListNode{Val: 2}
	head1.Next.Next = &ListNode{Val: 2}
	head1.Next.Next.Next = &ListNode{Val: 1}
	tests := []struct {
		name string
		head *ListNode
		want bool
	}{
		{"test1", head1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindrome(tt.head)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
