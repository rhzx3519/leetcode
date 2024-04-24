package Q2385;

public class Q2385 {
}

class TreeNode {
     int val;
     TreeNode left;
     TreeNode right;
     TreeNode() {}
     TreeNode(int val) { this.val = val; }
     TreeNode(int val, TreeNode left, TreeNode right) {
         this.val = val;
         this.left = left;
         this.right = right;
     }
}

/**
 * https://leetcode.cn/problems/amount-of-time-for-binary-tree-to-be-infected/?envType=daily-question&envId=2024-04-24
 */
class Solution {
    private int ans = 0;
    private int target = 0;

    private int dfs(TreeNode root) {
        if (root == null) return 0;
        int l = dfs(root.left);
        int r = dfs(root.right);
        if (root.val == target) {
            ans = -Math.min(l, r);
            return 1;
        }
        if (l > 0 || r > 0) {
            ans = Math.max(ans, Math.abs(l) + Math.abs(r));
            return Math.max(l, r) + 1;
        }
        return Math.min(l, r) - 1;
    }

    public int amountOfTime(TreeNode root, int start) {
        target = start;
        dfs(root);
        return ans;
    }
}




