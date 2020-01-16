class Solution {
public:
    int numberOfArithmeticSlices(vector<int>& A) {
        int n = A.size();
        if (n <= 2) return 0;
        int add = 0;
        int res = 0;
        for (int i = 2; i < n; i++) {
            if (A[i]-A[i-1] == A[i-1]-A[i-2])
                add += 1;
            else
                add = 0;
            res += add;
        }
        return res;
    }
};

/// https://blog.csdn.net/camellhf/article/details/52824234