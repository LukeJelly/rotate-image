package rotateimage

//RotateImage90 rotates the given square matrix by 90 degrees
func RotateImage90(a [][]int) [][]int {
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
            rArr[j][abs(i-(size-1))] = value
        }
    }
    return rArr
}

func abs(num int) int{
    if (num < 0){
        return -1 * num
    }else{
        return num
    }
}

