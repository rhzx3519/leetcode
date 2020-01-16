class Solution {
public:
    /**
     * @param ratings Children's ratings
     * @return the minimum candies you must give
     */
    int candy(vector<int>& ratings) {
        // Write your code here
        int ls = ratings.size();
        if (ls == 0) return 0;
        if (ls == 1) return 1;
        vector<int> candy(ls, 1);
        int res = 0;
        for (int i = 0; i < ls - 1; i ++) {
            if (ratings[i+1] > ratings[i])
                candy[i+1] = candy[i] + 1;
        }
        for (int i = ls-1; i > 0; i--) {
            if (ratings[i-1] > ratings[i])
                candy[i-1] = max(candy[i-1], candy[i]+1);
        }
        for (int i = 0; i < ls; i ++)
            res += candy[i];
        
        return res;
    }
};