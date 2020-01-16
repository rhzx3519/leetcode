class Solution {
public:
    int maxRotateFunction(vector<int>& A) {
        long f = 0,sum = 0;
        long len = A.size();
        for(int i=0;i<len;++i){
            sum+=A[i];
            f+=i*A[i];
        }
        long ans = f;
        for(int i=1;i<len;++i){
            f = f+sum-len*A[len-i];
            ans = max(ans,f);
        }
        return ans;

    }

};