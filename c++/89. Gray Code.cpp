class Solution {
public:
    vector<int> grayCode(int n) {
        vector<int> res;
        vector<string> v = grayCode1(n);
        for(int i = 0; i < v.size(); i++) {
            res.push_back(to_decimal(v[i]));
        }
        
        return res;
    }
    
    vector<string> grayCode1(int n) {
        vector<string> res;
        if (n == 0) {
            res.push_back(string("0"));
            return res;
        }
        
        if (n == 1){
            res.push_back(string("0"));
            res.push_back(string("1"));
            return res;
        }
        
        vector<string> v = grayCode1(n-1);
        vector<string> v1, v2;
        for(int i = 0; i < v.size(); i++) {
            v1.push_back('0' + v[i]);
            v2.push_back('1' + v[i]);
        }
        reverse(v2.begin(), v2.end());
        copy(v1.begin(), v1.end(), back_inserter(res));
        copy(v2.begin(), v2.end(), back_inserter(res));
        return res;
        
    }
    
    string to_binary(int x, int n) {
        string res;
        if (x == 0) return string("0");
        int t;
        while (x != 0) {
            t = x % 2;
            x /= 2;
            res += '0' + t;
        }
        reverse(res.begin(), res.end());
        while (res.size() < n) 
            res = '0' + res;
        return res;
    }
    
    int to_decimal(string b) {
        int res = 0, n = b.size();
        for (int i = 0; i < n; i++) {
            res += pow(2, n - i - 1) * (b[i] - '0');
        }
        return res;
    }
};