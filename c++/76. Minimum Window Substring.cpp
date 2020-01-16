class Solution {
public:
    string minWindow(string s, string t) {
        int data[255], now[255];
        memset(data, 0, sizeof(data));
        memset(now, 0, sizeof(now));
        for(int i = 0; i < t.size(); i++)
            data[t[i]]++;
        int i, j, left = 0, right = 0, min_len = 1<<29, num = 0;
        for (i = 0, j = 0; i < s.size(); i++) {
            if (num < t.size()) {
                if (now[s[i]] < data[s[i]]) num++;
                now[s[i]]++;
            }
            
            if (num == t.size()) {
                while (j <= i && now[s[j]] - 1 >= data[s[j]]) { // 去除不存在目标串中的多余的字符
                    now[s[j]]--;
                    j++;
                }
                if (min_len > i-j+1) {
                    min_len = i-j+1;
                    left = j;
                    right = i;
                }
                while (j <= i && num == t.size()) { // 去掉一位存在于目标串中的字符
                    now[s[j]]--;
                    if (now[s[j]] < data[s[j]]) num--;
                    j++;
                }
            }
        }
        
        return min_len == 1<<29 ? "" : s.substr(left, min_len);
    }
};