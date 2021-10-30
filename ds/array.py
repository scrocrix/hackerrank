if __name__ == "__main__":
    n = int(input())
    arr = list(map(int, input().split(" ")))
    for k, val in enumerate(arr):
        if (k + 1) > (len(arr) / 2):
            break

        currentCopy = arr[k]
        arr[k] = arr[len(arr) - (k + 1)]
        arr[len(arr) - (k + 1)] = currentCopy
    print(" ".join(list(map(str, arr))))