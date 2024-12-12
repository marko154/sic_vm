prog	START   0
first	JSUB    sinit
loop	JSUB	readNum
	COMP	#0
	JEQ	halt
	JSUB	fact
	JSUB	num
	JSUB	nl
	J	loop
halt	J       halt

readNum	STT	readT
	STS	readS
	LDT	#0
	LDS	#10
rLoop	CLEAR	A
	RD	#0xFA
	COMP	#10 .newline
	JEQ    readEnd
	SUB	#48
	MULR	S, T
	ADDR	A, T
	J	rLoop
readEnd	RMO	T, A
	LDT	readT
	LDS	readS
	RSUB

fact	COMP    #0
	JEQ     base

	STL     @sp
	JSUB    spush
	STA     @sp
	JSUB    spush

	SUB     #1
	JSUB    fact

	JSUB    spop
	LDB     @sp
	JSUB    spop
	LDL     @sp
	MULR    B, A
	J       factRet
base	LDA     #1
factRet	RSUB

sinit	STA     stackA
	LDA     #stack
	STA     sp
	LDA     stackA
	RSUB
spush	STA     stackA
	LDA     sp
	ADD     #3
	STA     sp
	LDA     stackA
	RSUB
spop	STA     stackA
	LDA     sp
	SUB     #3
	STA     sp
	LDA     stackA
	RSUB

.print char in register A
char	WD      #1
	RSUB

.print new line
nl	STA     nlA
	STL     nlL
	LDCH    #10
	JSUB    char
	LDA     nlA
	LDL	nlL
	RSUB

. A -> A % 10
mod10	STT	modT
	RMO	A, T
	DIV	#10
	MUL	#10
	SUBR	A, T
	RMO	T, A
	LDT	modT
	RSUB

. int print_num(int a) {
.     if (a < 10) {
.         print(a);
.     } else {
.         print_num(a / 10);
.         print(a % 10);
.     }
. }
.print number in register A
num	STL	@sp
	JSUB	spush
	STA	@sp
	JSUB	spush
	STT	@sp
	JSUB	spush

	COMP	#10
	JLT	numEnd
	RMO	A, T
	DIV	#10
	JSUB	num
	RMO	T, A
numEnd	JSUB	mod10
	ADD	#48
	WD	#1
	JSUB	spop
	LDT	@sp
	JSUB	spop
	LDA	@sp
	JSUB	spop
	LDL	@sp
	RSUB

nlA	RESW	1
nlL	RESW	1
modB	RESW	1
modT	RESW	1
numL	RESW	1
numA	RESW	1


readA	RESW    1
readT	RESW    1
readS	RESW    1

stack	RESW    1000
sp	WORD    0
stackA	WORD    0
	END     first
