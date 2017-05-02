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
    vector<TreeNode*> generateTrees(int n) {
        vector<TreeNode*> res;
        if (n == 0) return res;
        res = dfs(1, n);
        return res;
    }
    
    vector<TreeNode*> dfs(int l, int r) {
        vector<TreeNode*> res;
        if (l > r) {
            res.push_back(nullptr);
            return res;
        }
        if (l == r) {
            res.push_back(new TreeNode(l));
            return res;
        }
        
        for (int i = l; i <= r; i++) {
            vector<TreeNode*> left = dfs(l, i-1);
            vector<TreeNode*> right = dfs(i+1, r);
            for(auto t1 : left) {
                for (auto t2 : right) {
                    TreeNode *root = new TreeNode(i);
                    root->left = t1;
                    root->right = t2;
                    res.push_back(root);
                }
            }
        }
        
        return res;
    }
    

};