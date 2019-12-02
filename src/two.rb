#!/usr/bin/env ruby

def logic(input)
  input.each_with_index do |num, index|
    next unless (index % 4).zero?

    first = input[index+1]
    second = input[index+2]
    dest = input[index+3]

    if num == 1
      input[dest] = input[first] + input[second]
    elsif num == 2
      input[dest] = input[first] * input[second]
    elsif num == 99
      return input[0]
    else
      puts "ERROR"
    end
  end
end

def main
  input = File.read("input/two.txt")

  (0..99).each do |noun|
    (0..99).each do |verb|
      local_input = input.split(",").map(&:to_i)

      local_input[1] = noun
      local_input[2] = verb

      if logic(local_input) == 19_690_720
        puts 100 * noun + verb
        exit
      end
    end
  end
end

main
