/**
 * Definition for binary tree
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class BSTIterator {
public:
    BSTIterator(TreeNode *root) {
        TreeNode* p1 = root;
        while (p1) {
            st.push(p1);
            p1 = p1->left;
        }
    }

    /** @return whether we have a next smallest number */
    bool hasNext() {
        return !st.empty();
    }

    /** @return the next smallest number */
    int next() {
        TreeNode *p1 = st.top();
        st.pop();
        TreeNode *p2 = p1->right;
        while (p2) {
            st.push(p2);
            p2 = p2->left;
        }
        return p1->val;
    }
    
private:
    stack<TreeNode*> st;
};

/**
 * Your BSTIterator will be called like this:
 * BSTIterator i = BSTIterator(root);
 * while (i.hasNext()) cout << i.next();
 */