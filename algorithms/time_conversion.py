import re

if __name__ == "__main__":
    pm_convertion = {
        "12": "12",
        "01": "13",
        "02": "14",
        "03": "15",
        "04": "16",
        "05": "17",
        "06": "18",
        "07": "19",
        "08": "20",
        "09": "21",
        "10": "22",
        "11": "23"
    }

    am_convertion = {
        "12": "00",
        "01": "01",
        "02": "02",
        "03": "03",
        "04": "04",
        "05": "05",
        "06": "06",
        "07": "07",
        "08": "08",
        "09": "09",
        "10": "10",
        "11": "11"
    }

    hour = input().split(":")

    seconds = re.sub("\D", "", hour[2])

    if "PM" in hour[len(hour) - 1]:
        print(pm_convertion[hour[0]] + ":" + hour[1] + ":" + seconds)

    if "AM" in hour[len(hour) - 1]:
        print(am_convertion[hour[0]] + ":" + hour[1] + ":" + seconds)