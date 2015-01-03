
END = object()
TERM = object()


def off_by_one(a, b):
    off = False
    for i in xrange(len(a)):
        if a[i] != b[i]:
            if off is False:
                off = True
            else:
                return False
    return off


def ladder_length(start, end, dictionary):
    dictionary = set(dictionary) | {end}
    count = 1
    frame = {start}
    while True:
        print "#{} frame: {}, dictionary: {}".format(count, len(frame), len(dictionary))
        if not frame:
            return 0
        dictionary -= frame
        count += 1
        related = set()
        for word in frame:
            for dword in dictionary:
                if off_by_one(word, dword):
                    if dword == end:
                        return count
                    related.add(dword)
        frame = related
