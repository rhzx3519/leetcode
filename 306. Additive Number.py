class Solution {
public:
    bool isAdditiveNumber(string num) {
        vector<long> path;
        return dfs(num, path);
    }
    
    bool dfs(string num, vector<long> path) {
        if (path.size() >= 3)  {
            int n = path.size();
            if (path[n-1] != path[n-2] + path[n-3])
                return false;
        }
        if (num.empty() && path.size()>=3)
            return true;
        
        for (int i = 0; i < num.size(); i++) {
            string s = num.substr(0, i+1);
            if (s[0] == '0' && s.size()>1)
                continue;
            vector<long> v(path);
            v.push_back(atol(s.c_str()));
            if (dfs(num.substr(i+1), v))
                return true;
        }
        return false;
    }
};

// path中存储了当前所有的切片，每当path的size > 3，判断末尾三位是否符合累加数