class Solution {
public:
    vector<int> maxSlidingWindow(vector<int>& nums, int k) {
        vector<int> res;
        int n = nums.size();
        if (n == 0) return res;
        
        k = min(n, k);
        map<int, int> window;
        for (int i = 0; i < k; i++)
            ++window[nums[i]];
        res.push_back(window.rbegin()->first);
        
        for (int i = k; i < n; i++) {
            --window[nums[i-k]];
            if (window[nums[i-k]] == 0)
                window.erase(nums[i-k]);
            ++window[nums[i]];
            res.push_back(window.rbegin()->first);
        }
        window.clear();
        
        return res;
    }
};