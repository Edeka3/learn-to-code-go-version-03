package main

import (
	"fmt"
	"unsafe"
)

func main() {
	example1()
	example2()
	example3()
	example4()
}

func example1() {
	type Example struct {
		A bool    // 1 byte
		B int16   // 2 bytes
		C float32 // 4 bytes
	}
	// padding: 1 byte for A

	a := Example{}
	fmt.Println("EXAMPLE 1")
	fmt.Println("Alignment of bool:", unsafe.Alignof(bool(true)))
	fmt.Println("Alignment of int32:", unsafe.Alignof(int16(0)))
	fmt.Println("Alignment of float64:", unsafe.Alignof(float32(0.0)))
	fmt.Println("Alignment of Example:", unsafe.Alignof(a))
	fmt.Println("Size of Example:", unsafe.Sizeof(a))
	fmt.Println()
}

func example2() {
	type Example struct {
		A bool    // 1 byte
		B int32   // 4 bytes
		C float32 // 4 bytes
	}

	a := Example{}
	fmt.Println("EXAMPLE 2")
	fmt.Println("Alignment of bool:", unsafe.Alignof(bool(true)))
	fmt.Println("Alignment of int32:", unsafe.Alignof(int32(0)))
	fmt.Println("Alignment of float64:", unsafe.Alignof(float32(0.0)))
	fmt.Println("Alignment of Example:", unsafe.Alignof(a))
	fmt.Println("Size of Example:", unsafe.Sizeof(a))
	fmt.Println()
}

func example3() {
	type Example struct {
		A bool    // 1 byte
		B int32   // 4 bytes
		C float64 // 8 bytes
	}
	// Padding: 3 bytes for A

	a := Example{}
	fmt.Println("EXAMPLE 3")
	fmt.Println("Alignment of bool:", unsafe.Alignof(bool(true)))
	fmt.Println("Alignment of int32:", unsafe.Alignof(int32(0)))
	fmt.Println("Alignment of float64:", unsafe.Alignof(float64(0.0)))
	fmt.Println("Alignment of Example:", unsafe.Alignof(a))
	fmt.Println("Size of Example:", unsafe.Sizeof(a))
	fmt.Println()
}

func example4() {
	// Without optimization
	type Unoptimized struct {
		A bool    // 1 byte
		B float64 // 8 bytes
		C int32   // 4 bytes
		// Padding: 7 bytes for A, 4 bytes for C (assuming 64-bit machine)
	}
	a := Unoptimized{}
	fmt.Println("EXAMPLE 4 - Unoptimized")
	fmt.Println("Alignment of bool:", unsafe.Alignof(bool(true)))
	fmt.Println("Alignment of float64:", unsafe.Alignof(float64(0.0)))
	fmt.Println("Alignment of int32:", unsafe.Alignof(int32(0)))
	fmt.Println("Alignment of Unoptimized:", unsafe.Alignof(a))
	fmt.Println("Size of Unoptimized:", unsafe.Sizeof(a))
	fmt.Println()

	// With optimization
	type Optimized struct {
		B float64 // 8 bytes
		C int32   // 4 bytes
		A bool    // 1 byte
		// Padding: None or minimal (assuming 64-bit machine)
	}
	b := Optimized{}
	fmt.Println("EXAMPLE 4 - Optimized")
	fmt.Println("Alignment of float64:", unsafe.Alignof(float64(0.0)))
	fmt.Println("Alignment of int32:", unsafe.Alignof(int32(0)))
	fmt.Println("Alignment of bool:", unsafe.Alignof(bool(true)))
	fmt.Println("Alignment of Optimized:", unsafe.Alignof(b))
	fmt.Println("Size of Optimized:", unsafe.Sizeof(b))
	fmt.Println()
}
