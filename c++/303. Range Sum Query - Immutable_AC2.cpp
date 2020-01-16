class NumArray {
public:
    NumArray(vector<int> nums) {
        accu.push_back(0);
        for (auto &i : nums)
            accu.push_back(accu.back() + i);
    }
    
    int sumRange(int i, int j) {
        return accu[j+1] - accu[i];
    }
    
private:
    vector<int> accu;
};

/**
 * Your NumArray object will be instantiated and called as such:
 * NumArray obj = new NumArray(nums);
 * int param_1 = obj.sumRange(i,j);
 */