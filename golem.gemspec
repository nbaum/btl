lib = File.expand_path('../lib', __FILE__)
$LOAD_PATH.unshift(lib) unless $LOAD_PATH.include?(lib)
require "golem/version"

Gem::Specification.new do |s|
  s.name        = "golem"
  s.version     = Golem::VERSION
  s.date        = Time.now.strftime("%Y-%m-%d")
  s.authors     = ["Nathan Baum"]
  s.email       = "n@p12a.org.uk"
  s.executables = ["golem", "igolem"]
  s.files       = Dir["lib/**/*.rb"]
  s.homepage    = "http://www.github.org/nbaum/golem"
  s.license     = "MIT"
end
