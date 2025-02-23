package main

import (
	"bytes"
	"fmt"
	"os"
	"time"
)

const (
	NOTHING = 0
	WALL    = 1
	PLAYER  = 69

	MAX_SAMPLES = 100
)

type position struct {
	x, y int
}

type player struct {
	pos   position
	level *level
}

func (p *player) update() {
  p.pos.x += 1
  // p.level.data[p.pos.y][p.pos.x]
}

type stats struct {
	start  time.Time
	frames int
	fps    float64
}

func newStats() *stats {
	return &stats{
		start: time.Now(),
	}
}

func (s *stats) update() {
	s.frames++
	if s.frames == MAX_SAMPLES {
		s.fps = float64(s.frames) / time.Since(s.start).Seconds()
		s.frames = 0
		s.start = time.Now()
	}
}

type level struct {
	width, height int
	data          [][]byte
}

func newLevel(width, height int) *level {
	data := make([][]byte, height)

	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			data[h] = make([]byte, width)
		}
	}

	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			if h == 0 {
				data[h][w] = WALL
			}
			if w == 0 {
				data[h][w] = WALL
			}
			if w == width-1 {
				data[h][w] = WALL
			}
			if h == height-1 {
				data[h][w] = WALL
			}
		}
	}

	return &level{
		width:  width,
		height: height,
		data:   data,
	}
}

type game struct {
	isRunning bool
	level     *level
	stats     *stats
  player    *player

	drawBuf *bytes.Buffer
}

func newGame(width, height int) *game {
	lvl := newLevel(width, height)
	return &game{
		level:   lvl,
		drawBuf: new(bytes.Buffer),
		stats:   newStats(),
    player: &player{
      level: lvl,
      pos: position{x: 2, y: 5},
    },
	}
}

func (g *game) start() {
	g.isRunning = true
	g.loop()
}

func (g *game) loop() {
	for g.isRunning {
		g.update()
		g.render()
		g.stats.update()
		time.Sleep(time.Millisecond * 16) // Limit FPS
	}
}

func (g *game) update() {
  g.player.update()
}

func (g *game) renderPlayer() {
  g.level.data[g.player.pos.y][g.player.pos.x] = PLAYER
}

func (g *game) renderLevel() {
	for h := 0; h < g.level.height; h++ {
		for w := 0; w < g.level.width; w++ {
			if g.level.data[h][w] == NOTHING {
				g.drawBuf.WriteString(" ")
			} else if g.level.data[h][w] == WALL {
				g.drawBuf.WriteString("▢")
			} else if g.level.data[h][w] == PLAYER {
        g.drawBuf.WriteString("P")
      }
		}
		g.drawBuf.WriteString("\n")
	}
}

func (g *game) render() {
	g.drawBuf.Reset() // Xóa buffer trước khi render
	// đây là cách xóa màn hình terminal: https://stackoverflow.com/questions/22891644/how-do-i-clear-the-terminal-screen
	// \033: ký tự escape (thoát) | [2J: lệnh xóa toàn bộ màn hình | [1;1H: lệnh đưa con trỏ về góc trái trên cùng (tọa độ 1, 1)
	fmt.Fprint(os.Stdout, "\033[2J\033[1;1H")

	g.renderLevel()
  g.renderPlayer()
	g.renderStats()

	fmt.Fprint(os.Stdout, g.drawBuf.String())
}

func (g *game) renderStats() {
	g.drawBuf.WriteString("-- STATS\n")
	g.drawBuf.WriteString(fmt.Sprintf("FPS: %.2f\n", g.stats.fps))
}

func main() {
	width := 80
	height := 18
	g := newGame(width, height)
	g.start()
}
