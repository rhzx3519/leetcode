class Solution {
public:
    int titleToNumber(string s) {
        int res = 0, ls = s.size();
        for (int i = 0; i < ls; i++) {
            res += pow(26, ls - i - 1) * (s[i]-'A' + 1);
        }
        
        return res;
    }
};