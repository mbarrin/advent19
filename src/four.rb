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

def too_many?(numbers)
  tmp = {}
  numbers.each do |number|
    tmp.key?(number) ? tmp[number] += 1 : tmp[number] = 1
  end
  return true if tmp.values.include?(2)
  false
end

def main
  first = 137683
  last = 596253

  passwords = []
  (first..last).each do |password|
    split = password.to_s.chars.map(&:to_i)

    next unless increasing?(split)
    next unless adjacent?(split)
    next unless too_many?(split)

    passwords << split.join.to_i
  end

  puts passwords.count
end

main
