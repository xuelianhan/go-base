package binarysearch 

func search(nums []int, target int) int {
    var low, high = 0, len(nums) - 1
    for low <= high {
        var mid = low + (high - low) / 2
        if (nums[mid] == target) {
            return mid;
        } else if (nums[mid] < target) {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return -1
}
