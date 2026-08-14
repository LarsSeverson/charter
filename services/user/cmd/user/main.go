package user

// import (
// 	"context"
// 	"log"
// 	"os"
// 	"os/signal"
// 	"syscall"

// 	"github.com/LarsSeverson/charter/services/user/internal/application/command"
// 	"github.com/LarsSeverson/charter/services/user/internal/infrastructure/persistence/postgres"
// )

func main() {
	// ctx, stop := signal.NotifyContext(
	// 	context.Background(),
	// 	syscall.SIGINT,
	// 	syscall.SIGTERM,
	// )
	// defer stop()

	// cfg, err := config.Load()
	// if err != nil {
	// 	log.Printf("load configuration: %v", err)
	// 	os.Exit(1)
	// }

	// app, err := bootstrap.New(ctx, cfg)
	// if err != nil {
	// 	log.Printf("initialize application: %v", err)
	// 	os.Exit(1)
	// }
	// defer app.Close()

	// if err := app.Run(ctx); err != nil {
	// 	log.Printf("run application: %v", err)
	// 	os.Exit(1)
	// }
}
