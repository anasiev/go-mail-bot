# go-mail-bot 
The bot is designed on Golang for IoT device as Orange Pi, Raspberry Pi, etc..
Bot is the basis for the further development of remote monitoring via email.

### Functions:
1) Bot works on the local network and determines the external address using the Yandex service http//yandex.ru/internet.
2) Bot send email with ip adress from yandex mail service to any email.

### Requirements
golang 1.10

### Build and install
go build getIp.go

Add background task in cron:
crontab -e

@reboot sleep 30 && /home/john/getIp > /home/john/iot.log 2>&1 &
