package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "my-key", "pollo con arroz")
	viewContext(ctx)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go myProcess(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("Se a exedido la  deadline, ya tu sabe")
		fmt.Println(ctx.Err())
	}
}

func viewContext(ctx context.Context) {
	fmt.Printf("My valor es %s\n", ctx.Value("my-key"))
}

func myProcess(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Santos Bacalaos, Se acabo el tiempo(time out)")
			return
		default:
			fmt.Println("aun en el proceso")
		}
		time.Sleep(500 * time.Millisecond)
	}
}
