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
        int add = 0, t = 0;
        while (p1 && p2) {
            t = (p1->val + p2->val + add);
            add = t/10;
 
            if (p != NULL) {
                p->next = new ListNode(t%10);
                p = p->next;
            } else {
                p = new ListNode(t%10);
                head = p;
            }

            p1 = p1->next;
            p2 = p2->next;
        }
        
        while (p1) {
            t = (p1->val + add);
            add = t/10;
            p->next = new ListNode(t%10);
            p = p->next;
            p1 = p1->next;
        }
        
        while (p2) {
            t = (p2->val + add);
            add = t/10;
            p->next = new ListNode(t%10);
            p = p->next;
            p2 = p2->next;
        }
        
        while (add) {
            p->next = new ListNode(add%10);
            add /= 10;
            p = p->next;
        }
        
        return head;
    }
};