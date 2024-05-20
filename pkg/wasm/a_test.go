package wasm_test

import (
	"context"
	_ "embed"
	"log"
	"testing"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

//go:embed testdata/aster_bg.wasm
var addWasm []byte

func TestXxx(t *testing.T) {

	ctx := context.Background()

	// Create a new WebAssembly Runtime.
	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx) // This closes everything this Runtime created.

	// Instantiate WASI, which implements host functions needed for TinyGo to
	// implement `panic`.
	wasi_snapshot_preview1.MustInstantiate(ctx, r)

	// Instantiate the guest Wasm into the same runtime. It exports the `add`
	// function, implemented in WebAssembly.
	mod, err := r.Instantiate(ctx, addWasm)
	if err != nil {
		log.Panicf("failed to instantiate module: %v", err)
	}

	// Call the `add` function and print the results to the console.
	encry := mod.ExportedFunction("encry")

	results, err := encry.Call(ctx, 96)
	if err != nil {
		log.Panicf("failed to call encry: %v", err)
	}

	t.Log(results)
	// result := wasm.Call(instance, "encry", "a")
	// '9Z2cSdY7v1d/2Ntile/5UA=='
	// 报错
	// t.Log(result)

}
