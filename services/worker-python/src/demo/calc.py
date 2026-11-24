def add(left: int, right: int) -> int:
    return left + right


def divide(left: int, right: int) -> float:
    if right == 0:
        raise ValueError("cannot divide by zero")
    return left / right


def unused_helper(value: int) -> int:
    return value * 2
