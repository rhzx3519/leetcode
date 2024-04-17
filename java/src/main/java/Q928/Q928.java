package Q928;

import java.util.*;

/**
 * https://leetcode.cn/problems/minimize-malware-spread-ii/description/?envType=daily-question&envId=2024-04-17
 */
public class Q928 {
}

class Solution {
    private void dfs(int i, int[][] graph, boolean[] visited, boolean[] contaminated) {
        for (int j = 0; j < graph[i].length; j++) {
            if (graph[i][j] == 1 && !contaminated[j] && !visited[j]) {
                visited[j] = true;
               dfs(j, graph, visited, contaminated);
            }
        }
    }

    public int minMalwareSpread(int[][] graph, int[] initial) {
        int n = graph.length;
        boolean[] contaminated = new boolean[n];
        for (int i : initial) {
            contaminated[i] = true;
        }
        List<Integer>[] infectedBy = new ArrayList[n];
        Arrays.setAll(infectedBy, i -> new ArrayList<>());

        for (int i = 0; i < n; i++) {
            if (!contaminated[i]) {
                continue;
            }
            boolean[] visited = new boolean[n];
            visited[i] = true;
            dfs(i, graph, visited, contaminated);
            for (int j = 0; j < n; j++) {
                if (visited[j]) {
                    infectedBy[j].add(i);
                }
            }
        }

        Map<Integer, Integer> unclean = new HashMap<>();
        for (int i = 0; i < n; i++) {
            if (infectedBy[i].size() == 1) {
                int x = infectedBy[i].get(0);
                unclean.put(x, unclean.getOrDefault(x, 0) + 1);
            }
        }

        int ans = 0;
        int maxCount = 0;
        for (int i = 0; i < n; i++) {
            int c = unclean.getOrDefault(i, 0);
            if (contaminated[i] && c > maxCount || (c == maxCount && i < ans)) {
                ans = i;
                maxCount = c;
            }
        }
        return ans;
    }
}