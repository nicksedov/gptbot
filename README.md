# gptbot
Web-клиент для GPTBot (серверный компонент).

## Системные требования

- Go >= 1.20

## Установка 

1. git clone git@github.com:nicksedov/gptbot.git
2. cd gptbot
3. 

## Настройка файла конфигурации
Параметры прокси опциональны и требуются для успешного доступа к серверам api.openai.com из России, прокси-хост должен географически располагаться за ее пределами.
Шаблон файла конфигурации 
```yaml
server:
  host: 
  port: 3445

proxy:
  host: proxy.com
  port: 9999
  user: username
  password: ********
  
database:
  host:     localhost
  port:     5432
  db_name:  gptbot
  user:     postgres
  password: ********
  ssl_mode: disable

openai:
  api_token: ************
  model: gpt-3.5-turbo-0125
  
telegram:
  bot_token: ************
  service_chat: 12345678901234

logger:
  filename:   /var/log/gptbot/gptbot.log 
  maxsize:    10
  maxbackups: 10
  maxage:     30
  compress:   true
```
## Запуск из командной строки (режим разработки)
Для запуска требуется предварительно созданный конфигурационный файл `settings.yaml` в директории проекта, рядом с текущим файлом README.md.
1. go build -ldflags="-s -w"
2. ./gptbot
3. В случае нестандартного расположения или названия файла `settings.yaml`: ./gptbot -config=</path/to/settings.xml>

## Запуск юнит-тестов
Для успешного запуска юнит-тестов также требуется предварительно созданный конфигурационный файл `settings.yaml` в директории проекта, рядом с текущим файлом README.md.

## Развертывание в продакшн

1. cd /opt
2. mkdir /opt/gptbot-server
3. mkdir /opt/gptbot-database (если ранее не ставился gptbot-server)
4. cd /opt/gptbot-server 
5. mcedit settings.yaml

6. cd /etc/systemd/system
7. mcedit gptbot-server.service
```properties
[Unit]
Description=ChatGPT backend server

[Install]
WantedBy=multi-user.target
After=network.target

[Service]
Type=simple
Environment=GIN_MODE=release
ExecStart=/opt/gptbot-server/gptbot -bot=<bot_token> -openai=<openai_token> -proxy=<proxy_host:port> -proxy.user=<proxy_user> -proxy.password=<proxy_pass>
WorkingDirectory=/opt/gptbot-server
Restart=always
RestartSec=5
StandardOutput=syslog
StandardError=syslog
SyslogIdentifier=%n
```
8. systemctl daemon-reload
9. systemctl enable gptbot-server
10. systemctl start gptbot-server
11. Проверки:
   - systemctl status gptbot-server

## Обновление сервиса
Быстрое обновление производится запуском скрипта из рабочей директории проекта
```bash
#!/bin/bash

PROJECT=gptbot-server
PROJECT_SOURCEDIR=<project_dir>
PROJECT_DEPLOY_PATH=/opt/$PROJECT

cd $PROJECT_SOURCEDIR

echo "Building $PROJECT"

go build -ldflags="-s -w"

if [ "$EUID" -ne 0 ]; then
    echo "Unable to deploy. Please run as root"
    exit
else
    systemctl stop $PROJECT

    echo "Deploying $PROJECT"
    cp gptbot $PROJECT_DEPLOY_PATH

    systemctl start $PROJECT 
fi
```
