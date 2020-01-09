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
    vector<string> binaryTreePaths(TreeNode* root) {
        vector<string> res;
        if (root == NULL) return res;
        vector<int> path;
        path.push_back(root->val);
        dfs(root, path, res);
        return res;
    }
    
    void dfs(TreeNode* root, vector<int> &path, vector<string> &res) {
        if (root == NULL) return;
        
        
        if (root->left == NULL && root->right == NULL) {
            string s;
            for (int i = 0; i < path.size(); i++) 
                s += i == path.size() - 1 ? to_string(path[i]) : to_string(path[i]) + "->";
            res.push_back(s);
        }
        
        if (root->left) {
            path.push_back(root->left->val);
            dfs(root->left, path, res);
            path.pop_back();
        }
        
        if (root->right) {
            path.push_back(root->right->val);
            dfs(root->right, path, res);
            path.pop_back();
        }
    }
};