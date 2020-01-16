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
    int countNodes(TreeNode* root) {
        if (root == NULL) return 0;
        int l1 = 0, l2 = 0;
        TreeNode* p;
        
        p = root;
        while (p) {
            l1++;
            p = p->left;
        }
        p = root;
        while (p) {
            l2++;
            p = p->right;
        }
        if (l1 == l2) 
            return (2<<(l1-1))-1;
        
        return countNodes(root->left) + countNodes(root->right) + 1;
    }
};