

def first_missing_positive(data):
    x = 1
    while True:
        for y in data:
            if x == y:
                x += 1
                break
        else:
            return x
