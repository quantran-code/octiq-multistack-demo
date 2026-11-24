module Billing
  def self.add(left, right)
    left + right
  end

  def self.divide(left, right)
    raise ArgumentError, 'cannot divide by zero' if right.zero?

    left.to_f / right
  end
end
