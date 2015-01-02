from solution import Solution


class DescribeSolution:
    def it_computes_spiral_for_square_input(self):
        input = [
            [1, 2, 3],
            [4, 5, 6],
            [7, 8, 9],
        ]
        expected_output = [1, 2, 3, 6, 9, 8, 7, 4, 5]
        assert Solution().spiralOrder(input) == expected_output

    def it_computes_spiral_for_rectangle_input(self):
        input = [
            [1, 2, 3],
            [4, 5, 6],
            [7, 8, 9],
            [10, 11, 12],
        ]
        expected_output = [1, 2, 3, 6, 9, 12, 11, 10, 7, 4, 5, 8]
        assert Solution().spiralOrder(input) == expected_output

    def it_handles_empty_input(self):
        assert Solution().spiralOrder([]) == []
