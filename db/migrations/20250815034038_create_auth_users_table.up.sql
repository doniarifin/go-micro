begin;
create table if not exists auth_users (
   id varchar(100) not null,
   email varchar(200) not null unique,
   password varchar(255) not null,
   role varchar(500),
   primary key (id)
);
commit;