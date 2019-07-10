# -*- coding: utf-8 -*-
from __future__ import unicode_literals, absolute_import


def _get_row(s: str) -> list:
    return [int(x) for x in list(s.rstrip().split('	', )) if x.isdigit()]


def _check_row(row: str) -> int:
    row = _get_row(row)
    return max(row) - min(row)


def _check_row_2(row: str) -> int:
    row = _get_row(row)
    for i in range(len(row)):
        for j in range(len(row)):
            if i != j:
                if row[i] % row[j] == 0:
                    return int(row[i] / row[j])


def checksum(a: list, func) -> int:
    s = 0
    for row in a:
        s += func(row)

    return s


if __name__ == '__main__':
    # tests 1
    assert _check_row('5	1	9	5\n') == 8
    assert _check_row('7	5	3\n') == 4
    assert _check_row('2	4	6	8\n') == 6

    # tests 2
    assert _check_row_2('5	9	2	8\n') == 4
    assert _check_row_2('9	4	7	3\n') == 3
    assert _check_row_2('3	8	6	5\n') == 2

    # problems
    f = open('input.txt').readlines()
    print('sum 1: {}'.format(checksum(f, _check_row)))
    print('sum 2: {}'.format(checksum(f, _check_row_2)))
