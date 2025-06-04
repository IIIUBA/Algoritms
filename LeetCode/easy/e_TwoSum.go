package main

// My solution
// func twoSum(nums []int, target int) []int {

// 	var res []int

// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i]+nums[j] == target {
// 				res = append(res, i, j)
// 				return res
// 			}
// 		}

// 	}
// 	return []int{}
// }

func twoSum(nums []int, target int) []int {

	// создаём хэш-мапу
	m := make(map[int]int)

	// записываем значения в мапу
	for i, key := range nums {
		m[key] = i
	}

	// перебираем массив
	for i, num := range nums {
		// x = target - current
		complete := target - num
		// если целевое число найдено, которое с i в сумме даёт j то возвращаем его
		if j, ok := m[complete]; ok && j != i {
			return []int{i, j}
		}
	}

	return []int{}
}

// func main() {
// 	targer := 6
// 	nums := []int{3, 3, 4, 6, 5}

// 	fmt.Println(twoSum(nums, targer))
// }
