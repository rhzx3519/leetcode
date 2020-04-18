/*
1. 这道题给了我们一个数组，问我们这个数组能不能分成两个非空子集合，
使得两个子集合的元素之和相同。那么我们想，原数组所有数字和一定是偶数，
不然根本无法拆成两个和相同的子集合，那么我们只需要算出原数组的数字之和，
然后除以2，就是我们的target，那么问题就转换为能不能找到一个非空子集合，
使得其数字之和为target。开始我想的是遍历所有子集合，算和，但是这种方法无法通过OJ的大数据集合。
于是乎，动态规划DP就是我们的不二之选。我们定义一个一维的dp数组，其中dp[i]表示数字i是否是原数组的任意个子集合之和，
那么我们我们最后只需要返回dp[target]就行了。我们初始化dp[0]为true，由于题目中限制了所有数字为正数，
那么我们就不用担心会出现和为0或者负数的情况。那么关键问题就是要找出递归公式了，我们需要遍历原数组中的数字，
对于遍历到的每个数字nums[i]，我们需要更新我们的dp数组，要更新[nums[i], target]之间的值，那么对于这个区间中的任意一个数字j，
如果dp[j - nums[j]]为true的话，那么dp[j]就一定为true，于是地推公式如下：

dp[j] = dp[j] || dp[j - nums[i]]         (nums[i] <= j <= target)

2. 这道题还可以用bitset来做，感觉也十分的巧妙，bisets的大小设为5001，为啥呢，
因为题目中说了数组的长度和每个数字的大小都不会超过100，那么最大的和为10000，
那么一半就是5000，前面再加上个0，就是5001了。我们初始化把最低位赋值为1，然后我们算出数组之和，
然后我们遍历数字，对于遍历到的数字num，我们把bits向左平移num位，然后再或上原来的bits，
这样所有的可能出现的和位置上都为1。举个例子来说吧，比如对于数组[2,3]来说，初始化bits为1，
然后对于数字2，bits变为101，我们可以看出来bits[2]标记为了1，然后遍历到3，bits变为了101101，
我们看到bits[5],bits[3],bits[2]都分别为1了，正好代表了可能的和2，3，5，这样我们遍历玩整个数组后，
去看bits[sum >> 1]是否为1即可，参见代码如下：
*/

class Solution {
public:
    bool canPartition(vector<int>& nums) {
        int sum = accumulate(nums.begin(), nums.end(), 0);
        if (sum&1) return false;
        int n = nums.size();
        int target = sum>>1;
        vector<bool> dp(target+1, false);
        dp[0] = true;
        for (int i = 0; i < n; i++)
        {
            for (int j = target; j >= nums[i]; j--)
            {
                dp[j] = dp[j] || dp[j - nums[i]];
            }
        }
        
        return dp.back();
    }

    bool canPartition(vector<int>& nums)
    {
    	bitset<5001> bs(1);
    	int sum = accumulate(nums.begin(), nums.end(), 0);
    	for (auto &num : nums)
    	{
    		bs |= bs<<num;
    	}
    	return (!(sum&1)) && (bs[sum>>1]);
    }
<<<<<<< HEAD

    bool canPartition(vector<int>& nums) {
        sort(nums.begin(), nums.end());
        int sum = 0;
        for (int i: nums)
            sum += i;
        if (sum&1) return false;

        vector<int> dp(sum+1, 0);
        dp[0] = 1;
        for (int n : nums) {
            for (int i = sum>>1; i>=0; i--) {
                if (dp[i])
                    dp[i+n] = 1;
            }
        }

        return dp[sum>>1]==1;
    }    
=======
>>>>>>> 837bde34d06bb470ffcea088785f591f742870d7
};