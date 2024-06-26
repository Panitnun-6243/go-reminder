package main

import (
	"fmt"
	"net/http"

	arrayandslice "github.com/Panitnun-6243/go-basic/arrayAndSlice"
	"github.com/Panitnun-6243/go-basic/calculator"
	"github.com/Panitnun-6243/go-basic/channel"
	controlflow "github.com/Panitnun-6243/go-basic/controlFlow"
	"github.com/Panitnun-6243/go-basic/function"
	"github.com/Panitnun-6243/go-basic/generics"
	"github.com/Panitnun-6243/go-basic/goroutine"
	"github.com/Panitnun-6243/go-basic/interfaces"
	"github.com/Panitnun-6243/go-basic/loop"
	mapobject "github.com/Panitnun-6243/go-basic/mapobject"
	"github.com/Panitnun-6243/go-basic/mutex"
	"github.com/Panitnun-6243/go-basic/oop"
	"github.com/Panitnun-6243/go-basic/oop/abstraction"
	"github.com/Panitnun-6243/go-basic/oop/encapsulation"
	"github.com/Panitnun-6243/go-basic/oop/inheritance"
	"github.com/Panitnun-6243/go-basic/oop/polymorphism"
	"github.com/Panitnun-6243/go-basic/pointer"
	"github.com/Panitnun-6243/go-basic/structs"
	"github.com/labstack/echo/v4"
)

func main() {
	// print
	fmt.Println("Hello Tanny")
	// my package
	calculator.Add(6, 5)
	calculator.Sub(6, 5)
	calculator.Multiply(6, 5)
	calculator.Divide(6, 5)
	// implicit type -> กำหนด type ตั้งแต่แรก
	var implicitType int = 10
	fmt.Println(implicitType)
	// dynamic type -> เราให้ค่ากับตัวแปร ละมันจะหา type ของค่านั้น ๆ ให้อัตโนมัติ
	dynamicType := "Hello Golang"
	fmt.Println(dynamicType)
	// control flow
	fmt.Println(controlflow.OddEvenNumber(3))
	// function = input -> process -> output
	function.MyName("Panitnun", "Suvannabun")
	// anonymous function ใช้ตอนจะให้ผลลัพธ์ของ function เก็บในตัวแปร
	yourNameResult := func(name string) string {
		return "Your name is " + name
	}("Tanny")
	fmt.Println(yourNameResult)
	// loop
	loop.BasicLoop(5)
	loop.WhileLoop()
	loop.EachLoop()
	// pointer
	normalVariable := 2
	var pointerVarialble *int = &normalVariable
	// pass by pointer -> ต้องสร้างตัวแปร pointer มารับค่า address
	pointer.UsePointer(pointerVarialble)
	fmt.Println(*pointerVarialble)
	// pass by reference -> ไม่ต้องสร้างตัวแปร pointer แต่ใช้ address ตรง ๆ ของตัวแปรนั้นเลย
	pointer.UsePointer(&normalVariable)
	fmt.Println(normalVariable)
	// array and slice
	arrayandslice.ArraySlice()
	// map
	mapobject.MapObject()
	// struct
	structs.UseStruct()
	// interface
	interfaces.UseInterface()
	// generic
	generics.UseGeneric()
	// goroutine
	goroutine.UseGouroutine()
	// channel
	channel.UseChannel()
	// mutex
	mutex.UseMutex()
	// oop
	oop.UseOop()
	abstraction.UseAbstraction()
	encapsulation.UseEncapsulation()
	inheritance.UseInheritance()
	polymorphism.UsePolymorphism()

	// other package
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
