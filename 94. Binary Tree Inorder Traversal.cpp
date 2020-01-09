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
    vector<int> inorderTraversal(TreeNode* root) {
        vector<int> res;
        if (root == nullptr) return res;
        vector<int> lv = inorderTraversal(root->left);
        copy(lv.begin(), lv.end(), back_inserter(res));
        res.push_back(root->val);
        vector<int> rv = inorderTraversal(root->right);
        copy(rv.begin(), rv.end(), back_inserter(res));
        return res;
    }
};