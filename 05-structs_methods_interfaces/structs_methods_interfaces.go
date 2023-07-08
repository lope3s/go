package main

import "math"

//func Perimeter(width float64, height float64) float64 {
//	return 2 * (width + height)
//}
//
//func Area(width float64, height float64) float64 {
//	return width * height
//}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

//func Area(rectangle Rectangle) float64 {
//	return rectangle.Width * rectangle.Height
//}

// Struct:
// A struct is just a named collection of fields where you can store data.
type Rectangle struct {
	Width  float64
	Height float64
}

// Methods:
// A method is a function with a receiver. A method declaration binds an
// identifier, the method name, to a method, and associates the method
// with the receiver's base type.
// Methods are called by invoking them on an instance of a particular type.

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Syntax: func (receiverName ReceiverType) MethodName(args)
// It is a convention in Go to have the receiver variable be the first letter of the type.
// When your method is called on a variable of that type, you get your reference to its data
// via the receiverName variable. In other languages this is done implicity and you access the receiver via this.

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
    Base float64
    Height float64
}

func (t Triangle) Area() float64 {
    return (t.Base * t.Height) * 0.5
}

// Interfaces:
// are a very powerful concept in statically typed languages like Go because
// they allow you to make functions that can be used with different types and
// create highly-decoupled code whilst still maintaining type-safety.

type Shape interface {
	Area() float64
}

// Note that:
// Rectangle has a method called Area that returns a float64 so it satisfies the Shape interface.
// Circle has a method called Area that returns a float64 so it satisfies the Shape Interfaces.
// string does not have such a method, so it doesn't satisfy the interface.

// In Go interface resolution is implicit. If the type you pass in matches what the interface
// is asking for, it will compile.
