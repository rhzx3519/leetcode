class Solution {
public:
    bool pyramidTransition(string bottom, vector<string>& allowed) {
        unordered_set<string> st(allowed.begin(), allowed.end());
        return dfs(bottom, "", st);
    }

    bool dfs(string bottom, string nextStr, unordered_set<string>& allowed) {
        if (bottom.size() == 1) return true;

        if (bottom.size() - nextStr.size() > 1) {
            string str = bottom.substr(nextStr.size(), 2);
            for (char c = 'A'; c < 'H'; c++) {
                auto it = allowed.find(str + c);
                if (it!=allowed.end() && dfs(bottom, nextStr + c, allowed)) {
                    return true;
                }
            }
            return false;
        } else {
            return dfs(nextStr, "", allowed);
        }
    }

};

// 复杂一点的dfs， 通过判断上一层字符串的长度来做搜索分支