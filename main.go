// Package wasm_md5
// @program: archer-all
// @author: ygt
// @create: 2022/7/12 18:51

//go:build js && wasm

//  ../tinygo/bin/tinygo build -o main.wasm -target wasm ./main.go
package main

// GOOS=js GOARCH=wasm go build -o main.wasm main.go
import (
	"crypto/md5"
	"fmt"
	"syscall/js"
	"time"
)

func main() {

	fmt.Println("okkkkkkkkkkkkk")
	// 声明一个函数，用来导出到js端，供js端调用
	//calcMd5 := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	//	// 声明一个和文件大小一样的切片
	//	buffer := make([]byte, args[0].Length())
	//	// 将文件的bytes数据复制到切片中，这里传进来的是一个Uint8Array类型
	//	nano1 := time.Now().UnixMilli()
	//	js.CopyBytesToGo(buffer, args[0])
	//	nano2 := time.Now().UnixMilli()
	//	println("copy:", nano2-nano1)
	//	// 计算md5的值
	//	res := md5.Sum(buffer)
	//	nano3 := time.Now().UnixMilli()
	//	println("sum:", nano3-nano2)
	//	buffer = nil
	//	args = nil
	//	this = js.Null()
	//	// 调用js端的方法，将结构返回给js端
	//	js.Global().Get("target").Get("callback").Invoke(fmt.Sprintf("%x", res))
	//	return nil
	//})
	//defer calcMd5.Release()
	// 获取js端全局对象中的target对象，设置target的calcMd5方法为上面的calcMd5实现
	//js.Global().Get("target").Set("calcMd5", js.FuncOf(calcMd5))
	//js.Global().Set("calcMd5", js.FuncOf(calcMd5))
	js.Global().Set("receiveB", js.FuncOf(receiveB))
	js.Global().Set("getR", js.FuncOf(getR))
	// 阻止go程序退出，因为退出了，js端就不能再调用了
	signal := make(chan int)
	<-signal
}

func calcMd5(this js.Value, args []js.Value) interface{} {
	// 声明一个和文件大小一样的切片
	buffer := make([]byte, args[0].Length())
	// 将文件的bytes数据复制到切片中，这里传进来的是一个Uint8Array类型
	nano1 := time.Now().UnixMilli()
	js.CopyBytesToGo(buffer, args[0])
	nano2 := time.Now().UnixMilli()
	println("copy:", nano2-nano1)
	// 计算md5的值
	res := md5.Sum(buffer)
	nano3 := time.Now().UnixMilli()
	println("sum:", nano3-nano2)
	buffer = nil
	args = nil
	this = js.Null()
	return js.ValueOf(fmt.Sprintf("%x", res))
}

var hash = md5.New()

func receiveB(this js.Value, args []js.Value) interface{} {
	buffer := make([]byte, args[0].Length())
	js.CopyBytesToGo(buffer, args[0])
	hash.Write(buffer)
	return nil
}

func getR(this js.Value, args []js.Value) interface{} {
	defer hash.Reset()
	nano1 := time.Now().UnixMilli()
	r := fmt.Sprintf("%x", hash.Sum(nil))
	nano2 := time.Now().UnixMilli()
	println("sum:", nano2-nano1)
	return js.ValueOf(r)
}
