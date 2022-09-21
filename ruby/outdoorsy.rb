# frozen_string_literal: true

require 'logger'
require 'optparse'

# set up logging
@logger = Logger.new($stdout)
@logger.level = Logger::INFO
@logger.datetime_format = '%Y-%m-%d %H:%M'

# parse CLI arguments
@options = {}
optparse = OptionParser.new do |opts|
  opts.banner = 'Usage: outdoorsy.rb [options] filename1 filename2 ...'
  opts.on('-v', '--verbose', 'Show extra information') do
    @options[:verbose] = true
  end
  opts.on('-n', '--sort_by_name', 'Sort by full name') do
    @options[:sort_by_name] = true
  end
  opts.on('-t', '--sort_by_vehicle_type', 'Sort by vehicle type') do
    @options[:sort_by_vehicle_type] = true
  end
end
optparse.parse!

# Empty customer_list for us to populate
@customer_list = []

def determine_delimiter(line)
  return ',' if line.include?(',') && !line.include?('|')
  return '|' if line.include?('|') && !line.include?(',')
  raise(Exception.new, 'Both delimiter types found') if line.include?(',') && line.include?('|')

  raise(Exception.new, 'no delimiters found')
end

def load_customer(line)
  customer = line.split(determine_delimiter(line))
  raise(Exception.new, "Customer did not contain exactly 6 values: #{customer}") unless customer.length == 6

  @customer_list.append(customer)
end

def load_file(filename)
  File.open(filename).each do |line|
    @logger.debug("Loading customer from line #{line.strip}")
    load_customer(line)
  end
end

def load_files(filenames)
  filenames.each do |filename|
    if File.exist?(filename)
      @logger.debug("Loading file #{filename}")
      load_file(filename)
    else
      @logger.info("Skipping invalid file: #{filename}")
    end
  end
end

def sort_customers_by_field(field_number)
  @customer_list = @customer_list.sort_by { |c| c[field_number].downcase }
end

def output_customers
  @customer_list.each do |customer|
    puts "#{customer[0]} #{customer[1]} #{customer[2]} #{customer[3]} #{customer[4]} #{customer[5]}"
  end
  puts "Total customers: #{@customer_list.length}"
end

@logger.level = Logger::DEBUG if @options[:verbose]
@logger.debug('Debug logging is on.')

filenames = ARGV
load_files(filenames)

if @options[:sort_by_name]
  @logger.debug('Sorting customers by full name')
  sort_customers_by_field(0)
elsif @options[:sort_by_vehicle_type]
  @logger.debug('Sorting customers by vehicle type')
  sort_customers_by_field(3)
end

output_customers
