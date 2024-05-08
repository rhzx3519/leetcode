package Q2079;

public class Q2079 {
}

/**
 * https://leetcode.cn/problems/watering-plants/?envType=daily-question&envId=2024-05-08
 */
class Solution {
    public int wateringPlants(int[] plants, int capacity) {
        int sum = 0;
        int left = capacity;
        int tot = 0;
        for (int i = 0; i < plants.length; i++) {
            left -= plants[i];
            if (left < 0) {
                tot += i*2 + 1;
                left = capacity - plants[i];
            } else {
                tot++;
            }
        }
        return tot;
    }

    public static void main(String[] args) {
        Solution s = new Solution();
//        System.out.println(s.wateringPlants(new int[]{2,2,3,3}, 5));
//        System.out.println(s.wateringPlants(new int[]{1,1,1,4,2,3}, 4));
        System.out.println(s.wateringPlants(new int[]{3,2,4,2,1}, 6));
    }
}
