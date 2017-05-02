class Solution {
public:
    int nthSuperUglyNumber(int n, vector<int>& primes) {
        int k = primes.size();
        vector<int> res(1, 1);
        vector<int> term(k);
        vector<int> idx(k, 0);
        
        int i, min_val;
        while (res.size() < n) {
            min_val = INT_MAX;
            for (i = 0; i < k; i++) {
                term[i] = primes[i]*res[idx[i]];
                min_val = min(min_val, term[i]);
            }
            res.push_back(min_val);
            for (i = 0; i < k; i++) {
                if (min_val == term[i])
                    idx[i]++;
            }
        }
        
        return res.back();
    }
};