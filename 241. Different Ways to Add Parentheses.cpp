class Solution {
public:
    vector<int> diffWaysToCompute(string input) {
        vector<int> res;
        for (int i = 0; i < input.size(); i++) {
            char ch = input[i];
            if (ch == '+' || ch == '-' || ch == '*') {
                vector<int> lv = diffWaysToCompute(input.substr(0, i));
                vector<int> rv = diffWaysToCompute(input.substr(i+1));
                
                for (auto &l : lv) {
                    for (auto &r : rv) {
                        if (ch == '+')
                            res.push_back(l+r);
                        else if (ch == '-')
                            res.push_back(l-r);
                        else if( ch == '*')
                            res.push_back(l*r);
                    }
                }
            }
        }
        
        if (res.empty())
            res.push_back(atoi(input.c_str()));
        
        return res;
    }
};