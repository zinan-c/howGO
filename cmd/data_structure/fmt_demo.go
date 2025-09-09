// cmd/api/main.go
package main

/*
 * not sure the import items could change to from local rather than github.com/zinan-c/howGO
 * because the go.mod module name is changed to howgo
 */
import (
    "fmt"
)

func main() {
    a := 10

    // simple output the variable a
    fmt.Println(a)

    // you can also do some string operations and then output
    fmt.Println("Hello" + ", " + "World!")

    s := "abc"
    // which is more convenient
    fmt.Println(s, a)

    // also you can use formatted output
    fmt.Printf("%s %d\n", s, a)

    a := 10
    
    //demo for if-else statement
    if a > 5 {
        fmt.Println("a > 5")
    } else if a == 5 {
        fmt.Println("a == 5")
    } else {
        fmt.Println("a < 5")
    }

    //demo for switch statement
    switch a {
    case 1:
        fmt.Println("a is 1")
    case 5:
        fmt.Println("a is 5")
    case 10:
        fmt.Println("a is 10")
    default:
        fmt.Println("a is not 1, 5, or 10")
    }

    //demo for for loop
    for i := 0; i < 5; i++ {
        fmt.Print(i, " ")
    }
    fmt.Println()

    num := 100
    for num > 0 {
        fmt.Print(num, " ")
        num /= 2
    }
    fmt.Println()

    //demo for slice initialization
	var nums []int

	// initialize a slice nums with length 7, all values are 0
	nums = make([]int, 7)

	// initialize a slice nums with values 1, 3, 5
	nums = []int{1, 3, 5}

	// initialize a slice nums with length 7, all values are 2
	nums = make([]int, 7)
	for i := range nums {
		nums[i] = 2
	}

	fmt.Println(nums)

	// initialize a 2D slice dp with 3 rows and 3 columns, all values are true
	var dp [][]bool
	dp = make([][]bool, 3)
	for i := 0; i < len(dp); i++ {
		row := make([]bool, 3)
		for j := range row {
			row[j] = true
		}
		dp[i] = row
	}

	fmt.Println(dp)

    n := 10
	// initialize a slice nums with length n, all values are 0
	nums := make([]int, n)

	// output：false
	fmt.Println(len(nums) == 0)

	// output ：10
	fmt.Println(len(nums))

	// in the slice nums, add an element 20 at the end
	// append function will return a new slice with the new element added
	nums = append(nums, 20)
	// output：11
	fmt.Println(len(nums))

	// get the last element of the slice
	// output：20
	fmt.Println(nums[len(nums)-1])

	// delete the last element of the slice
	nums = nums[:len(nums)-1]
	// output：10
	fmt.Println(len(nums))

	// also you can modify the elements in the slice
	nums[0] = 11
	// output：11
	fmt.Println(nums[0])

	// insert 99 at index 3
	// ... is used to unpack the slice, because append only accepts variadic parameters
	nums = append(nums[:3], append([]int{99}, nums[3:]...)...)

	// delete the element at index 2
	nums = append(nums[:2], nums[3:]...)

	// exchange the nums[0] and nums[1]
	nums[0], nums[1] = nums[1], nums[0]

	// interate the slice
	for _, num := range nums {
		fmt.Print(num, " ")
	}
	fmt.Println()
}