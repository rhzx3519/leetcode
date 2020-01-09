class Solution {
public:
    string removeKdigits(string num, int k) {
        string s;
        for (int i = 0; i < num.size(); i++) {
            char ch = num[i];
            while (k!=0 && !s.empty() && s.back()>ch) {
                s.pop_back();
                k--;
            }
            if (s.empty() && ch == '0')
                continue;
            s.push_back(ch);
        }

        while (!s.empty() && k!= 0) {
            s.pop_back();
            k--;
        }

        return s.empty() ? "0" : s;
    }
};

// 构造一个递增的序列，同时开头不能为0