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
    ListNode* rotateRight(ListNode* head, int k) {
		if (head == nullptr || k == 0) return head;

		ListNode *p = head, *tail = nullptr;
		int len = 0;
		while (p) {
			len++;
			tail = p;
			p = p->next;
		}
		k %= len;
		int cnt = len - k;
		p = head;
		while (cnt > 1) {
			p = p->next;
			cnt--;
		}

		tail->next = head;
		if (p) {
			head = p->next;
			p->next = nullptr;
		}
			
		return head;
    }
};