package main

import "sort"

type pair struct {
    id  int
    val int
}

type SnapshotArray struct {
    snapshots [][]pair
    snapId    int
}

func Constructor(length int) SnapshotArray {
    snapshots := make([][]pair, length)
    return SnapshotArray{
        snapshots: snapshots,
    }
}

func (this *SnapshotArray) Set(index int, val int) {
    this.snapshots[index] = append(this.snapshots[index], pair{this.snapId, val})
}

func (this *SnapshotArray) Snap() int {
    this.snapId++
    return this.snapId - 1
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
    snap := this.snapshots[index]
    var j = sort.Search(len(snap), func(i int) bool {
        return snap[i].id >= snap_id+1
    }) - 1
    if j >= 0 {
        return snap[j].val
    }
    return 0
}

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */
