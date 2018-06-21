package logger

import (
	"fmt"
	"io"
	"os"
	"time"
	// "github.com/aybabtme/color/brush"
)

var stdout io.Writer = os.Stdout

// Color the level string
type colorLevelString string

func (c colorLevelString) String() (str string) {

	return fmt.Sprintf("%s", string(c))
	/*
		switch c {
		case "FNST":
			str = fmt.Sprintf("%s", brush.DarkGreen(c))
		case "FINE":
			str = fmt.Sprintf("%s", brush.Green(c))
		case "DEBG":
			str = fmt.Sprintf("%s", brush.Purple(c))
		case "TRAC":
			str = fmt.Sprintf("%s", brush.LightGray(c))
		case "INFO":
			str = fmt.Sprintf("%s", brush.Blue(c))
		case "WARN":
			str = fmt.Sprintf("%s", brush.Yellow(c))
		case "EROR":
			str = fmt.Sprintf("%s", brush.Red(c))
		case "CRIT":
			str = fmt.Sprintf("%s", brush.DarkRed(c))
		default:
			str = string(c)
		}
		return
	*/

}

// This is the standard writer that prints to standard output.
type ConsoleLogWriter chan *LogRecord

// This creates a new ConsoleLogWriter
func NewConsoleLogWriter() ConsoleLogWriter {
	records := make(ConsoleLogWriter, LogBufferLength)
	go records.run(stdout)
	return records
}

func (w ConsoleLogWriter) run(out io.Writer) {
	var timestr string
	var timestrAt int64

	for rec := range w {
		if at := rec.Created.UnixNano() / 1e9; at != timestrAt {
			timestr, timestrAt = rec.Created.Format("2006/01/02 15:04:05"), at
		}
		fmt.Fprintf(out, "[%s] [%s] (%s) %s\n",
			timestr,
			colorLevelString(levelStrings[rec.Level]),
			/*fmt.Sprintf("%s", brush.LightGray(rec.Source)),
			fmt.Sprintf("%s", brush.DarkBlue(rec.Message)),*/
			fmt.Sprintf("%s", rec.Source),
			fmt.Sprintf("%s", rec.Message),
		)
	}
}

// This is the ConsoleLogWriter's output method. This will block if the output
// buffer is full.
func (w ConsoleLogWriter) LogWrite(rec *LogRecord) {
	w <- rec
}

// Close stops the logger from sending messages to standard output. Attempts to
// send log messages to this logger after a Close have undefined behavior.
func (w ConsoleLogWriter) Close() {
	close(w)
	time.Sleep(50 * time.Millisecond) // Try to give console I/O time to complete
}
