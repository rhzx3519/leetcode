class Solution {
public:
    int longestCommonSubsequence(string text1, string text2) {
        int n1 = text1.size(), n2 = text2.size();
        vector<vector<int>> dp(n1+1, vector<int>(n2+1, 0));
        for (int i = 1; i <= n1; i++) {
            for (int j = 1; j <= n2; j++) {
                if (text1[i-1] == text2[j-1]) 
                    dp[i][j] = dp[i-1][j-1] + 1;
                else
                    dp[i][j] = max(dp[i-1][j], dp[i][j-1]);
            }
        }

        return dp[n1][n2];
    }
};

//    设A=“a0，a1，…，am”，B=“b0，b1，…，bn”，且Z=“z0，z1，…，zk”为它们的最长公共子序列。不难证明有以下性质：
//        如果am=bn，则zk=am=bn，且“z0，z1，…，z(k-1)”是“a0，a1，…，a(m-1)”和“b0，b1，…，b(n-1)”的一个最长公共子序列；
//        如果am!=bn，则若zk!=am，蕴涵“z0，z1，…，zk”是“a0，a1，…，a(m-1)”和“b0，b1，…，bn”的一个最长公共子序列；
//        如果am!=bn，则若zk!=bn，蕴涵“z0，z1，…，zk”是“a0，a1，…，am”和“b0，b1，…，b(n-1)”的一个最长公共子序列。
// ————————————————
// 版权声明：本文为CSDN博主「Running07」的原创文章，遵循 CC 4.0 BY-SA 版权协议，转载请附上原文出处链接及本声明。
// 原文链接：https://blog.csdn.net/hrn1216/article/details/51534607