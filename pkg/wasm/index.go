package wasm

// func New(wasmFilepath string) *wasmer.Instance {
// 	wasmBytes, _ := os.ReadFile(wasmFilepath)

// 	engine := wasmer.NewEngine()
// 	store := wasmer.NewStore(engine)

// 	// Compiles the module
// 	module, _ := wasmer.NewModule(store, wasmBytes)

// 	// Instantiates the module
// 	importObject := wasmer.NewImportObject()
// 	instance, _ := wasmer.NewInstance(module, importObject)

// 	return instance
// }

// func Call(instance *wasmer.Instance, funcName string, args ...interface{}) any {

// 	funcInstance, _ := instance.Exports.GetFunction(funcName)
// 	result, _ := funcInstance(args...)

// 	return result
// }
