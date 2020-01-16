class Solution {
public:
    int longestValidParentheses(string s) {
        int res = 0;
        int i = 0, ls = s.size();
        unordered_map<int, int> um;
        um[0] = -1;
        int cnt = 0;
        while (i < ls)
        {
            if (s[i] == '(')
            {
                cnt++;
                if (um.count(cnt) == 0)
                    um[cnt] = i;
            }
            else
            {
                if (um.count(cnt) != 0)
                    um.erase(cnt);
                cnt--;
                if (um.count(cnt) == 0)
                    um[cnt] = i;
                else
                    res = max(res, i - um[cnt]);
            }
            i++;
        }
        
        return res;
    }
};