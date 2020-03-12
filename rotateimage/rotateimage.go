package rotateimage

/* Prompt: You are given an n x n 2D matrix that represents an image. Rotate 
 * the image by 90 degrees (clockwise).
 *
 * This was given to me as an example coding interview question. I solved it 
 * a differnet way but couldn't shake how much this seemed like it needed a 
 * linearly algebra approach. I solved it using a 2x2 rotation matrix to find
 * the new location for every given point. 
 *
 * I will use a 5x5 matrix to explain the algorithm but it works for any nxn
 * matrix.  
 *
 *     Matrix A                   Rotation Matrix B 
 *  [[a,b,c,d,e],             [[cos(theta), -sin(theta)],
 *   [f,g,h,i,j],              [sin(theta), cos(theta)]]
 *   [k,l,m,n,o],                Evaluated at -90 degrees
 *   [p,q,r,s,t],                      [[0, 1],
 *   [u,v,w,x,y]]                      [-1, 0]] 
 *  
 *  Give any point in the matrix A (eq, (0,0)), if we put multiply that point by
 *  matrix B we get a new point (0,0).  Of course this rotates around the origin,
 *  and we want to rotate in place so we have to shift that up by the size of the
 *  matrix - 1.  So our next location because (0,0 + (len(a)-1)) = (0,4).  Which
 *  is where the 'a' would go after the rotation.  We then pull out what is 
 *  currently in that location and put the a there.  We can then take our point
 *  (0,4) and put it back through the rotation matrix to find where the next point
 *  will be.  This will happen 4 times before we get back to where we start.   *
 *  Now that we have defined how to use the rotation matrix, we can abstract that 
 *  into rotateThroughFour.  
 *
 *  If we loop through all values in the first row, minus
 *  the last index, we have affectively rotated the entire outer "ring".  We can 
 *  now move into the inner ring [[g,h,i],[l,m,n],[q,r,s]] and do the same thing
 *  over again.  Keep rotating all outer rings until we there are no more rings.
 *
 * 
 * Version 1.0
 * Luke.Kelly
 */


//RotateImage90 rotates the given square matrix by 90 degrees
func RotateImage90(arr [][]int) [][]int {
    size := len(arr)
    
    if(size == 1){
        return arr
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

    for i:=0; i < 4; i++ {
        newPos := findNewPostionInMatrix(xPos,yPos, size)
        xPos = newPos[0]
        yPos = newPos[1]
        temp := arr[xPos][yPos]
        arr[xPos][yPos] = movedValue
        movedValue = temp
    }
}

func findNewPostionInMatrix(xPos int, yPos int, size int) [2]int {
    var location [2]int
    location[0] = yPos
    location[1] = (-1 * xPos) + (size-1)
    return location
}

