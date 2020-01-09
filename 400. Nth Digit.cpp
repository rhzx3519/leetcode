class Solution {
public:
	int findNthDigit(int n) {
		if (n < 10) return n;
		int i = 1;
		int cnt = 0;
		while (n - 9 * pow(10, i - 1)*i > 0) {
			n -= 9 * pow(10, i - 1)*i;
			cnt += 9 * pow(10, i - 1);
			i++;
		}
		int a = (n - 1) / i;
		int b = (n - 1) % i;
		int num = pow(10, i-1) + a;
		return foo(num, i, b);
	}

	int foo(int n, int i, int b) {
		for (int j = 0; j < i - 1 - b; j++)
			n /= 10;
		return n % 10;
	}
};