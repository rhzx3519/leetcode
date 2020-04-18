class Solution {
public:
    vector<int> nextGreaterElements(vector<int>& nums) {
        int n = nums.size();
        vector<int> res(n, -1);
        stack<int> st;

        for (int i = 0; i < 2*n; i++) {
            int idx = i%n;
            while (!st.empty() && nums[st.top()] < nums[idx]) {
                res[st.top()] = nums[idx];
                st.pop();
            }
            if (i < n) st.push(i);
        }
        return res;
    }
};

// 使用栈记录i之前的数字

