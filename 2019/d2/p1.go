package main

import "fmt"

func main() {
	inp1 := []int{1, 12, 2, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 13, 1, 19, 1, 6, 19, 23, 2, 23, 6, 27, 1, 5, 27, 31, 1, 10, 31, 35, 2, 6, 35, 39, 1, 39, 13, 43, 1, 43, 9, 47, 2, 47, 10, 51, 1, 5, 51, 55, 1, 55, 10, 59, 2, 59, 6, 63, 2, 6, 63, 67, 1, 5, 67, 71, 2, 9, 71, 75, 1, 75, 6, 79, 1, 6, 79, 83, 2, 83, 9, 87, 2, 87, 13, 91, 1, 10, 91, 95, 1, 95, 13, 99, 2, 13, 99, 103, 1, 103, 10, 107, 2, 107, 10, 111, 1, 111, 9, 115, 1, 115, 2, 119, 1, 9, 119, 0, 99, 2, 0, 14, 0}
	f2(inp1, 19690720)
}

func f2(clrMem []int, g int) {

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			//i := 0
			mem := append([]int(nil), clrMem...)
			mem[1] = noun
			mem[2] = verb

			r := f1(mem)
			fmt.Printf("%d:%d == %d\n", noun, verb, r)
			if r == g {
				fmt.Printf("noun: %d verb: %d", noun, verb)
				fmt.Printf("res: %d", 100*noun+verb)

				return
			}
		}
	}

}

func f1(l []int) int {

	for i := 0; i < len(l); i = i + 4 {

		switch op := l[i]; op {
		case 1:
			a1 := l[l[i+1]]
			a2 := l[l[i+2]]
			goalID := l[i+3]
			l[goalID] = a1 + a2
		case 2:
			a1 := l[l[i+1]]
			a2 := l[l[i+2]]
			goalID := l[i+3]
			l[goalID] = a1 * a2
		case 99:
			break
		default:
			break
		}

	}
	return l[0]
}
