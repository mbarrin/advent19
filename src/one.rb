#!/usr/bin/env ruby

def required_fuel(mass)
  (mass / 3.0).floor - 2
end

def fuel_for_module(mass)
  fuel = required_fuel(mass)

  subtotal = fuel

  until fuel <= 0
    fuel = required_fuel(fuel)
    subtotal += fuel if fuel.positive?
  end

  subtotal
end

def main
  input = File.readlines("input/one.txt")

  total = 0

  input.each do |mass|
    total += fuel_for_module(mass.to_i)
  end

  puts total
end

main
