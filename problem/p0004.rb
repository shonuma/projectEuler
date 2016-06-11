# ruby 2.0.0
def palindrome?(v)
  str = v.to_s.split('')
  last = str.length
  for i in 0..last do
    if str[i] != str[last-1-i]
      return false
    elsif i >= last-i
      break
    end
  end
  true
end

max = 0
output = ""

[*(100..999)].each do |r|
  [*(100..999)].each do |l|
    if r*l > max && palindrome?(r*l)
      max = r*l
      output = "#{r} * #{l} = #{r*l}"
    end
  end
end

puts output

