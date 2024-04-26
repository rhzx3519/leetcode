package Q1146;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.TreeMap;

public class Q1146 {
}

/**
 * https://leetcode.cn/problems/snapshot-array/?envType=daily-question&envId=2024-04-26
 */
class SnapshotArray {

    private int snapId = 0;
    private List<TreeMap<Integer, Integer>> snaps = new ArrayList<>();

    public SnapshotArray(int length) {
        for (int i = 0; i < length; i++) {
            snaps.add(new TreeMap<Integer, Integer>());
        }
    }

    public void set(int index, int val) {
        snaps.get(index).put(snapId, val);
    }

    public int snap() {
        return snapId++;
    }

    public int get(int index, int snap_id) {
        TreeMap<Integer, Integer> tm = snaps.get(index);
        Integer key = tm.floorKey(snap_id);
        return key == null ? 0 : tm.get(key);
    }
}
/**
 * Your SnapshotArray object will be instantiated and called as such:
 * SnapshotArray obj = new SnapshotArray(length);
 * obj.set(index,val);
 * int param_2 = obj.snap();
 * int param_3 = obj.get(index,snap_id);
 */