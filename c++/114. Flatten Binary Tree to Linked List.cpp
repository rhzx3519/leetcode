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
    void flatten(TreeNode* root) {
        dfs(root);
    }
    
    TreeNode* dfs(TreeNode* root) {
        if (root == nullptr) return nullptr;
        TreeNode *left = dfs(root->left), *right = dfs(root->right);
        if (left) {
            left->right = root->right;
            root->right = root->left;
        }
        root->left = nullptr;
        if (right)
            return right;
        else if (left)
            return left;
        else
            return root;
        
    }
};