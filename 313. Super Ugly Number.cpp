class Solution {
public:
    int nthSuperUglyNumber(int n, vector<int>& primes) {
		int k = primes.size();
		vector<int> nums(k, 0);
		vector<int> v = { 1 };
		int i = 1;
		int res = 1;
		while (i < n) {
			int min_val = INT_MAX;
			for (int j = 0; j < k; j++) {
				min_val = min(min_val, v[nums[j]] * primes[j]);
			}
			v.push_back(min_val);
			// 更新质因数索引
			for (int j = 0; j < k; j++) {
				while (v[nums[j]] * primes[j] <= min_val)
					nums[j]++;
			}

			i++;
		}

		return v.back();
    }
};