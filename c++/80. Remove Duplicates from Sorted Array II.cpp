class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        int i = 0, j = 0, k = 0;
        while (i < nums.size()) {
            nums[k++] = nums[i];
            j = i + 1;
            while (j < nums.size() && nums[i] == nums[j])
                j++;
            if (j - i >= 2)
                nums[k++] = nums[i];
            i = j;
        }
        return k;
    }
};

class Solution(object):
    def removeDuplicates(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        i = 0
        for num in nums:
            if i < 2 or num != nums[i-2]:
                nums[i] = num
                i += 1
        return i 