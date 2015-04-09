class Serializers::Post < Serializers::Base
  structure(:default) do |arg|
    {
      created_at: arg.created_at.try(:iso8601),
      id:         arg.uuid,
      title:      arg.title,
      body:       arg.body,
      updated_at: arg.updated_at.try(:iso8601),
    }
  end
end
