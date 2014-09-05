gopinbasdfsdfsd (https://travis-ci.org/kshvakov/gopinba.svg?branch=master)](https://travis-ci.org/kshvakov/gopinba)
=======

Клиент на [Go](http://golang.org) для  [Pinba](http://pinba.org/)

Отсылает данные о запросе (время выполнения, название скрипта, хостнейм и т.д. + таймеры [тыц](/Example/web.go#L19), [тыц](https://github.com/tony2001/pinba_engine/wiki/PHP-extension#pinba_timer_start))  в "пинбасервер" *

*"пинбасервер" 
Pinba is a MySQL storage engine that acts as a realtime monitoring/statistics server for PHP (и не только PHP, но и Python, Node.js, Ruby и все что может посылать запросы по UDP)

"Пинбасервер" позволяет просматривать данные в реальном времени, строить отчеты/репорты  [тыц](/Example/pinba_reports.sql), [тыц](https://github.com/tony2001/pinba_engine/wiki/Reports), [ТЫЦ](https://github.com/tony2001/pinba_engine/wiki/Usage-examples)

Но и этого ["маловато будет"](http://www.youtube.com/watch?v=ZlxJ0jtSF1Y), для тех кому "маловато" есть [Graphite](http://graphite.wikidot.com/) и [InfluxDB](http://influxdb.com/) в которых можно прекрасно эту самую статистику агрегировать,
а к ним и [Grafana](http://grafana.org/) прикрутить для наглядности