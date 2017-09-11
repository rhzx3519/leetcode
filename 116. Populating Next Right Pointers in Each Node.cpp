/**
 * Definition for binary tree with next pointer.
 * struct TreeLinkNode {
 *  int val;
 *  TreeLinkNode *left, *right, *next;
 *  TreeLinkNode(int x) : val(x), left(NULL), right(NULL), next(NULL) {}
 * };
 */
class Solution {
public:
    void connect(TreeLinkNode *root) {
        while (root) {
            TreeLinkNode *next = nullptr; // the first node of next level  
            TreeLinkNode *pre = nullptr; // previous node on the same level  
            for (; root; root = root->next) {
                if (next == nullptr) next = root->left ? root->left : root->right;
                if (root->left) {
                    if (pre) pre->next = root->left;
                    pre = root->left;
                }
                if (root->right) {
                    if (pre) pre->next = root->right;
                    pre = root->right;
                }
            }
            root = next;// turn to next level  
        }
    }
};