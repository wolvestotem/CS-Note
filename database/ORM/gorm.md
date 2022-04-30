# GORM

[toc]

Document: [doc](https://gorm.io/docs/)
API: [api](https://pkg.go.dev/gorm.io/gorm#DB.Assign)

Create:
```go
//default values
type User struct {
  ID   int64
  Name string `gorm:"default:galeone"`
  Age  int64  `gorm:"default:18"`
}

// create by struct pointer
user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
result := db.Create(&user) // pass pointer of data to Create
result.Error        // returns error

var users = []User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
db.Create(&users)

// create by map[string]interface{}
db.Model(&User{}).Create(map[string]interface{}{
  "Name": "jinzhu", "Age": 18,
})

// batch insert from `[]map[string]interface{}{}`
db.Model(&User{}).Create([]map[string]interface{}{
  {"Name": "jinzhu_1", "Age": 18},
  {"Name": "jinzhu_2", "Age": 20},
})

```

Query
```go
// Get the first record ordered by primary key
db.First(&user)
// SELECT * FROM users ORDER BY id LIMIT 1;

// Get one record, no specified order
db.Take(&user)
// SELECT * FROM users LIMIT 1;

// Get last record, ordered by primary key desc
db.Last(&user)
// SELECT * FROM users ORDER BY id DESC LIMIT 1;
```

#### [约定](https://gorm.io/zh_CN/docs/conventions.html)

Gorm喜欢约定而不是配置

默认下蛇形命名就是表名/列名，但是也可以用户指定

#### embeded

[模型定义](https://gorm.io/zh_CN/docs/models.html)

- 对 tag 描述见 doc
- 还有关联 tag

在JOIN中用得上

```golang
type UserDALModel struct {
    gorm.Model `gorm:"embedded"`
    models.User `gorm:"embedded"`
}
```

```go
type Author struct {
	Name  string
	Email string
}

type Blog struct {
	ID      int
	Author  Author `gorm:"embedded"`
	Upvotes int32
}
// 等效于
type Blog struct {
	ID    int64
	Name  string
	Email string
	Upvotes  int32
}
```

You can also specify a prefix if you want with `embedded_prefix`

```go
type Blog struct {
   ID      int
   Author  Author `gorm:"embedded;embeddedPrefix:author_"`
   Upvotes int32
}
// 等效于
type Blog struct {
   ID          int64
   AuthorName  string
   AuthorEmail string
   Upvotes     int32
}
```

相同字段待验证
