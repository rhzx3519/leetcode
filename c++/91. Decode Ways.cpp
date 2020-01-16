class Solution {
public:
    int numDecodings(string s) {
        int ls = s.size();
        if (ls == 0) return 0;
        int LOW = 1, HIGH = 26;
        
        int f1 = 0, f2 = 1, f3 = 0, d;
    
        for (int i = 0; i < ls; i++) {
            f3 = 0;
            d = s[i] - '0';
            if (d >= LOW && d <= HIGH)
                f3 += f2;
            if (i > 0 && (s[i - 1]) != '0') {
                d += (s[i-1] - '0')*10;
                if (d >= LOW && d <= HIGH)
                    f3 += f1;
            }
            
            f1 = f2;
            f2 = f3;
        }
        
        return f3;
    }
};

/*
作者：沙粒牛
链接：https://www.nowcoder.com/discuss/43937?toCommentId=827018
来源：牛客网

我们来做一个简单的密码破译游戏。破译的规则很简单，将数字转换为字母，1转化为a，2转化为b，依此类推，26转化为z。现在输入的密码是一串数字，输出的破译结果是该数字串通过转换规则所能产生的所有字符串。

输入：
1
12
123
输出：

a
ab l
abc aw lc

这题是leetcode 91.Decode Ways原题的改进版，原题只要输出编码数量，此题需要输出所有编码，代码如下：

作者：沙粒牛
链接：https://www.nowcoder.com/discuss/43937?toCommentId=827018
来源：牛客网
*/
// 1. dp
const int LOW = 1, HIGH = 26;
vector<string> dp;
void decode(string s)
{
    vector res;
    int ls = s.size();
    if (ls == 0) return;
    for (int i = 0; i < ls; i++)
    {
        char ch1 = '0', ch2 = '0';
        int d = s[i] - '0';
        if (d >= LOW && d <= HIGH)
        {
            ch1 = 'a' + d - 1;
        }
        if (i > 0 && s[i-1] != '0')
        {
            d += (s[i-1] - '0')*10;
            if (d >= LOW && d <= HIGH)
            {
                ch2 = 'a' + d - 1;
            }
        }
        if (ch1 != '0')
        {
            if (i > 0)
            {
                vector t = dp[i-1];
                for (auto &str : t)
                {
                    string s1 = str + ch1;
                    dp[i].push_back(s1);
                }
            }
            else
            {
                string s1; s1.push_back(ch1);
                dp[i].push_back(s1);
            }
        }
        if (ch2 != '0')
        {
            if (i > 1)
            {
                vector t = dp[i-2];
                for (auto &str : t)
                {
                    string s1 = str + ch2;
                    dp[i].push_back(s1);
                }
            }
            else
            {
                string s1; s1.push_back(ch2);
                dp[i].push_back(s1);
            }
        }
    }
}
int main()
{
    string s;
    while (cin >> s)
    {
        dp.clear();
        dp.resize(s.size() + 1);
        decode(s);
        vector res = dp[s.size() - 1];
        for (int i = 0; i < res.size(); i++)
            i == res.size() - 1 ? cout << res[i] << endl : cout << res[i] << " ";
    }
    return 0;
}

// 2. dfs
#include <stdio.h>
#include <vector>
#include <string>
#include <iostream>
using namespace std;

void dfs(int idx, string& s, vector<string>& res, string& str)
{
    if (idx >= s.size())
    {
        res.push_back(str); return;
    }
    int d = s[idx] - '0';
    char ch = 'a' + d - 1;
    str.push_back(ch);
    dfs(idx+1, s, res, str);
    str.pop_back();
    if (idx >= s.size() - 1) return;
    if (s[idx] <= '2' && s[idx + 1] <= '6' || s[idx + 1] <= '1')
    {
        ch = ((s[idx] - '0') * 10 + s[idx + 1] - '0') + 'a' - 1;
        str.push_back(ch);
        dfs(idx + 2, s, res, str);
        str.pop_back();
    }
}

int main()
{
    string s;
    while (cin >> s)
    {
        vector<string> res;
        string str;
        dfs(0, s, res, str);
        for (int i = 0; i < res.size(); i++)
            i == res.size() - 1 ? cout << res[i] << endl : cout << res[i] << " ";
    }
    return 0;
}