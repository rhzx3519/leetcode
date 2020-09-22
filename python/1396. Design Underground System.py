class UndergroundSystem(object):

    def __init__(self):
        self.routes = {}
        self.cost = {}


    def checkIn(self, id, stationName, t):
        """
        :type id: int
        :type stationName: str
        :type t: int
        :rtype: None
        """
        if id not in self.routes:
            self.routes[id] = []
        self.routes[id].append([stationName, t])


    def checkOut(self, id, stationName, t):
        """
        :type id: int
        :type stationName: str
        :type t: int
        :rtype: None
        """
        start, t1 = self.routes[id][-1]

        if (start, stationName) not in self.cost:
            self.cost[(start, stationName)] = [0, 0]
        self.cost[(start, stationName)][0] += t - t1
        self.cost[(start, stationName)][1] += 1


    def getAverageTime(self, startStation, endStation):
        """
        :type startStation: str
        :type endStation: str
        :rtype: float
        """
        if (startStation, endStation) not in self.cost:
            return 0
        s, c = self.cost[(startStation, endStation)]
        return float(s) /c




# Your UndergroundSystem object will be instantiated and called as such:
# obj = UndergroundSystem()
# obj.checkIn(id,stationName,t)
# obj.checkOut(id,stationName,t)
# param_3 = obj.getAverageTime(startStation,endStation)