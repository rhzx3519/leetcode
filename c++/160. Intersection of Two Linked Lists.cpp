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
    ListNode *getIntersectionNode(ListNode *headA, ListNode *headB) {
        ListNode* pa = headA, *pb = headB;
        int len1 = 0, len2 = 0;
        while (pa)
        {
            len1++;
            pa = pa->next;
        }
        while (pb)
        {
            len2++;
            pb = pb->next;
        }
        
        pa = headA; pb = headB;
        if (len1 > len2)
        {
            swap(pa, pb);
            swap(len1, len2);
        } 
        
        while (len2 > len1 && pb)
        {
            pb = pb->next;
            len2--;
        }
  
        while (pa && pb)
        {
            if (pa==pb)
                return pa;
            pa = pa->next;
            pb = pb->next;
        }
        return NULL;
    }
};