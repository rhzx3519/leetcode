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
    int findBottomLeftValue(TreeNode* root) {
        int res = 0;
        queue<TreeNode*> que;
        que.push(root);
        while (!que.empty()) {
            int n = que.size();
            res = que.front()->val;
            while (n-->0) {
                TreeNode*  p = que.front();
                que.pop();
                if (p->left)
                    que.push(p->left);
                if (p->right)
                    que.push(p->right);
            }
        }
        return res;
    }
};