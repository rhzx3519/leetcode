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
    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
        if (root == NULL || p == NULL || q == NULL)
            return NULL;
        
        unordered_map<TreeNode*, int> depth;
        unordered_map<TreeNode*, TreeNode*> parent;
        depth[NULL] = NULL;
        parent[root] = NULL;
        
        traverse(root, 1, depth, parent);
        
        while (depth[p] > depth[q]) 
            p = parent[p];
        while (depth[q] > depth[p]) 
            q = parent[q];
            
        while (p != q){
            p = parent[p];
            q = parent[q];
        }
        
        depth.clear();
        parent.clear();
        
        return p;
    }
    
    void traverse(TreeNode* root, int d, unordered_map<TreeNode*, int> &depth, unordered_map<TreeNode*, TreeNode*> &parent) {
        depth[root] = d;
        
        if (root->left) {
            parent[root->left] = root;
            traverse(root->left, d+1, depth, parent);
        }

        if (root->right) {
            parent[root->right] = root;
            traverse(root->right, d+1, depth, parent);
        }
    }
};