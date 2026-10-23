package main
import "syscall/js"
func aplicar(this js.Value, a []js.Value) any { return a[0].Int()+1 }
func main(){ js.Global().Set("aplicar", js.FuncOf(aplicar)); select{} }
