class Solution {
public:
    bool canJump(vector<int>& nums) {
        if (nums.size() < 2)
            return true;
        int i = 0, pos = 0, new_pos = 0;
        while (true) {
            new_pos = pos;
            while (i <= pos) {
                new_pos = max(new_pos, i+nums[i]);
                i++;
            }
            if (new_pos == pos) return false;
            pos = new_pos;
            if (pos >= nums.size() - 1)
                return true;
        }
        
        return false;
    }
};