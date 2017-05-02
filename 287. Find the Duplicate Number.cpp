class Solution {
public:
    int findDuplicate(vector<int>& nums) {
        int n = nums.size();
        if (n <= 1) return -1;
        int slow, fast;
        slow = fast = nums[0];
        while (fast != slow) {
            slow = nums[slow];
            fast = nums[nums[fast]];
        }
        
        fast = 0;
        while (fast != slow) {
            slow = nums[slow];
            fast = nums[fast];
        }
        
        return fast;
    }
};