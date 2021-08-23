if __name__ == "__main__":
    input = list(map(int, input().rstrip().split()))

    if input[3] > input[1] or input[3] == input[1]:
        print("NO")
        exit(0)

    x = input[0] + input[1]
    y = input[2] + input[3]

    if x > y:
        print("NO")
        exit(0)

    while (x != y):
        x = x + input[1]
        y = y + input[3]

        if x > y:
            print("NO")
            exit(0)

    print("YES")
