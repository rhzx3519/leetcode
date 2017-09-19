class Solution {
public:
    typedef long long LL;
    int reverse(int x) {
        LL xx = x;
        LL res = 0;
        while (xx != 0)
        {
            res = res * 10 + xx%10;
            if (res > INT_MAX || res < INT_MIN) return 0;
            xx /= 10;
        }

        return res;
    }
};