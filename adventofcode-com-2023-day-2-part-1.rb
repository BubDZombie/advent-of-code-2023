#!/usr/bin/ruby

class MissingColorException < Exception
end

class LowCountException < Exception
end

class Set
  attr_accessor :counts
  def initialize
    @counts = {}
  end

  def power
    power = 1
    ['red', 'green', 'blue'].each do |color|
      if @counts.has_key?(color)
        power *= @counts[color]
      end
    end
    return(power)
  end
end

class Game
  attr_accessor :game_number
  attr_accessor :sets

  def check(bag)
    puts "Checking game #{@game_number} with sets #{@sets} against bag #{bag}."
    @sets.each do |set|
      set.counts.each do |color, count|
        puts "Checking if there are #{count} #{color} in #{bag}."
        unless bag.has_key?(color)
          raise MissingColorException.new("No #{color} in the bag for game #{@game_number}.")
        end
        unless bag[color] >= count
          raise LowCountException.new("Not enough #{color} in the bag for game #{@game_number}.")
        end
      end
    end
    return(@game_number)
  end

  def minimum
    min_set = Set.new
    @sets.each do |set|
      set.counts.each do |color, count|
        unless min_set.counts.has_key?(color)
          min_set.counts[color] = count
        end
        if count > min_set.counts[color]
          min_set.counts[color] = count
        end
      end
    end
    puts("The minimum set for game #{@game_number} is #{min_set.counts} and its power is #{min_set.power}.")
    return(min_set)
  end
end

bag = {}
ARGV.each_slice(2) do |count, color|
  puts "Adding #{count} #{color} to bag"
  bag[color] = count.to_i
end

sum = 0
power_sum = 0
$stdin.each_line do |line|
  game = Game.new
  title, game_string = line.split(':')
  game.game_number = title.split[1].to_i
  game.sets = []
  game_string.split(';').each do |set_string|
    set = Set.new
    set_string.split(',').each do |count_string|
      count, color = count_string.split
      count = count.strip.to_i
      color.strip!
      puts "Adding #{count} #{color} to set in game #{game.game_number}."
      set.counts[color] = count
    end
    game.sets << set
    puts(set.counts)
  end
  begin
    min_set = game.minimum
    power_sum += min_set.power
    game.check(bag)
    sum += game.game_number
  rescue Exception => e
    puts e
    puts e.backtrace
  end
  puts("Game sum: #{sum}, power sum: #{power_sum}")
end
