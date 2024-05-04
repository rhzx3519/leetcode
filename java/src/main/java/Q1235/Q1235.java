package Q1235;

import java.util.Arrays;

public class Q1235 {
}


/**
 * https://leetcode.cn/problems/maximum-profit-in-job-scheduling/?envType=daily-question&envId=2024-05-04
 */
class Job {
    public int start;
    public int end;
    public int profit;
    public Job(int start, int end, int profit) {
        this.start = start;
        this.end = end;
        this.profit = profit;
    }
}
class Solution {
    public int jobScheduling(int[] startTime, int[] endTime, int[] profit) {
        Job[] jobs = new Job[startTime.length];
        for (int i = 0; i < startTime.length; i++) {
            jobs[i] = new Job(startTime[i], endTime[i], profit[i]);
        }
        Arrays.sort(jobs, (a, b) -> a.end - b.end);
        int[] f = new int[jobs.length + 1];
        for (int i = 0; i < jobs.length; i++) {
            int j = search(jobs, i, jobs[i].start);
            f[i + 1] = Math.max(f[i], f[j+1] + jobs[i].profit);
        }
        return f[jobs.length];
    }

    private int search(Job[] jobs, int right, int upper) {
        int left = -1;
        while (left + 1 < right) {
           int mid = left + (right - left) / 2;
           if (jobs[mid].end <= upper) {
               left = mid;
           } else {
               right = mid;
           }
        }
        return left;
    }
}