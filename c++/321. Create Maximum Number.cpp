class Solution {
public:
    // 贪心求数组中的最大的大小为k的子数组 
    void maxkseq(vector<int> &a, vector<int> &res, int k) {
        if (k == 0) return;
        int n = a.size();
        int i, j;
        res.resize(k, 0);
        
        for (i = 0, j = 0; i < n; i++) {
            while (j > 0 && n-i+j > k && a[i] > res[j-1]) // 出栈
                j--;
            if (j < k) // 入栈
                res[j++] = a[i];        
        }
    }
    
    void merge(const vector<int>& v1, const vector<int>& v2, vector<int> &v3) {
        int lv1 = v1.size(), lv2 = v2.size(), lv3 = lv1 + lv2;
        v3.resize(lv3, 0);
        int i, j, k;
        i = j = k = 0;
        while (i < lv1 || j < lv2) {
            v3[k++] = compare(v1, v2, i, j) ? v1[i++] : v2[j++];
        }
    }
    
    bool compare(const vector<int> &v1, const vector<int> &v2, int i, int j) {
        int lv1 = v1.size(), lv2 = v2.size();
        while (i < lv1 && j < lv2 && v1[i] == v2[j]) {
            i++;
            j++;
        }
        return j == lv2 || (i < lv1 && v1[i] > v2[j]);
    }

    vector<int> maxNumber(vector<int>& nums1, vector<int>& nums2, int k) {
        vector<int> res;
        if (k == 0) return res;
        int ls1 = nums1.size(), ls2 = nums2.size();
        vector<int> v1, v2, v3;
        for (int i = max(0, k - ls2); i <= min(k, ls1); i++) {
            v1.clear(); v2.clear();
            maxkseq(nums1, v1, i);
            maxkseq(nums2, v2, k-i);
            merge(v1, v2, v3);
            if (compare(v3, res, 0, 0))
                res = v3;
        }
        return res;
    }
};