if __name__ == "__main__":
    n = int(input().strip())

    arr = list(map(int, input().rstrip().split()))

    positive = 0
    zero = 0
    negative = 0

    for item in arr:
        if item > 0:
            positive = positive + 1
            continue

        if item < 0:
            negative = negative + 1
            continue

        zero = zero + 1

    print(positive / n)
    print(negative / n)
    print(zero / n)
