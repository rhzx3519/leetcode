package Q518;

public class Q518 {
    public int change(int amount, int[] coins) {
        int[] f = new int[amount+1];
        f[0] = 1;
        for (int coin : coins) {
            for (int i = coin; i <= amount; i++) {
                f[i] += f[i - coin];
            }
        }

        System.out.println(f);
        return f[amount];
    }

    public static void main(String[] args) {
        Q518 q = new Q518();
        q.change(5, new int[]{1,2,5});
    }
}
