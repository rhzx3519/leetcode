class Solution {
public:
    int fourSumCount(vector<int>& A, vector<int>& B, vector<int>& C, vector<int>& D) {
        unordered_map<int, int> um;
        for (int i = 0; i < A.size(); i++) {
            for (int j = 0; j < B.size(); j++) {
                um[A[i]+B[j]]++;
            }
        }

        int res = 0;
        for (int i = 0; i < C.size(); i++) {
            for (int j = 0; j < D.size(); j++) {
                int s = -(C[i] + D[j]);
                if (um.count(s))
                    res += um[s];
            }
        }

        return res;
    }
};