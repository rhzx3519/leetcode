class Solution {
public:
	string getPermutation(int n, int k) {
		string s("1");
		vector<int> frac(n + 1, 1);
		for (int i = 1; i < n; i++) {
			s += '1' + i;
			frac[i] = frac[i - 1] * i;
		}

		k--;
		string res;
		for (int i = n - 1; i >= 0; i--) {
			int pos = k / frac[i];
			res += s[pos];
			s.erase(s.begin() + pos);
			k %= frac[i];
		}

		return res;
	}
};