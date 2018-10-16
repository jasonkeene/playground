
1. The text says it is a one to one relationship, however it isn't. There are
   multiple, redundant encodings for the same instruction in assembly.
2. A symbol table is created. Optionally, during the first pass the assembler
   can start assembling instructions.
   # forgot macro expansion and evaluation of constant expressions
3. Using the symbol table the instructions are all assembled.
4. The linker takes object files and creates an executable binary out of them.
   It does this by rewriting addresses to point to the right places in memory.
5. The loader sets up memory and loads the program's contents into memory for
   execution to begin.
6. PI equ 3.14
7. source -> assembler -> object --> linker -> exectuable -> loader -> memory
                          object /               
8. At runtime.
9. The symbol names and the locations in memory where they point to.
