<<<<<<< HEAD
class Solution {
public:
    string removeKdigits(string num, int k) {
        string s;
        for (int i = 0; i < num.size(); i++) {
            char ch = num[i];
            while (k!=0 && !s.empty() && s.back()>ch) {
                s.pop_back();
                k--;
            }
            if (s.empty() && ch == '0')
                continue;
            s.push_back(ch);
        }

        while (!s.empty() && k!= 0) {
            s.pop_back();
            k--;
        }

        return s.empty() ? "0" : s;
    }
};

// 构造一个递增的序列，同时开头不能为0
=======
/*
	1. 如果第二位是0，则删除第一位数字
	2. 需找第一个下降的数字删除
*/
class Solution {
public:
	string removeKdigits(string num, int k) {
		while (1)
		{
			int n = num.size();
			if (n <= k || n == 0) return "0";
			if (k-- == 0) return num;
			if (num.at(1) == '0')
			{
				size_t pos = num.find_first_not_of('0', 1);
				if (pos != string::npos)
					num = num.substr(pos);
				else
					return "0";
			}
			else
			{
				int idx = 0;
				while (idx < n - 1 && num[idx] <= num[idx + 1]) idx++;
				num.erase(num.begin() + idx);
			}
		}

		return num;
	}
};
>>>>>>> 837bde34d06bb470ffcea088785f591f742870d7
