- Admin data
```shell
create /expense_tracker
create /expense_tracker/admin
create /expense_tracker/admin/server_port
create /expense_tracker/admin/server_port/server "expense_tracker_admin"
create /expense_tracker/admin/server_port/port ":8000"
```

- Common data
```shell
create /expense_tracker/common
create /expense_tracker/common/mongodb
create /expense_tracker/common/mongodb/uri "mongodb://localhost:27017"
create /expense_tracker/common/mongodb/db_name "expense_tracker"
```

```shell
create /expense_tracker/common/auth
create /expense_tracker/common/auth/secret_key "secret_key_admin"
```