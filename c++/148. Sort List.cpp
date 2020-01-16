/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    ListNode* sortList(ListNode* head) {
        quick_sort(head, nullptr);
        return head;
    }
    
	void quick_sort(ListNode *begin, ListNode *end) {
		if (begin != end) {
			ListNode *p = partition(begin, end);
			quick_sort(begin, p);
			quick_sort(p->next, end);
		}
	}
    
    ListNode* partition(ListNode *head, ListNode *end) {
        ListNode *p = head, *q = head;
        int key = p->val;
        while (p != end) {
            if (p->val < key) {
                q = q->next;
                swap(p->val, q->val);
            }
            p = p->next;
        }
        swap(head->val, q->val);
        return q;
    }
};