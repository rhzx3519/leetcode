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
    ListNode *detectCycle(ListNode *head) {
        unordered_set<ListNode*> us;
        while (head) {
            if (us.find(head) == us.end())
                us.insert(head);
            else
                return head;
            head = head->next;
        }
        return head;
    }
};

class Solution {
public:
	ListNode * detectCycle(ListNode *head) {
		if (head == NULL || head->next == NULL) return NULL;
		ListNode* slow = head->next, *fast = head->next->next;
		while (fast && fast->next) {
			if (slow == fast)
				break;
			slow = slow->next;
			fast = fast->next->next;
		}

		fast = head;
		while (slow != fast) {
			slow = slow->next;
			fast = fast->next;
		}

		return fast;
	}
};