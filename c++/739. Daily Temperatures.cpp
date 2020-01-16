class Solution {
public:
    vector<int> dailyTemperatures(vector<int>& T) {
        int n = T.size();
        stack<int> st;
        vector<int> res(n, 0);
        for (int i = 0; i < n; i++) {
            if (!st.empty()) {
                while (!st.empty() && T[st.top()] < T[i]) {
                    res[st.top()] = i - st.top();
                    st.pop();
                }
            }
            st.push(i);
        }
        return res;
    }
};

// 维护递减栈，后入栈的元素总比栈顶元素小。

// 比对当前元素与栈顶元素的大小
//      若当前元素 < 栈顶元素：入栈
//      若当前元素 > 栈顶元素：弹出栈顶元素，记录两者下标差值即为所求天数
// 这里用栈记录的是 T 的下标。