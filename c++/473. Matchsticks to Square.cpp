class Solution {
public:
    bool makesquare(vector<int>& nums) {
        int n = nums.size();
        int sum = accumulate(nums.begin(), nums.end(), 0);
        int k = 4;
        if (sum < 4 || sum%4!=0) return false;
        vector<int> vis(n, 0);
        return dfs(0, k, sum/4, sum/4, nums, vis);
    }

    bool dfs(int idx, int k, int cur_len, int len, vector<int> &nums, vector<int> &vis) {
        if (k==0) return true;
        if (cur_len==0)
            return dfs(0, k-1, len, len, nums, vis);
        for (int i = idx; i < nums.size(); i++) {
            if (vis[i]==1) continue;

            vis[i] = 1;
            if (cur_len>=nums[i] && dfs(i+1, k, cur_len-nums[i], len, nums, vis))
                return true;
            vis[i] = 0;
        }
        return false;
    }
};