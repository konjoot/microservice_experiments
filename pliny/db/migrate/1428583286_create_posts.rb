Sequel.migration do
  change do
    create_table(:posts) do
      uuid         :uuid, default: Sequel.function(:uuid_generate_v4), primary_key: true
      String       :title, null: false
      String       :body,  null: false
      timestamptz  :created_at, default: Sequel.function(:now), null: false
      timestamptz  :updated_at
    end
  end
end
