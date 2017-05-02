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
        ListNode *p = head, *pre = nullptr, *t;
        while (p) {
            if (pre && pre->val == p->val) {
                t = p;
                pre->next = p->next;
                p = p->next;
                delete t;
            } else {
                pre = p;
                p = p->next;
            }
        }
        
        return head;
    }
};