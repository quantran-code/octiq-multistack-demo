require 'simplecov'
require 'simplecov_json_formatter'
SimpleCov.start do
  add_filter '/spec/'
  formatter SimpleCov::Formatter::JSONFormatter
end

require_relative '../lib/calc'
