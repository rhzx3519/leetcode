package Q322;

public class Q322 {

    private final int INF = 1_000_00;
    public int coinChange(int[] coins, int amount) {
        int n = coins.length;
        int[] f= new int[amount+1];
        for (int i = 1; i < f.length; i++) {
            f[i] = INF;
        }
        for (int i = 1;i <= amount; i++) {
            for (int coin : coins) {
                for (int j = 1; i - coin*j >= 0; j++) {
                    f[i] = Math.min(f[i], f[i - coin*j] + j);
                }
            }
        }
        if (f[amount] == INF) {
            return -1;
        }
        return f[amount];
    }

    public static void main(String[] args) {
        Q322 q =  new Q322();
        System.out.println(q.coinChange(new int[]{1,2,5}, 11));
        System.out.println(q.coinChange(new int[]{2}, 3));
        System.out.println(q.coinChange(new int[]{1}, 0));
    }
}
