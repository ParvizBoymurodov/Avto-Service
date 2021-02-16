--TODO:Чадвали бакайдгири
create table if not exists registered (
                                          id bigserial not null,
                                          name text not null,
                                          surname text not null,
                                          login text unique  not null,
                                          password text not null,
                                          remove bool default false
);

--TODO:Протсеси бакайдгири ба барнома
insert into registered (name, surname, login, password) VALUES ('Ном', 'Насаб', 'Логин', md5('Парол') );


--TODO:Протсеси дохилшави ба барнома
select login, password from registered where login = 'Логин';

--TODO:Ивазкунии парол
update registered set password = md5('Парол') where id= 1;

--TODO:Удалит кардани истифодабар
update registered set remove = true where id = 1;

--TODO:Руйхати истифода барандагон
select name, surname, login from registered where remove = false;

--TODO:Чадвали сервисхо
create table if not exists services (
                                        id bigserial not null,
                                        service_name text not null,
                                        region text not null,
                                        remove bool default false
);


--TODO:Илова кардани сервис
insert into services (service_name, region) values ('Номи сервис', 'Адрес');

--TODO:Руйхати сервисхо
select id, service_name from services where remove = false;

--TODO:Тағир додани сервисхо
update services set service_name = 'Номи нав ё дурусти сервис', region='Макони нав ё иваз кардан' where id = 1;

--TODO:Удалит кардани сервис
update services set remove = true where id=1;


--TODO:Чадвали магозахои сервисхо
create table if not exists servicemarket (
                                             id bigserial not null,
                                             market_name text not null,
                                             service_id INTEGER not null,
                                             title text not null,
                                             remove bool default false
);

--TODO:Илова кардани магоза барои сервисхо
insert into  servicemarket (market_name, service_id, title) VALUES ('Номи магоза', 1, 'Маьлумот дар бораи магоза');

--TODO:Руйхати магозахо ва сервисхо
select sm.market_name,  sm.title, s.service_name from servicemarket sm inner join services s on s.id = sm.service_id where sm.remove = false and s.remove = false;

--TODO:Удалит кардани магоза
update servicemarket set remove = true where id= 'id';






