echo	START   0
first	JSUB	sinit
	+LDA	#12345
	JSUB	num
	JSUB	nl
	+LDA	#txt
	JSUB	string
	JSUB	nl
halt	J       halt

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

.print a string from memory address stored in register A
string	STL	stringL
	STA	stringA
	STA	sAddr
loop	CLEAR	A
	+LDCH    @sAddr
	COMP	#0
	JEQ	break
	JSUB    char
	LDA	sAddr
	ADD	#1
	STA	sAddr
	J     	loop
break	LDL	stringL
	LDA	stringA
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
. initialize sp to the start of stack
sinit	STA     stackA
	LDA     #stack
	STA     sp
	LDA     stackA
	RSUB
. inc sp
spush	STA     stackA
	LDA     sp
	ADD     #3
	STA     sp
	LDA     stackA
	RSUB
. dec sp
spop	STA     stackA
	LDA     sp
	SUB     #3
	STA     sp
	LDA     stackA
	RSUB

stack	RESW    1000
sp	WORD    0
stackA	WORD    0

txt	BYTE    C'hello world!'
	BYTE    0
nlA	RESW	1
nlL	RESW	1
sAddr	RESW	1
stringA	RESW	1
stringL	RESW	1
modB	RESW	1
modT	RESW	1
numL	RESW	1
	END     first
