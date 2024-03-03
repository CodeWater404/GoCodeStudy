package main

import "fmt"

/**
  @author: CodeWater
  @since: 2024/1/28
  @desc: $
**/

func main() {
	fmt.Println("Hello World!")
	/*可以使用下面的命令将 Go 语言的源代码编译成汇编语言，然后通过汇编语言分析程序具体的执行过程：
	1. go build -gcflags -S main.go
		# command-line-arguments
	main.main STEXT size=89 args=0x0 locals=0x40 funcid=0x0 align=0x0
	        0x0000 00000 (F:/Code/GoCode/exercise/compile/1_compile_example.go:11)  TEXT    main.main(SB), ABIInternal, $64-0
	        0x0000 00000 (F:/Code/GoCode/exercise/compile/1_compile_example.go:11)  CMPQ    SP, 16(R14)
	        0x0004 00004 (F:/Code/GoCode/exercise/compile/1_compile_example.go:11)  PCDATA  $0, $-2
	        0x0004 00004 (F:/Code/GoCode/exercise/compile/1_compile_example.go:11)  JLS     82
	        0x0006 00006 (F:/Code/GoCode/exercise/compile/1_compile_example.go:11)  PCDATA  $0, $-1
	        0x0006 00006 (F:/Code/GoCode/exercise/compile/1_compile_example.go:11)  PUSHQ   BP
	        0x0007 00007 (F:/Code/GoCode/exercise/compile/1_compile_example.go:11)  MOVQ    SP, BP
	        0x000a 00010 (F:/Code/GoCode/exercise/compile/1_compile_example.go:11)  SUBQ    $56, SP
	        0x000e 00014 (F:/Code/GoCode/exercise/compile/1_compile_example.go:11)  FUNCDATA        $0, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
	        0x000e 00014 (F:/Code/GoCode/exercise/compile/1_compile_example.go:11)  FUNCDATA        $1, gclocals·EaPwxsZ75yY1hHMVZLmk6g==(SB)
	        0x000e 00014 (F:/Code/GoCode/exercise/compile/1_compile_example.go:11)  FUNCDATA        $2, main.main.stkobj(SB)
	        0x000e 00014 (F:/Code/GoCode/exercise/compile/1_compile_example.go:12)  MOVUPS  X15, main..autotmp_8+40(SP)
	        0x0014 00020 (F:/Code/GoCode/exercise/compile/1_compile_example.go:12)  LEAQ    type:string(SB), DX
	        0x001b 00027 (F:/Code/GoCode/exercise/compile/1_compile_example.go:12)  MOVQ    DX, main..autotmp_8+40(SP)
	        0x0020 00032 (F:/Code/GoCode/exercise/compile/1_compile_example.go:12)  LEAQ    main..stmp_0(SB), DX
	        0x0027 00039 (F:/Code/GoCode/exercise/compile/1_compile_example.go:12)  MOVQ    DX, main..autotmp_8+48(SP)
	        0x002c 00044 (D:\Go1.21\src\fmt\print.go:314)   MOVQ    os.Stdout(SB), BX
	        0x0033 00051 (<unknown line number>)    NOP
	        0x0033 00051 (D:\Go1.21\src\fmt\print.go:314)   LEAQ    go:itab.*os.File,io.Writer(SB), AX
	        0x003a 00058 (D:\Go1.21\src\fmt\print.go:314)   LEAQ    main..autotmp_8+40(SP), CX
	        0x003f 00063 (D:\Go1.21\src\fmt\print.go:314)   MOVL    $1, DI
	        0x0044 00068 (D:\Go1.21\src\fmt\print.go:314)   MOVQ    DI, SI
	        0x0047 00071 (D:\Go1.21\src\fmt\print.go:314)   PCDATA  $1, $0
	        0x0047 00071 (D:\Go1.21\src\fmt\print.go:314)   CALL    fmt.Fprintln(SB)
	        0x004c 00076 (F:/Code/GoCode/exercise/compile/1_compile_example.go:19)  ADDQ    $56, SP
	        0x0050 00080 (F:/Code/GoCode/exercise/compile/1_compile_example.go:19)  POPQ    BP
	        0x0051 00081 (F:/Code/GoCode/exercise/compile/1_compile_example.go:19)  RET
	        0x0052 00082 (F:/Code/GoCode/exercise/compile/1_compile_example.go:19)  NOP
	        0x0052 00082 (F:/Code/GoCode/exercise/compile/1_compile_example.go:11)  PCDATA  $1, $-1
	        0x0052 00082 (F:/Code/GoCode/exercise/compile/1_compile_example.go:11)  PCDATA  $0, $-2
	        0x0052 00082 (F:/Code/GoCode/exercise/compile/1_compile_example.go:11)  CALL    runtime.morestack_noctxt(SB)
	        0x0057 00087 (F:/Code/GoCode/exercise/compile/1_compile_example.go:11)  PCDATA  $0, $-1
	        0x0057 00087 (F:/Code/GoCode/exercise/compile/1_compile_example.go:11)  JMP     0
	        0x0000 49 3b 66 10 76 4c 55 48 89 e5 48 83 ec 38 44 0f  I;f.vLUH..H..8D.
	        0x0010 11 7c 24 28 48 8d 15 00 00 00 00 48 89 54 24 28  .|$(H......H.T$(
	        0x0020 48 8d 15 00 00 00 00 48 89 54 24 30 48 8b 1d 00  H......H.T$0H...
	        0x0030 00 00 00 48 8d 05 00 00 00 00 48 8d 4c 24 28 bf  ...H......H.L$(.
	        0x0040 01 00 00 00 48 89 fe e8 00 00 00 00 48 83 c4 38  ....H.......H..8
	        0x0050 5d c3 e8 00 00 00 00 eb a7                       ]........
	        rel 2+0 t=23 type:string+0
	        rel 2+0 t=23 type:*os.File+0
	        rel 23+4 t=14 type:string+0
	        rel 35+4 t=14 main..stmp_0+0
	        rel 47+4 t=14 os.Stdout+0
	        rel 54+4 t=14 go:itab.*os.File,io.Writer+0
	        rel 72+4 t=7 fmt.Fprintln+0
	        rel 83+4 t=7 runtime.morestack_noctxt+0
	go:cuinfo.producer.main SDWARFCUINFO dupok size=0
	        0x0000 72 65 67 61 62 69                                regabi
	go:cuinfo.packagename.main SDWARFCUINFO dupok size=0
	        0x0000 6d 61 69 6e                                      main
	go:info.fmt.Println$abstract SDWARFABSFCN dupok size=44
	        0x0000 05 66 6d 74 2e 50 72 69 6e 74 6c 6e 00 01 b9 02  .fmt.Println....
	        0x0010 01 13 61 00 00 00 00 00 00 13 6e 00 01 00 00 00  ..a.......n.....
	        0x0020 00 13 65 72 72 00 01 00 00 00 00 00              ..err.......
	        rel 0+0 t=22 type:[]interface {}+0
	        rel 0+0 t=22 type:error+0
	        rel 0+0 t=22 type:int+0
	        rel 21+4 t=31 go:info.[]interface {}+0
	        rel 29+4 t=31 go:info.int+0
	        rel 39+4 t=31 go:info.error+0
	go:itab.*os.File,io.Writer SRODATA dupok size=32
	        0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	        0x0010 5a 22 ee 60 00 00 00 00 00 00 00 00 00 00 00 00  Z".`............
	        rel 0+8 t=1 type:io.Writer+0
	        rel 8+8 t=1 type:*os.File+0
	        rel 24+8 t=-32767 os.(*File).Write+0
	main..inittask SNOPTRDATA size=8
	        0x0000 00 00 00 00 00 00 00 00                          ........
	        rel 0+0 t=81 fmt..inittask+0
	go:string."Hello World!" SRODATA dupok size=12
	        0x0000 48 65 6c 6c 6f 20 57 6f 72 6c 64 21              Hello World!
	main..stmp_0 SRODATA static size=16
	        0x0000 00 00 00 00 00 00 00 00 0c 00 00 00 00 00 00 00  ................
	        rel 0+8 t=1 go:string."Hello World!"+0
	runtime.nilinterequal·f SRODATA dupok size=8
	        0x0000 00 00 00 00 00 00 00 00                          ........
	        rel 0+8 t=1 runtime.nilinterequal+0
	runtime.memequal64·f SRODATA dupok size=8
	        0x0000 00 00 00 00 00 00 00 00                          ........
	        rel 0+8 t=1 runtime.memequal64+0
	runtime.gcbits.0100000000000000 SRODATA dupok size=8
	        0x0000 01 00 00 00 00 00 00 00                          ........
	type:.namedata.*[1]interface {}- SRODATA dupok size=18
	        0x0000 00 10 2a 5b 31 5d 69 6e 74 65 72 66 61 63 65 20  ..*[1]interface
	        0x0010 7b 7d                                            {}
	type:*[1]interface {} SRODATA dupok size=56
	        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	        0x0010 a8 0e 57 36 08 08 08 36 00 00 00 00 00 00 00 00  ..W6...6........
	        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	        0x0030 00 00 00 00 00 00 00 00                          ........
	        rel 24+8 t=1 runtime.memequal64·f+0
	        rel 32+8 t=1 runtime.gcbits.0100000000000000+0
	        rel 40+4 t=5 type:.namedata.*[1]interface {}-+0
	        rel 48+8 t=1 type:[1]interface {}+0
	runtime.gcbits.0200000000000000 SRODATA dupok size=8
	        0x0000 02 00 00 00 00 00 00 00                          ........
	type:[1]interface {} SRODATA dupok size=72
	        0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	        0x0010 6e 20 6a 3d 02 08 08 11 00 00 00 00 00 00 00 00  n j=............
	        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	        0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	        0x0040 01 00 00 00 00 00 00 00                          ........
	        rel 24+8 t=1 runtime.nilinterequal·f+0
	        rel 32+8 t=1 runtime.gcbits.0200000000000000+0
	        rel 40+4 t=5 type:.namedata.*[1]interface {}-+0
	        rel 44+4 t=-32763 type:*[1]interface {}+0
	        rel 48+8 t=1 type:interface {}+0
	        rel 56+8 t=1 type:[]interface {}+0
	type:.importpath.fmt. SRODATA dupok size=5
	        0x0000 00 03 66 6d 74                                   ..fmt
	gclocals·g2BeySu+wFnoycgXfElmcg== SRODATA dupok size=8
	        0x0000 01 00 00 00 00 00 00 00                          ........
	gclocals·EaPwxsZ75yY1hHMVZLmk6g== SRODATA dupok size=9
	        0x0000 01 00 00 00 02 00 00 00 00                       .........
	main.main.stkobj SRODATA static size=24
	        0x0000 01 00 00 00 00 00 00 00 f0 ff ff ff 10 00 00 00  ................
	        0x0010 10 00 00 00 00 00 00 00                          ........
	        rel 20+4 t=5 runtime.gcbits.0200000000000000+0
	*/
	/*
		2. Go 语言更详细的编译过程，我们可以通过下面的命令获取汇编指令的优化过程：
		   GOSSAFUNC=main go build main.go
		   上述命令会在当前文件夹下生成一个 ssa.html 文件，我们打开这个文件后就能看到汇编代码优化的每一个步骤.
		todo:win10下会失败，暂时不知道原因。
	*/

}
