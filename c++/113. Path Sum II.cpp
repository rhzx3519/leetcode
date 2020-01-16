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
    vector<vector<int>> pathSum(TreeNode* root, int sum) {
        vector<vector<int>> res;
        vector<int> path;
        if (root == nullptr) return res;
        
        path.push_back(root->val);
        dfs(root, sum, path, res);
        return res;
    }
    
    void dfs(TreeNode* root, int sum, vector<int> &path, vector<vector<int>> &res) {
        if (root->left == nullptr && root->right == nullptr) {
            if (root->val == sum) {
                res.push_back(path);
                return;
            }
        }
        
        if (root->left) {
            path.push_back(root->left->val);
            dfs(root->left, sum - root->val, path, res);
            path.pop_back();
        }
        
        if (root->right) {
            path.push_back(root->right->val);
            dfs(root->right, sum - root->val, path, res);
            path.pop_back();
        }    
    }
};