
1. A picture of the VonNeumann Architexture:

   [ cpu ]    [ main memory ]    [ secondary storage ]    [ devices ]
      |              |                     |                  |
      ----------------------- memory bus ----------------------

2. The memory bus.
3. Some form of secondary storage, ie disk drives.
4. Main memory.
5. The latency for reads and writes to on die cache is much smaller vs reading
   from main memory.
6. 4 bytes
7. 1 byte, memory is byte addressable.
8. 1. LSB is 40
   2. MSB is 00
9. 40 4B 4C 00
10. A picture of RAX register:

    [                    rax 64 bit                  ]
                             [       eax 32 bit      ]
                                       [  ax 16 bit  ]
                                              [  al  ]
                                       [  ah  ]

11. 1. 8
    2. 64
    3. 16
    4. 32
    5. 64
    6. 32 # wrong it was 8!
    7. 8
    8. 32 # wrong it was 16!
12. rip
13. rsp
14. 00 00 00 00  00 00 00 00
    00 00 00 00  00 00 00 05
    00 00 00 00  00 00 00 07
    00 00 00 00  00 00 00 20
    00 00 00 00  00 00 00 00
15. 01 23 45 67  89 AB CD EF
    al: EF
    ax: CD EF
    eax: 89 AB CD EF
    rax: 01 23 45 67  89 AB CD EF
