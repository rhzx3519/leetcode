package Q924;

import java.util.HashMap;
import java.util.Map;

public class Q924 {
}

class Solution {
    public int minMalwareSpread(int[][] graph, int[] initial) {
        int n = graph.length;
        boolean[] vis = new boolean[n];
        boolean[] isInitial = new boolean[n];
        int mn = Integer.MAX_VALUE;
        for (int x : initial) {
            isInitial[x] = true;
            mn = Math.min(mn, x);
        }

        int ans = -1;
        int maxSize = 0;
        for (int x : initial) {
            if (vis[x]) {
                continue;
            }
            nodeId = -1;
            size = 0;
            dfs(x, graph, vis, isInitial);
            if (nodeId >= 0 && (size > maxSize || size == maxSize && nodeId < ans)) {
                ans = nodeId;
                maxSize = size;
            }
        }
        return ans < 0 ? mn : ans;
    }

    private int nodeId, size;

    private void dfs(int x, int[][] graph, boolean[] vis, boolean[] isInitial) {
        vis[x] = true;
        size++;
        // 按照状态机更新 nodeId
        if (nodeId != -2 && isInitial[x]) {
            nodeId = nodeId == -1 ? x : -2;
        }
        for (int y = 0; y < graph[x].length; y++) {
            if (graph[x][y] == 1 && !vis[y]) {
                dfs(y, graph, vis, isInitial);
            }
        }
    }
}