arithr	START   0
start	LDS     x .sum
	LDT     y
	RMO	T, A
	ADDR    S, A
	STA     sum
	.diff
	RMO	S, A
	SUBR    T, A
	STA     diff
	.prod
	RMO	T, A
	MULR    S, A
	STA     prod
	.quot
	LDA     x
	RMO	S, A
	DIVR    T, A
	STA     quot
	.mod
	MULR	T, A
	SUBR	A, S
	. SUBR	x*y - (x/y) * y
	STS	mod
halt	J      halt
. podatki
x	WORD    9
y	WORD    2
sum	WORD    0
diff	WORD    0
prod	WORD    0
quot	WORD    0
mod	WORD    0
	END    start
