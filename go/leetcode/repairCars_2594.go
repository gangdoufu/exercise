package leetcode

import "math"

/*
给你一个整数数组 ranks ，表示一些机械工的 能力值 。ranksi 是第 i 位机械工的能力值。能力值为 r 的机械工可以在 r * (n的平方) 分钟内修好 n 辆车。

同时给你一个整数 cars ，表示总共需要修理的汽车数目。

请你返回修理所有汽车 最少 需要多少时间。
注意：所有机械工可以同时修理汽车。
示例 1：

输入：ranks = [4,2,3,1], cars = 10
输出：16
解释：
- 第一位机械工修 2 辆车，需要 4 * 2 * 2 = 16 分钟。
- 第二位机械工修 2 辆车，需要 2 * 2 * 2 = 8 分钟。
- 第三位机械工修 2 辆车，需要 3 * 2 * 2 = 12 分钟。
- 第四位机械工修 4 辆车，需要 1 * 4 * 4 = 16 分钟。
16 分钟是修理完所有车需要的最少时间。

工人是可以同时修理汽车的.所以总共的时间不会超过任意一个人单独维修的时间。
我们可以通过二分的方式来匹配时间。下限是1 上线时任意一个人单独维修全部汽车的时间
计算在给定的时间内 机械工修理的量的总和,如果大于目标值 则左移。如果小于目标值则右移
*/
func repairCars(ranks []int, cars int) int64 {
	left, right := 1, ranks[0]*cars*cars
	for left < right {
		m := (left + right) >> 1
		cnt := 0
		for _, x := range ranks {
			cnt += int(math.Sqrt(float64(m / x)))
		}
		// 如果修的车超过了右边就移动右边
		if cnt >= cars {
			right = m
		} else {
			left = m + 1
		}

	}
	return int64(left)
}
