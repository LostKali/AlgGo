package matrix

/*
You are given an m x n matrix accounts where accounts[i][j] is the amount of money the i​​​​​​​​​​​th​​​​ customer
has in the j​​​​​​​​​​​th​​​​ bank.

Return the wealth that the richest customer has.

Imagine every customer has multiple bank accounts, with each account holding a certain amount of money.
The total wealth of a customer is calculated by summing all the money across all their multiple.
*/
import "fmt"

type Solution struct{}

func (s Solution) maximumWealth(accounts [][]int) int {
	maxWealth := 0

	for _, account := range accounts {
		wealth := 0
		for _, money := range account {
			wealth += money
		}
		if wealth > maxWealth {
			maxWealth = wealth
		}
	}

	return maxWealth
}

func main() {
	solution := Solution{}

	fmt.Println(solution.maximumWealth([][]int{{5, 2, 3}, {0, 6, 7}}))            // 13
	fmt.Println(solution.maximumWealth([][]int{{1, 2}, {3, 4}, {5, 6}}))          // 11
	fmt.Println(solution.maximumWealth([][]int{{10, 100}, {200, 20}, {30, 300}})) // 330
}
