package basic

import (
	"bufio"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"testing"
	"time"
)

// 对迷宫游戏中的方向进行抽象（Direction abstraction）
type Direction int

const (
	N            int = 0
	NE               = 1
	E                = 2
	SE               = 3
	S                = 4
	SW               = 5
	W                = 6
	NW               = 7
	NotAvailable     = 8
)

func (d Direction) String() string {
	switch d {
	case 0:
		return "north"
	case NE:
		return "north-east"
	case E:
		return "east"
	case SE:
		return "south-east"
	case S:
		return "south"
	case SW:
		return "south-west"
	case W:
		return "west"
	case NW:
		return "north-west"
	case NotAvailable:
		return "not available"
	}
	return "unknown"
}
func (d Direction) PrintDirection() {
	fmt.Println("direction: ", d)
}

// Point abstraction
type Point struct {
	x, y int
}

var None = Point{-1, -1}

func (p Point) Equals(other Point) bool {
	return p.x == other.x && p.y == other.y
}
func (p Point) PrintPoint() {
	fmt.Printf("<%d, %d>\n", p.x, p.y)
}

// Path abstraction
type Path struct {
	point          Point
	moveDirection  Direction
	movesAvailable []Direction
}

func NewPath(point Point, origin Direction) Path {
	if origin < 0 {
		return Path{point: point,
			moveDirection: NotAvailable,
			//初始化为所有方向都可以获得,但是来自的点的方向不可获得
			movesAvailable: []Direction{0, NE, E, SE, S, SW, W, NW},
		}
	}
	var movesDirAvailable []Direction
	switch origin {
	case Direction(N):
		movesDirAvailable = []Direction{0, NE, E, SE, NotAvailable, SW, W, NW}
	case NE:
		movesDirAvailable = []Direction{0, NE, E, SE, S, NotAvailable, W, NW}
	case E:
		movesDirAvailable = []Direction{0, NE, E, SE, S, SW, NotAvailable, NW}
	case SE:
		movesDirAvailable = []Direction{0, NE, E, SE, S, SW, W, NotAvailable}
	case S:
		movesDirAvailable = []Direction{8, NE, E, SE, S, SW, W, NW}
	case SW:
		movesDirAvailable = []Direction{0, NotAvailable, E, SE, S, SW, W, NW}
	case W:
		movesDirAvailable = []Direction{0, NE, NotAvailable, SE, S, SW, W, NW}
	case NW:
		movesDirAvailable = []Direction{0, NE, E, NotAvailable, S, SW, W, NW}
	}
	return Path{point: point,
		moveDirection: NotAvailable,
		//初始化为所有方向都可以获得,但是来自的点的方向不可获得
		movesAvailable: movesDirAvailable,
	}
}

// RandomMove在当前可用的方向中随机选择一个可用的方向，如果没有可用方向，则返回
// 特殊的方向NotAvailable。
func (path *Path) RandomMove() Direction {
	//可得到的移动方向的序号集合，以便随机选取一个方向
	indicesAvailable := []int{}
	//任何一个path都有可以移动的方向列表（初始化为全部方向），查找该列表中不是NotAvailable
	//的方向，添加其index到可选择方向的序号列表中。
	for i := 0; i < len(path.movesAvailable); i++ {
		if path.movesAvailable[i] != NotAvailable {
			indicesAvailable = append(indicesAvailable, i)
		}
	}
	//在path当前可选择方向的序号列表中随机选择一个序号（也就是一个方向）
	// 然后将Path的移动方向（moveDirection）设置为该方向
	//并将该方向（移动过的方形）设置为不可获得的方向，避免回退到该path时再重新走该方向。
	count := len(indicesAvailable)
	//fmt.Printf("点%v的可用方向数量为%v,包括：%v\n", path.point, count, path.movesAvailable)
	if count > 0 {
		randomIndex := rand.IntN(count)
		indexAvailable := indicesAvailable[randomIndex]
		path.moveDirection = path.movesAvailable[indexAvailable]
		path.movesAvailable[indexAvailable] = NotAvailable //走过的方向不能重走，设置为NotAvailable
		return path.moveDirection
	} else {
		return NotAvailable
	}

}

func TestPath(t *testing.T) {

	myDirection := Direction(6)
	myDirection.PrintDirection()
	myPoint := Point{3, 4}
	myPoint.PrintPoint()
	result := myPoint.Equals(Point{3, 4})
	fmt.Println(result)

	myPath := NewPath(Point{3, 4}, Direction(N))
	randomMove := myPath.RandomMove()
	fmt.Println(randomMove)
	fmt.Println(myPath)
}

type Maze struct {
	rows, cols int
	start, end Point
	mazefile   string
	barriers   [][]bool //表示Point是否阻塞，true表示阻塞（无法进入），false表示可进入
	current    Path
	movecoutn  int
	pathStack  SliceStackAny[Path]
	gameover   bool
}

// NewMaze按照给定迷宫矩阵的行列数，启点、终点和迷宫矩阵中各位置点的阻塞\畅通的配置文件来初始化一个迷宫
func NewMaze(rows, cols int, start, end Point, mazeFile string) (maze Maze) {
	maze.rows = rows
	maze.cols = cols
	maze.start = start
	maze.end = end
	maze.mazefile = mazeFile
	//从文件初始化barriers
	maze.barriers = make([][]bool, rows)
	for i := 0; i < rows; i++ {
		maze.barriers[i] = make([]bool, cols)
	}
	file, err := os.Open(maze.mazefile)
	if err != nil {
		log.Fatal(err) //相当于print完日志后代用os.Exit()	}
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var textLines []string
	//Scan方法扫描（把读取的字节放在内部变量中，由其他方法返回，比如Text方法）
	//到给定分割方法所设定的位置，这里是行分割方法，因此，是按行扫描。
	for scanner.Scan() {
		//Text方法返回Scan方法所扫描到的文本
		textLines = append(textLines, scanner.Text())
	}
	for row := 0; row < rows; row++ {
		line := textLines[row]
		for col := 0; col < cols; col++ {
			s := string(line[col])
			maze.barriers[row][col] = (s == "1")
		}
	}
	maze.current = NewPath(start, -1)
	maze.pathStack = SliceStackAny[Path]{}
	maze.pathStack.Push(maze.current)
	maze.barriers[start.x][start.y] = true //起点设置为阻塞
	return maze
}

// NewPostion 根据现有的点（oldPosition）和移动方向（move），得到移动的目标点,也就是函数的返回值
func NewPostion(oldPosition Point, move Direction) Point {
	switch move {
	case Direction(N):
		return Point{x: oldPosition.x, y: oldPosition.y - 1}
	case NE:
		return Point{x: oldPosition.x + 1, y: oldPosition.y - 1}
	case E:
		return Point{x: oldPosition.x + 1, y: oldPosition.y}
	case SE:
		return Point{x: oldPosition.x + 1, y: oldPosition.y + 1}
	case S:
		return Point{x: oldPosition.x, y: oldPosition.y + 1}
	case SW:
		return Point{x: oldPosition.x - 1, y: oldPosition.y + 1}
	case W:
		return Point{x: oldPosition.x - 1, y: oldPosition.y}
	case NW:
		return Point{x: oldPosition.x - 1, y: oldPosition.y - 1}
	default:
		panic("error move")
	}
}

// StepAhead向前前进一步，随机选择可获得的方向前进，
// 如果存在新位置，就前进到该位置，并且形成新的Path，
// 如果没有可以前进的新位置，就回退到出发的位置。
func (m *Maze) StepAhead() (Point, Point) {
	validMove := false     //表示是否已经找到了下一个移动路径
	backTracePoint := None //表示回退点。
	newPos := None
	//下面for循环是弹出栈中的当前路径，为其寻找下一个前进路径，如果当前路径无法找到前进路径，就再弹出路径
	for {

		if m.gameover || validMove || m.pathStack.IsEmpty() {
			break
		}
		validMove = false             //表示是否找到了下一个移动路径
		m.current = m.pathStack.Pop() //弹出当前的路径
		m.movecoutn += 1
		nextMove := m.current.RandomMove() //随机选择下一个移动方向
		//下面for循环是从当前的点随机寻找可行方向（未尝试过的方向），如果该方向不阻塞，则作为下一个路径。
		// 如果找到下一个路径（即：validMove==true），就压栈，并通知外层循环结束寻找。
		// 如果循环尝试了所有可能的方向（最多8个），都找不到不阻塞的点，那就以下一个路径的方向无法获得来结束循环
		//,即：nextMove==NotAvailable，
		for {
			if validMove || nextMove == NotAvailable {
				break //如果找不到下一个可移动方向，或者移动到了非正确的位置都停止寻找下一个方向的循环
			}
			//确定新位置点
			newPos = NewPostion(m.current.point, nextMove)
			if nextMove != m.current.moveDirection {
				panic("我理解的程序逻辑理解错误！")
			}
			if m.barriers[newPos.x][newPos.y] == true {
				nextMove = m.current.RandomMove() //新位置是死路，那么就随机再选方向
				continue
			}
			//如果新位置不是死路，那么说明本次移动是正确的。
			validMove = true //找到了下一个位置
			if newPos.Equals(m.end) {
				for {
					if m.pathStack.IsEmpty() {
						break
					}
					m.pathStack.Pop()
				}
				m.gameover = true
			}
			//m.barriers[m.current.point.x][m.current.point.y] = true
			m.pathStack.Push(m.current)                         //当前路径压栈
			newPath := NewPath(newPos, m.current.moveDirection) //创建新路径
			m.pathStack.Push(newPath)                           //新路径压栈
			break
		} //for循环结束从当前点成功找下一个前进方向，或确认没有可前进的方向就循环结束。
		//如果确认没有可前进的方向，就退回到上一个路径重新寻找下一个可行路径
		if !validMove && !m.pathStack.IsEmpty() {
			fmt.Printf("从%v回退到%v\n", m.current.point, m.pathStack.Top().point)
			backTracePoint = m.pathStack.Top().point
		}
	} //for 循环结束

	if m.pathStack.IsEmpty() {
		fmt.Printf("从位置%v处找不到可以前进的路径\n", m.current.point)
		return None, None
	}
	return newPos, backTracePoint
}
func TestMaze(t *testing.T) {
	t.Deadline()
	start := Point{1, 1}
	end := Point{38, 38}
	maze := NewMaze(40, 40, start, end, "maze.txt")
	fmt.Printf("start(1,1) is :%v,end(38,38) is : %v\n", maze.barriers[1][1], maze.barriers[38][38])
	newPos, _ := maze.StepAhead()
	time.Sleep(1 * time.Second)
	if newPos != None {
		fmt.Println("前进到", newPos)
	}
	for {
		if newPos == None || newPos.Equals(end) {
			break
		}
		newPos, _ = maze.StepAhead()
		time.Sleep(10 * time.Millisecond)
		if newPos != None {
			fmt.Printf("前进到点:%v\n", newPos)
		}
	}
	if newPos.Equals(end) {
		fmt.Printf("在走了%v步之后，成功找到出口点%v ", maze.movecoutn, end)
	}
}
