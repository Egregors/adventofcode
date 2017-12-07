# -*- coding: utf-8 -*-
from __future__ import unicode_literals, absolute_import


class Grid:
    U = '^'
    R = '>'
    D = 'v'
    L = '<'

    def __init__(self, x: int):
        self.goal = x
        self.s = list()

        self.s.append({
            'n': 1,
            'n2': 1,
            'd': self.R,
            'x': 0,
            'y': 0
        })

        self.exist = set()
        self.n2s = dict()
        self.exist.add('0;0')
        self.n2s['0;0'] = 1

        self._alloc()

    def _forward(self, i: int) -> dict:
        current = self.s[i - 1]
        next_ = {'n': i + 1}

        if current['d'] == self.U:
            next_['x'] = current['x']
            next_['y'] = current['y'] + 1
            next_['d'] = self.U if '{};{}'.format(next_['x'] - 1, next_['y']) in self.exist else self.L

        elif current['d'] == self.R:
            next_['x'] = current['x'] + 1
            next_['y'] = current['y']
            next_['d'] = self.R if '{};{}'.format(next_['x'], next_['y'] + 1) in self.exist else self.U

        elif current['d'] == self.D:
            next_['x'] = current['x']
            next_['y'] = current['y'] - 1
            next_['d'] = self.D if '{};{}'.format(next_['x'] + 1, next_['y']) in self.exist else self.R

        elif current['d'] == self.L:
            next_['x'] = current['x'] - 1
            next_['y'] = current['y']
            next_['d'] = self.L if '{};{}'.format(next_['x'], next_['y'] - 1) in self.exist else self.D

        else:
            raise ValueError()

        l = list()
        if '{};{}'.format(next_['x'], next_['y'] + 1) in self.exist:
            l.append(self.n2s['{};{}'.format(next_['x'], next_['y'] + 1)])

        if '{};{}'.format(next_['x'] + 1, next_['y'] + 1) in self.exist:
            l.append(self.n2s['{};{}'.format(next_['x'] + 1, next_['y'] + 1)])

        if '{};{}'.format(next_['x'] + 1, next_['y']) in self.exist:
            l.append(self.n2s['{};{}'.format(next_['x'] + 1, next_['y'])])

        if '{};{}'.format(next_['x'] + 1, next_['y'] - 1) in self.exist:
            l.append(self.n2s['{};{}'.format(next_['x'] + 1, next_['y'] - 1)])

        if '{};{}'.format(next_['x'], next_['y'] - 1) in self.exist:
            l.append(self.n2s['{};{}'.format(next_['x'], next_['y'] - 1)])

        if '{};{}'.format(next_['x'] - 1, next_['y'] - 1) in self.exist:
            l.append(self.n2s['{};{}'.format(next_['x'] - 1, next_['y'] - 1)])

        if '{};{}'.format(next_['x'] - 1, next_['y']) in self.exist:
            l.append(self.n2s['{};{}'.format(next_['x'] - 1, next_['y'])])

        if '{};{}'.format(next_['x'] - 1, next_['y'] + 1) in self.exist:
            l.append(self.n2s['{};{}'.format(next_['x'] - 1, next_['y'] + 1)])

        next_['n2'] = sum(l)

        if next_['n2'] > self.goal:
            print('p2: {}'.format(next_['n2']))
            exit(0)

        new_exist = '{};{}'.format(next_['x'], next_['y'])
        self.exist.add(new_exist)
        self.n2s[new_exist] = next_['n2']
        return next_

    def _alloc(self):
        for i in range(1, self.goal):
            if i % 10000 == 0:
                print(i)
            self.s.append(self._forward(i))

    def solve(self):
        return abs(0 - self.s[self.goal - 1]['x']) + abs(0 - self.s[self.goal - 1]['y'])


if __name__ == '__main__':
    # tests
    # assert Grid(1).solve() == 0
    # assert Grid(12).solve() == 3
    # assert Grid(23).solve() == 2
    # assert Grid(1024).solve() == 31

    # problems
    print('steps: {}'.format(Grid(312051).solve()))
