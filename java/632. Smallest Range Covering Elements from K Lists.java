class Solution {
    public int[] smallestRange(List<List<Integer>> nums) {
        PriorityQueue<int[]> que = new PriorityQueue<>((n1, n2) -> n1[0] - n2[0]);
        int N = nums.size();
        int[] index = new int[N];
        int end = -100001;
        for(int i = 0; i < N; i++) {
            int val = nums.get(i).get(0);
            end = Math.max(end, val);
            que.offer(new int[]{val, i});
        }
        
        int start = que.peek()[0];
        int l = start, r = end;
        while (true) {
            int grp = que.poll()[1];
            if (index[grp] + 1 == nums.get(grp).size()) {
                break;
            }
            index[grp]++;
            int[] newGap = new int[]{nums.get(grp).get(index[grp]), grp};
            que.offer(newGap);
            if (newGap[0] > r) {
                r = newGap[0];
            }
            l = que.peek()[0];
            if (r - l < end - start) {
                end = r;
                start = l;
            }
        }
        return new int[]{start, end};
    }
}
// 优先队列 + 贪心