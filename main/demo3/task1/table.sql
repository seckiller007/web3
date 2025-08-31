-- auto-generated definition
create table students
(
    id    bigint unsigned auto_increment comment '主键'
        primary key,
    name  varchar(64)      not null comment '学生姓名',
    age   tinyint unsigned not null comment '年龄',
    grade varchar(32)      not null comment '年级'
)
    comment '学生信息表';

