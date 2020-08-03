package code;

import java.util.List;
import java.util.Map;
import java.util.PriorityQueue;
import java.util.Set;

class Twitter {

    private static int timeStamp = 0;

    private static class User {
        private int id;
        public Set<Integer> followed;
        public Tweet head;

        public User(int id) {
            this.id = id;
            this.head = null;
            this.followed = new HashSet<>();
            this.followed.add(this.id);
        }

        public void follow(int userId) {
            this.followed.add(userId);
        }

        public void unfollow(int userId) {
            if (this.id!=userId)
                this.followed.remove(userId);
        }

        public void post(int tweetId) {
            Tweet twt = new Tweet(tweetId, timeStamp);
            timeStamp++;
            twt.next = head;
            head = twt;
        }

    }

    private static class Tweet {
        private int id;
        private int time;
        private Tweet next;

        public Tweet(int id, int time) {
            this.id = id;
            this.time = time;
        }
    }

    private Map<Integer, User> userMap = new HashMap<>();

    /** Initialize your data structure here. */
    public Twitter() {

    }

    /** Compose a new tweet. */
    public void postTweet(int userId, int tweetId) {
        if (!userMap.containsKey(userId)) {
            userMap.put(userId, new User(userId));
        }
        User user = userMap.get(userId);
        user.post(tweetId);
    }

    /** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
    public List<Integer> getNewsFeed(int userId) {
        if (!userMap.containsKey(userId)) {
            userMap.put(userId, new User(userId));
        }
        User user = userMap.get(userId);
        Set<Integer> users = user.followed;

        PriorityQueue<Tweet> pq = new PriorityQueue<>(users.size(), (a, b)->(b.time - a.time));
        for (int followeeId: user.followed) {
            Tweet twt = userMap.get(followeeId).head;
            if (twt==null) {
                continue;
            }
            pq.add(twt);
        }

        List<Integer> res = new ArrayList<>();
        while (!pq.isEmpty()) {
            if (res.size()==10) {
                break;
            }
            Tweet twt = pq.poll();
            if (twt.next!=null) {
                pq.add(twt.next);
            }
            res.add(twt.id);
        }
        return res;

    }

    /** Follower follows a followee. If the operation is invalid, it should be a no-op. */
    public void follow(int followerId, int followeeId) {
        if (!userMap.containsKey(followerId)) {
            userMap.put(followerId, new User(followerId));
        }
        if (!userMap.containsKey(followeeId)) {
            userMap.put(followeeId, new User(followeeId));
        }
        User user = userMap.get(followerId);
        user.follow(followeeId);
    }

    /** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
    public void unfollow(int followerId, int followeeId) {
        if (userMap.containsKey(followerId)) {
            User user = userMap.get(followerId);
            user.unfollow(followeeId);
        }

    }
}
/**
 * Your Twitter object will be instantiated and called as such:
 * Twitter obj = new Twitter();
 * obj.postTweet(userId,tweetId);
 * List<Integer> param_2 = obj.getNewsFeed(userId);
 * obj.follow(followerId,followeeId);
 * obj.unfollow(followerId,followeeId);
 */