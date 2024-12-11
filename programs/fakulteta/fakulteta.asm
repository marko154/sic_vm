prog	START   0
	JSUB    sinit
	LDA     #6
	JSUB    fact
halt	J       halt

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

stack	RESW    1000
sp	WORD    0
stackA	WORD    0
	END     prog
