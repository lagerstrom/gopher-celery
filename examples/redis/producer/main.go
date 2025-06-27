// Program producer sends two "myproject.mytask" tasks to "important" queue.
package main

import (
	"os"

	"github.com/go-kit/log"
	celery "github.com/marselester/gopher-celery"
)

func main() {
	logger := log.NewJSONLogger(log.NewSyncWriter(os.Stderr))

	app := celery.NewApp(
		celery.WithLogger(logger),
		celery.WithTaskProtocol(2),
	)
	err := app.Delay("myproject.mytask", "important", "fizz", "bazz")
	logger.Log("msg", "task was sent using protocol v2", "err", err)

	app = celery.NewApp(
		celery.WithLogger(logger),
		celery.WithTaskProtocol(1),
	)
	err = app.Delay("myproject.mytask", "important", "fizz", "bazz")
	logger.Log("msg", "task was sent using protocol v1", "err", err)
}
