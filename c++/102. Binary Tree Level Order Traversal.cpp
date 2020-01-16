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
    vector<vector<int>> levelOrder(TreeNode* root) {
        vector<vector<int>> ret;
        queue<TreeNode*> que;
        vector<int> vec;
        if(root == NULL) return ret;
        
        que.push(root);
        que.push(NULL);
        
        while(!que.empty()) {
            TreeNode *node = que.front();
            que.pop();
            if (node) {
                vec.push_back(node->val);
                if(node->left)
                    que.push(node->left);
                if(node->right)
                    que.push(node->right);
            } else {
                if(!vec.empty()) {
                    que.push(NULL);
                    ret.push_back(vec);
                    vec.clear();
                }
            }
        }
        
        return ret;
    }
};