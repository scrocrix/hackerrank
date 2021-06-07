if __name__ == "__main__":
    students = int(input())
    while students > 0:
        grade = int(input())

        if grade < 38:
            print(grade)
        else:
            next_multiple_of_five = grade

            while next_multiple_of_five % 5 != 0:
                next_multiple_of_five = next_multiple_of_five + 1

            if next_multiple_of_five - grade < 3:
                print(next_multiple_of_five)
            else:
                print(grade)

        students = students - 1
