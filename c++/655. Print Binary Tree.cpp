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
    vector<vector<string>> printTree(TreeNode* root) {
        int m = 0;
        int n = 0;
        n = dfs(root, 1, m);
        vector<vector<string>> res(m, vector<string>(n, ""));
        fill(root, 0, n-1, 0, res);
        return res;
    }

    void fill(TreeNode* root, int l , int r, int d, vector<vector<string>>& res) {
        int mid = l + (r-l)/2;
        res[d][mid] = to_string(root->val);
        if (root->left)
            fill(root->left, l, mid-1, d+1, res);
        if (root->right)
            fill(root->right, mid+1, r, d+1, res);
    }

    int dfs(TreeNode* root, int cur_depth, int &depth) {
        if (!root) return 0;
        if (!root->left && !root->right) {
            depth = max(cur_depth, depth);
            return 1;
        }
            
        int l1 = dfs(root->left, cur_depth+1, depth);
        int l2 = dfs(root->right, cur_depth+1, depth);
        return 1 + max(l1, l2)*2;
    }
};

// 1. dfs确定高度和宽度
// 2. dfs 二分填充