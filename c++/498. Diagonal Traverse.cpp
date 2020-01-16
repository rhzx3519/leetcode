class Solution {
public:
	vector<int> findDiagonalOrder(vector<vector<int>>& matrix) {
    	vector<int> res;
        if (matrix.empty()) return res;

		int m = matrix.size(), n = matrix[0].size();
		bool f = true;
		int i = 0, j = 0;
		while (i<m || j < n) {
			if (res.size() >= m * n) break;
			if (j < n)
				res.push_back(matrix[i][j]);

			if (i == 0) {
				if (j < n - 1)
					j++;
				else
					i++;
			}

			while (j > 0) {
				if (i < m && j < n)
					res.push_back(matrix[i][j]);
				i++;
				j--;
			}
			if (i < m)
				res.push_back(matrix[i][j]);

			if (j == 0) {
				if (i < m - 1)
					i++;
				else
					j++;
			}

			while (i > 0) {
				if (i < m && j < n)
					res.push_back(matrix[i][j]);
				i--;
				j++;
			}
		}

		return res;
	}
};

// 简单遍历，边界条件有点多