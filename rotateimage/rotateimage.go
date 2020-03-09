package rotateimage

//RotateImage90 rotates the given square matrix by 90 degrees
func RotateImage90(a [][]int) [][]int {
    size := len(a)
    
    if(size == 1){
        return a;
    }
    
    xPos := 0
    yPos := 0
    for i := 0; i < size * size; i++ {
        movedValue := a[xPos][yPos]
        newPos := rotationMatrix90Degrees(xPos,yPos, size)
        xPos = newPos[0]
        yPos = newPos[1]
        temp := a[xPos][yPos]
        a[xPos][yPos] = movedValue
        movedValue = temp
    }
    return a
}

func rotationMatrix90Degrees(xPos int, yPos int, size int) [2]int {
    var location [2]int
    location[0] = yPos
    location[1] = (-1 * xPos) + (size - 1)
    return location
}

