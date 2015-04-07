module Posts
  class API < Grape::API
    prefix :api
    format :json
    formatter :json, Grape::Formatter::Jbuilder
    version 'v1', using: :header, vendor: 'maksimov'

    resources :posts do
      desc "index posts."
      get jbuilder: 'posts/index' do
        @posts = Post.all
      end

      desc "show post."
      params do
        requires :id, type: Integer, desc: "Post's ID."
      end
      get ':id', jbuilder: 'posts/show' do
        @post = Post.find(params[:id])
      end

      desc "create post."
      params do
        requires :title, type: String, desc: "Post's title."
        requires :body, type: String, desc: "Post's body."
      end
      post jbuilder: 'posts/show' do
        @post = Post.create(declared(params))
      end

      desc "update post."
      params do
        requires :id, type: Integer, desc: "Post's ID."
        optional :title, type: String, desc: "Post's title."
        optional :body, type: String, desc: "Post's body."
      end
      put ':id', jbuilder: 'posts/show' do
        @post = Post.find(params[:id])
        @post.update_attributes(params)
      end

      desc "destroy post."
      params do
        requires :id, type: Integer, desc: "Post's ID."
      end
      delete ':id' do
        Post.find(params[:id]).destroy!
        body status: 'ok'
      end
    end
  end
end
