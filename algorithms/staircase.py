if __name__ == "__main__":
    n = int(input().strip())
    for k, item in enumerate(range(n)):
        blank_spaces = " " * (n - (k + 1))
        hashtags = "#" * (k + 1)
        print(blank_spaces + hashtags)
