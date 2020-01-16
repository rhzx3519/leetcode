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
    int kthSmallest(TreeNode* root, int k) {
        int res;
        priority_queue<int, vector<int>, less<int>>  pq;
        
        stack<TreeNode*> st;
        TreeNode *p = root;
        while (!st.empty() || p) {
            while (p) {
                st.push(p);
                p = p->left;
            }
            
            if (!st.empty()) {
                p = st.top();
                st.pop();
                pq.push(p->val);
                if (pq.size() > k)
                    pq.pop();
                p = p->right;
            }
        }
        
        return pq.top();
    }
    
};