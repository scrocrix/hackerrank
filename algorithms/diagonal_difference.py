if __name__ == "__main__":
    matrix_size = int(input())
    lr_position = 0
    rl_position = matrix_size - 1
    lr = 0
    rl = 0

    for i in range(matrix_size):
        items = list(map(int, input().split(" ")))

        lr = lr + items[lr_position]
        rl = rl + items[rl_position]

        if i == 0:
            lr_position = 1
            rl_position = rl_position - 1
        else:
            lr_position = lr_position + 1
            rl_position = rl_position - 1

    print(abs(lr - rl))