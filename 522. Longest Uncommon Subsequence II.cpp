/**
 * 
 * 思路，先按长度降序，从最长的字符串i开始，检查i是不是至少相同长度的字符串j的子序列，如果不是，就找到了
 * 
 * 
*/
class Solution {
public:
    int findLUSlength(vector<string>& strs) {
        int n = strs.size();
        sort(strs.begin(), strs.end(), [](const string &s1, const string &s2){
            return s1.size() > s2.size();
        });
        for (int i = 0; i < strs.size(); i++) {
            bool f = true;
            for (int j = 0; j < n && strs[j].size() >= strs[i].size(); j++) {
                if (i != j && isSubsequence(strs[i], strs[j])) {
                    f = false;
                    break;
                }
            }
            if (f) return strs[i].size();
        }
        return -1;
    }

    bool isSubsequence(string &a, string &b) {
        auto ia = a.begin();
        auto ib = b.begin();
        while (ia!=a.end() && ib!=b.end()) {
            if (*ia == *ib) 
                ia++;
            ib++;
        }
        return ia==a.end();
    }
};