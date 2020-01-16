struct Tweet{
    int id;
    int order;  
    Tweet(int id_, int order_):id(id_), order(order_) {
        
    }
};

class Twitter {
public:
    /** Initialize your data structure here. */
    Twitter() {
        order = 0;
    }
    
    /** Compose a new tweet. */
    void postTweet(int userId, int tweetId) {
        user_tweets[userId].push_back(Tweet(tweetId, order++));
    }
    
    /** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
    vector<int> getNewsFeed(int userId) {
        vector<int> res;
        vector<Tweet> user_t;
        if (user_tweets.count(userId)) {
            vector<Tweet> v = user_tweets[userId];
            user_t = getLastTen(v);
        }
        
        if (friends.count(userId)) {
            for (int fid : friends[userId]) {
                if (user_tweets.count(fid)) {
                    vector<Tweet> v = user_tweets[fid];
                    user_t.insert(user_t.end(), v.begin(), v.end());
                }
            }
        }
        
        sort(user_t.begin(), user_t.end(), [](const Tweet &t1, const Tweet &t2){
            return t1.order > t2.order;    
        });
        
        vector<Tweet> v = getFirstTen(user_t);
        for (int i = 0; i < v.size(); i++)
            res.push_back(v[i].id);
        return res;
    }
    
    /** Follower follows a followee. If the operation is invalid, it should be a no-op. */
    void follow(int followerId, int followeeId) {
        if (followerId == followeeId) return;
        friends[followerId].insert(followeeId);
    }
    
    /** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
    void unfollow(int followerId, int followeeId) {
        if (followerId == followeeId) return;
        friends[followerId].erase(followeeId);
    }
    
private:
    vector<Tweet> getFirstTen(vector<Tweet> &tweets) {
        int last = 10;
        last = min((int)tweets.size(), last);
        return vector<Tweet>(tweets.begin(), tweets.begin() + last);
    }
    
    vector<Tweet> getLastTen(vector<Tweet> &tweets) {
        int last = 10;
        last = min((int)tweets.size(), last);
        return vector<Tweet>(tweets.end() - last, tweets.end());
    }

    unordered_map<int, vector<Tweet>> user_tweets;
    unordered_map<int, set<int>> friends;
    int order;
};

/**
 * Your Twitter object will be instantiated and called as such:
 * Twitter obj = new Twitter();
 * obj.postTweet(userId,tweetId);
 * vector<int> param_2 = obj.getNewsFeed(userId);
 * obj.follow(followerId,followeeId);
 * obj.unfollow(followerId,followeeId);
 */