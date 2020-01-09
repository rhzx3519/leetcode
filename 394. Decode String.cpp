class Solution {
public:
	string decodeString(string s) {
		if (s.empty()) return s;
		int idx = 0;
		char ch = s[idx];
		string res;
		while (idx < s.size()) {
			if (s[idx] >= '0' && s[idx] <= '9')
				res.append(parse_repeat(s, idx));
			else {
				res.push_back(s[idx]);
				idx++;
			}
				
		}
		return res;
	}

	string parse_repeat(const string &s, int& idx) {
		int n = s.size();
		char ch = s[idx];
		int num = 0;
		string t;

		if (is_digit(s[idx])) {
			while (idx < n && is_digit(s[idx])) {
				num = num * 10 + s[idx] - '0';
				idx++;
			}

			idx++; // eat [

			while (idx < n && s[idx] != ']') {
				if (is_digit(s[idx]))
					t.append(parse_repeat(s, idx));                
				else
                {
					t.push_back(s[idx]);
				    idx++;
                }
			}

            while (idx < n && s[idx] != ']') {
                t.push_back(s[idx]);
                idx++;
            }

			idx++; // eat ]
		} 

		string res;
		for (int i = 0; i < num; i++)
			res.append(t);

		return res;
	}

	bool is_digit(char ch) {
		return ch >= '0' && ch <= '9';
	}
};