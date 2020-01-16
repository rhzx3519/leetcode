/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    TreeNode* sortedListToBST(ListNode* head) {
        if (head == nullptr) return nullptr;
        ListNode *p = head;
        int len = 0;
        while (p) {
            p = p->next;
            len++;
        }
        if (len == 1)
            return new TreeNode(head->val);
        else if (len == 2) {
            TreeNode *root = new TreeNode(head->val);
            root->right = new TreeNode(head->next->val);
            return root;
        } else {
            int n = (len)/2, cnt = 0;
            ListNode *p = head, *pre = nullptr;
            while (p && cnt != n) {
                pre = p;
                p = p->next;
                cnt++;
            }
            pre->next = nullptr;
            TreeNode *root = new TreeNode(p->val);
            root->left = sortedListToBST(head);
            root->right = sortedListToBST(p->next);
            return root;
        }
    }
};