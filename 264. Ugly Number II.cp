class Solution {
public:
    int nthUglyNumber(int n) {
        vector<long long> v = {1LL};
        int i = 0, j = 0, k = 0;
        int cur = 1, cnt = 1;
        while (cnt < n) {
            cur = min(v[i]*2, min(v[j]*3, v[k]*5));
            v.push_back(cur);
            while (v[i]*2 <= cur) i++;
            while (v[j]*3 <= cur) j++;
            while (v[k]*5 <= cur) k++;
            cnt++;
        }
        
        return v.back();
    }
};