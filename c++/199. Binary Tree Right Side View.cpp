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
    vector<int> rightSideView(TreeNode* root) {
        vector<int> res;
        if (root == NULL) return res;
        
        dfs(root, 0, res);
        return res;
    }
    
private:
    void dfs(TreeNode* root, int depth, vector<int> &res) {
        if (res.size() <= depth)
            res.push_back(root->val);
        
        if (root->right != NULL)
            dfs(root->right, depth + 1, res);    
        
        if (root->left != NULL)
            dfs(root->left, depth + 1, res);
    }
};


// 基于层次遍历更好理解
class Solution {
public:
    vector<int> rightSideView(TreeNode* root) {
        queue<TreeNode*> que;
        que.push(root);
        que.push(NULL);
        vector<int> res;
        vector<int> v;
        
        while (!que.empty()) {
            TreeNode* p = que.front();
            que.pop();
            if (p)
            {
                v.push_back(p->val);
                if (p->left)
                    que.push(p->left);
                if (p->right)
                    que.push(p->right);
            } else {
                if (!v.empty()) {
                    que.push(NULL);
                    res.push_back(v.back());
                    v.clear();
                }
            }

        }
        return res;
    }
};