class Solution {
public:
    int totalNQueens(int n) {
        int res = 0;
		vector<int> A(n, INT_MIN);
		dfs(0, n, res, A);
		return res;
    }
    
    void dfs(int row, int n, int &res, vector<int> A) {
		if (row == n) {
    		res++;
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