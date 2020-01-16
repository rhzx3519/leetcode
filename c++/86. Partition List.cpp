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
    ListNode* partition(ListNode* head, int x) {
        ListNode *p = head, p1(0), *q = &p1, *pre = nullptr, *t;
        while (p) {
            if (p->val < x) {
                t = p;
                p = p->next;
                if (pre)
                    pre->next = p;
                else
                    head = p;
                q->next = t;
                q = q->next;
            } else {
                pre = p;
                p = p->next;
            }
        }
        
        q->next = head;
        
        return p1.next;
    }

};