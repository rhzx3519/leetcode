package main

/**
https://leetcode.cn/problems/student-attendance-record-i/?envType=daily-question&envId=2024-08-18
*/
func checkRecord(s string) bool {
    var cl, ca int
    for i := range s {
        switch s[i] {
        case 'A':
            ca++
            cl = 0
        case 'L':
            cl++
        case 'P':
            cl = 0
        }
        if cl >= 3 {
            return false
        }
        if ca >= 2 {
            return false
        }
    }
    return true
}
