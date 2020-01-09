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
    bool isPalindrome(ListNode* head) {
        if (head == NULL) return true;
        ListNode *fast, *slow;
        fast = head; slow = head;
        while (fast->next && fast->next->next) {
            slow = slow->next;
            fast = fast->next->next;
        }
        slow->next = reverse(slow->next);
        slow = slow->next;
        
        while (slow) {
            if (head->val != slow->val)
                return false;
            head = head->next;
            slow = slow->next;
        }
        
        return true;
    }
    
    ListNode* reverse(ListNode* head) {
        ListNode *next = NULL, *pre = NULL ;
        while (head) {
            next = head->next;
            head->next = pre;
            pre = head;
            head = next;
        }
        return pre;
    }
};