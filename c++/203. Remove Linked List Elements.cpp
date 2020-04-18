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
    ListNode* removeElements(ListNode* head, int val) {
        ListNode *p = head, *pre = NULL, *t;
        while (p) {
            if (p->val == val) {
                if (pre) {
                    t = p;
                    p = p->next;
                    pre->next = p;
                } else {
                    t = p;
                    p = p->next;
                    head = p;
                }
                delete t;
                continue;
            }
            pre = p;
            p = p->next;
           
        }
        
        return head;
    }
};