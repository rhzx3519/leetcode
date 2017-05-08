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
    int rob(TreeNode* root) {
        vector<int> v = dfs(root);
        return max(v[0], v[1]);
    }
    
    vector<int> dfs(TreeNode* root) {
        vector<int> res = {0, 0};
        if (root == NULL) return res;
        
        vector<int> left = dfs(root->left);
        vector<int> right = dfs(root->right);
        
        //rob当前结点
        int a1 = root->val + left[1] + right[1];
        
        // not rob当前结点
        int a2 = max(left[0], left[1]) + max(right[0], right[1]);
        res[0] = a1;
        res[1] = a2;
        
        return res;
    }
};