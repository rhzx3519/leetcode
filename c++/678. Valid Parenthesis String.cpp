class Solution {
public:
    bool checkValidString(string s) {
        stack<int> left, star;
        for (int i = 0; i < s.size(); i++) {
            char ch =s[i];
            if (ch == '(') left.push(i);
            else if (ch == '*') star.push(i);
            else {
                if (left.empty() && star.empty()) return false;
                if (!left.empty()) left.pop();
                else star.pop();
            }
        }

        while (!left.empty() && !star.empty()) {
            if (left.top() > star.top()) return false;
            left.pop();
            star.pop();
        }

        return left.empty();
    }
};

// 两个栈分别将"("和"*"的序号压入栈中，每次遇到右括号，首先检测左括号栈中是否为空，
// 不为空则弹出元素，否则弹出star栈，最后考虑left和star栈可能存在元素，判断此时栈中元素的序号大小，
// 如果left栈中的序号大于star中的则表明 存在"*("此种情况，直接false