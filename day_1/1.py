# -*- coding: utf-8 -*-
from __future__ import unicode_literals, absolute_import


def solve_captcha(code: str) -> int:
    c = list(code)
    c.append(c[0])

    s = 0
    for i in range(len(c) - 1):
        s += int(c[i]) if c[i] == c[i + 1] else 0

    return s


def solve_captcha_2(code: str) -> int:
    c = list(code)
    len_ = len(c)
    c = c + c[:int(len_ / 2)]

    s = 0
    for i in range(len_):
        s += int(c[i] if c[i] == c[i + int(len_ / 2)] else 0)

    return s


if __name__ == '__main__':
    # tests 1
    assert solve_captcha('1122') == 3
    assert solve_captcha('1111') == 4
    assert solve_captcha('1234') == 0
    assert solve_captcha('91212129') == 9

    # tests 2
    assert solve_captcha_2('1212') == 6
    assert solve_captcha_2('1221') == 0
    assert solve_captcha_2('123425') == 4
    assert solve_captcha_2('123123') == 12
    assert solve_captcha_2('12131415') == 4

    # problems
    f = open("input.txt").readline()
    print('part 1: {}'.format(solve_captcha(f)))
    print('part 2: {}'.format(solve_captcha_2(f)))
