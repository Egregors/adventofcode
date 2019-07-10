# -*- coding: utf-8 -*-
from __future__ import unicode_literals, absolute_import


def is_valid(s: str) -> bool:
    s = s.split(' ')
    for w in s:
        if len([x for x in s if w == x]) > 1:
            return False
    return True


def is_valid_2(s: str) -> bool:
    row_d = list()
    for word in s.split(' '):
        word_d = dict()
        for l in word:
            word_d[l] = len([x for x in word if x == l])
        row_d.append(word_d)

    for word in row_d:
        if len([x for x in row_d if x == word]) > 1:
            return False
    return True


if __name__ == '__main__':
    # tests
    assert is_valid('aa bb cc dd ee') is True
    assert is_valid('aa bb cc dd aa') is False
    assert is_valid('aa bb cc dd aaa') is True

    assert is_valid_2('abcde fghij') is True
    assert is_valid_2('abcde xyz ecdab') is False
    assert is_valid_2('a ab abc abd abf abj') is True
    assert is_valid_2('iiii oiii ooii oooi oooo') is True
    assert is_valid_2('oiii ioii iioi iiio') is False

    # problems
    f = open('input.txt').readlines()

    res_1, res_2 = 0, 0

    for row in f:
        res_1 += 1 if is_valid(row.rstrip()) else 0
        res_2 += 1 if is_valid_2(row.rstrip()) else 0

    print('valid 1: {}'.format(res_1))
    print('valid 2: {}'.format(res_2))
