import bisect

class ExamRoom(object):

    def __init__(self, N):
        """
        :type N: int
        """
        self.N = N
        self.students = []

    def seat(self):
        """
        :rtype: int
        """
        pos = 0
        if not self.students:
            pos = 0
        else:
            dist, pos = self.students[0], 0
            for i, j in enumerate(self.students):
                if i==0: continue
                prev = self.students[i-1]
                d = (self.students[i] - prev)/2
                if d > dist:
                    dist = d
                    pos = prev + d
            d = self.N - 1 - self.students[-1]
            if d > dist:
                pos = self.N - 1

        bisect.insort(self.students, pos)
        print self.students
        return pos


    def leave(self, p):
        """
        :type p: int
        :rtype: None
        """
        print self.students
        self.students.remove(p)
        
        
        
        
if __name__ == '__main__':
    N = 10
    obj = ExamRoom(N)
    obj.seat()
    obj.seat()
    obj.seat()
    obj.seat()
    obj.leave(4)
    obj.seat()











