class Solution {
public:
    int findDuplicate(vector<int>& nums) {
<<<<<<< HEAD
        int slow = nums[0], fast = nums[nums[0]];
        while (slow != fast) {
            slow = nums[slow];
            fast = nums[nums[fast]];
        }
        fast = 0;
        while (slow != fast) {
            slow = nums[slow];
            fast = nums[fast];
        }

        return fast;
    }
};

// 重复数字就是环的入口， 将题目转换为求链表环的入口
=======
        int n = nums.size();
        if (n <= 1) return -1;
        int slow = nums[0], fast = nums[nums[0]];
        
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
>>>>>>> 837bde34d06bb470ffcea088785f591f742870d7
