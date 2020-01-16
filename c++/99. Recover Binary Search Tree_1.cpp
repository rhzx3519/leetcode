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
    void recoverTree(TreeNode* root) {
        prev = cur = NULL;
        t1 = t2 = NULL;
        stack<TreeNode*> st;
        while(root || !st.empty()) {
            if (root) {
                st.push(root);
                root = root->left;
            } else {
                root = st.top();
                st.pop();
                if (prev && prev->val > root->val) {
                    if (t1 == NULL)
                        t1 = prev;
                    t2 = root;
                }
                prev = root;
                root = root->right;
            }
        }
        
        if (t1 != NULL && t2 != NULL)
            swap(t1->val, t2->val);
    }

private:
    TreeNode *prev, *cur;
    TreeNode *t1, *t2;
};