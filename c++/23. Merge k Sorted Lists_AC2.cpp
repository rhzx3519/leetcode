/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */

struct cmp
{
    bool operator()(ListNode* l1, ListNode*l2)  
    {
        return l1->val > l2->val;
    }
};

class Solution {
public:
    ListNode* mergeKLists(vector<ListNode*>& lists) {
        priority_queue<ListNode*, vector<ListNode*>, cmp> que;
        for (auto &p : lists)
        {
            if (p)
                que.push(p);
        }
            
        if (que.empty()) return NULL;
        
        ListNode* head = NULL, *tail = NULL;
        while (!que.empty())
        {
            ListNode* p = que.top();
            que.pop();
            if (!p) continue;
            
            if (p->next)
                que.push(p->next);
            p->next = NULL;

            
            if (!tail)
                head = p;
            else
                tail->next = p;
            
            tail = p;
        }
        return head;        
    }
};