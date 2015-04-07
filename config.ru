# This file is used by Rack-based servers to start the application.

require 'grape/jbuilder'
require ::File.expand_path('../config/environment', __FILE__)
use Rack::Config do |env|
  env['api.tilt.root'] = 'app/views/api'
end
run Rails.application
