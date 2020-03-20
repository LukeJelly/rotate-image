# rotate-image
Prompt: You are given an n x n 2D matrix that represents an image. Rotate 
 the image by 90 degrees (clockwise).

 This was given to me as an example coding interview question. I solved it 
 a different way but couldn't shake how much this seemed like it needed a 
 linearly algebra approach. I solved it using a 2x2 rotation matrix to find
 the new location for every given point. 
  
  Give any point in a Matrix (ex, (0,0)), if we multiply that point by
  a rotation matrix for -90 degrees, we get a new point (0,0). Of course this
  rotates around the origin, and we want to rotate in place so we have to shift 
  that up by the size of the matrix - 1.  So our next location because (0,0 + (len(a)-1)) = (0,4).
  Which is where the (0,0) item would go after the rotation.  We then pull out what is 
  currently in that location and put the a there.  We can then take our point
  (0,4) and put it back through the rotation matrix to find where the next point
  will be.  This will happen 4 times before we get back to where we start.   
  Now that we have defined how to use the rotation matrix, we can abstract that 
  into rotateThroughFour.  

  If we loop through all values in the first row, minus
  the last index, we have affectively rotated the entire outer "ring".  We can 
  now move into the inner ring and do the same thing
  over again.  Keep rotating all outer rings until we there are no more rings.

 
 Version 1.0
 Luke.Kelly
