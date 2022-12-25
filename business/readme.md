## user

| Field    | Type        | Null | Key | Default | Extra          |
|----------|-------------|------|-----|---------|----------------|
| id       | int         | NO   | PRI | NULL    | auto_increment |
| name     | varchar(20) | NO   |     | NULL    |                |
| password | varchar(45) | NO   |     | NULL    |                |

## warehose
货物通过一个数组储存，元素是map[string]int类型，以“货物名:数量”的形式储存

| Field        | Type         | Null | Key | Default                         | Extra          |
|--------------|--------------|------|-----|---------------------------------|----------------|
| id           | int          | NO   | PRI | NULL                            | auto_increment |
| warehoseName | varchar(20)  | NO   |     | newWarehose                     |                |
| goods        | varchar(100) | NO   |     | [{"goods1":100} {"goods2":999}] |                |