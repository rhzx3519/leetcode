/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    vector<ListNode*> splitListToParts(ListNode* root, int k) {
        int n = 0;
        ListNode *p = root;
        while (p) {
            n++;
            p = p->next;
        }

        vector<ListNode*> res;
        if (n <= k) {
            p = root;
            int cnt = 0;
            while (cnt < k) {
                if (p) {                
                    ListNode* next = p->next;
                    p->next = NULL;
                    res.push_back(p);
                    p = next;
                } else
                    res.push_back(NULL);
                cnt++;
            }

            return res;
        }

        int cnt = 0;
        int num = n/k;
        int remain = n%k;
        p = root;
        ListNode* pre = p;
        while (cnt < k) {
            int c = 0;
            int t = remain > 0? 1 : 0;
            res.push_back(p);
            while (c < num + t) {
                pre = p;
                p = p->next;
                c++;
            }
            pre->next = NULL;
            
            remain--;
            cnt++;
        }

        return res;
    }
};



