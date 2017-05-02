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
    vector<vector<int>> levelOrderBottom(TreeNode* root) {
        vector<vector<int>> res;
        vector<int> v;
        queue<TreeNode*> que;
        que.push(root);
        que.push(NULL);
        
        while (!que.empty()) {
            TreeNode *t = que.front();
            que.pop();
            if (t) {
                v.push_back(t->val);
                if (t->left)
                    que.push(t->left);
                if (t->right)
                    que.push(t->right);
            } else {
                if (!v.empty()) {
                    res.push_back(v);
                    v.clear();
                    que.push(NULL);
                }
            }
        }
        
        reverse(res.begin(), res.end());
        return res;
    }
};