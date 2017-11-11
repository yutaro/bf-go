package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/go-llvm/llvm"
)

func initializeLLVM() {
	llvm.LinkInMCJIT()
	llvm.InitializeNativeTarget()
	llvm.InitializeNativeAsmPrinter()
}

const (
	BUF_LENGTH = 10000
)

type bfB struct {
	llvm.Builder
}

func (b *bfB) allocaInt(name string) llvm.Value {
	val := b.CreateAlloca(llvm.Int32Type(), name)
	b.CreateStore(llvm.ConstInt(llvm.Int32Type(), 0, false), val)
	return val
}

func (b *bfB) allocaArrayInt(length uint64, name string) llvm.Value {
	return b.CreateArrayAlloca(llvm.Int32Type(), llvm.ConstInt(llvm.Int32Type(), length, false), name)
}

func (b *bfB) incVal(target llvm.Value) {
	val := b.CreateLoad(target, "val")
	res := b.CreateAdd(val, llvm.ConstInt(llvm.Int32Type(), 1, false), "res")
	b.CreateStore(res, target)
}

func (b *bfB) decVal(target llvm.Value) {
	val := b.CreateLoad(target, "val")
	res := b.CreateSub(val, llvm.ConstInt(llvm.Int32Type(), 1, false), "res")
	b.CreateStore(res, target)
}

func declearInc(b *bfB, mod llvm.Module) {
	inc := llvm.FunctionType(llvm.VoidType(), []llvm.Type{llvm.Int32Type()}, false)
	llvm.AddFunction(mod, "inc", inc)
	llvm.AddBasicBlock(mod.NamedFunction("inc"), "entry")

	b.CreateRetVoid()
}

func main() {
	initializeLLVM()

	tape, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	mod := llvm.NewModule("bf")
	b := &bfB{llvm.NewBuilder()}
	defer b.Dispose()

	declearInc(b, mod)

	main := llvm.FunctionType(llvm.Int32Type(), []llvm.Type{}, false)
	llvm.AddFunction(mod, "main", main)
	entry := llvm.AddBasicBlock(mod.NamedFunction("main"), "entry")
	b.SetInsertPoint(entry, entry.FirstInstruction())

	ptr := b.allocaInt("ptr")
	b.allocaArrayInt(BUF_LENGTH, "buf")

	for _, op := range tape {
		switch op {
		case '>':
			b.incVal(ptr)
		case '<':
			b.decVal(ptr)
		case '+':

		case '-':

		case '.':

		case ',':

		case '[':

		case ']':

		}
	}

	b.CreateRet(llvm.ConstInt(llvm.IntType(32), 0, false))
	mod.Dump()

	//	file, _ := os.Create("./result")
	//	llvm.WriteBitcodeToFile(mod, file)
}
