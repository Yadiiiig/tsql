# tsql - The sequel of ydb (04/20/2023)

database structure:
```tsql
# defining the structure of a table:
tstr://tid/field_name:field_type/..:..

# indices (no clue yet)
tind://


# notes:
# structure will be static, data is being stored in the correct order
# tid: required default field, acts as identifier
```

run database creation: go run main.go -setup=true

