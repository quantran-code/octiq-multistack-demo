package com.example;

public final class Calc {
    private Calc() {}

    public static int add(int left, int right) {
        return left + right;
    }

    public static double divide(int left, int right) {
        if (right == 0) {
            throw new IllegalArgumentException("cannot divide by zero");
        }
        return (double) left / right;
    }
}
