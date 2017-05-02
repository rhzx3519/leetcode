class SegmentTreeNode {
public:
    int start, end, sum;
    SegmentTreeNode *left, *right;
    SegmentTreeNode(int start, int end, int sum) {
        this->start = start;
        this->end = end;
        this->sum = sum;
        this->left = this->right = NULL;
    }
};

class NumArray {
public:
    NumArray(vector<int> nums) {
        root = generate(0, nums.size() - 1, nums);
    }
    
    int sumRange(int i, int j) {
        return query(root, i, j);
    }
    
private:
    SegmentTreeNode *generate(int start, int end, vector<int> &A) {
        if (start > end) return nullptr;
        if (start == end) {
            SegmentTreeNode *root = new SegmentTreeNode(start, end, A[start]);
            return root;
        }
        
        SegmentTreeNode *root = new SegmentTreeNode(start, end, 0);
        root->left = generate(start, (start + end) / 2, A);
        root->right = generate((start+end)/2 + 1, end, A);
        root->sum = root->left->sum + root->right->sum;
        return root;
    }
    
    int query(SegmentTreeNode *root, int start, int end) {
        if (!root) return 0;
        if (start <= root->start && root->end <= end) {
            return root->sum;   
        }
        int mid = (root->start + root->end) / 2;
        if (start > mid)
            return query(root->right, start, end);
        else if (end < mid + 1)
            return query(root->left, start, end);
        return query(root->left, start, mid) + query(root->right, mid + 1, end);
    }
    
private:
    SegmentTreeNode *root;
};

/**
 * Your NumArray object will be instantiated and called as such:
 * NumArray obj = new NumArray(nums);
 * int param_1 = obj.sumRange(i,j);
 */