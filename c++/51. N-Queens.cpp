class Solution {
public:
	vector<vector<string>> solveNQueens(int n) {
		vector<vector<string>> res;
		vector<int> A(n, INT_MIN);
		dfs(0, n, res, A);
		return res;
	}

	void dfs(int row, int n, vector<vector<string>> &res, vector<int> A) {
		if (row == n) {
			vector<string> rows;
			for (int i = 0; i < n; i++) {
				string str;
				for (int j = 0; j < n; j++) {
					if (j == A[i])
						str += 'Q';
					else
						str += '.';
				}
				rows.push_back(str);
			}
			res.push_back(rows);
			return;
		}

		for (int i = 0; i < n; i++) {
			if (place(row, i, A, n)) {
				A[row] = i;
				dfs(row + 1, n, res, A);
				A[row] = INT_MIN;
			}
		}
	}

	bool place(int row, int col, vector<int> &A, int n) {
		for (int i = 0; i < n; i++)
			if (A[i] == col || abs(A[i] - col) == abs(row - i))
				return false;
		return true;
	}

};
