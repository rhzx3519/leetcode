class TweetCounts(object):

    def __init__(self):
        self.user = collections.defaultdict(list)


    def recordTweet(self, tweetName, time):
        """
        :type tweetName: str
        :type time: int
        :rtype: None
        """
        self.user[tweetName].append(time)


    def getTweetCountsPerFrequency(self, freq, tweetName, startTime, endTime):
        """
        :type freq: str
        :type tweetName: str
        :type startTime: int
        :type endTime: int
        :rtype: List[int]
        """
        endTime += 1
        if freq == 'minute':
            length = 60
        elif freq == 'hour':
            length = 60 * 60
        else:
            length = 60 * 60 * 24
        ans = [0] * ((endTime - startTime - 1) // length + 1)
        for t in self.user[tweetName]:
            if endTime > t >= startTime:
                ans[(t - startTime) // length] += 1
        return ans




# Your TweetCounts object will be instantiated and called as such:
# obj = TweetCounts()
# obj.recordTweet(tweetName,time)
# param_2 = obj.getTweetCountsPerFrequency(freq,tweetName,startTime,endTime)