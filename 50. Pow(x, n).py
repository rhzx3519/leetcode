class Solution {
public:
    double myPow(double x, int n) {
        if (x == 0.0 || x == 1.0)
            return x;
        if (n == 0)
            return 1.0;
        double y;
        if (n == -n) { // 用于判断n是否溢出
            y = myPow(x, n/2);
            return y*y;
        }
        
        if (n < 0)
            return 1.0/myPow(x, -n);
        double res = 1.0f;
        y = myPow(x, n/2);
        res = y*y;
        if (n&1)
            res *= x;
        return res;
    }
};