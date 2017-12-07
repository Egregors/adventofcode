def _redistribution(a: list) -> list:
    bank_vol = max(a)
    bank_id = a.index(bank_vol)

    i = bank_id + 1 if bank_id != len(a) - 1 else 0
    a[bank_id] = 0
    while bank_vol > 0:
        a[i] += 1
        bank_vol -= 1
        i = i + 1 if i != len(a) - 1 else 0
    return a


def solve(bank: str):
    b = [int(x) for x in bank.rstrip().split('\t')]
    log = set()
    log2 = dict()

    iter_ = ''.join(str(b))
    iters = 0

    while iter_ not in log:
        log.add(iter_)
        log2[iter_] = iters
        iter_ = ''.join(str(_redistribution(b)))
        iters += 1

    return iters, len(log2) - log2[iter_]


if __name__ == '__main__':
    # tests
    input_ = '0	2	7	0'

    assert solve(input_) == (5, 4)

    # problems
    f = open('input.txt').readline()
    print('meh: {}'.format(solve(f)))
