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
    TreeNode* addOneRow(TreeNode* root, int v, int d) {
        if (d==1) {
            TreeNode* newRoot = new TreeNode(v);
            newRoot->left = root;
            return newRoot;
        }
        TreeNode *p = NULL;
        queue<TreeNode*> que;
        que.push(root);

        int cnt = 0;
        while (!que.empty()) {
            int sz = que.size();
            vector<TreeNode*> vec;
            cnt++;
            while (sz--) {
                p = que.front();
                que.pop();
                vec.push_back(p);
                if (p->left)
                    que.push(p->left);
                if (p->right)
                    que.push(p->right);
            }

            if (cnt==d-1)
            {
                for (TreeNode* node : vec) {
                    TreeNode *p1 = node->left;
                    TreeNode *p2 = node->right;

                    node->left = new TreeNode(v);
                    node->right = new TreeNode(v);
                    node->left->left = p1;
                    node->right->right = p2;
                }
                break;
            }
            
        }

        return root;
    }
};