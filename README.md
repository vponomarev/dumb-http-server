# dumb-http-server

Сервер-попугай для отладки работы клиентов.

Принимает на вход HTTP(S) запрос, на выход отдаёт JSON с заголовками полученного запроса.

# Параметры
* `-port XXX` (либо переменная окружения `LISTEN_PORT=XXX`) - порт, на котором сервер будет принимать входящие запросы
* `-ssl` (либо переменная окружения `LISTEN_SSL=true`) - включение режима SSL
* `-cert FN` (либо переменная окружения `SSL_CERT=FN`) - имя файла с SSL сертификатом
* `-key FN` (либо переменная окружения `SSL_KEY=FN`) - имя файла с закрытым ключом SSL сертификата