create table images(tag varchar(100),
 id integer,
  path varchar(100),
   size integer);

create table homepagenews(imgurl varchar(200),
linkurl varchar(200),
title varchar(200),
id serial);

create table homepageartical (id serial,
imgurl varchar(200),
linkurl varchar(200),

date date);




insert into homepagenew(imgurl,linkurl,title) values(
    'http://129.204.193.192:4400/get?tag=&name=5.jpg',
    'http://129.204.193.192:4400/get?tag=&name=5.jpg',
    'http://129.204.193.192:4400/get?tag=&name=5.jpg'
);

insert into homepageartical(imgurl,linkurl,date)values(
    'http://129.204.193.192:4400/get?tag=&name=11.jpg',
    'http://129.204.193.192:4400/get?tag=&name=11.jpg',
    CURRENT_DATE
);


select imgurl, linkurl, title from homepagenew;

select imgurl, linkurl, date from homepageartical order by id asc offset 1 limit 10;