class Solution {
public:
    int lengthOfLIS(vector<int>& nums) {
<<<<<<< HEAD
        vector<int> dp;
        int n = nums.size();
        for (int i = 0; i < nums.size(); i++) {
            if (dp.empty() || dp.back() < nums[i]) {
                dp.push_back(nums[i]);
                continue;
            }
            int j = lower_bound(dp.begin(), dp.end(), nums[i]) -dp.begin();
            dp[j] = nums[i];
        }
        return dp.size();
    }
};

// dp存储了当前最长上升子序列，遍历数组的时候，
// 将元素插入到dp中合适的位置，遍历完数组之后，dp即是当前最长上升子序列
=======
        vector<int> seq;
        int n = nums.size();
        if (n < 2) {
            return n;
        }
        seq.push_back(nums[0]);
        
        int i;
        int j;
        for (i = 1; i < n; ++i) {
            if (nums[i] > seq.back()) {
                seq.push_back(nums[i]);
                continue;
            }
            j = lower_bound(seq.begin(), seq.end(), nums[i]) - seq.begin();
            seq[j] = nums[i];
        }
        
        int res = seq.size();
        seq.clear();
        return res;
    }
};
>>>>>>> 837bde34d06bb470ffcea088785f591f742870d7
