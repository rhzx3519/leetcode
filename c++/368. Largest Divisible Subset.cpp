class Solution {
public:
    vector<int> largestDivisibleSubset(vector<int>& nums) {
        int n = nums.size();
        if (n < 2) return nums;
        vector<int> dp(n, 1);
        vector<int> parent(n, -1); // 辅助数组，记录最大子集数字的索引
        sort(nums.begin(), nums.end()); // 升序排序

        int max_len = 0, max_idx = 0;
        for (int i = 0; i < n; i++) {
            for (int j = i-1; j >= 0; j--) { // 往前找
                if (nums[i]%nums[j]==0) {
                    if (dp[j] + 1 > dp[i]) { 
                        dp[i] = dp[j] + 1;
                        parent[i] = j;      // 更新最大子集长度、用辅助数组记录前一个数的索引

                        if (dp[i] > max_len) {
                            max_len = dp[i];
                            max_idx = i;
                        }
                    }
                }
            }
        }

        vector<int> res;
        while (max_idx != -1) {
            res.push_back(nums[max_idx]);
            max_idx = parent[max_idx];
        }
        
        return res;
    }
};

	// 这道题给了我们一个数组，让我们求这样一个子集合，集合中的任意两个数相互取余均为0，而且提示中说明了要使用DP来解。那么我们考虑，较小数对较大数取余一定为0，那么问题就变成了看较大数能不能整除这个较小数。那么如果数组是无序的，处理起来就比较麻烦，所以我们首先可以先给数组排序，这样我们每次就只要看后面的数字能否整除前面的数字。定义一个动态数组dp，其中dp[i]表示到数字nums[i]位置最大可整除的子集合的长度，还需要一个一维数组parent，来保存上一个能整除的数字的位置，两个整型变量mx和mx_idx分别表示最大子集合的长度和起始数字的位置，我们可以从后往前遍历数组，对于某个数字再遍历到末尾，在这个过程中，如果nums[j]能整除nums[i], 且dp[i] < dp[j] + 1的话，更新dp[i]和parent[i]，如果dp[i]大于mx了，还要更新mx和mx_idx，最后循环结束后，我们来填res数字，根据parent数组来找到每一个数字
