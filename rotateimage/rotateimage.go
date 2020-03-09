package rotateimage

import "math"

func RotateImage(a [][]int) [][]int {
    size := len(a)
    rArr := make([][]int, size)
    for k := range(a){
        rArr[k] = make([]int, size)
    }
    
    if(size == 1){
        return a;
    }
    for i := 0; i < len(a); i++ {
        for j:=0; j< len(a); j++ {
            value := a[i][j]
            rArr[j][math.Abs(i-(size-1))] = value
        }
    }
    return rArr
}

