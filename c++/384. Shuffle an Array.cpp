class Solution {
public:
    Solution(vector<int> nums) : nums(nums) {
        
    }
    
    /** Resets the array to its original configuration and return it. */
    vector<int> reset() {
        return nums;
    }
    
    /** Returns a random shuffling of the array. */
    vector<int> shuffle() {
        vector<int> res(nums);
        int n = res.size();
        for (int i = 0; i < n; i++) {
            int j = rand()%(n-i);
            swap(res[i+j], res[i]);
        }
        return res;
    }

private:
    vector<int> nums;
};

/**
 * Your Solution object will be instantiated and called as such:
 * Solution obj = new Solution(nums);
 * vector<int> param_1 = obj.reset();
 * vector<int> param_2 = obj.shuffle();
 */