#mysql -u root < 

for file in `ls migration/*.sql`; do
  mysql < ${file}
done
