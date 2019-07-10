def go(cmd_list: list) -> int:
    cmd_list = [int(x) for x in cmd_list]
    i, steps = 0, -1
    try:
        while True:
            steps += 1
            cmd_list[i] += 1
            i = i + cmd_list[i] - 1
    except IndexError:
        return steps


def go_2(cmd_list: list) -> int:
    cmd_list = [int(x) for x in cmd_list]
    i, steps = 0, -1
    try:
        while True:
            steps += 1
            if cmd_list[i] >= 3:
                cmd_list[i] -= 1
                i += cmd_list[i] + 1
            else:
                cmd_list[i] += 1
                i += cmd_list[i] - 1

    except IndexError:
        return steps


if __name__ == '__main__':
    # tests
    test_data = [
        '0',
        '3',
        '0',
        '1',
        '-3'
    ]

    assert go(test_data) == 5
    assert go_2(test_data) == 10

    # problems
    print('steps: {}'.format(go(open('input.txt').readlines())))
    print('steps: {}'.format(go_2(open('input.txt').readlines())))
