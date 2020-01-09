class Solution {
public:
    uint32_t reverseBits(uint32_t n) {
        uint32_t rn = 0;
        for (int i = 0; i < 32; i++) {
            rn |= (n>>i&1);
            if (i == 31)
				break;
            rn <<= 1;
            
        }
        return rn;
    }
};