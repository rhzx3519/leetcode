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
    int maxPathSum(TreeNode* root) {
        int res = INT_MIN;
        dfs(root, res);
        return res;
    }
    
    int dfs(TreeNode* root, int &res) {
        if (!root) return 0;
        int l = max(0, dfs(root->left, res));
        int r = max(0, dfs(root->right, res));
        int sum = l + r + root->val;
        res = max(res, sum);
        return max(0, max(l, r)) + root->val;
    }
};