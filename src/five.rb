#!/usr/bin/env ruby
require 'pry'

def logic(input)
  index = 0
  while index < input.length do
    local = input[index].rjust(5, "0")

    opcode = local[-2..-1].to_i

    param1 = immediate?(local[-3].to_i) ? index + 1 : input[index + 1].to_i
    param2 = immediate?(local[-4].to_i) ? index + 2 : input[index + 2].to_i
    param3 = immediate?(local[-5].to_i) ? index + 3 : input[index + 3].to_i

    case opcode
    when 1
      input[param3] = (input[param1].to_i + input[param2].to_i).to_s
      index += 4
    when 2
      input[param3] = (input[param1].to_i * input[param2].to_i).to_s
      index += 4
    when 3
      print "enter number: "
      input[param1] = gets.chomp
      index += 2
    when 4
      puts input[param1].to_i
      index += 2
    when 5 # jump if true
      if input[param1].to_i.zero?
        index += 3
      else
        index = input[param2].to_i
      end
    when 6
      if input[param1].to_i.zero?
        index = input[param2].to_i
      else
        index += 3
      end
    when 7
      input[param3] = if input[param1] < input[param2]
                        "1"
                      else
                        "0"
                      end
      index += 4
    when 8
      input[param3] = if input[param1] == input[param2]
                        "1"
                      else
                        "0"
                      end
      index += 4
    when 99
      return input[0]
    end
  end
end

def immediate?(number)
  return true if number == 1
  false
end

def main
  input = File.read("input/five.txt")

  logic(input.split(","))
end

main
