/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */


/// 递归
class Solution {
public:
    TreeNode* insertIntoBST(TreeNode* root, int val) {
        if (!root) return new TreeNode(val);
        if (root->val > val)
            root->left = insertIntoBST(root->left, val);
        else
            root->right = insertIntoBST(root->right, val);
        return root;
    }
};

// 迭代
class Solution {
public:
    TreeNode* insertIntoBST(TreeNode* root, int val) {
        if (!root) return new TreeNode(val);

        TreeNode* parent = root, *p = root;
        while (p) {
            parent = p;
            p = p->val > val ? p->left : p->right;
        }

        if (parent->val > val)
            parent->left = new TreeNode(val);
        else
            parent->right = new TreeNode(val);
        return root;
    }
};