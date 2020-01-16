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
    vector<int> preorderTraversal(TreeNode* root) {
        vector<int> res;
        dfs(root, res);
        
        return res;
    }
    
    void iterate(TreeNode* root, vector<int>& res)
    {
        if (!root) return;
        stack<TreeNode*> st;
        TreeNode* p = root;
        
        while (true)
        {
            while (p)
            {
                res.push_back(p->val);
                st.push(p);
                p = p->left;
            }
            if (st.empty()) break;
            p = st.top()->right;
            st.pop();
        }
    }
    
    void dfs(TreeNode* root, vector<int>& res)
    {
        if (!root) return;
        res.push_back(root->val);
        dfs(root->left, res);
        dfs(root->right, res);
    }
};