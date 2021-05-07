package main

type Employee struct {
	Id int
	Importance int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	ans := 0
	for _, e := range employees {
		if e.Id == id {
			ans += e.Importance
			for _, sub := range e.Subordinates {
				ans += getImportance(employees, sub)
			}
		}
	}

	return ans
}
