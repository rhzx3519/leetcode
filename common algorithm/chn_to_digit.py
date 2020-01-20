#!usr/bin/env python  
#-*- coding:utf-8 _*-  

cn_sum = {
    '一': 1, '二': 2, '三': 3, '四': 4, '五': 5, '六': 6, '七': 7, '八': 8, '九': 9, '零': 0,
}

cn_unit = {
    '十': 10,
    '拾': 10,
    '百': 100,
    '佰': 100,
    '千': 1000,
    '仟': 1000,
    '万': 10000,
    '萬': 10000,
    '亿': 100000000,
    '億': 100000000,
    '兆': 1000000000000,
    '角': 0.1,
    '分': 0.01
}

def chn_to_sum(chn):
    # 传入字符串
    arr = []
    stack = []
    res = 0
    for ch in chn:
        if ch in cn_unit:
            tmp = 0
            for i in range(len(stack)):
                tmp = tmp*10 + stack[i]
            stack = []
            res += tmp * cn_unit[ch]
            print(res)
        else:
            stack.append(cn_sum[ch])

    tmp = 0
    for i in range(len(stack)):
        tmp = tmp*10 + stack[i]

    res += tmp
    print(res)



if __name__ == '__main__':
    chn_to_sum(u"一百零五") 
