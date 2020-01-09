class Solution {
public:
    bool find132pattern(vector<int>& nums) {
        int min_val = INT_MIN, n = nums.size();
        stack<int> st;
        for (int i = n-1; i >= 0; i--) {
            if (nums[i] < min_val) return true;
            while (!st.empty() && st.top() < nums[i]) {
                min_val = st.top();
                st.pop();
            }
            st.push(nums[i]);
        }
        return false;
    }
};