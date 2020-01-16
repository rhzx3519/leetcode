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
    TreeNode *buildTree(vector<int> &inorder, vector<int> &postorder) {
        // write your code here
        TreeNode *root = build(inorder, 0, inorder.size() - 1, postorder, 0, postorder.size() - 1);
        return root;
    }
    
    TreeNode* build(vector<int> &inorder, int li, int hi, vector<int> &postorder, int lp, int hp) {
        if (li > hi) return nullptr;
        int mid = postorder[hp];
        TreeNode *root = new TreeNode(mid);
        int i = li;
        while(i < hi && inorder[i] != mid) i++;
        root->left = build(inorder, li, i-1, postorder, lp, lp+i-1-li);
        root->right = build(inorder, i+1, hi, postorder, hp - (hi-i), hp - 1);
        return root;
    }
};