Percona Server + Pinba Engine
====

* Добавляем репозиторий [Percona](http://www.percona.com/doc/percona-server/5.5/installation.html#using-percona-software-repositories?id=repositories:start)

* Устанавливаем percona-server 5.5/5.6
* Скачиваем исходный код (важно, source code должен быть тойже версии что и установленный сервер). Для Ubuntu apt-get source ...

собираем (пример)

```
cd percona-server-5.6.20-68.0/
./BUILD/autorun.sh
./configure #получаем список зависимостей и ставим...ставим...ставим
make #тут немного ждем)
```

скачиваем pinba engine https://github.com/tony2001/pinba_engine/tree/devel

конфигурим ([зависимости](https://github.com/tony2001/pinba_engine/wiki/Installation#some-ubuntu-specific-hints))

```
./configure --with-mysql=../percona-server-5.6.20-68.0/ --libdir=/usr/lib/mysql/plugin

make 
make install
```

далее по [инструкции](https://github.com/tony2001/pinba_engine/wiki/Installation#pinba-engine-installation) пользуемся 

PS: лучше использовать отдельный сервер/ВМ mysql для pinba, т.к. при обновлении плагин придется пересобирать снова, а так: собрали и забыли (для pinba фиксы в самом mysql особо параллельны)