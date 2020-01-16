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
    vector<int> largestValues(TreeNode* root) {
        vector<int> res;
        if (!root) return res;
        queue<TreeNode*> que;
        que.push(root);
        while (!que.empty()) {
            int max_val = INT_MIN;
            int n = que.size();
            while (--n>=0) {
                TreeNode* p = que.front();
                que.pop();
                max_val = max(max_val, p->val);
                if (p->left)
                    que.push(p->left);
                if (p->right)
                    que.push(p->right);
            }
            res.push_back(max_val);
        }

        return res;
    }
};