class Solution {
public:
    int climbStairs(int n) {
        if (n == 1) return 1;
        int f1 = 1, f2 = 2, f3;
        int i = 3;
        while (i <= n)
        {
            f3 = f1 + f2;
            f1 = f2;
            f2 = f3;
            i++;
        }
        return f2;
    }
};