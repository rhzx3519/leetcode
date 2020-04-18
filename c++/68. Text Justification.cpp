class Solution {
public:
	vector<string> fullJustify(vector<string>& words, int maxWidth) {
		vector<string> res;
		if (words.empty()) return words;

		if (words.size() == 1) {
			string str = format(0, 0, words, maxWidth, true);
			res.push_back(str);
			return res;
		}
		int len = words[0].size(), start = 0;
		for (int i = 1; i < words.size(); i++) {
			if (len + words[i].size() + 1 > maxWidth) {
				string str = format(start, i - 1, words, maxWidth, false);
				res.push_back(str);
				start = i; len = words[i].size();
			}
			else {
				len += words[i].size() + 1;
			}
		}

		if (start <= words.size() - 1) {
			string str = format(start, words.size() - 1, words, maxWidth, true);
			res.push_back(str);
		}

		return res;
	}

	string format(int start, int end, vector<string> &words, int maxWidth, bool is_last) {
		string res;
		if (start == end) {
			res = words[start];
			res += string(maxWidth - words[start].size(), ' ');
			return res;
		}
		int words_size = 0;
		for (int i = start; i <= end; i++)
			words_size += words[i].size();
		int space_size = maxWidth - words_size;
		int each_space_size = space_size / (end - start), remain = space_size % (end - start);

		if (!is_last) {
			for (int i = start; i <= end; i++) {
				res += words[i];
				if (i != end) {
					res += string(each_space_size, ' ');
					if (remain != 0) {
						res += ' ';
						remain--;
					}
				}

			}
		}
		else {
			for (int i = start; i <= end; i++) {
				res += words[i];
				if (i != end) 
					res += ' ';
			}
			res += string(maxWidth - res.size(), ' ');
		}
		

		return res;
	}
};