class Solution {
public:
    bool searchMatrix(vector<vector<int>>& matrix, int target) {
        if (matrix.empty()) return false;
        int m = matrix.size(), n = matrix[0].size();
        if (n == 0) return false;
        
        int l = 0, r = m-1, mid;
        while (l <= r) {
            mid = (l+r)>>1;
            if (matrix[mid][0] == target)
                return true;
            else if (matrix[mid][0] > target)
                r = mid - 1;
            else
                l = mid + 1;
        }
        
        int row = max(0, r);
        l = 0; r = n-1;
        while (l <= r) {
            mid = (l+r)>>1;
            if (matrix[row][mid] == target)
                return true;
            else if (matrix[row][mid] > target)
                r = mid - 1;
            else
                l = mid + 1;
        }
        
        return false;
    }
};