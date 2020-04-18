class Solution {
public:
    int change(int amount, vector<int>& coins) {
        int n = coins.size();
        vector<int> dp(amount+1, 0);
        dp[0] = 1;
        for (int i = 0; i < n; i++) {
            for (int j = 1; j < amount+1; j++) {
                // 在上一钟零钱状态的基础上增大
                // 例如对于总额5，当只有面额为1的零钱时，只有一种可能 5x1
                // 当加了面额为2的零钱时，除了原来的那一种可能外
                // 还加上了组合了两块钱的情况，而总额为5是在总额为3的基础上加上两块钱来的
                // 所以就加上此时总额为3的所有组合情况                
                if (j >= coins[i])
                    dp[j] = dp[j] + dp[j-coins[i]];
            }
        }

        return dp[amount];
    }
};