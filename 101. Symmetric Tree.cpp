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
    bool isSymmetric(TreeNode* root) {
        if (root == nullptr) return true;
        
        return foo(root->left, root->right);
    }
    
    bool foo(TreeNode* left, TreeNode* right) {
        if (left == nullptr)
            return right == nullptr;
        if (right == nullptr)
            return left = nullptr;
            
        return left->val == right->val &&
            foo(left->left, right->right) && foo(left->right, right->left);
    }
};