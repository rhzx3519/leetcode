inline bool ISDIGIT(char ch) {  return((ch) >= '0' && (ch) <= '9'); }
inline bool ISDIGIT1TO9(char ch) { return((ch) >= '1' && (ch) <= '9');}

class Solution {
public:
	bool isNumber(string s) {
		int ls = s.size();
		// space
		int i = 0, sign = 1;
		int before = 0;
		while (i < ls && s[i] == ' ') i++;
		if (i == s.size()) return false;
		if (i < ls && s[i] == '-') {
			sign = 0;
			i++;
		}
		else if (i < ls && s[i] == '+')
			i++;
		if (i < ls && s[i] == '0') {
			before = 1;
			i++;
		}

		while (i < ls && ISDIGIT(s[i])) {
			before = 1;
			i++;
		}

		if (i < ls && s[i] == '.') {
			i++;
			if (!before && (i >= ls || !ISDIGIT(s[i]))) return false;
			before = 1;
			while (i < ls && ISDIGIT(s[i])) i++;
		}
		if (i < ls && (s[i] == 'e' || s[i] == 'E')) {
			if (!before) return false;
			i++;
			if (i < ls && (s[i] == '+' || s[i] == '-'))
				i++;
			if (i >= ls || (i < ls && !ISDIGIT(s[i]))) return false;
			i++;
			while (i < ls && ISDIGIT(s[i])) i++;
		}

		while (i < ls && s[i] == ' ') i++;

		return (i >= s.size());
	}
};