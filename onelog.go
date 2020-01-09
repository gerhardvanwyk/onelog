// Package onelogplus is a fast, low allocation and modular JSON logger.
//
// It uses github.com/francoispqt/gojay as JSON encoder.
//
// Basic usage:
// 	import "github.com/gerhardvanwyk/onelogplus/log"
//
//	log.Info("hello world !") // {"level":"info","message":"hello world !", "time":1494567715}
//
// You can create your own logger:
//	import "github.com/gerhardvanwyk/onelogplus
//
//	var logger = onelog.New(os.Stdout, onelog.ALL)
//
//	func main() {
//		logger.Info("hello world !") // {"level":"info","message":"hello world !"}
//	}
package onelogplus
