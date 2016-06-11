# ruby 2.0

sum = 0

[*(1..999)].each do |i|
  sum += i if i % 3 == 0 || i % 5 == 0
end
p sum
