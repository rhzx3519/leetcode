/**
 * Definition for singly-linked list with a random pointer.
 * struct RandomListNode {
 *     int label;
 *     RandomListNode *next, *random;
 *     RandomListNode(int x) : label(x), next(NULL), random(NULL) {}
 * };
 */
class Solution {
public:
    RandomListNode *copyRandomList(RandomListNode *head) {
       // write your code here
        RandomListNode *p = head, *pre = nullptr;
        RandomListNode *new_head = nullptr;
        unordered_map<RandomListNode*, RandomListNode*> mp;
        while(p) {
            RandomListNode *node = new RandomListNode(p->label);
            if (pre) {
                pre->next = node;
            } else
                new_head = node;    
            mp[p] = node;
            pre = node;
            p = p->next;
        }
        
        p = head;
        RandomListNode *q = new_head;
        while(p && q) {
            q->random = mp[p->random];
            p = p->next;
            q = q->next;
        }
        
        return new_head;
    }
};