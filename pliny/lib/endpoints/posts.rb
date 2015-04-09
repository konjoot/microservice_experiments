module Endpoints
  class Posts < Base
    namespace "/posts" do
      before do
        content_type :json, charset: 'utf-8'
      end

      get do
        encode serialize(Post.all)
      end

      post do
        param :title, String, required: true
        param :body, String, required: true

        post = Post.new(post_params)
        post.save
        status 201
        encode serialize(post)
      end

      get "/:id" do |id|
        post = Post.first(uuid: id) || halt(404)
        encode serialize(post)
      end

      patch "/:id" do |id|
        param :title, String, blank: false
        param :body, String, blank: false

        post = Post.first(uuid: id) || halt(404)
        post.update(post_params)
        encode serialize(post)
      end

      delete "/:id" do |id|
        post = Post.first(uuid: id) || halt(404)
        post.destroy
        encode serialize(post)
      end

      private

      def serialize(data, structure = :default)
        Serializers::Post.new(structure).serialize(data)
      end

      def post_params
        # pending
      end
    end
  end
end
