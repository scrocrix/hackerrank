if __name__ == "__main__":
    alice = input().split(" ")
    bob = input().split(" ")
    alice = list(alice)
    bob = list(bob)
    alice_points = 0
    bob_points = 0

    for i in range(len(alice)):
        if int(alice[i]) > int(bob[i]):
            alice_points = alice_points + 1
        elif int(bob[i]) > int(alice[i]):
            bob_points = bob_points + 1

    print(" ".join([str(alice_points), str(bob_points)]))