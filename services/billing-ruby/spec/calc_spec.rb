require 'spec_helper'

RSpec.describe Billing do
  it 'adds numbers' do
    expect(Billing.add(2, 3)).to eq(5)
  end

  it 'divides numbers' do
    expect(Billing.divide(10, 2)).to eq(5.0)
    expect { Billing.divide(1, 0) }.to raise_error(ArgumentError)
  end
end
