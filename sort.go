package main

import "fmt"
func sort(nums...int ){
	fmt.Println("排序前：",nums)
	for i:=0;i< len(nums);i++{
		for j := 0; j < len(nums)-1; j++ {
			t:=0
			if nums[j]>nums[j+1]{
				t=nums[j]
				nums[j]=nums[j+1]
				nums[j+1]=t
			}
		}
	}
	fmt.Println("排序后：",nums)
}
func main() {
	nums:=make([]int,5)
	fmt.Println("input:")
	for i := 0; i < len(nums); i++ {
		fmt.Scan(&nums[i])
	}
	sort(nums...)

}
