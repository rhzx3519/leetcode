package Q2642;

import java.util.ArrayList;
import java.util.Comparator;
import java.util.List;
import java.util.PriorityQueue;

/**
 * https://leetcode.cn/problems/design-graph-with-shortest-path-calculator/description/?envType=daily-question&envId=2024-03-26
 */
public class Q2642 {
    private List<List<int[]>> adj;

    public Q2642(int n, int[][] edges) {
        adj = new ArrayList<>(n);
        for (int i = 0; i < n; i++) {
            adj.add(new ArrayList<>());
        }

        for (int[] edge : edges) {
            int a = edge[0];
            int b = edge[1];
            int c = edge[2];
            adj.get(a).add(new int[]{b, c});
        }
    }

    public void addEdge(int[] edge) {
        int a = edge[0];
        int b = edge[1];
        int c = edge[2];
        adj.get(a).add(new int[]{b, c});
    }

    public int shortestPath(int node1, int node2) {
        int[] dis = djisktra(node1, adj.size());
        if (dis[node2] == Integer.MAX_VALUE>>1) {
            return -1;
        }
        return dis[node2];
    }

    private int[] djisktra(int from, int n) {
        int[] dis = new int[n];
        for (int i = 0; i < dis.length; i++) {
            dis[i] = Integer.MAX_VALUE>>1;
        }
        dis[from] = 0;
        PriorityQueue<int[]> pq = new PriorityQueue<>(Comparator.comparingInt(a -> a[1]));
        pq.offer(new int[]{from, 0});
        while (!pq.isEmpty()) {
            int[] t = pq.poll();
            if (dis[t[0]] < t[1]) {
                continue;
            }
            for (int[] edge : adj.get(t[0])) {
                int ni = edge[0];
                int cost = edge[1];
                if (dis[ni] > dis[t[0]] + cost) {
                    dis[ni] = dis[t[0]] + cost;
                    pq.offer(new int[]{ni, dis[ni]});
                }
            }
        }
        return dis;
    }
}
