package algorithms

//Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.

func mergeKListsV0(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	end := len(lists) - 1
	r := lists[end]
	for i := end; i > 0; i-- {
		r = mergeTwoLists(r, lists[i-1])
	}
	return r
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return mergeKListsHelper(lists)
}

func mergeKListsHelper(lists []*ListNode) *ListNode {
	length := len(lists)
	half := length / 2

	if length == 1 {
		return lists[0]
	} else if length == 2 {
		var (
			c1, c2   = lists[0], lists[1]
			res, cur *ListNode
		)

		if c1 == nil {
			return c2
		}
		if c2 == nil {
			return c1
		}

		if c1.Val < c2.Val {
			res, cur, c1 = c1, c1, c1.Next
		} else {
			res, cur, c2 = c2, c2, c2.Next
		}

		for c1 != nil && c2 != nil {
			if c1.Val < c2.Val {
				cur.Next, c1 = c1, c1.Next
			} else {
				cur.Next, c2 = c2, c2.Next
			}
			cur = cur.Next
		}

		if c1 != nil {
			cur.Next = c1
		}
		if c2 != nil {
			cur.Next = c2
		}
		return res
	}
	return mergeKLists([]*ListNode{mergeKLists(lists[half:]), mergeKLists(lists[:half])})
}