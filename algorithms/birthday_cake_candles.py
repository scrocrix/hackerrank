if __name__ == "__main__":
    n = int(input())
    candles = input()
    candles_arr = list(map(int, candles.rstrip().split()))
    highest_candle = str(max(candles_arr))
    print(candles.count(highest_candle))
