# -*- coding: utf-8 -*-
from __future__ import unicode_literals, absolute_import


def solve_captcha(code: str) -> int:
    c = list(code)
    c.append(c[0])

    s = 0
    for i in range(len(c) - 1):
        s += int(c[i]) if c[i] == c[i + 1] else 0

    return s


if __name__ == '__main__':
    # tests
    assert solve_captcha('1122') == 3
    assert solve_captcha('1111') == 4
    assert solve_captcha('1234') == 0
    assert solve_captcha('91212129') == 9

    # problem
    f = open("input.txt").readline()
    print(solve_captcha(f))
