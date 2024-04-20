package Q39;

import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

public class Q39 {
}

/**
 * https://leetcode.cn/problems/combination-sum/?envType=daily-question&envId=2024-04-20
 * @return
 */
class Solution {
    private List<List<Integer>> ans = new ArrayList<>();

    private void dfs(int i, int cur, List<Integer> seq, int[] candidates, int target) {
        if (i == candidates.length) {
            if (cur == target) {
                ans.add(new ArrayList<>(seq));
            }
            return;
        }
        for (int k = 0; cur + candidates[i]*k <= target; k++) {
            cur += candidates[i]*k;
            for (int j = 0; j < k; j++) {
                seq.add(candidates[i]);
            }
            dfs(i + 1, cur, seq, candidates, target);
            for (int j = 0; j < k; j++) {
                cur -= candidates[i];
                seq.remove(seq.size() - 1);
            }
        }
    }

    public List<List<Integer>> combinationSum(int[] candidates, int target) {
        dfs(0, 0, new ArrayList<>(), candidates, target);
        return ans;
    }
}


