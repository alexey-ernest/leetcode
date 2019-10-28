import "fmt"

/**
 * // This is the robot's control interface.
 * // You should not implement it, or speculate about its implementation
 * type Robot struct {
 * }
 * 
 * // Returns true if the cell in front is open and robot moves into the cell.
 * // Returns false if the cell in front is blocked and robot stays in the current cell.
 * func (robot *Robot) Move() bool {}
 *
 * // Robot will stay in the same cell after calling TurnLeft/TurnRight.
 * // Each turn will be 90 degrees.
 * func (robot *Robot) TurnLeft() {}
 * func (robot *Robot) TurnRight() {}
 *
 * // Clean the current cell.
 * func (robot *Robot) Clean() {}
 */

var (
    up = [2]int{-1,0}
    right = [2]int{0,1}
    down = [2]int{1,0}
    left = [2]int{0,-1}
)

func cleanRoom(robot *Robot) {
    // clean initial cell
    roombaClean(robot, nil, up, 0, 0)
}

func roombaClean(robot *Robot, visited map[int]map[int]bool, vec [2]int, i,j int) [2]int {
    
    // first call
    isFirstCall := false
    if visited == nil {
        isFirstCall = true
        visited = make(map[int]map[int]bool)
    }
    
    //fmt.Printf("%d, %d\n", 1+i, 3+j)
    // mark current cell as visited
    if visited[i] == nil {
        visited[i] = make(map[int]bool)
    }
    visited[i][j] = true
    
    // clean current cell
    robot.Clean()
    
    dirs := [][2]int{
        up,
        right,
        down,
        left,
    }
    lastdir := vec
    for _, d := range dirs {
        if !isFirstCall &&
           vec == up && d == down ||
           vec == down && d == up ||
           vec == right && d == left ||
           vec == left && d == right {
               // not going back
            continue
        }
        
        ia := i + d[0]
        ja := j + d[1]
        if visited[ia] != nil && visited[ia][ja] {
            continue
        }
        
        move := moveTo(robot, lastdir, d)
        lastdir = d
        if move {
            lastdir = roombaClean(robot, visited, lastdir, ia, ja)
        }
    }
    
    // going back
    dirback := [2]int{-vec[0],-vec[1]}
    moveTo(robot, lastdir, dirback)
    
    return dirback
}

func moveTo(robot *Robot, vec [2]int, dir [2]int) bool {
    
    // move
    if vec == up {
        // up
        if dir == right {
            robot.TurnRight()
        } else if dir == left {
            robot.TurnLeft()    
        } else if dir == down {
            robot.TurnRight()
            robot.TurnRight()
        }
    } else if vec == down {
        // down
        if dir == right {
            robot.TurnLeft()
        } else if dir == left {
            robot.TurnRight()
        } else if dir == up {
            robot.TurnRight()
            robot.TurnRight()
        }
    } else if vec == right {
        // right
        if dir == up {
            robot.TurnLeft()
        } else if dir == down {
            robot.TurnRight()
        } else if dir == left {
            robot.TurnRight()
            robot.TurnRight()
        }
    } else if vec == left {
        // left
        if dir == up {
            robot.TurnRight()
        } else if dir == down {
            robot.TurnLeft()
        } else if dir == right {
            robot.TurnRight()
            robot.TurnRight()
        }   
    } else {
        panic("invalid direction vec")
    }
    
    return robot.Move()
}
