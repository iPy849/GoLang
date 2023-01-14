package main

import (
	"context"
	"fmt"
)

type aKey string

func searchKey(ctx context.Context, k aKey) {
	v := ctx.Value(k) // NOTE: Lo extrae y limpia el almacenamiento
	if v != nil {
		fmt.Println("found value:", v)
		return
	} else {
		fmt.Println("key not found:", k)
	}
}

func main() {
	myKey := aKey("mySecretValue")
	// NOTE: Guarda un par de valores k,v -- Es el único contexto que hace eso
	ctx := context.WithValue(context.Background(), myKey, "mySecret")
	searchKey(ctx, myKey)

	searchKey(ctx, aKey("notThere"))
	emptyCtx := context.TODO() // NOTE: Implementa la función Value pero no devuelve nada
	searchKey(emptyCtx, aKey("notThere"))
}
