typedef long long LL;

class Solution {
public:
    string fractionToDecimal(int numerator, int denominator) {
        string res;
        LL a = numerator, b = denominator, r;
        unordered_map<LL, int> um;
        
        int sign = 1;
        if ((a < 0 && b > 0) || (a > 0 && b < 0))
            sign = 0;
        a = a >= 0 ? a : -a;
        b = b >= 0 ? b : -b;
        if (a % b == 0)
            res = to_string(a/b);
        else
            res = to_string(a/b) + '.';
        
        r = a%b;
        int idx = res.length();
        while (r != 0 && um.find(r) == um.end()) {
            um[r] = idx;
            r *= 10;
            res += to_string(r/b);
            r %= b;
            idx++;
        }
        
        if (r != 0) {
            res.insert(um[r], 1, '(');
            res.insert(res.length(), 1, ')');
        }
        
        return sign ? res : '-'+res;
        
    }
};