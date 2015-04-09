require "spec_helper"

describe Endpoints::Posts do
  include Committee::Test::Methods
  include Rack::Test::Methods

  def app
    Routes
  end

  def schema_path
    "./schema/schema.json"
  end

  before do
    @post = Post.create

    # temporarily touch #updated_at until we can fix prmd
    @post.updated_at
    @post.save
  end

  describe 'GET /Posts' do
    it 'returns correct status code and conforms to schema' do
      get '/Posts'
      assert_equal 200, last_response.status
      assert_schema_conform
    end
  end

=begin
  describe 'POST /Posts' do
    it 'returns correct status code and conforms to schema' do
      header "Content-Type", "application/json"
      post '/Posts', MultiJson.encode({})
      assert_equal 201, last_response.status
      assert_schema_conform
    end
  end
=end

  describe 'GET /Posts/:id' do
    it 'returns correct status code and conforms to schema' do
      get "/Posts/#{@post.uuid}"
      assert_equal 200, last_response.status
      assert_schema_conform
    end
  end

  describe 'PATCH /Posts/:id' do
    it 'returns correct status code and conforms to schema' do
      header "Content-Type", "application/json"
      patch "/Posts/#{@post.uuid}", MultiJson.encode({})
      assert_equal 200, last_response.status
      assert_schema_conform
    end
  end

  describe 'DELETE /Posts/:id' do
    it 'returns correct status code and conforms to schema' do
      delete "/Posts/#{@post.uuid}"
      assert_equal 200, last_response.status
      assert_schema_conform
    end
  end
end
