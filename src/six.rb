#!/usr/bin/env ruby
require 'pry'

class Node
  attr_accessor :name, :left, :right

  def initialize(name:)
    @name = name
    @left = nil
    @right = nil
  end
end

def find_or_create_parent(root, name)
  return Node.new(name: name) if root.left.nil? && root.right.nil?
  return root.left if root.left&.name == name
  return root.right if root.right&.name == name

  find_or_create_parent(root.left, name)
  find_or_create_parent(root.right, name)
end

def main
  input = ["COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"]

  root, child = input[0].split(")")

  root_node = Node.new(name: root)
  root_node.left = Node.new(name: child)

  input[1..-1].each do |orbit|
    parent, child = orbit.split(")")

    parent_node = find_or_create_parent(root_node, parent)
    parent_node.left = Node.new(name: child)
  end

  puts root_node.left.lef
end

main
