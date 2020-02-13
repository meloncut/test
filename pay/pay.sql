CREATE DATABASE IF NOT EXISTS trade DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;

use trade;
create table orders
(
    id int unsigned auto_increment,
    order_code varchar(40) not null comment '订单号',
    user_id int unsigned not null default 0 comment '用户id',
    original_price int unsigned not null default 0 comment '原价',
    receipt_amount int unsigned not null default 0 comment '实收',
    status tinyint unsigned not null default 0 comment '订单状态',
    wealth_id int unsigned default 0 not null comment '对应的财富ID',
    amount int unsigned default 0 not null comment '充值财富数量',
    unit_price int unsigned default 0 not null  comment '商品单价',
    quantity int unsigned default 0 not null comment '商品数量',
    created_at timestamp default CURRENT_TIMESTAMP null comment '创建时间',
    updated_at timestamp default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP null,
    paid_at timestamp null,
    closed_at timestamp null,
    constraint orders_pk
        primary key (id)
)
    comment '订单表';

create unique index orders_order_code_uindex
    on orders (order_code);

create table wealth
(
    id         int unsigned auto_increment
        primary key,
    name       varchar(60)      default ''                not null,
    type       tinyint unsigned default '0'               not null,
    created_at timestamp        default CURRENT_TIMESTAMP not null
)
    comment '财富表';
create table wealth_accounts
(
    id         int unsigned auto_increment comment '主键id'
        primary key,
    wealth_id  int unsigned default '0'               not null comment '关联的财富id',
    user_id    int unsigned default '0'               not null comment '用户id',
    amount     int unsigned default '0'               not null,
    created_at timestamp    default CURRENT_TIMESTAMP not null,
    updated_at timestamp    default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP
)
    comment '用户财富账户表';

alter table wealth_accounts add unique index wealth_accounts_user_id_wealth_id(user_id,wealth_id);


