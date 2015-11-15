class Array
  def disp
    s = "(" + map(&:disp).join(" ")
    s += " . " + tail.to_s if tail
    s += ")"
  end
  attr_accessor :tail
end
