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
    vector<TreeNode*> findDuplicateSubtrees(TreeNode* root) {
        unordered_map<string, int> mp;
        vector<TreeNode*> res;
        dfs(root, mp, res);
        return res;
    }

    string dfs(TreeNode* root, unordered_map<string, int> &mp, vector<TreeNode*> &res) {
        if (!root) return "#";
        string str = to_string(root->val) + " " + dfs(root->left, mp, res) + " " + dfs(root->right, mp, res);
        if (mp[str]==1) res.push_back(root);
        mp[str]++;
        return str;
    }
};

// 二叉树序列化