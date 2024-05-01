package Q2462;

import java.util.Arrays;
import java.util.PriorityQueue;
import java.util.Queue;

public class Q2462 {
}

/**
 * https://leetcode.cn/problems/total-cost-to-hire-k-workers/?envType=daily-question&envId=2024-05-01
 */

class Worker implements Comparable<Worker> {
    public int value;
    public int idx;

    public Worker(int value, int idx) {
        this.value = value;
        this.idx = idx;
    }

    @Override
    public int compareTo(Worker o) {
        if (this.value != o.value) {
            return Integer.compare(this.value, o.value);
        }
        return Integer.compare(this.idx, o.idx);
    }
}

/**
 * https://leetcode.cn/problems/total-cost-to-hire-k-workers/?envType=daily-question&envId=2024-05-01
 */
class Solution {
    public long totalCost(int[] costs, int k, int candidates) {
        int n = costs.length;
        long tot = 0;
        if (n <= candidates*2) {
            Arrays.sort(costs);
            for (int i = 0; i < k; i++) {
                tot += costs[i];
            }
            return tot;
        }

        Queue<Integer> leftPQ = new PriorityQueue<Integer>();
        Queue<Integer> rightPQ = new PriorityQueue<Integer>();
        for (int i = 0; i < candidates; i++) {
            leftPQ.offer(costs[i]);
            rightPQ.offer(costs[n-1-i]);
        }
        int i = candidates;
        int j = n - 1 - candidates;
        for (; k != 0; k--) {
            int v = 0;
            if (!leftPQ.isEmpty() && !rightPQ.isEmpty()) {
                if (leftPQ.peek() <= rightPQ.peek()) {
                    tot += leftPQ.poll();
                    if (i <= j) {
                        leftPQ.offer(costs[i]);
                        i++;
                    }
                } else {
                    tot += rightPQ.poll();
                    if (i <= j) {
                        rightPQ.offer(costs[j]);
                        j--;
                    }
                }
            } else if (!leftPQ.isEmpty()) {
                tot += leftPQ.poll();
            } else {
                tot += rightPQ.poll();
            }

        }
        return tot;
    }
}



















