class Solution {
public:
    int arrangeCoins(long n) {
        return std::floor(sqrt(2*n+0.25) - 0.5);
    }
};