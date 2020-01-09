class Solution {
public:
    vector<string> addOperators(string num, int target) {
        vector<string> res;
        string exp;
        dfs(0, exp, num, target, 0, 0, res);
        return res;
    }
    
private:
    void dfs(int idx, string &exp, string &num, int target, 
            int64_t prev, int64_t sum, vector<string> &res) {
        int ld = num.size();
        if (idx >= ld) {
            if (sum == target)
                res.push_back(exp);
            return;
        }
        
        int n = num[idx] == '0' ? 1 : ld - idx;
        int64_t cur = 0;
        int le = exp.size();
        string s;
        
        int i;
        for (i = 0; i < n; i++) {
            s.push_back(num[idx+i]);
            cur = cur * 10 + (num[idx+i] - '0');
            
            if (idx == 0) {
                exp += s;
                dfs(idx+i+1, exp, num, target, cur, sum + cur, res);
                exp.resize(le);
            } else {
                exp.push_back('+');
                exp += s;
                dfs(idx+i+1, exp, num, target, cur, sum + cur, res);
                exp.resize(le);
                
                exp.push_back('-');
                exp += s;
                dfs(idx+i+1, exp, num, target, -cur, sum-cur, res);
                exp.resize(le);
                
                exp.push_back('*');
                exp += s;
                dfs(idx+i+1, exp, num, target, prev*cur, sum - prev + prev*cur, res);
                exp.resize(le);
            }
        }
    }
};