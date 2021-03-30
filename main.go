package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/fsmiamoto/amfm/fft"
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
const redrawInterval = 300 * time.Millisecond

const numOfSamples = 10000

//TODO: Make these adjustable through the UI
var carrier = signal.Cos(200, numOfSamples).Scale(3)
var message = signal.Sin(10, numOfSamples).Scale(1)

var dsb = modulator.NewDSB(1.5)
var modulated = dsb.Modulate(carrier, message)
var spectrum = fft.ForSignal(modulated)

func playFreqDomainChart(ctx context.Context, lc *linechart.LineChart, delay time.Duration) {
	ticker := time.NewTicker(delay)
	defer ticker.Stop()

	// Let's show only the upper part of the spectrum
	middle := len(spectrum) / 2

	//TODO: Convert indices from the spectrum to frequency values
	for {
		select {
		case <-ticker.C:
			if err := lc.Series("spectrum", spectrum[middle:],
				linechart.SeriesCellOpts(cell.FgColor(cell.ColorNumber(33))),
			); err != nil {
				exitWithErr(err)
			}
		case <-ctx.Done():
			return
		}
	}
}

func playTimeDomainChart(ctx context.Context, lc *linechart.LineChart, delay time.Duration) {
	ticker := time.NewTicker(delay)
	defer ticker.Stop()

	for i := 0; ; {
		select {
		case <-ticker.C:
			i = (i + 1) % len(modulated)
			rotatedCarrier := append(modulated[i:], modulated[:i]...)
			if err := lc.Series("modulated", rotatedCarrier,
				linechart.SeriesCellOpts(cell.FgColor(cell.ColorNumber(33))),
			); err != nil {
				exitWithErr(err)
			}
			rotatedMsg := append(message[i:], message[:i]...)
			if err := lc.Series("message", rotatedMsg,
				linechart.SeriesCellOpts(cell.FgColor(cell.ColorWhite)),
			); err != nil {
				exitWithErr(err)
			}
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	t, err := termbox.New()
	if err != nil {
		exitWithErr(err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	quitter := func(k *terminalapi.Keyboard) {
		if k.Key == 'q' || k.Key == 'Q' || k.Key == keyboard.KeyCtrlC {
			cancel()
		}
	}

	timeDomain, err := linechart.New(
		linechart.AxesCellOpts(cell.FgColor(cell.ColorRed)),
		linechart.YLabelCellOpts(cell.FgColor(cell.ColorGreen)),
		linechart.XLabelCellOpts(cell.FgColor(cell.ColorCyan)),
	)
	if err != nil {
		exitWithErr(err)
	}
	go playTimeDomainChart(ctx, timeDomain, redrawInterval/3)

	freqDomain, err := linechart.New(
		linechart.AxesCellOpts(cell.FgColor(cell.ColorRed)),
		linechart.YLabelCellOpts(cell.FgColor(cell.ColorGreen)),
		linechart.XLabelCellOpts(cell.FgColor(cell.ColorCyan)),
	)
	if err != nil {
		exitWithErr(err)
	}
	go playFreqDomainChart(ctx, freqDomain, redrawInterval/3)

	c, err := container.New(
		t,
		container.Border(linestyle.Light),
		container.BorderTitle(title),
		container.SplitHorizontal(
			container.Top(
				container.Border(linestyle.Double),
				container.BorderTitle("Time domain"),
				container.PlaceWidget(timeDomain),
			),
			container.Bottom(
				container.Border(linestyle.Double),
				container.BorderTitle("Frequency domain"),
				container.PlaceWidget(freqDomain),
			),
		),
	)

	if err := termdash.Run(ctx, t, c, termdash.KeyboardSubscriber(quitter), termdash.RedrawInterval(redrawInterval)); err != nil {
		exitWithErr(err)
	}
}

func exitWithErr(err error) {
	fmt.Fprintln(os.Stderr, "error: ", err)
	os.Exit(1)
}
