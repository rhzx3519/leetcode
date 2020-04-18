/*
堆排序的实现
　　实现堆排序需要解决两个问题：
　　　　1.如何由一个无序序列建成一个堆？
　　　　2.如何在输出堆顶元素之后，调整剩余元素成为一个新的堆？
　　先考虑第二个问题，一般在输出堆顶元素之后，视为将这个元素排除，然后用表中最后一个元素填补它的位置，自上向下进行调整：首先将堆顶元素和它的左右子树的根结点进行比较，把最小的元素交换到堆顶；然后顺着被破坏的路径一路调整下去，直至叶子结点，就得到新的堆。
　　我们称这个自堆顶至叶子的调整过程为“筛选”。
　　从无序序列建立堆的过程就是一个反复“筛选”的过程。

构造初始堆
　　初始化堆的时候是对所有的非叶子结点进行筛选。
　　最后一个非终端元素的下标是[n/2]向下取整，所以筛选只需要从第[n/2]向下取整个元素开始，从后往前进行调整。
　　比如，给定一个数组，首先根据该数组元素构造一个完全二叉树。
　　然后从最后一个非叶子结点开始，每次都是从父结点、左孩子、右孩子中进行比较交换，交换可能会引起孩子结点不满足堆的性质，所以每次交换之后需要重新对被交换的孩子结点进行调整。

进行堆排序
　　有了初始堆之后就可以进行排序了。
　　堆排序是一种选择排序。建立的初始堆为初始的无序区。
　　排序开始，首先输出堆顶元素（因为它是最值），将堆顶元素和最后一个元素交换，这样，第n个位置（即最后一个位置）作为有序区，前n-1个位置仍是无序区，对无序区进行调整，得到堆之后，再交换堆顶和最后一个元素，这样有序区长度变为2。。。
　　不断进行此操作，将剩下的元素重新调整为堆，然后输出堆顶元素到有序区。每次交换都导致无序区-1，有序区+1。不断重复此过程直到有序区长度增长为n-1，排序完成。
*/


#include <stdio.h>
#include <iostream>
#include <string>
#include <vector>
using namespace std;

void heap_ajust(vector<int>& a, int start, int end)
{
	int t = a[start];
	for (int i = start * 2 + 1; i < end; i++) // 左右孩子分别为2*i+1 和2*i+2
	{
		if (i < end-1 && a[i + 1] > a[i])
			i++;
		if (t > a[i]) break; // 调整结束
		a[start] = a[i];
		start = i;
	}
	a[start] = t; // 插入第一个调整的元素
}

void heap_sort(vector<int>& a)
{
	int n = a.size();
	for (int i = n / 2; i >= 0; i--)
		heap_ajust(a, i, n);

	for (int i = n - 1; i >= 0; i--)
	{
		swap(a[i], a[0]);
		heap_ajust(a, 0, i);
	}
}

/*
归并排序
*/

void merge_sort(vector<pair<int, int>>& A, int l, int r)
{
    if (l < r)
    {
        int mid = l + (r - l)/2;
        merge_sort(A, l, mid);
        merge_sort(A, mid+1, r);
        merge(A, l, mid, mid+1, r);
    }
}

void merge(vector<pair<int, int>>& A, int l1, int r1, int l2, int r2)
{
    vector<pair<int, int>> tmp;
    int i = l1, j = l2;
    while (i <= r1 && j <= r2)
    {
        if (A[i].first <= A[j].first)
        {
            // res[A[i].second] += j - l2;
            tmp.push_back(A[i++]);
        }
        else
            tmp.push_back(A[j++]);
    }
    
    while (i <= r1)
    {
        // res[A[i].second] += j - l2;
        tmp.push_back(A[i++]);
    }
    while (j <= r2)
        tmp.push_back(A[j++]);
    
    for (i = l1; i <= r2; i++)
        A[i] = tmp[i - l1];
}

/*
	快速排序
*/
class Solution {
public:
    ListNode* sortList(ListNode* head) {
        quick_sort(head, nullptr);
        return head;
    }
    
    void quick_sort(ListNode *begin, ListNode *end) {
        if (begin != end) {
            ListNode *p = partition(begin, end);
            quick_sort(begin, p);
            quick_sort(p->next, end);
        }
    }
    
    ListNode* partition(ListNode *head, ListNode *end) {
        ListNode *p = head, *q = head;
        int key = p->val;
        while (p != end) {
            if (p->val < key) {
                q = q->next;
                swap(p->val, q->val);
            }
            p = p->next;
        }
        swap(head->val, q->val);
        return q;
    }
};
