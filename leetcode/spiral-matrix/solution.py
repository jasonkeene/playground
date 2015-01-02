#!/usr/bin/env python


class Solution:
    directions = ['right', 'down', 'left', 'up']

    def spiralOrder(self, matrix):
        if not matrix:
            return []
        self.matrix = matrix
        self.result = []
        self.position = (0, 0)
        self.direction = 0
        self.width, self.height = len(matrix[0]), len(matrix)
        self.visited = set()
        self.go()
        return self.result

    def go(self):
        # continue until you can't move in all four directions
        cant_move = 0
        while cant_move < 4:
            # append current position if it hasn't already been appended
            if self.position not in self.visited:
                self.append_element()

            if self.next_element_is_edge():
                cant_move += 1
                self.change_direction()
            else:
                cant_move = 0
                self.position = self.next_position()

    def change_direction(self):
        self.direction = (self.direction + 1) % len(self.directions)

    def append_element(self):
        self.result.append(self.lookup())
        self.visited.add(self.position)

    def next_position(self):
        return {
            'right': (self.position[0] + 1, self.position[1]),
            'down': (self.position[0], self.position[1] + 1),
            'left': (self.position[0] - 1, self.position[1]),
            'up': (self.position[0], self.position[1] - 1),
        }.get(self.directions[self.direction])

    def lookup(self, position=None):
        position = position or self.position
        return self.matrix[position[1]][position[0]]

    def next_element_is_edge(self):
        on_edge = {
            'right': self.position[0] + 1 == self.width,
            'down': self.position[1] + 1 == self.height,
            'left': self.position[0] - 1 == -1,
            'up': self.position[1] - 1 == -1,
        }.get(self.directions[self.direction])
        return on_edge or self.next_position() in self.visited
