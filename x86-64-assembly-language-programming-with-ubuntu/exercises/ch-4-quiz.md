
1. yasm
2. ;
3. .data
4. .bss
5. .text
6. 1. bNum db 10
   2. wNum dw 10291
   3. dwNum dd 2126010
   4. qwNum dq 10000000000
7. 1. bArr resb 100
   2. wArr resw 3000
   3. dwArr resd 200
   4. qArr resq 5000
8. global _start
   _start:
