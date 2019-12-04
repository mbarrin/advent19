#!/usr/bin/env ruby
require 'pry'

class Point
  attr_accessor :x, :y

  def initialize(x, y)
    @x = x
    @y = y
  end

  def distance
    x.to_i.magnitude + y.to_i.magnitude
  end

  def ==(other)
    return true if x == other.x && y == other.y
    false
  end
end

class Line
  attr_accessor :first, :last

  def initialize(first, last)
    @first = first
    @last = last
  end

  def equation
    top = last.y - first.y
    bottom = last.x - first.x

    return { x: last.x, y: 0 } if bottom.zero?
    return { x: 0, y: last.y } if top.zero?

    nil
  end

  def points
    if first.y == last.y
      if first.x < last.x
        (first.x..last.x).map { |x| Point.new(x, first.y) }
      else
        first.x.downto(last.x).map { |x| Point.new(x, first.y) }
      end
    elsif first.x == last.x
      if first.y < last.y
        (first.y..last.y).map { |y| Point.new(first.x, y) }
      else
        first.y.downto(last.y).map { |y| Point.new(first.x, y) }
      end
    end
  end

  def include?(point)
    points.include?(point)
  end

  def length
    points.size - 1
  end

  def fast_intersection(other)
    return nil if equation.nil? || other.equation.nil?

    x = equation[:x] + other.equation[:x]
    y = equation[:y] + other.equation[:y]

    return nil unless (first.x <= x && x <= last.x) || (last.x <= x && x <= first.x)
    return nil unless (other.first.x <= x && x <= other.last.x) || (other.last.x <= x && x <= other.first.x)

    return nil unless (first.y <= y && y <= last.y) || (last.y <= y && y <= first.y)
    return nil unless (other.first.y <= y && y <= other.last.y) || (other.last.y <= y && y <= other.first.y)

    Point.new(x, y)
  end
end

class Wire
  attr_accessor :lines

  def initialize(input)
    @lines = generate_path(input)
  end

  def generate_path(wire)
    path = []
    first = Point.new(0,0)
    coord = [0,0]
    route = wire.split(",")

    route.each do |r|
      direction = r[0]
      count = r[1..-1].to_i

      case direction
      when "U"
        coord[1] += count
      when "R"
        coord[0] += count
      when "D"
        coord[1] -= count
      when "L"
        coord[0] -= count
      end
      last = Point.new(coord[0], coord[1])
      path << Line.new(first, last)
      first = Point.new(coord[0], coord[1])
    end
    path
  end
end

def intersection(wire1, wire2)
  intersections = []
  wire1.lines.each do |line1|
    wire2.lines.each do |line2|
      intersections << line1.fast_intersection(line2)
    end
  end
  intersections.compact
end

def distance(wire, point)
  steps = 0
  wire.lines.each do |line|
    if line.include?(point)
      steps += line.points.index(point)
      return steps
    else
      steps += line.length
    end
  end
end

def main
  input = File.readlines("input/three.txt")

  wire1 = Wire.new(input[0])
  wire2 = Wire.new(input[1])

  intersections = intersection(wire1, wire2)

  puts "distance: #{intersections.map(&:distance).sort[1]}"

  total_distance = []
  intersections.each do |intersection|
    total_distance << distance(wire1, intersection) + distance(wire2, intersection)
  end

  puts "steps: #{total_distance.sort[1]}"
end

main
