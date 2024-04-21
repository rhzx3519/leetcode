package Q216;

import java.util.ArrayList;
import java.util.List;

public class Q216 {
}


class Solution {

    private List<List<Integer>> ans = new ArrayList<>();

    private void dfs(int i, int s, int k, List<Integer> arr) {
        if (i > 9 || s <= 0 || arr.size() >= k)  {
            if (s == 0 && arr.size() == k) {
                ans.add(new ArrayList<>(arr));
            }
            return;
        }
        if (s - i >= 0) {
            arr.add(i);
            dfs(i+1, s-i, k, arr);
            arr.remove(arr.size()-1);
        }
        dfs(i+1, s, k, arr);
    }

    public List<List<Integer>> combinationSum3(int k, int n) {
        dfs(1, n,  k, new ArrayList<>());
        return ans;
    }
}



