/**
python 前缀和思想 举个栗子： [ -1，-1，-1，1，-1，1，-1，1，-1，-1，-1，-1 ] 在扫描完到第四个元素时，前缀和为-2 且未记录过，则将值-2（作为字典的键）和下标3（作为字典的值）记录起来。当扫描到 [ -1，-1，-1，1，-1，-1，1，1，-1，-1，-1，-1 ] ， 此时得到的前缀和为-2，且知道标记中有记录过-2，则说明此刻下标到之前记录的下标的这段数组总和为0 [ -1，-1，-1，1，-1，-1，1，1，-1，-1，-1，-1 ] 。
*/
class Solution {
public:
    bool checkSubarraySum(vector<int>& nums, int k) {
        if (nums.size() < 2) return false;

        unordered_map<int, int> um;
        um[0] = -1;
        int sum = 0;
        for (int i = 0; i < nums.size(); i++) {
            sum += nums[i];
            if (k != 0)
                sum = sum % k;
            if (um.count(sum)==0)
                um[sum] = i;
            else {
                int last = um[sum];
                if (i - last > 1)
                    return true;
            }
        }
        return false;
    }
};