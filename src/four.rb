# It is a six-digit number.
# The value is within the range given in your puzzle input.
# Two adjacent digits are the same (like 22 in 122345).
# Going from left to right, the digits never decrease; they only ever increase or stay the same (like 111123 or 135679).
require 'pry'

def increasing?(numbers)
  (0...numbers.count-1).each do |i|
    return false if numbers[i] > numbers[i+1]
  end
  true
end

def adjacent?(numbers)
  (0...numbers.count-1).each do |i|
    return true if numbers[i] == numbers[i+1]
  end
  false
end

def main
  first = 137683
  last = 596253

  passwords = []
  (first..last).each do |password|
    split = password.to_s.chars.map(&:to_i)

    tmp = {}
    split.each do |number|
      tmp.key?(number) ? tmp[number] += 1 : tmp[number] = 1
    end

    binding.pry

    next unless increasing?(split)
    next unless adjacent?(split)

    passwords << split.join.to_i
  end

  puts passwords.count
end

main
