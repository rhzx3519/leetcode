class Solution {
public:
    int rangeBitwiseAnd(int m, int n) {
        int x = n;
        while (x > m)
            x = x&(x-1);
        return x;
    }
};