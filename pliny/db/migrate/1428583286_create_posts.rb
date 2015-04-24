Sequel.migration do
  change do
    create_table(:posts) do
      integer      :id, primary_key: true
      String       :title, null: false
      String       :body,  null: false
      timestamptz  :created_at, default: Sequel.function(:now), null: false
      timestamptz  :updated_at
    end
  end
end
