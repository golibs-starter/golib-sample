create table orders
(
    id int auto_increment,
    total_amount int not null,
    created_at timestamp default NOW() not null,
    constraint orders_pk
        primary key (id)
);
