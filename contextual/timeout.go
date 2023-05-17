package contextual

import (
	"context"
	"fmt"
	"time"
)

func StartWithTimeout() {
	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
	defer cancel()

	saySomething(ctx)
}

func saySomething(ctx context.Context) {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Time after seconds worked")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("context done: ", err)

	}
}
