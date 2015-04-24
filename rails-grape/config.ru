# This file is used by Rack-based servers to start the application.

require 'grape/jbuilder'
require ::File.expand_path('../config/environment', __FILE__)
# Unicorn self-process killer
require 'unicorn/worker_killer'

# Max requests per worker
use Unicorn::WorkerKiller::MaxRequests, 3072, 4096

# Max memory size (RSS) per worker
use Unicorn::WorkerKiller::Oom, (1000*(1024**2)), (1500*(1024**2))

use Rack::Config do |env|
  env['api.tilt.root'] = 'app/views/api'
end
run Rails.application
