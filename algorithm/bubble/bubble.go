package bubble

import "fmt"

func BubbleSort(list []int) {
    n := len(list)
    // 第一轮是否交换过
    didSwap := false

    for i := n - 1; i > 0; i-- {
        // 每次从第一位开始比较，比较到第i位， 第i位前一轮已排好序了。
        for j := 0; j < i; j++ {
            fmt.Println(list[j])
            fmt.Println(list[j+1])
            if list[j] > list[j+1] {
                list[j], list[j+1] = list[j+1], list[j]
                didSwap = true
            }            
        }
        
        if !didSwap { // 第一轮未交换说明，已排好顺序。
            return
        }
    }
} 


func main() {
    list := []int{5, 9, 1, 6, 14, 6, 49, 25, 4, 6, 3}
    BubbleSort(list)

    fmt.Println(list)
} 
