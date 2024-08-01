# Create a shell script to insert records
cat <<EOF > insert_records.sql
BEGIN TRANSACTION;
EOF

for i in $(seq 1 10000); do
  echo "INSERT INTO counts (id) VALUES ($i);" >> insert_records.sql
done

echo "COMMIT;" >> insert_records.sql

# Execute the script in SQLite
sqlite3 chinook.db < insert_records.sql
