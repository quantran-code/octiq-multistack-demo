import pytest

from demo.calc import add, divide


def test_add() -> None:
    assert add(2, 3) == 5


def test_divide() -> None:
    assert divide(10, 2) == 5
    with pytest.raises(ValueError):
        divide(1, 0)
