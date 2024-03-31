package Q331;

import java.util.ArrayList;
import java.util.List;
import java.util.Stack;

public class Q331 {
    public boolean isValidSerialization(String preorder) {
        String[] nodes = preorder.split(",");
        List<String> st = new ArrayList<>();
        for (String node : nodes) {
            st.add(node.trim());
            System.out.println( st.get(st.size()-1) == "#");
            while (st.size() >= 3 && "#".equals(st.get(st.size()-1)) && "#".equals(st.get(st.size() - 2))  && !"#".equals(st.get(st.size()-3))) {
                st.remove(st.size() - 1);
                st.remove(st.size() - 1);
                st.remove(st.size() - 1);
                st.add("#");
            }
        }
        return st.size() == 1 && "#".equals(st.get(0));
    }

    public static void main(String[] args) {
        Q331 q = new Q331();
        q.isValidSerialization("9,3,4,#,#,1,#,#,2,#,6,#,#");
    }
}
