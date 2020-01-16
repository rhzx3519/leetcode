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
    ListNode* deleteDuplicates(ListNode* head) {
        ListNode* new_head = new ListNode(0);
        ListNode* p = head, *pre = new_head; pre->next = p;
        while (p)
        {
            while (p->next && pre->next->val == p->next->val)
            {
                p = p->next;
            }
            
            if (pre->next == p)
                pre = pre->next;
            else
                pre->next = p->next;
            p = p->next;
        }
        
        return new_head->next;
    }
};