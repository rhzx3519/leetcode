/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * class NestedInteger {
 *   public:
 *     // Return true if this NestedInteger holds a single integer, rather than a nested list.
 *     bool isInteger() const;
 *
 *     // Return the single integer that this NestedInteger holds, if it holds a single integer
 *     // The result is undefined if this NestedInteger holds a nested list
 *     int getInteger() const;
 *
 *     // Return the nested list that this NestedInteger holds, if it holds a nested list
 *     // The result is undefined if this NestedInteger holds a single integer
 *     const vector<NestedInteger> &getList() const;
 * };
 */
class NestedIterator {
public:
    NestedIterator(vector<NestedInteger> &nestedList) {
        int i = nestedList.size() - 1;
        for (; i >= 0; i--)
            st.push(nestedList[i]);
    }

    int next() {
        NestedInteger ni = st.top();
        st.pop();
        return ni.getInteger();
    }

    bool hasNext() {
        while (!st.empty()) {
            NestedInteger ni = st.top();
            if (ni.isInteger())
                return true;
            st.pop();
            for (int i = ni.getList().size() - 1; i >= 0; i--) {
                st.push(ni.getList()[i]);
            }
        }
        
        return false;
    }
    
private:
    stack<NestedInteger> st;
};

/**
 * Your NestedIterator object will be instantiated and called as such:
 * NestedIterator i(nestedList);
 * while (i.hasNext()) cout << i.next();
 */