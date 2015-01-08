

def can_move_right(grid, position):
    new_position = move_right(position)
    return new_position[0] < len(grid[0])


def can_move_down(grid, position):
    new_position = move_down(position)
    return new_position[1] < len(grid)


def move_right(position):
    return position[0] + 1, position[1]


def move_down(position):
    return position[0], position[1] + 1


def min_path_sum(grid, position=(0, 0)):
    min_path_sums = []

    if can_move_right(grid, position):
        new_position = move_right(position)
        min_path_sums.append(min_path_sum(grid, new_position))
    if can_move_down(grid, position):
        new_position = move_down(position)
        min_path_sums.append(min_path_sum(grid, new_position))

    dsum = min(min_path_sums) if min_path_sums else 0
    return grid[position[1]][position[0]] + dsum
