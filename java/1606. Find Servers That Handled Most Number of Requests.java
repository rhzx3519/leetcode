class Solution {
    private static class Entry {
        int start;
        int t;
        int serverId;

        public Entry(int start, int t, int serverId) {
            this.start = start;
            this.t = t;
            this.serverId = serverId;
        }
    }

    private int[] servers;

    public List<Integer> busiestServers(int k, int[] arrival, int[] load) {
        List<Integer> freeServers = new ArrayList<>();
        for (int i = 0; i < k; i++) {
            freeServers.add(i);
        }
        servers = new int[k];
        int[] req = new int[k];
        List<Integer> ans = new ArrayList<>();
        int maxVal = 0;

        PriorityQueue<Entry> que = new PriorityQueue<>(Comparator.comparingInt(e -> (e.start + e.t)));
        int n = arrival.length;
        for (int i = 0; i < n; i++) {
            int t = arrival[i];
            int l = load[i];
            while (!que.isEmpty() && (que.peek().start + que.peek().t) <= t) {
                Entry entry = que.poll();
                servers[entry.serverId] = 0;
                int idx = binarySearch(freeServers, entry.serverId);
                freeServers.add(idx, entry.serverId);
            }
            if (freeServers.isEmpty()) {
                continue;
            }
            int idx = binarySearch(freeServers, i%k) % freeServers.size();
            int serverId = freeServers.get(idx);
            servers[serverId] = 1;
            freeServers.remove(idx);
            que.offer(new Entry(t, l, serverId));
            req[serverId] += 1;
            if (req[serverId] > maxVal) {
                maxVal = req[serverId];
                ans.clear();
                ans.add(serverId);
            } else if (req[serverId] == maxVal) {
                ans.add(serverId);
            }
        }

        return ans;
    }

    private int binarySearch(List<Integer> st, int target) {
        int l = 0, r = st.size() - 1;
        while (l <= r) {
            int mid = l + (r - l)/2;
            int val = st.get(mid);
            if (val == target) {
                return mid;
            } else if (val > target) {
                r = mid - 1;
            } else {
                l = mid + 1;
            }
        }
        return l;
    }
}