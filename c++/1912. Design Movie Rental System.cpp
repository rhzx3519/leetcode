//
// Created by LouZhengHao on 2021/6/28.
//

#include<vector>
#include<set>
#include<map>
using namespace std;

class MovieRentingSystem {
private:
    set<pair<int, int>> mv[100001]; // (price, shop)
    set<pair<int, pair<int, int>>> rented; // (price, (shop, movie))
    map<pair<int, int>, int> prices; // {(shop, movie): price}

public:
    MovieRentingSystem(int n, vector<vector<int>>& entries) {
        for (size_t i = 0; i < entries.size(); i++) {
            int shop = entries[i][0], movie = entries[i][1], price = entries[i][2];
            prices[make_pair(shop, movie)] = price;
            mv[movie].insert(make_pair(price, shop));
        }
    }

    vector<int> search(int movie) {
        vector<int> ans;
        auto it = mv[movie].begin();
        for (size_t i = 0; i < 5 && it != mv[movie].end(); i++, it++) {
            ans.push_back(it->second);
        }
        return ans;
    }

    void rent(int shop, int movie) {
        auto pr = make_pair(shop, movie);
        int price = prices[pr];
        rented.insert(make_pair(price, pr));
        mv[movie].erase(make_pair(price, shop));
    }

    void drop(int shop, int movie) {
        auto pr = make_pair(shop, movie);
        int price = prices[pr];
        rented.erase(make_pair(price, pr));
        mv[movie].insert(make_pair(price, shop));
    }

    vector<vector<int>> report() {
        vector<vector<int>> ans;
        for (auto it = rented.begin(); it != rented.end() && ans.size() < 5; it++) {
            auto pr = it->second;
            ans.push_back({pr.first, pr.second});
        }

        return ans;
    }
};


/**
 * Your MovieRentingSystem object will be instantiated and called as such:
 * MovieRentingSystem* obj = new MovieRentingSystem(n, entries);
 * vector<int> param_1 = obj->search(movie);
 * obj->rent(shop,movie);
 * obj->drop(shop,movie);
 * vector<vector<int>> param_4 = obj->report();
 */
