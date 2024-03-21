package Q2671;

import javafx.collections.transformation.SortedList;

import java.util.*;

class FrequencyTracker {

    private Map<Integer, Integer> nums = new HashMap();
    private HashMap<Integer, Integer> freq = new HashMap<>();

    public FrequencyTracker() {

    }

    public void add(int number) {
        int current = this.nums.getOrDefault(number, 0);
        freq.put(current, freq.getOrDefault(current, 0) - 1);
        nums.put(number, current + 1);
        freq.put(current + 1, freq.getOrDefault(current + 1, 0) + 1);
    }

    public void deleteOne(int number) {
        if (nums.getOrDefault(number, 0) == 0) {
            return;
        }
        int current = this.nums.getOrDefault(number, 0);
        freq.put(current, freq.getOrDefault(current, 0) - 1);
        nums.put(number, current - 1);
        freq.put(current - 1, freq.getOrDefault(current - 1, 0) + 1);
    }

    public boolean hasFrequency(int frequency) {
        return freq.getOrDefault(frequency, 0) > 0;
    }
}
