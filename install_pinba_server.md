Percona Server + Pinba Engine
====

* Добавляем репозиторий [Percona](http://www.percona.com/doc/percona-server/5.5/installation.html#using-percona-software-repositories?id=repositories:start)

* Устанавливаем percona-server 5.5/5.6
* Скачиваем исходный код (важно, source code должен быть тойже версии что и установленный сервер)

собираем 

```
cd percona-server-5.6.20-68.0/ #пример
./BUILD/autorun.sh
./configure
make #тут немного ждем)
```

скачиваем pinba engine https://github.com/tony2001/pinba_engine/tree/devel

конфигурим 

```
./configure --with-mysql=../percona-server-5.6.20-68.0/ --libdir=/usr/lib/mysql/plugin

make 
make install
```

далее по [инструкции](https://github.com/tony2001/pinba_engine/wiki/Installation#pinba-engine-installation) пользуемся 

