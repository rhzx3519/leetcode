class Solution {
public:
    int countNumbersWithUniqueDigits(int n) {
        if (n == 0) return 1;
        int res = 10;
        int a = 9, b = 9;
        while (--n > 0 && b > 0) {
            a *= b;
            res += a;
            b--;
        }
        
        return res;
    }
};