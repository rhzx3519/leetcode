#!usr/bin/env python  
#-*- coding:utf-8 _*-  

import collections


class Solution(object):
    def countOfAtoms(self, formula):
        """
        :type formula: str
        :rtype: str
        """
        if not formula: return ''
        formula = '(' + formula + ')'
        n = len(formula)
        num = 0
        preName = None
        is_last_list = False
        st = []
        i = 0
        last = None
        while i < n:
            if formula[i]=='(':
                st.append(collections.defaultdict(int))
                i += 1
            elif formula[i]==')':
                last = st[-1]
                i += 1
            elif formula[i].isalpha():
                name = formula[i]
                i += 1
                while i<n and self.islow(formula[i]):
                    name += formula[i]
                    i += 1
                st[-1][name] += 1
                last = name
                # print st[-1]
            elif formula[i].isdigit():
                num = 0
                while i<n and formula[i].isdigit():
                    num = num*10 + (ord(formula[i]) - ord('0'))
                    i += 1

                # print last
                if isinstance(last, str):
                    st[-1][last] += num - 1
                else:
                    for k, v in last.iteritems():
                        st[-1][k] *= num
                    tmp = st.pop()
                    for k, v in tmp.iteritems():
                        st[-1][k] += v
                    # print st[-1]


        # print st
        ans = []
        cnt = st.pop()
        # print cnt
        keys = cnt.keys()
        keys.sort()
        
        for key in keys:
            ans.append(key)
            num = cnt[key]
            if num > 1:
                ans.append(str(num))
        # print ans
        return ''.join(ans)
        
    def islow(self, ch):
        return 'a'<=ch<='z'
    
    def isupper(self, ch):
        return 'A'<=ch<='Z'

if __name__ == '__main__':
    formula = "(((U42Se42Fe10Mc31Rh49Pu49Sb49)49V39Tm50Zr44Og6)33((W2Ga48Tm14Eu46Mt12)23(RuRnMn11)7(Yb15Lu34Ra19CuTb2)47(Md38BhCu48Db15Hf12Ir40)7CdNi21(Db40Zr24Tc27SrBk46Es41DsI37Np9Lu16)46(Zn49Ho19RhClF9Tb30SiCuYb16)15)37(Cr48(Ni31)25(La8Ti17Rn6Ce35)36(Sg42Ts32Ca)37Tl6Nb47Rh32NdGa18Cm10Pt49(Ar37RuSb30Cm32Rf28B39Re7F36In19Zn50)46)38(Rh19Md23No22PoTl35Pd35Hg)41)50"
    # formula = "H2O" 
    # formula = "Mg(OH)2"
    # formula = "K4(ON(SO3)2)2"
    formula = "H11He49NO35B7N46Li20" # "B7H11He49Li20N47O35"
    su = Solution()
    print su.countOfAtoms(formula)





