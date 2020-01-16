// The rand7() API is already defined for you.
// int rand7();
// @return a random integer in the range 1 to 7

class Solution {
public:
    int rand10() {
        while (true) {
            int num = (rand7()-1)*7 + rand7()-1;
            if (num<40)
                return num%10 + 1;
        }
        
        
    }
};

// 利用已有的随机函数rand7将等概率空间扩大，然后再从生成的大等概率空间中取小范围的数。
// rand7 - 1的范围是0 ~ 6，每个数出现概率相等，为1/7，
// （rand7 - 1 ） * 7 的结果是[0, 7, 14, 21, 28, 35, 42]，每个数字出现的概率相等，为1/7，
// 所以 tmp = (rand7()-1)*7 + rand7()-1，得到的数字刚好能均匀覆盖0 ~ 48，每个数对应一种组合，出现的概率是1/49，
// 0 ~ 48 就是得到的大的等概率空间，
// 取0~39这一部分出来，返回%10 +1的结果就是rand10的答案，