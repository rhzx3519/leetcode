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
    TreeNode* constructMaximumBinaryTree(vector<int>& nums) {
        return build(nums, 0, nums.size()-1);
    }


    TreeNode* build(vector<int>& nums, int l, int r) {
        if (l==r) return new TreeNode(nums[l]);
        if (l > r) return NULL;
        
        int max_val = INT_MIN;
        int max_idx = l;
        for (int i = l; i <= r; i++) {
            if (nums[i] > max_val) {
                max_val = nums[i];
                max_idx = i;
            }
        }

        TreeNode* root = new TreeNode(nums[max_idx]);
        root->left = build(nums, l, max_idx-1);
        root->right = build(nums, max_idx+1, r);
        return root;
    }
};