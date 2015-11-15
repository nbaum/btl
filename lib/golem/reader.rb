module Golem
  class Reader < StringScanner
    def chomp ()
      scan(/(\s|;.+\n)+/)
    end
    def read ()
      chomp
      if scan(/\(/)
        chomp
        return nil if scan(/\)/)
        a = []
        until scan(/\)/)
          a << read()
          a.tail = read() if scan(/\.(?=[()'`,#";\s])/)
        end
        a
      elsif scan(/'/)
        [:"quote", read()]
      elsif s = scan(/"([^"\\]|\\.)*"/)
        s[1..-2]
      elsif f = scan(/\d+\.\d+/)
        f.to_f
      elsif i = scan(/\d+/)
        i.to_i
      elsif s = scan(/[^()'`,#";\s]+/)
        s.to_sym
      elsif scan(/`/)
        [:"quasiquote", read()]
      elsif scan(/,@/)
        [:"unquote-splicing", read()]
      elsif scan(/,/)
        [:unquote, read()]
      elsif c = scan(/#\\.[^()'`,#";\s]+/)
        c
      elsif scan(/#</)
        fail "Unreadable"
      else
        fail "Unexpected `#{getch}'"
      end
    ensure
      chomp
    end
  end
end
