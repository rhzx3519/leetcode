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
    bool hasPathSum(TreeNode* root, int sum) {
        if (root == nullptr) return false;
        if (root->left == nullptr && root->right == nullptr)
            if (root->val == sum)
                return true;
                
        bool res = false;
        if (root->left)
            res = res || hasPathSum(root->left, sum - root->val);
        if (root->right)
            res = res || hasPathSum(root->right, sum - root->val);
        
        return res;
    }
    
};