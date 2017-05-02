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
    ListNode* reverseBetween(ListNode* head, int m, int n) {
        int i = 1;
        ListNode *p = head, *pre = nullptr, new_head(0), *t, *q;
        while (p && i < m) {
            pre = p;
            p = p->next;
            i++;
        }
            
        if (pre){ 
            q = pre->next;
        } else {
            q = head;
        }
        
        ListNode* tail = q;
        while (q && i <= n) {
            t = q;
            q = q->next;
            i++;
            t->next = new_head.next;
            new_head.next = t;
        }
        if (pre) {
            pre->next = new_head.next;
            tail->next = q;
            return head;
        } 
        
        tail->next = q;
        return new_head.next;

    }
};