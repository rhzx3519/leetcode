class Solution {
public:
    int leastBricks(vector<vector<int>>& wall) {
        unordered_map<int, int> um;
        for (auto &w : wall) {
            int sum = 0;
            for (int i = 0; i < w.size()-1; i++) {
                sum += w[i];
                um[sum]++;
            }
        }

        int res = 0;
        for (auto it = um.begin(); it != um.end(); it++) {
           res = max(res, it->second);
        }

        return wall.size() - res;
    }
};

// C++ 统计边界数最大的位置即可