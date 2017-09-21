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
    ListNode* mergeTwoLists(ListNode* l1, ListNode* l2) {
        ListNode* p1 = l1, *p2 = l2, *p = NULL;
        if (!p1) return p2;
        if (!p2) return p1;
        ListNode* head = p1->val < p2->val ? p1 : p2;
        while (p1 || p2)
        {
            if (p1 && p2)
            {
                if (p1->val < p2->val)
                {
                    if (!p)
                    {
                        p = p1; p1 = p1->next;
                    }
                    else
                    {
                        p->next = p1;
                        p = p1;
                        p1 = p1->next;
                    }
                }
                else
                {
                    if (!p)
                    {
                        p = p2; p2 = p2->next;
                    }
                    else
                    {
                        p->next = p2;
                        p = p2;
                        p2 = p2->next;
                    }
                }
            }
            else if (p1)
            {
                p->next = p1;
                p = p->next;
                p1 = p1->next;
            }
            else if (p2)
            {
                p->next = p2;
                p = p->next;
                p2 = p2->next;
            }
        }
        
        return head;
    }
};