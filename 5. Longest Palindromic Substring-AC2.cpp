class Solution {
public:
    string longestPalindrome(string s) {
        int n = s.size();
        if (n == 1) return s;
        int start = 0, max_len = 1;
        
        for (int i = 0; i < n; i++)
        {
            int j = i - 1, k = i + 1;
            while (j >= 0 && k <= n - 1 && s[j] == s[k])
            {
                if (k - j + 1 > max_len)
                {
                    start = j;
                    max_len = k - j + 1;
                }
                j--;
                k++;
            }
        }
        
        for (int i = 0; i < n; i++)
        {
            int j = i, k = i + 1;
            while (j >= 0 && k <= n -1 && s[j] == s[k])
            {
                if (k - j + 1 > max_len)
                {
                    start = j;
                    max_len = k - j + 1;
                }
                j--;
                k++;
            }
        }
        
        return s.substr(start, max_len);
    }
};