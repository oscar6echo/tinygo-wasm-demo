package main

func main() {
	println("---START TEST WASM module")
	println("add(2, 3) =", add(2, 3))
	println("multiply(2,3) =", multiply(2, 3))
	println("square(5) =", square(5))
	println("custom description =", customDesc())
	println("custom(10, 5) =", custom(10, 5))
	println("---END TEST")
}

// This function is imported from JavaScript, as it doesn't define a body.
// You should define a function named 'main.add' in the WebAssembly 'env'
// module from JavaScript.
func add(x, y int) int

// These functions are exported to JavaScript, so can be called using
// exports.multiply() etc  in JavaScript.
// only ints and floats are supported

//export multiply
func multiply(x, y int) int {
	return x * y
}

//export square
func square(x int) int {
	return x * x
}

//export custom
func custom(x, y int) float32 {
	var a float32 = float32(add(square(x), multiply(x, y)))
	var b float32 = float32(a) / 100
	// return [2]float32{a, b}
	return b
}

func customDesc() string {
	return "f(x, y) = sum(x^2+x*y)"
}
