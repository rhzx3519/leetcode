#!usr/bin/env python  
#-*- coding:utf-8 _*-  

# Exponential Moving Average (EMA)
# https://www.investopedia.com/terms/e/ema.asp

from __future__ import division
import numpy as np
import matplotlib.pyplot as plt
import random
 
period = 20
que = []
x = []
y = []
sma = []
ema = []
N = 100
pre = None
smoothing = 10.0

for i in range(N):
    x.append(i)
    val = random.randint(100, 100 + i)
    y.append(val)

    que.append(val)
    if len(que) > period:
        que.pop(0)

    # print que
    if len(que) < period:
        sma.append(0)
    else:
        sma.append(sum(que)/len(que))
        
    if i == period:
        pre = sum(que)/len(que)

    if not pre:
        ema.append(0)
    else:
        ema.append(val*(smoothing/(1+period)) + pre*(1-(smoothing/(1+period))))
        pre = ema[-1]
        # print ema[-1]



# print sma
plt.figure()
plt.plot(x, y, color='black', label='origin')
plt.plot(x, sma, color='red', label='sma')
plt.plot(x, ema, color='yellow', label='ema')
plt.show()


# a = [5, 6, 7, 13, 43, 45, 46, 55, 56, 60, 61, 62, 65, 66, 66, 67, 90, 100, 104, 132]
# print np.percentile(a, 25, interpolation='linear')


