package Q2192;

import java.util.*;
import java.util.stream.Collectors;

public class Q2192 {
    private  List<List<Integer>> adj = new ArrayList<>();

    private void dfs(int i, boolean[] visited) {
        for (int ni : adj.get(i)) {
            if (!visited[ni]) {
                visited[ni] = true;
                dfs(ni, visited);
            }
        }
    }

    public List<List<Integer>> getAncestors(int n, int[][] edges) {
        List<List<Integer>> ans = new ArrayList<>();

        for (int i = 0; i < n; i++) {
            adj.add(new ArrayList<>());
            ans.add(new ArrayList<>());
        }
        int[] ind = new int[n];
        for (int[] edge : edges) {
            adj.get(edge[1]).add(edge[0]);
        }


        for (int i = 0; i < n; i++) {
            boolean[] visited = new boolean[n];
            dfs(i, visited);
            for (int j = 0; j < n; j++) {
                if (visited[j]) {
                    ans.get(i).add(j);
                }
            }
        }

        return ans;
    }

}
