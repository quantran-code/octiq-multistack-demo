package com.example;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;

import org.junit.jupiter.api.Test;

class CalcTest {
    @Test
    void adds() {
        assertEquals(5, Calc.add(2, 3));
    }

    @Test
    void divides() {
        assertEquals(5.0, Calc.divide(10, 2));
        assertThrows(IllegalArgumentException.class, () -> Calc.divide(1, 0));
    }
}
