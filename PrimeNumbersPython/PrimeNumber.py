#  python 3.8.3

import numba
from datetime import datetime


@numba.njit()
def prime(n):
    n += 1
    PrimeNumbers = []  # Numba refers arry element type from append function below
    for i in range(2, n):
        c = 2
        while True:

            if i % c == 0:
                break

            c += 1

            if c == i:
                PrimeNumbers.append(i)
                break

    # print(PrimeNumbers)
    return PrimeNumbers


TestCount = 10

for i in range(TestCount):

    start = datetime.now()
    prime(100000)

    print(f"ExCount #{i+1} | Execution time = {(datetime.now()-start).total_seconds()}s")
