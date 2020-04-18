// 这里我用dp[i]表示以nums[i]为结尾的递推序列的长度，用cnt[i]表示以nums[i]为结尾的递推序列的个数，初始化都赋值为1，只要有数字，那么至少都是1。
// 然后我们遍历数组，对于每个遍历到的数字nums[i]，我们再遍历其之前的所有数字nums[j]，当nums[i]小于等于nums[j]时，不做任何处理，因为不是递增序列。
// 反之，则判断dp[i]和dp[j]的关系，如果dp[i]等于dp[j] + 1，说明nums[i]这个数字可以加在以nums[j]结尾的递增序列后面，并且以nums[j]结尾的递增序列
// 个数可以直接加到以nums[i]结尾的递增序列个数上。如果dp[i]小于dp[j] + 1，说明我们找到了一条长度更长的递增序列，那么我们此时奖dp[i]更新为dp[j]+1，
// 并且原本的递增序列都不能用了，直接用cnt[j]来代替。维护一个全局最长的子序列长度mx，每次都进行更新，到最后遍历一遍每个节点，如果长度等于mx,res+=cnt[i];
// ————————————————
// 版权声明：本文为CSDN博主「jarrysmith」的原创文章，遵循 CC 4.0 BY-SA 版权协议，转载请附上原文出处链接及本声明。
// 原文链接：https://blog.csdn.net/xuxuxuqian1/article/details/81071975

class Solution {
public:
    int findNumberOfLIS(vector<int>& nums) {
        int n = nums.size();
        vector<int> dp(n, 1); // 以i结尾的最长递增子序列长度
        vector<int> cnt(n, 1); // 以i结尾的最长递增子序列个数
        int max_val = 1;
        for (int i = 1; i < n; i++) {
            for (int j = 0; j < i; j++) {
                if (nums[j] < nums[i]) {
                    if (dp[j] + 1 > dp[i]) {
                        dp[i] = dp[j] + 1;
                        cnt[i] = cnt[j];
                    } else if (dp[j] + 1 == dp[i])
                        cnt[i] += cnt[j];
                }
            }
            max_val = max(dp[i], max_val);
        }

        int res = 0;
        for (int i = 0; i < n; i++) {
            if (dp[i] == max_val) {
                res += cnt[i];
            }
        }

        return res;
    }
};