func twoSum(nums []int, target int) []int {
    // ハッシュ作成
    hashmap := make(map[int]int)
    for i,num := range nums {
        key := target - num
        // 二つ目の返り値がokであれば次の処理
        if j,ok := hashmap[key];ok{
            return []int{j, i}
        }
        hashmap[num] = i
    }
    return nil
}