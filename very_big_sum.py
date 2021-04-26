if __name__ == "__main__":
    args = int(input())
    n = list(input().split(" "))
    v = 0
    for i in n:
        v = v + int(i)
    print(v)