# gptbot
Web-клиент для GPTBot (серверный компонент).

## Системные требования

- Go >= 1.20

## Установка и запуск из командной строки (режим разработки)

1. git clone git@github.com:nicksedov/gptbot.git
2. cd gptbot
3. go build -ldflags="-s -w"
4. ./gptbot -bot=<*bot_token*> -openai=<*openai_token*> -proxy=<*proxy_host:port*> -proxy.user=<*proxy_user*> -proxy.password=<*proxy_pass*>

Параметры прокси опциональны и требуются для успешного доступа к серверам api.openai.com из России, прокси-хост должен географически располагаться за ее пределами. 

## Запуск юнит-тестов
Для успешного запуска юнит-тестов требуется создание конфигурационного файла `secrets.yaml` в директории проекта.
```yaml
Proxy: <proxy_host:port>
ProxyUser: <proxy_user>
ProxyPassword: <proxy_pass>
BotToken: <bot_token>
OpenAIToken: <openai_token>
TelegramChatID: <test_chat_id>
```
## Развертывание в продакшн

1. cd /opt
2. mkdir /opt/gptbot-server
3. mkdir /opt/gptbot-database (если ранее не ставился gptbot-server)
4. cd /opt/gptbot-server 
5. mcedit settings.yaml
```yaml
server:
  host: 
  port: 3445

database:
  path: /opt/gptbot-database/calendar.sqlite
```
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