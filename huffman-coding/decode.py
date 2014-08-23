#!/usr/bin/env python
""" Decode Huffman Coding

Example coding defined here: https://www.youtube.com/watch?v=ZdooBTdW5bM
"""


CODES = {
    '00': 'I',
    '01': 'S',
    '100': 'P',
    '101': 'R',
    '1100': 'M',
    '1101': 'V',
    '1110': 'E',
    '1111': ' ',
}


def decode(encoded, codes=CODES):
    message = ''
    part = ''
    max_code_len = max(len(c) for c in codes)

    while len(encoded):
        encoded, char = encoded[1:], encoded[0]
        part += char
        if part in codes:
            message += codes[part]
            part = ''
        elif len(part) > max_code_len:
            raise ValueError("Invalid code detected: {}".format(part))
    if part:
        raise ValueError("Partial code remaining: {}".format(part))

    return message


if __name__ == '__main__':
    print decode('1100000101000101001001000011111010011011110101')
