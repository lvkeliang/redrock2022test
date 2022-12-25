## 登录注册表：

| Field    | Type        | Null | Key | Default | Extra          |
|----------|-------------|------|-----|---------|----------------|
| id       | bigint      | NO   | PRI | NULL    | auto_increment |
| name     | varchar(20) | NO   |     | \0      |                |
| password | varchar(20) | NO   |     | \0      |                |
| email    | varchar(20) | NO   |     | \0      |                |
| phone    | varchar(20) | NO   |     | \0      |                |


##借阅记录表：
注：status的值：1为已规划，0为未归还

| Field       | Type     | Null | Key | Default             | Extra          |
|-------------|----------|------|-----|---------------------|----------------|
| id          | bigint   | NO   | PRI | NULL                | auto_increment |
| user_id     | bigint   | NO   |     | -1                  |                |
| book_id     | bigint   | NO   |     | -1                  |                |
| status      | TINYINT  | NO   |     | 0                   |                |
| borrow_time | DATETIME | NO   |     | 1000-01-01 00:00:00 |                |
| return_time | DATETIME | NO   |     | 1000-01-01 00:00:00 |                |

##书籍库表：

| Field        | Type        | Null | Key | Default             | Extra          |
|--------------|-------------|------|-----|---------------------|----------------|
| id           | bigint      | NO   | PRI | NULL                | auto_increment |
| book_name    | varchar(20) | NO   |     | \0                  |                |
| author       | varchar(20) | NO   |     | \0                  |                |
| publisher    | varchar(20) | NO   |     | \0                  |                |
| publish_time | DATETIME    | NO   |     | 1000-01-01 00:00:00 |                |
| price        | varchar(20) | NO   |     | $0                  |                |
| amount       | bigint      | NO   |     | 0                   |                |