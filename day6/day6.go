package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"os"
	"runtime"
	"sort"
	"strings"
	"sync"

	"github.com/fogleman/gg"
)

type frameData struct {
	index int
	frame *image.Paletted
}

type renderData struct {
	index int
	data  string
}

func renderFrame(block string, width, height int, fontSize float64) *image.Paletted {
	// Create a new image
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	white := color.RGBA{255, 255, 255, 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{white}, image.Point{}, draw.Src)

	// Draw the text block on the image
	dc := gg.NewContextForRGBA(img)
	dc.SetRGB(0, 0, 0)                     // Black color for text
	dc.LoadFontFace("arial.ttf", fontSize) // Ensure you have this font file
	lines := strings.Split(block, "\n")
	y := fontSize
	for _, line := range lines {
		dc.DrawString(line, 10, y)
		y += fontSize + 2 // Line spacing
	}

	// Render the text to the image
	dc.Fill()

	// Convert to paletted image for GIF
	palettedImg := image.NewPaletted(img.Bounds(), color.Palette{white, color.Black})
	draw.FloydSteinberg.Draw(palettedImg, img.Bounds(), img, image.Point{})

	return palettedImg
}

func worker(jobs <-chan renderData, results chan<- frameData, width, height int, fontSize float64, progressChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		frame := renderFrame(job.data, width, height, fontSize)
		results <- frameData{index: job.index, frame: frame}
		progressChan <- 1 // Update progress
	}
}

type Position struct {
	y int
	x int
}

type PathTile struct {
	pos Position
	dir int
}

func (p *PathTile) isContained(l []PathTile) bool {
	for _, pl := range l {
		if pl.pos.x == p.pos.x && pl.pos.y == p.pos.y && pl.dir == p.dir {
			return true
		}
	}
	return false
}

type Game struct {
	Grid          [][]string
	max_x         int
	max_y         int
	position      Position
	direction     int // 0 up | 1 right | 2 down | 3 left
	start         Position
	steps         int
	render_buffer []string
	path          []PathTile
	obstacles     []Position
}

func (g *Game) init(path string) {

	readFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var line []string
	var a = 0
	for fileScanner.Scan() {
		line = strings.Split(fileScanner.Text(), "")
		g.Grid = append(g.Grid, line)
		for b, letter := range line {
			switch letter {
			case "^":
				g.start = Position{a, b}
				g.direction = 0
			case ">":
				g.start = Position{a, b}
				g.direction = 1
			case "v":
				g.start = Position{a, b}
				g.direction = 2
			case "<":
				g.start = Position{a, b}
				g.direction = 3
			}
		}
		a++
	}
	readFile.Close()
	g.max_x = len(g.Grid[0]) - 1
	g.max_y = len(g.Grid) - 1
	g.position = g.start
}

func (g *Game) forward(trace bool) bool {
	switch g.direction {
	case 0:
		if g.position.y <= 0 {
			return false
		}
		if g.Grid[g.position.y-1][g.position.x] == "#" {
			g.direction = 1
			return true
		} else {
			if trace {
				g.Grid[g.position.y][g.position.x] = "|"
			}
			g.position.y -= 1
			g.path = append(g.path, PathTile{g.position, g.direction})
			g.steps += 1
			return true
		}
	case 1:
		if g.position.x >= g.max_x {
			return false
		}
		if g.Grid[g.position.y][g.position.x+1] == "#" {
			g.direction = 2
			return true
		} else {
			if trace {
				g.Grid[g.position.y][g.position.x] = "-"
			}
			g.position.x += 1
			g.path = append(g.path, PathTile{g.position, g.direction})
			g.steps += 1
			return true
		}
	case 2:
		if g.position.y >= g.max_y {
			return false
		}
		if g.Grid[g.position.y+1][g.position.x] == "#" {
			g.direction = 3
			return true
		} else {
			if trace {
				g.Grid[g.position.y][g.position.x] = "|"
			}
			g.position.y += 1
			g.path = append(g.path, PathTile{g.position, g.direction})
			g.steps += 1
			return true
		}
	case 3:
		if g.position.x <= 0 {
			return false
		}
		if g.Grid[g.position.y][g.position.x-1] == "#" {
			g.direction = 0
			return true
		} else {
			if trace {
				g.Grid[g.position.y][g.position.x] = "-"
			}
			g.position.x -= 1
			g.path = append(g.path, PathTile{g.position, g.direction})
			g.steps += 1
			return true
		}
	default:
		panic("This case should not be triggered")
	}
}

func (g *Game) detectLoop(trace bool) (next, loop bool) {
	var pt PathTile
	switch g.direction {
	case 0:
		if g.position.y <= 0 {
			return false, false
		}
		if g.Grid[g.position.y-1][g.position.x] == "#" {
			g.direction = 1
			return true, false
		} else {
			if trace {
				g.Grid[g.position.y][g.position.x] = "|"
			}
			g.position.y -= 1
			pt = PathTile{g.position, g.direction}
			if pt.isContained(g.path) {
				return false, true
			}
			g.path = append(g.path, pt)
			g.steps += 1
			return true, false
		}
	case 1:
		if g.position.x >= g.max_x {
			return false, false
		}
		if g.Grid[g.position.y][g.position.x+1] == "#" {
			g.direction = 2
			return true, false
		} else {
			if trace {
				g.Grid[g.position.y][g.position.x] = "-"
			}
			g.position.x += 1
			pt = PathTile{g.position, g.direction}
			if pt.isContained(g.path) {
				return false, true
			}
			g.path = append(g.path, pt)
			g.steps += 1
			return true, false
		}
	case 2:
		if g.position.y >= g.max_y {
			return false, false
		}
		if g.Grid[g.position.y+1][g.position.x] == "#" {
			g.direction = 3
			return true, false
		} else {
			if trace {
				g.Grid[g.position.y][g.position.x] = "|"
			}
			g.position.y += 1
			pt = PathTile{g.position, g.direction}
			if pt.isContained(g.path) {
				return false, true
			}
			g.path = append(g.path, pt)
			g.steps += 1
			return true, false
		}
	case 3:
		if g.position.x <= 0 {
			return false, false
		}
		if g.Grid[g.position.y][g.position.x-1] == "#" {
			g.direction = 0
			return true, false
		} else {
			if trace {
				g.Grid[g.position.y][g.position.x] = "-"
			}
			g.position.x -= 1
			pt = PathTile{g.position, g.direction}
			if pt.isContained(g.path) {
				return false, true
			}
			g.path = append(g.path, pt)
			g.steps += 1
			return true, false
		}
	default:
		panic("This case should not be triggered")
	}
}

func (g *Game) generateGif(path string) {
	// Settings
	width := 930
	height := 1040
	fontSize := 6.0
	totalFrames := len(g.render_buffer) / 10
	numCores := runtime.NumCPU()

	// Create channels
	jobs := make(chan renderData, totalFrames)
	results := make(chan frameData, totalFrames)
	progressChan := make(chan int, totalFrames)

	// Add frames to the job queue
	go func() {
		for i, block := range g.render_buffer {
			if i%10 == 0 {
				jobs <- renderData{index: i, data: block}
			}
		}
		close(jobs)
	}()

	// Launch workers
	var wg sync.WaitGroup
	for i := 0; i < numCores*2; i++ {
		wg.Add(1)
		go worker(jobs, results, width, height, fontSize, progressChan, &wg)
	}

	// Progress display
	go func() {
		completed := 0
		for progress := range progressChan {
			completed += progress
			percentage := float64(completed) / float64(totalFrames) * 100
			fmt.Printf("\rRendering frames: %d/%d (%.2f%%)", completed, totalFrames, percentage)
		}
		fmt.Println()
	}()

	// Wait for workers to finish
	go func() {
		wg.Wait()
		close(results)
		close(progressChan)
	}()

	// Collect and sort results
	var frames []frameData
	for frame := range results {
		frames = append(frames, frame)
	}
	fmt.Println()

	sort.Slice(frames, func(i, j int) bool {
		return frames[i].index < frames[j].index
	})

	// Assemble frames into the GIF
	outGIF := &gif.GIF{}
	for a, frame := range frames {
		outGIF.Image = append(outGIF.Image, frame.frame)
		outGIF.Delay = append(outGIF.Delay, 1) // Delay in 100ths of a second
		percentage := float64(a) / float64(totalFrames) * 100
		fmt.Printf("\rStitching frames: %d/%d (%.2f%%)", a, totalFrames, percentage)
	}
	fmt.Printf("\nsaving Gif to: %s", path)

	// Save the GIF to a file
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = gif.EncodeAll(f, outGIF)
	if err != nil {
		panic(err)
	}
}

func (g *Game) render() {
	var str = ""
	for a, line := range g.Grid {
		if a == g.position.y {
			for b, char := range line {
				if b == g.position.x {
					str += "O"
				} else {
					str += char
				}
			}
		} else {
			str += strings.Join(line, "")
		}
		str += "\n"
	}
	g.render_buffer = append(g.render_buffer, str)
}

func main() {
	var game = Game{}
	game.init("C:\\Users\\julia\\Documents\\projects\\advent-of-code-24\\day6")
	var retry = true
	for retry {
		retry = game.forward(true)
		game.render()
	}
	fmt.Println(game.steps)
	//game.generateGif("day6/visu.gif")
	fmt.Println(game.path)

	var tmp string
	var ret = true
	var loop bool
	for _, pos := range game.path {
		tmp = game.Grid[pos.pos.y][pos.pos.x]
		game.Grid[pos.pos.y][pos.pos.x] = "#"
		for ret {
			ret, loop = game.detectLoop(false)
			if loop {
				game.obstacles = append(game.obstacles, pos.pos)
			}
		}
		game.Grid[pos.pos.y][pos.pos.x] = tmp

	}
	fmt.Println(len(game.obstacles))

}
