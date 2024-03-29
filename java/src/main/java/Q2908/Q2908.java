package Q2908;

public class Q2908 {
    public int minimumSum(int[] nums) {
        int n = nums.length;
        int[] left = new int[n];
        int min = nums[0];
        for (int i = 1; i < n; i++) {
            if (min < nums[i]) {
                left[i] = min;
            }
            min = Math.min(min, nums[i]);
        }

        int ans = Integer.MAX_VALUE;
        min = nums[n-1];
        for (int i = n-2; i >= 0; i--) {
            if (min < nums[i] && left[i] != 0) {
                ans = Math.min(ans, left[i] + nums[i] + min);
            }
            min = Math.min(min, nums[i]);
        }

        if (ans == Integer.MAX_VALUE) {
            return -1;
        }
        return ans;
    }
}
