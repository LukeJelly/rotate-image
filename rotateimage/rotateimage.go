package rotateimage

//RotateImage90 rotates the given square matrix by 90 degrees
func RotateImage90(arr [][]int) [][]int {
    size := len(arr)
    
    if(size == 1){
        return arr;
    }
    
    for i := 0; i <= size / 2; i++ {
        for j := i; j <  (size - (i + 1)); j++{
            rotateThroughFour(i,j,size,arr)
        }
    }
    return arr
}

func rotateThroughFour(row int, col int, size int, arr [][]int){
    startValue := arr[row][col]
    movedValue := startValue;
    arr[row][col] *= -1

    xPos := row
    yPos := col
    doneFullCircle := false

    for !doneFullCircle {
        newPos := rotationMatrix90Degrees(xPos,yPos, size)
        xPos = newPos[0]
        yPos = newPos[1]
        temp := arr[xPos][yPos]
        arr[xPos][yPos] = movedValue
        movedValue = temp
        if (temp < 0){
            doneFullCircle = true
        }
    }
}

func rotationMatrix90Degrees(xPos int, yPos int, size int) [2]int {
    var location [2]int
    location[0] = yPos
    location[1] = (-1 * xPos) + (size-1)
    return location
}

