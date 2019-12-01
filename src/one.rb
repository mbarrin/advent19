#!/usr/bin/env ruby

input = File.readlines("input/one.txt")

total = 0

input.each do |number|
  total += (number.to_i / 3.0).floor - 2
end

puts total
