# Glow

It's a tiny wrapper aroung [Gorm](https://github.com/jinzhu/gorm).

#### Features

It provides a utility functions to create a new `*gorm.DB` connection:

```sh
export GORM_DIALECT=sqlite
export GORM_CONNECTION=giraffe.db
```

```go
db, err := glow.OpenDB()
```

**MIT License**
