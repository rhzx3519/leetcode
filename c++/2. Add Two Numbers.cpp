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
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        ListNode *head, *p = NULL;
        ListNode *p1 = l1, *p2 = l2;
        int add = 0;
        while (p1 || p2 || add) {
            int t = 0;
            if (p1)
                t += p1->val;
            if (p2)
                t += p2->val;
            t += add;
            add = t/10;
 
            if (p != NULL) {
                p->next = new ListNode(t%10);
                p = p->next;
            } else {
                p = new ListNode(t%10);
                head = p;
            }

            if (p1) p1 = p1->next;
            if (p2) p2 = p2->next;
        }
        
        return head;
    }
};