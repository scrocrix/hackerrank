if __name__ == "__main__":
    n = input()
    arr = input().split(" ")
    arr = list(arr)
    s = 0
    for item in arr:
        item = int(item)
        s = s + item
    print(s)