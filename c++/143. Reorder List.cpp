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
    void reorderList(ListNode* head) {
        if (!head) return;
        
        ListNode *p = head;
        int len = 0, cnt = 0;
        while (p) {
            len++;
            p = p->next;
        }
        
        p = head;
        while (p->next && cnt < len/2 - 1) {
            cnt++;
            p = p->next;
        }
        
        ListNode* new_head = new ListNode(0), *p1, *t, *q, *r;
        p1 = p->next;
        p->next = NULL;
        while (p1) {
            t = p1->next;
            p1->next = new_head->next;
            new_head->next = p1;
            p1 = t;
        }
        
        //merge
        p = head; q = new_head->next; r = new_head;
        while (p && q) {
            r->next = p;
            t = p->next;
            p->next = q;
            p = t;
            r = q;
            q = q->next;
        }
        
        head = new_head->next;
        delete(new_head);
    }
};