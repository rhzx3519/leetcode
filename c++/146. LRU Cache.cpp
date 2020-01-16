class LRUCache {
public:
    LRUCache(int capacity):size(capacity) {
    }
    
    int get(int key) {
        auto it = hash.find(key);
        if (it != hash.end()) {
            cache.splice(cache.begin(), cache, it->second); //splice 从cache中去除迭代器指向的元素，添加到cache头部
            return it->second->second;
        }
        return -1;
    }
    
    void put(int key, int value) {
        auto it = hash.find(key);
        if (it != hash.end()) {
            it->second->second = value;
            cache.splice(cache.begin(), cache, it->second);
            return;
        }
        
        cache.insert(cache.begin(), make_pair(key, value));
        hash[key] = cache.begin();
        if (cache.size() > size) {
            hash.erase(cache.back().first);
            cache.pop_back();
        }
        
    }

private:
    unordered_map<int, list<pair<int, int>>::iterator> hash;
    list<pair<int, int>> cache;
    int size;
};

/**
 * Your LRUCache object will be instantiated and called as such:
 * LRUCache obj = new LRUCache(capacity);
 * int param_1 = obj.get(key);
 * obj.put(key,value);
 */