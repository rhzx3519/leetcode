class Solution {
public:
    bool isUgly(int num) {
        int &a = num;
        if (a <= 0) return false;
        while (a % 2 == 0) a /= 2;
        while (a % 3 == 0) a /= 3;
        while (a % 5 == 0) a /= 5;
        return (a == 1);
    }
};	