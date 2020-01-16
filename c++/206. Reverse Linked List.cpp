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
    ListNode* reverseList(ListNode* head) {
        ListNode new_head(1), *p = head, *t;
        while (p) {
            t= p;
            p = p->next;
            t->next = new_head.next;
            new_head.next = t;
        }
        
        return new_head.next;
    }
};