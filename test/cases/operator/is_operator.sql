-- @suit

-- @case
-- @desc:test for is operator in select where
-- @label:bvt
drop table if exists tbl_int;
CREATE TABLE tbl_int (col1 INT);
INSERT INTO tbl_int VALUES (1), (2), (3), (4), (5), (6), (7), (8), (NULL), (NULL);
SELECT * FROM tbl_int WHERE col1 IS NULL;
SELECT * FROM tbl_int WHERE col1 IS NOT NULL;
drop table if exists tbl_int;

drop table if exists tbl_double;
CREATE TABLE tbl_double (col1 DOUBLE);
INSERT INTO tbl_double VALUES (-1.1), (0.0), (1.1), (2.2), (3.3), (4.4), (5.5), (6.6), (NULL), (NULL);
SELECT * FROM tbl_double WHERE col1 IS NULL;
SELECT * FROM tbl_double WHERE col1 IS NOT NULL;
drop table if exists tbl_double;

drop table if exists tbl_datetime;
CREATE TABLE tbl_datetime (col1 DATETIME(6));
INSERT INTO tbl_datetime VALUES
  ("1000-01-01 00:00:01"), ("9999-12-31 23:59:59.999998"),
  ("2017-01-01 00:00:00"), ("2017-01-01 00:00:00.000001"),
  ("2017-02-01 00:00:00"), ("2018-01-01 00:00:00.999999"),
  ("2018-01-01 00:00:01"), ("3019-01-01 10:10:10.101010"), (NULL), (NULL);
SELECT * FROM tbl_datetime WHERE col1 IS NULL;
SELECT * FROM tbl_datetime WHERE col1 IS NOT NULL;
drop table if exists tbl_datetime;

drop table if exists tbl_decimal;
CREATE TABLE tbl_decimal (col1 DECIMAL(65, 30));
drop table if exists tbl_decimal;

drop table if exists t1;
create table t1 (id int not null, str char(10));
insert into t1 values (1, null),(2, null),(3, "foo"),(4, "bar");
select * from t1 where str is null;
select * from t1 where str is not null;
drop table if exists t1;

drop table if exists t1;
drop table if exists t2;
create table t1 (dt datetime not null, t datetime not null);
create table t2 (dt datetime not null);

insert into t1 values ('2001-01-01 1:1:1', '2001-01-01 1:1:1'),
('2001-01-01 1:1:1', '2001-01-01 1:1:1');
insert into t2 values ('2001-01-01 1:1:1'), ('2001-01-01 1:1:1');
SELECT outr.dt FROM t1 AS outr WHERE outr.dt IN (SELECT innr.dt FROM t2 AS innr WHERE outr.dt IS NULL );

drop table if exists t1;
drop table if exists t2;

create table t1 (id int not null, str char(10), index(str));
insert into t1 values (1, null), (2, null), (3, "foo"), (4, "bar");
select * from t1 where str is not null order by id;
select * from t1 where str is null;
drop table if exists t1;

create table t1 (a int, key (a));
insert into t1 values (NULL), (0), (1), (2), (3), (4), (5), (6), (7), (8), (9),
(10), (11), (12), (13), (14), (15), (16), (17), (18), (19);
select * from t1 where not(a is null);
select * from t1 where not(a is not null);
drop table if exists t1;

-- @case
-- @desc:test for is operator in update where
-- @label:bvt
drop table if exists t1;
CREATE TABLE t1 (a INT);
INSERT INTO t1 VALUES (1),(NULL);
-- @ignore{
UPDATE t1 SET a = 2 WHERE a IS NULL;
select * from t1;
-- @ignore}
drop table if exists t1;

drop table if exists t1;
drop table if exists t2;
drop table if exists t3;
create table t1 (a int, b int);
create table t2 (a int, b int);
insert into t1 values (1,1),(2,1),(3,1);
insert into t2 values (1,1), (3,1);
-- @ignore{
select t1.a, t1.b,t2.a, t2.b from t1 left join t2  on t1.a=t2.a where t1.b=1 and t2.b=1 or t2.a is NULL;
-- @ignore}
drop table if exists t1;
drop table if exists t2;
drop table if exists t3;

-- @case
-- @desc:test for is operator in join
-- @label:bvt
drop table if exists t1;
drop table if exists t2;
drop table if exists t3;
CREATE TABLE t1 (
  grp int(11) default NULL,
  a bigint(20) unsigned default NULL,
  c char(10) NOT NULL default ''
);
INSERT INTO t1 VALUES (1,1,'a'),(2,2,'b'),(2,3,'c'),(3,4,'E'),(3,5,'C'),(3,6,'D'),(NULL,NULL,'');
create table t2 (id int, a bigint unsigned not null, c char(10), d int, primary key (a));
insert into t2 values (1,1,"a",1),(3,4,"A",4),(3,5,"B",5),(3,6,"C",6),(4,7,"D",7);
select t1.*,t2.* from t1 left join t2 on (t1.a=t2.a) where t2.id is null;
select t1.*,t2.* from t1 left join t2 on (t1.a=t2.a and t2.id is null);

drop table if exists t1;
drop table if exists t2;
drop table if exists t3;
CREATE TABLE t1 (
  id smallint(5) unsigned NOT NULL,
  name char(60) DEFAULT '' NOT NULL,
  PRIMARY KEY (id)
);
INSERT INTO t1 VALUES (1,'Antonio Paz');
INSERT INTO t1 VALUES (2,'Lilliana Angelovska');
INSERT INTO t1 VALUES (3,'Thimble Smith');

CREATE TABLE t2 (
  id smallint(5) unsigned NOT NULL,
  owner smallint(5) unsigned DEFAULT 0 NOT NULL,
  name char(60),
  PRIMARY KEY (id)
);
INSERT INTO t2 VALUES (1,1,'El Gato');
INSERT INTO t2 VALUES (2,1,'Perrito');
INSERT INTO t2 VALUES (3,3,'Happy');
select t1.name, t2.name, t2.id from t1 left join t2 on (t1.id = t2.owner) where t2.id is null;
select t1.name, t2.name, t2.id from t1 left join t2 on (t1.id = t2.owner) where t2.name is null;
select t1.name, t2.name, t2.id from t2 right join t1 on (t1.id = t2.owner) where t2.id is null;
select t1.name, t2.name, t2.id from t2 right join t1 on (t1.id = t2.owner) where t2.name is null;

drop table if exists t1;
drop table if exists t2;
drop table if exists t3;
CREATE TABLE t1 (id1 INT NOT NULL PRIMARY KEY, dat1 CHAR(1), id2 INT);
INSERT INTO t1 VALUES (1,'a',1);
INSERT INTO t1 VALUES (2,'b',1);
INSERT INTO t1 VALUES (3,'c',2);

CREATE TABLE t2 (id2 INT NOT NULL PRIMARY KEY, dat2 CHAR(1));
INSERT INTO t2 VALUES (1,'x');
INSERT INTO t2 VALUES (2,'y');
INSERT INTO t2 VALUES (3,'z');
SELECT t2.id2 FROM t2 LEFT OUTER JOIN t1 ON t1.id2 = t2.id2 WHERE id1 IS NULL;
SELECT t2.id2 FROM t2 NATURAL LEFT OUTER JOIN t1 WHERE id1 IS NULL;
drop table if exists t1;
drop table if exists t2;
drop table if exists t3;

CREATE TABLE t1 (
  id int(11),
  pid int(11),
  rep_del tinyint(4)
);
INSERT INTO t1 VALUES (1,NULL,NULL);
INSERT INTO t1 VALUES (2,1,NULL);
select * from t1 LEFT JOIN t1 t2 ON (t1.id=t2.pid) AND t2.rep_del IS NULL;
select * from t1 LEFT JOIN t1 t2 ON (t1.id=t2.pid) AND t2.rep_del IS NULL;
drop table if exists t1;

drop table if exists t1;
drop table if exists t2;
drop table if exists t3;
CREATE TABLE t1 (a DATE NOT NULL, b INT);
INSERT INTO t1 VALUES ('1999-05-100',1), ('1999-05-10',2);

CREATE TABLE t2 (a DATETIME NOT NULL, b INT);
INSERT INTO t2 VALUES ('1999-05-10 00:01:01',1), ('1999-05-10 00:00:00',2);

SELECT * FROM t1 WHERE a IS NULL;
SELECT * FROM t2 WHERE a IS NULL;
SELECT * FROM t1 LEFT JOIN t1 AS t1_2 ON 1 WHERE t1_2.a IS NULL;
SELECT * FROM t2 LEFT JOIN t2 AS t2_2 ON 1 WHERE t2_2.a IS not NULL;
SELECT * FROM t1 JOIN t1 AS t1_2 ON 1 WHERE t1_2.a IS NULL;
SELECT * FROM t2 JOIN t2 AS t2_2 ON 1 WHERE t2_2.a IS not NULL;
drop table if exists t1;
drop table if exists t2;
drop table if exists t3;

-- @case
-- @desc:test for is operator in function
-- @label:bvt
drop table if exists t1;
create table t1 (col1 datetime);
insert into t1 values("2004-10-31 15:30:00");
insert into t1 values("2004-12-12 11:22:33");
insert into t1 values("2004-12-12 10:22:59");
insert into t1 values(null);
select count(*) from t1 where YEAR(col1) IS NULL;
select count(*) from t1 where YEAR(col1) IS not NULL;
drop table if exists t1;

-- @case
-- @desc:test for is operator in subquery
-- @label:bvt
drop table if exists t1;
drop table if exists t2;
drop table if exists t3;
create table t1 (id int(10) not null, cur_date datetime not null);
create table t2 (id int(10) not null, cur_date date not null);
insert into t1 (id, cur_date) values (1, '2007-04-25 18:30:22');
insert into t2 (id, cur_date) values (1, '2007-04-25');
select * from t1
where id in (select id from t1 as x1 where (t1.cur_date is null));

select * from t2
where id in (select id from t2 as x1 where (t2.cur_date is null));

insert into t1 (id, cur_date) values (2, '2007-04-26 18:30:22');
insert into t2 (id, cur_date) values (2, '2007-04-26');

select * from t1
where id in (select id from t1 as x1 where (t1.cur_date is null));

select * from t2
where id in (select id from t2 as x1 where (t2.cur_date is null));

-- @case
-- @desc:test for is operator in having
-- @label:bvt
drop table if exists t1;
drop table if exists t2;
drop table if exists t3;
CREATE TABLE `t1` (
  `numeropost` int(8) unsigned NOT NULL,
  `maxnumrep` int(10) unsigned NOT NULL default 0,
  PRIMARY KEY  (`numeropost`)
) ;

INSERT INTO t1 (numeropost,maxnumrep) VALUES (40143,1),(43506,2);

CREATE TABLE `t2` (
      `mot` varchar(30) NOT NULL default '',
      `topic` int(8) unsigned NOT NULL default 0,
      `dt` date,
      `pseudo` varchar(35) NOT NULL default '',
       PRIMARY KEY  (`topic`)
 ) ;

INSERT INTO t2 (mot,topic,dt,pseudo) VALUES ('joce','40143','2002-10-22','joce'), ('joce','43506','2002-10-22','joce');
SELECT * from t2 where topic IN (SELECT topic FROM t2 GROUP BY topic HAVING topic is null);
-- @bvt:issue#3307
SELECT * from t2 where topic IN (SELECT SUM(topic) FROM t1);
-- @bvt:issue
SELECT * from t2 where topic IN (SELECT topic FROM t2 GROUP BY topic HAVING topic is not null);
SELECT * from t2 where topic NOT IN (SELECT topic FROM t2 GROUP BY topic HAVING topic is null);

-- @case
-- @desc:test for is operator in case-when
-- @label:bvt
drop table if exists t1;
drop table if exists t2;
drop table if exists t3;
CREATE TABLE t1 (a varchar(10), PRIMARY KEY (a));
CREATE TABLE t2 (a varchar(10), b date, PRIMARY KEY(a));
INSERT INTO t1 VALUES ('test1');
INSERT INTO t2 VALUES('test1','2016-12-13'),('test2','2016-12-14'),('test3','2016-12-15'),('test4',NULL),('test5',NULL);
SELECT b,
       CASE  WHEN b is NULL then 'found' ELSE 'not found' END FROM t2;
SELECT b,
       CASE  WHEN b is not NULL then 'found' ELSE 'not found' END FROM t2;