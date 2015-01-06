

def pascals_triangle(rows=5):
    triangle = []
    previous_row = []
    for i in range(rows):
        current_row = []
        for j in range(i + 1):
            if j == 0 or j == i:
                current_row.append(1)
            else:
                current_row.append(previous_row[j - 1] + previous_row[j])
        triangle.append(current_row)
        previous_row = current_row
    return triangle
