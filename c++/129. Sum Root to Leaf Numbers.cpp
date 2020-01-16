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
    int sumNumbers(TreeNode* root) {
        int res = 0;
        vector<int> path;
        dfs(root, path, res);
        return res;
    }
    
     void dfs(TreeNode* root, vector<int> path, int &sum) {
        if (root == nullptr) return;
        if (root->left == nullptr && root->right == nullptr) {
            path.push_back(root->val);
            int t = 0;
            for (int i = 0; i < path.size(); i++)
                t += pow(10, path.size() - i - 1)*path[i];
            sum += t;
            return;
        }
        path.push_back(root->val);
        
        if (root->left != nullptr) 
            dfs(root->left, path, sum);
        
        if (root->right != nullptr)
            dfs(root->right, path, sum);
    }
};