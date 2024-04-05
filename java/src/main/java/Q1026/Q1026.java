package Q1026;

class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;
    TreeNode(int val) { this.val = val; }
}

public class Q1026 {
    private int ans = 0;

    private void dfs(TreeNode root, int max, int min) {
        if (root == null) {
            return;
        }
        ans = Math.max(ans, Math.max(Math.abs(root.val - max), Math.abs(root.val - min)));
        max = Math.max(root.val, max);
        min = Math.min(root.val, min);

        if (root.left != null) {
            dfs(root.left, max, min);
        }
        if (root.right != null) {
            dfs(root.right, max, min);
        }
    }

    public int maxAncestorDiff(TreeNode root) {
        dfs(root, root.val, root.val);
        return ans;
    }
}
