class Solution {
public:
    int nthSuperUglyNumber(int n, vector<int>& primes) {
<<<<<<< HEAD
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
=======
        int k = primes.size();
        vector<int> res(1, 1);
        vector<int> term(k);
        vector<int> idx(k, 0);
        
        int i, min_val;
        while (res.size() < n) {
            min_val = INT_MAX;
            for (i = 0; i < k; i++) {
                term[i] = primes[i]*res[idx[i]];
                min_val = min(min_val, term[i]);
            }
            res.push_back(min_val);
            for (i = 0; i < k; i++) {
                if (min_val == term[i])
                    idx[i]++;
            }
        }
        
        return res.back();
>>>>>>> 837bde34d06bb470ffcea088785f591f742870d7
    }
};