class Solution {
public:
    int mySqrt(int x) {
        int l = 0, r = x;
        while (l <= r) {
            long mid = (long)((l+r)>>1);
            if (mid*mid == x)
                return (int)mid;
            else if (mid*mid > x)
                r = (int)mid - 1;
            else
                l = (int)mid + 1;
        }
        return r;
    }
};