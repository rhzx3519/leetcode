/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * class NestedInteger {
 *   public:
 *     // Constructor initializes an empty nested list.
 *     NestedInteger();
 *
 *     // Constructor initializes a single integer.
 *     NestedInteger(int value);
 *
 *     // Return true if this NestedInteger holds a single integer, rather than a nested list.
 *     bool isInteger() const;
 *
 *     // Return the single integer that this NestedInteger holds, if it holds a single integer
 *     // The result is undefined if this NestedInteger holds a nested list
 *     int getInteger() const;
 *
 *     // Set this NestedInteger to hold a single integer.
 *     void setInteger(int value);
 *
 *     // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 *     void add(const NestedInteger &ni);
 *
 *     // Return the nested list that this NestedInteger holds, if it holds a nested list
 *     // The result is undefined if this NestedInteger holds a single integer
 *     const vector<NestedInteger> &getList() const;
 * };
 */
class Solution {
public:
    NestedInteger deserialize(string s) {
<<<<<<< HEAD
        int n = s.size();
        int idx = 0;
        char ch = s[0];
        if (ch == '[')
            return parse_list(s, idx);
        else
            return parse_int(s, idx);
=======
        int idx = 0;
        char ch = s[idx];
        while (idx < s.size()) {
            if (ch == '[')
                return parse_list(s, idx);
            else
                return parse_number(s, idx);
        }
>>>>>>> 837bde34d06bb470ffcea088785f591f742870d7
    }

private:
    NestedInteger parse_list(const string &s, int &idx) {
        NestedInteger root;
        idx++;// eat '['
        
        while (idx < s.size()) {
            char ch = s[idx];
            if (ch == '[') 
                root.add(parse_list(s, idx));
            else if (is_number(ch) || ch == '-')
                root.add(parse_number(s, idx));
            else if (ch == ',')
                idx++; // skip
            else if (ch == ']')
                break;
        }
        idx++; // eat ']'
        
        return root;
    }
    
    NestedInteger parse_number(const string &s, int &idx) {
        NestedInteger root;
        int n = 0;
        int sign = 1;
        if (s[idx] == '-') {
            sign = -1;
            idx++;
        }
            
        while (idx < s.size()) {
            char ch = s[idx];
            if (is_number(ch)) {
                n = n*10 + (ch - '0');
                idx++;
            } else
                break;
        }
        
        root.setInteger(sign*n);
        return root;
    }
    
    inline bool is_number(char ch) {
        return ch >= '0' && ch <= '9';
    }
    
};