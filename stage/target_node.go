package stage

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quasilyte/ge"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/sinecord/gamedata"
	"github.com/quasilyte/sinecord/styles"
)

type targetNode struct {
	board      *Board
	pos        gmath.Vec
	instrument gamedata.InstrumentKind
	disposed   bool
	outline    bool
	r          float32
	color      ge.ColorScale
	hp         int
}

func newTargetNode(b *Board, t gamedata.Target) *targetNode {
	var size float64
	hp := 1
	switch t.Size {
	case gamedata.SmallTarget:
		size = 15
	case gamedata.NormalTarget:
		size = 30
	case gamedata.BigTarget:
		size = 60
		hp = 3
	default:
		panic("unexpected target size")
	}

	var colorScale ge.ColorScale
	colorScale.SetColor(styles.TargetColor)
	return &targetNode{
		board:      b,
		pos:        t.Pos,
		instrument: t.Instrument,
		r:          float32(size / 2),
		color:      colorScale,
		hp:         hp,
		outline:    t.Outline,
	}
}

func (n *targetNode) IsDisposed() bool { return n.disposed }

func (n *targetNode) Dispose() { n.disposed = true }

func (n *targetNode) Update(delta float64) {
}

func (n *targetNode) Draw(screen *ebiten.Image) {
	shape := gamedata.InstrumentShape(n.instrument)
	r := float32(n.r)
	if n.outline {
		n.board.canvas.drawShape(screen, shape, float32(n.pos.X), float32(n.pos.Y), r, 0, n.color)
	} else {
		n.board.canvas.drawFilledShape(screen, shape, float32(n.pos.X), float32(n.pos.Y), r, 0, n.color)
	}
}

func (n *targetNode) OnDamage() bool {
	n.hp--
	destroyed := n.hp <= 0
	if destroyed {
		n.Dispose()
	}
	return destroyed
}
