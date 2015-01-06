from solution import pascals_triangle


class DescribePascalsTriangle:
    def it_builds_a_valid_triangle(self):
        assert pascals_triangle(5) == [
            [1],
            [1, 1],
            [1, 2, 1],
            [1, 3, 3, 1],
            [1, 4, 6, 4, 1],
        ]
