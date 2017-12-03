# -*- coding: utf-8 -*-
from __future__ import unicode_literals, absolute_import


def _check_row(row: str) -> int:
    row = [int(x) for x in list(row.rstrip().split('	', )) if x.isdigit()]
    return max(row) - min(row)


def checksum(a: list) -> int:
    s = 0
    for row in a:
        s += _check_row(row)

    return s


if __name__ == '__main__':
    # tests
    assert _check_row('5	1	9	5\n') == 8
    assert _check_row('7	5	3\n') == 4
    assert _check_row('2	4	6	8\n') == 6

    # problem
    f = open('input.txt').readlines()
    print('sum: {}'.format(checksum(f)))
