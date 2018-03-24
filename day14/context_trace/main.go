package main

import (
	"context"
	"fmt"
)

func add(ctx context.Context, a, b int) int {
	traceId := ctx.Value("trace_id").(string)
	fmt.Printf("add trace_id:%v\n", traceId)
	return a + b
}
func calc(ctx context.Context, a, b int) int {
	traceId := ctx.Value("trace_id").(string)
	fmt.Printf("calc trace_id:%v\n", traceId)
	return add(ctx, a, b)
}
func main() {
	ctx := context.WithValue(context.Background(), "trace_id", "123456")
	calc(ctx, 388, 200)
}
