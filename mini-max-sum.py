if __name__ == "__main__":
    arr = list(map(int, input().rstrip().split()))
    arr.sort()
    m = 0
    mx = 0

    for k in range(len(arr)-1):
        if k == len(arr):
            break
        m += arr[k]

    arr.sort(reverse=True)

    for k in range(len(arr)-1):
        if k == len(arr):
            break
        mx += arr[k]

    print(str(m) + " " + str(mx))


