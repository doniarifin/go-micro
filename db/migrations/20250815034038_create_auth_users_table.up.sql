begin;

create table auth_users (
   id       varchar(100) not null,
   email    varchar(200) not null unique,
   password varchar(255) not null,
   primary key ( id )
);

commit;