create table if not exists anime
(
    id          INTEGER
        constraint subject_pk
            primary key autoincrement,
    bangumi_id  INTEGER default 0  not null, -- bangumi的ID
    category_id INTEGER default 0  not null, -- 分类
    title       TEXT    default '' not null, -- 标题
    season      INTEGER default 0  not null, -- 第几季
    cover       TEXT    default '' not null, -- 封面图
    total       INTEGER default 0  not null, -- 总集数
    rss_url     TEXT    default '' not null, -- RSS链接
    play_time   INTEGER default 0  not null, -- 放送时间
    create_time INTEGER default 0  not null,
    update_time INTEGER default 0  not null
);

create unique index if not exists anime_title_season_uindex
    on anime (title, season);


create table if not exists category
(
    id          INTEGER
        constraint category_pk
            primary key autoincrement,
    name        TEXT    default '' not null,
    origin      INTEGER default 2  not null,
    create_time INTEGER default 0  not null,
    update_time INTEGER default 0  not null
);

create unique index if not exists category_name_uindex
    on category (name);

create table if not exists rule
(
    id               INTEGER
        constraint rule_pk
            primary key autoincrement,
    name             TEXT    default '' not null,
    must_contain     TEXT    default '' not null,
    must_not_contain TEXT    default '' not null,
    use_regex        INTEGER default 2  not null,
    episode_filter   TEXT    default '' not null,
    smart_filter     INTEGER default 2  not null,
    create_time      INTEGER default 0  not null,
    update_time      INTEGER default 0  not null
);

create unique index if not exists rule_name_uindex
    on rule (name);

CREATE VIEW if not exists anime_info_view as
select a.ID,
       a.bangumi_id,
       a.category_id,
       c.name as category_name,
       a.title,
       a.season,
       a.cover,
       a.total,
       a.rss_url,
       a.play_time,
       a.create_time,
       a.update_time
from anime a
         inner join category c on a.category_id = c.id;
