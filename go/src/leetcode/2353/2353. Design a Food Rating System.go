/**
@author ZhengHao Lou
Date    2022/07/25
*/
package main

import "sort"

type FoodRatings struct {
	Cuisines map[string][]string
	Foods    map[string]string
	Ratings  map[string]int
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	mc := map[string][]string{}
	mf := map[string]string{}
	mr := map[string]int{}
	for i := range foods {
		mc[cuisines[i]] = append(mc[cuisines[i]], foods[i])
		mf[foods[i]] = cuisines[i]
		mr[foods[i]] = ratings[i]
	}
	fr := FoodRatings{
		Cuisines: mc,
		Foods:    mf,
		Ratings:  mr,
	}
	for k := range fr.Cuisines {
		sort.Slice(fr.Cuisines[k], func(i, j int) bool {
			if fr.Ratings[fr.Cuisines[k][i]] != fr.Ratings[fr.Cuisines[k][j]] {
				return fr.Ratings[fr.Cuisines[k][i]] > fr.Ratings[fr.Cuisines[k][j]]
			}
			return fr.Cuisines[k][i] < fr.Cuisines[k][j]
		})
	}
	return fr
}

func (this *FoodRatings) ChangeRating(food string, newRating int) {
	this.Ratings[food] = newRating
	k := this.Foods[food]
	sort.Slice(this.Cuisines[k], func(i, j int) bool {
		if this.Ratings[this.Cuisines[k][i]] != this.Ratings[this.Cuisines[k][j]] {
			return this.Ratings[this.Cuisines[k][i]] > this.Ratings[this.Cuisines[k][j]]
		}
		return this.Cuisines[k][i] < this.Cuisines[k][j]
	})
}

func (this *FoodRatings) HighestRated(cuisine string) string {
	return this.Cuisines[cuisine][0]
}

func main() {
	obj := Constructor([]string{"kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"},
		[]string{"korean", "japanese", "japanese", "greek", "japanese", "korean"}, []int{9, 12, 8, 15, 14, 7})
	obj.ChangeRating("sushi", 16)
	obj.ChangeRating("ramen", 16)
	obj.HighestRated("japanese")
}

/**
 * Your FoodRatings object will be instantiated and called as such:
 * obj := Constructor(foods, cuisines, ratings);
 * obj.ChangeRating(food,newRating);
 * param_2 := obj.HighestRated(cuisine);
 */
