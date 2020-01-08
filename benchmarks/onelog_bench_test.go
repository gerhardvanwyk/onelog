package benchmarks

import (
	"github.com/gerhardvanwyk/onelogplus"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestPrint(t *testing.T) {
	logger := onelogplus.New(os.Stdout, onelogplus.FINEST).
		Hook(func(e onelogplus.Entry) {
			e.Int64("time", time.Now().Unix())
		})
	logger.InfoWith("message").
		String("test", "test").
		String("test", "test").
		String("test", "test").
		String("test", "test").
		String("test", "test").
		String("test", "test").
		String("test", "test").
		Write()
}

func BenchmarkOnelog(b *testing.B) {
	b.Run("with-fields", func(b *testing.B) {
		logger := onelogplus.New(ioutil.Discard, onelogplus.ALL).
			Hook(func(e onelogplus.Entry) {
				e.Int64("time", time.Now().Unix())
			})
		s := struct {
			i int
		}{i: 0}
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.InfoWithFields("message", func(e onelogplus.Entry) {
					e.Int("test", s.i)
					e.String("test", "test")
					e.String("test", "test")
					e.String("test", "test")
					e.String("test", "test")
					e.String("test", "test")
					e.String("test", "test")
					e.String("test", "test")
				})
			}
		})
	})
	b.Run("message-only", func(b *testing.B) {
		logger := onelogplus.New(ioutil.Discard, onelogplus.ALL).
			Hook(func(e onelogplus.Entry) {
				e.Int64("time", time.Now().Unix())
			})
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info("message")
			}
		})
	})
	b.Run("entry-message-only", func(b *testing.B) {
		logger := onelogplus.New(ioutil.Discard, onelogplus.ALL).
			Hook(func(e onelogplus.Entry) {
				e.Int64("time", time.Now().Unix())
			})
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.InfoWith("message").Write()
			}
		})
	})
	b.Run("entry-fields", func(b *testing.B) {
		logger := onelogplus.New(ioutil.Discard, onelogplus.ALL).
			Hook(func(e onelogplus.Entry) {
				e.Int64("time", time.Now().Unix())
			})
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.InfoWith("message").
					String("test", "test").
					String("test", "test").
					String("test", "test").
					String("test", "test").
					String("test", "test").
					String("test", "test").
					String("test", "test").
					Write()
			}
		})
	})

	b.Run("accumulated context", func(b *testing.B) {
		logger := onelogplus.New(ioutil.Discard, onelogplus.ALL).
			With(func(e onelogplus.Entry) {
				e.Int("int", 1)
			})
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info("message")
			}
		})
	})
}
