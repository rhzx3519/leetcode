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
    ListNode* insertionSortList(ListNode* head) {
        ListNode *p = head, *new_head = new ListNode(0), *q, *t;
        while (p) {
            t = p;
            p = p->next;
            q = new_head;
            while(q->next && q->next->val < t->val)
                q = q->next;
            t->next = q->next;
            q->next = t;
        }
        
        head = new_head->next;
        delete(new_head);
        return head;
    }
};