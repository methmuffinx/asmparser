/* ARM example program which moves 0x2 into r4 and then xors 0x1 to it */

.global main /* 'main' is our entry point and must be global */

main:
	MOV r4, #2
	EOR r4, r4, #1
	B main