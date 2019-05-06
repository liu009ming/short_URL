create table short_url(
	shortUrl bigint(64) not null auto_increment,
	longUrl varchar(10240) not null,
	primary key (shortUrl)
)engine=InnoDB default charset=utf8;