class Solution {
public:
    int lengthOfLastWord(string s) {
        int i = s.size() - 1;
        while (i >= 0 && s[i] == ' ') i--;
        int res = 0;
        while (i >= 0 && s[i] != ' ') {
            i--;
            res++;
        }
        return res;
    }
};