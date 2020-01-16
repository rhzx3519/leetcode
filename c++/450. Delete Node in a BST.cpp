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
    TreeNode* deleteNode(TreeNode* root, int key) {
        TreeNdoe* pre = NULL;
        TreeNdoe* targetNode = search(root, pre, key);
        if (!targetNode) return root;
        TreeNode* left = targetNode->left, *right = targetNode->right;
        if (left && right) {

        } else if (left) {

        } else if (right) {

        } else {
            if (is_left)
                pre->left = NULL;
            else
                pre->right = NULL;
            delete targetNode;
        }
    }

    TreeNode* search(TreeNode* root, TreeNdoe* pre, int key) {
        if (!root) return NULL;
        if (root->val == key) return root;
        else if (root->val > key)
            return search(root->left, root, key);
        return search(root->right, root, key);
    }
};