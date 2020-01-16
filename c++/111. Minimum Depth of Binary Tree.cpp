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
    int minDepth(TreeNode* root) {
        if (root == nullptr) return 0;
        if (root->left == nullptr && root->right == nullptr) return 1;
        
        int res = INT_MAX;
        if (root->left) 
            res = min(res, 1 + minDepth(root->left));
        
        if (root->right)
            res = min(res, 1 + minDepth(root->right));
        
        return res;
    }

};