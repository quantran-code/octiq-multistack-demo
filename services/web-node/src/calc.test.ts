import { describe, expect, it } from 'vitest'
import { add, divide } from './calc'

describe('calc', () => {
  it('adds', () => expect(add(2, 3)).toBe(5))
  it('divides', () => expect(divide(10, 2)).toBe(5))
  it('rejects division by zero', () => expect(() => divide(1, 0)).toThrow())
})
