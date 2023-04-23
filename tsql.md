# tsql - The sequel of ydb (04/20/2023)

database structure:
```tsql
# defining the structure of a table:
tstr://tid/field_name:field_type/..:..

# example:
tid:name:string/url:string/employees:int/enabled:bool/score:float32/
0:foo0/foo0.com/0/false/0.331657/

# indices (no clue yet)
tind://


# notes:
# structure will be static, data is being stored in the correct order
# tid: required default field, acts as identifier
```

```go
// code for tid retrieval
end := 7
for end < len(table.Layout) && table.Layout[end] != ':' {
   end++
}

tid := str[7:end]
result := table.Layout[end+1:]
```
run database creation: go run main.go -setup=true

