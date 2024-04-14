package Q705;

import java.util.HashMap;
import java.util.Map;
import java.util.Set;

/**
 * https://leetcode.cn/problems/design-hashset/?envType=daily-question&envId=2024-04-14
 */
public class Q705 {
}
class MyHashSet {
    private Set<Integer> hash = new HashSet<>();

    public MyHashSet() {

    }

    public void add(int key) {
        hash.add(key);
    }

    public void remove(int key) {
        hash.remove(key);
    }

    public boolean contains(int key) {
        return hash.contains(key);
    }
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * MyHashSet obj = new MyHashSet();
 * obj.add(key);
 * obj.remove(key);
 * boolean param_3 = obj.contains(key);
 */