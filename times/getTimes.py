# Get median times for each day of AoC 2020-2023, save to CSV file
#
# Usage: python getTimes.py > times.csv
#
# Then, edit the file, add header row, and change spaces to commas
#
# AK 15/12/2023

import requests, re
import pandas as pd

# Convert 'hh:mm:ss' to minutes
def convertToMins(time):
    h, m, s = time.split(':')
    return float(h) * 60 + float(m) + float(s)/60

# Get the average time for a given day/year, in minutes
def getAvgTime(y, d):

    # Get all 'hh:mm:ss' times from AoC leaderboard
    r = requests.get(f'https://adventofcode.com/{y}/leaderboard/day/{d}')
    times = re.findall(r'\d\d:\d\d:\d\d', r.text)

    # Convert the first 100 (i.e., gold stars) to minutes
    t = [convertToMins(time) for time in times[:100]]

    # Return the mean time for each day, in minutes
    return sum(t) / len(t)

# Do the last four years
for y in [2020, 2021, 2022, 2023]:
    for d in range(1, 25):
        if y == 2023 and d > 15:  # only on day 15 this year
            break
        avg = getAvgTime(y, d)
        print(y, d, avg)  # print to console

