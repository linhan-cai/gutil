golang simple storage

```sql

create database kv_storage;
use kv_storage;

drop table kv_1;
create table kv_1 (
    `id` int primary key auto_increment,
    `k` varchar(255) not null,
    `v` text,
	`version` int,
    unique index (k)
) default charset=utf8mb4;


TODO


```
- 分库
- 分表