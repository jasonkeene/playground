import pytest

from solution import first_missing_positive


class DescribeFirstMissingPositive:
    @pytest.mark.parametrize(['data', 'result'], [
        ([1, 2, 0], 3),
        ([3, 4, -1, 1], 2),
    ])
    def it_finds_first_missing_positive_integer(self, data, result):
        assert first_missing_positive(data) == result
