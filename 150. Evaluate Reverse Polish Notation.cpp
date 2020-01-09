class Solution {
public:
    int evalRPN(vector<string>& tokens) {
        // Write your code here
        stack<int> st;
        for (int i = 0; i < tokens.size(); i++) {
            string ch = tokens[i];
            if (ch == "+" || ch == "-" ||
                ch == "*" || ch == "/") {
                int b = st.top();
                st.pop();
                int a = st.top();
                st.pop();
                int c = 0;
                if (ch == "+")
                     c = a + b;
                else if (ch == "-")
                    c = a - b;
                else if (ch == "*")
                    c = a * b;
                else if (ch == "/")
                    c = a / b;

                st.push(c);
            } else {
                int tmp;
                stringstream ss;
                ss<<ch;
                ss>>tmp;
                st.push(tmp);
            }
        }
        int ret = st.top();
        return ret;
    }
    
};