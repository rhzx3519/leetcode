package Q1766;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

/**
 * https://leetcode.cn/problems/tree-of-coprimes/description/
 * 思路：回溯查找
 */
public class Q1766 {
}

class Solution {
    private static final int MX = 51;
    private static final int[][] coprime = new int[MX][MX];

    static {
        // coprime 保存 [1, MX] 中与i互质的所有元素
        for (int i = 1; i < MX; i++) {
            int k = 0;
            for (int j = 1; j < MX; j++) {
                if (gcd(i, j) == 1) {
                    coprime[i][k++] = j;
                }
            }
        }
    }

    private static int gcd(int a, int b) {
        return b == 0 ? a : gcd(b, a % b);
    }

    public int[] getCoprimes(int[] nums, int[][] edges) {
        int n = nums.length;
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, i -> new ArrayList<>());
        for (int [] e : edges) {
            int a = e[0];
            int b = e[1];
            g[a].add(b);
            g[b].add(a);
        }

        int[] ans = new int[n];
        Arrays.fill(ans, -1);
        int[] valDepth = new int[MX];
        int[] valNodeId = new int[MX];
        dfs(0, -1, 1, g, nums, ans, valDepth, valNodeId);
        return ans;
    }

    private void dfs(int x, int fa, int depth, List<Integer>[] g,
                     int[] nums, int[] ans, int[] valDepth, int[] valNodeId) {
        int val = nums[x];

        int maxDepth = 0;
        // 计算与 val 互质的祖先节点中，其中节点深度最大
        for (int j : coprime[val]) {
            if (j == 0) {
                break;
            }
            if (valDepth[j] > maxDepth) {
                maxDepth = valDepth[j];
                ans[x] = valNodeId[j];
            }
        }

        // 回溯时恢复
        int tmpDepth = valDepth[val];
        int tmpNodeId = valNodeId[val];

        valDepth[val] = depth;  // 保存节点值等于val的最近祖先深度
        valNodeId[val] = x;     // 保存节点值等于val的最近祖先编号

        // 向下递归
        for (int y : g[x]) {
            if (y != fa) {
                dfs(y, x, depth + 1, g, nums, ans, valDepth, valNodeId);
            }
        }

        valDepth[val] = tmpDepth;
        valNodeId[val] = tmpNodeId;
    }
}

















