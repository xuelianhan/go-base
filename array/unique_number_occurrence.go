package array

func uniqueOccurrences(arr []int) bool {
     freq_map := frequency_map(arr)
     set := make(map[int]bool)
     for _, cnt := range freq_map {
        if _, ok := set[cnt]; ok {
            return false
        }
        set[cnt] = true
     }
     return true
}


func frequency_map( arr []int) map[int]int{
    freq := make(map[int]int)
    for _ , num :=  range arr {
        freq[num] = freq[num]+1
    }
    return freq
}