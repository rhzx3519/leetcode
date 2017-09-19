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