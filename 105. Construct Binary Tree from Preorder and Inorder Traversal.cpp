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
    TreeNode *buildTree(vector<int> &preorder, vector<int> &inorder) {
        // write your code here
        TreeNode *root = build(inorder, 0, inorder.size() - 1, preorder, 0, preorder.size() - 1);
        return root;
    }
    
    TreeNode* build(vector<int> &inorder, int li, int hi, vector<int> &preorder, int lp, int hp) {
        if (li > hi) return nullptr;
        int mid = preorder[lp];
        TreeNode *root = new TreeNode(mid);
        int i = li;
        while(i < hi && inorder[i] != mid) i++;
        
        root->left = build(inorder, li, i-1, preorder, lp + 1, lp + i- li);
        root->right = build(inorder, i+1, hi, preorder, hp - (hi -i - 1), hp);
        return root;
    }
};