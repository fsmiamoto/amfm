package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/fsmiamoto/amfm/modulator"
	"github.com/fsmiamoto/amfm/signal"
	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/keyboard"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/termbox"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgets/linechart"
)

const title = "AM FM"
const redrawInterval = 250 * time.Millisecond

func errorAndExit(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

const numOfSamples = 10000

// playLineChart continuously adds values to the LineChart, once every delay.
// Exits when the context expires.
func playLineChart(ctx context.Context, lc *linechart.LineChart, delay time.Duration) {
	am := modulator.AM{Sensitivity: 2}
	carrier := signal.Cos(100, numOfSamples).Scale(3)
	message := signal.Sin(10, numOfSamples).Scale(1)
	modulated, _ := am.Modulate(carrier, message)

	ticker := time.NewTicker(delay)
	defer ticker.Stop()
	for i := 0; ; {
		select {
		case <-ticker.C:
			i = (i + 1) % len(modulated)
			rotated := append(modulated[i:], modulated[:i]...)
			if err := lc.Series("modulated", rotated,
				linechart.SeriesCellOpts(cell.FgColor(cell.ColorNumber(33))),
				linechart.SeriesXLabels(map[int]string{
					0: "zero",
				}),
			); err != nil {
				panic(err)
			}
			rotated2 := append(message[i:], message[:i]...)
			if err := lc.Series("message", rotated2,
				linechart.SeriesCellOpts(cell.FgColor(cell.ColorWhite)),
			); err != nil {
				panic(err)
			}
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	t, err := termbox.New()
	if err != nil {
		errorAndExit(err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	quitter := func(k *terminalapi.Keyboard) {
		if k.Key == 'q' || k.Key == 'Q' || k.Key == keyboard.KeyCtrlC {
			cancel()
		}
	}

	lc, err := linechart.New(
		linechart.AxesCellOpts(cell.FgColor(cell.ColorRed)),
		linechart.YLabelCellOpts(cell.FgColor(cell.ColorGreen)),
		linechart.XLabelCellOpts(cell.FgColor(cell.ColorCyan)),
	)
	if err != nil {
		panic(err)
	}
	go playLineChart(ctx, lc, redrawInterval/3)

	c, err := container.New(
		t,
		container.Border(linestyle.Light),
		container.BorderTitle(title),
		container.PlaceWidget(lc),
	)

	if err := termdash.Run(ctx, t, c, termdash.KeyboardSubscriber(quitter), termdash.RedrawInterval(redrawInterval)); err != nil {
		errorAndExit(err)
	}
}
