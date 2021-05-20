# GORM
- stands for Goland ORM (object relational model -- tying an OOP language and relational db together)
- GORM supports:
  - associations
    - has one
    - has many
    - belongs to
    - many to many
    - polymorphism
  - hooks
    - before
    - after create
    - save
    - update
    - delete
    - find
  - preloading (eager loading)
  - transactions
  - Composite Primary Key
  - SQL Builder (db.Table("users").Where("name=?", "steph").Select("name, age").Row())
  - Auto Migrations
  - Plugins based on GORM callbacks
  - GORM provides official support for SQLite, MySQL, Postgres and SQL Server
  