arith	START   0
start	LDA     x
	ADD     y
	STA     sum
	.diff
	LDA     x
	SUB     y
	STA     diff
	.prod
	LDA     x
	MUL     y
	STA     prod
	.quot
	LDA     x
	DIV     y
	STA     quot
	.mod
	LDA     quot
	MUL     y
	STA     temp
	LDA     x
	SUB     temp
	STA     mod
halt	J      halt
. podatki
x	WORD    9
y	WORD    2
sum	WORD    0
diff	WORD    0
prod	WORD    0
quot	WORD    0
mod	WORD    0
temp	WORD    0
	END    start
