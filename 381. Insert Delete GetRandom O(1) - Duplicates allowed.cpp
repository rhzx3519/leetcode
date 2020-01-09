class RandomizedCollection {
public:
    /** Initialize your data structure here. */
    RandomizedCollection() {
        
    }
    
    /** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
    bool insert(int val) {
        bool res = um.find(val) == um.end();
        um[val].push_back(nums.size()); // um stores the pos of val in nums;
        nums.push_back(make_pair(val, um[val].size() - 1)); // nums stores the pos of val in um;
        
        return res;
    }
    
    /** Removes a value from the collection. Returns true if the collection contained the specified element. */
    bool remove(int val) {
        if (um.find(val) == um.end()) return false;
        
        pair<int, int> last = nums.back();
        um[last.first][last.second] = um[val].back();
        nums[um[val].back()] = last;
        um[val].pop_back();
        if (um[val].empty())
            um.erase(val);
        nums.pop_back();
        
        return true;
    }
    
    /** Get a random element from the collection. */
    int getRandom() {
        return nums[rand() % nums.size()].first;
    }
    
private:
    vector<pair<int, int>> nums;
    unordered_map<int, vector<int>> um;
};

/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * RandomizedCollection obj = new RandomizedCollection();
 * bool param_1 = obj.insert(val);
 * bool param_2 = obj.remove(val);
 * int param_3 = obj.getRandom();
 */