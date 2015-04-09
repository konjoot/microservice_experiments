require "spec_helper"

describe Endpoints::Posts do
  include Rack::Test::Methods

  describe "GET /Posts" do
    it "succeeds" do
      get "/Posts"
      assert_equal 200, last_response.status
    end
  end
end
