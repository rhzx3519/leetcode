class Solution {
public:
	bool canPartitionKSubsets(vector<int>& nums, int k) {
		int sum = accumulate(nums.begin(), nums.end(), 0);
		int n = nums.size();
		if (n < k || sum < k || sum % k != 0) return false;
		vector<int> vis(n, 0);

		return dfs(0, k, sum / k, sum / k, nums, vis);
	}

	bool dfs(int idx, int k, int len, int cur_len, vector<int> &nums, vector<int> &vis) {
		if (k == 0) return true;
		if (cur_len == 0)
			return dfs(0, k - 1, len, len, nums, vis);

		for (int i = idx; i < nums.size(); i++) {
			if (vis[i] == 1)
				continue;
			vis[i] = 1;
			if (nums[i] <= cur_len && dfs(i + 1, k, len, cur_len - nums[i], nums, vis))
				return true;
			vis[i] = 0;
		}
		return false;
	}
};