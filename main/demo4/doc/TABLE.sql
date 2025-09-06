-- auto-generated definition
create table users
(
    id         bigint auto_increment
        primary key,
    name       varchar(50)                null,
    email      varchar(191)               not null,
    post_count int                        null,
    created_at datetime(3)                null,
    updated_at datetime(3)                null,
    deleted_at datetime(3)                null,
    password   varchar(100)               not null,
    role       varchar(32) default 'user' null,
    constraint uni_users_email
        unique (email)
);

create index idx_users_deleted_at
    on users (deleted_at);
-- auto-generated definition
create table posts
(
    id             bigint auto_increment
        primary key,
    user_id        bigint       null,
    title          varchar(100) null,
    content        varchar(300) null,
    comment_status varchar(10)  null,
    created_at     datetime     null,
    updated_at     datetime     null,
    constraint posts_users_id_fk
        foreign key (user_id) references users (id)
);

-- auto-generated definition
create table comments
(
    id         bigint auto_increment
        primary key,
    post_id    bigint       null,
    user_id    bigint       null,
    content    varchar(300) null,
    created_at datetime     null,
    updated_at datetime     null,
    constraint comments_posts_id_fk
        foreign key (post_id) references posts (id),
    constraint comments_users_id_fk
        foreign key (user_id) references users (id)
);




