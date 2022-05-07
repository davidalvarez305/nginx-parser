Parse NGINX Access logs and print count of crawls by bot.

Download and unzip files from your web server:

```
cd nginx-parser && ./unzip.sh -f SOME_FOLDER -k KEY_NAME.pem -ip YOUR_SERVER_IP
```

Run go program to parse Nginx & print results:

```
go run main.go PATH_TO_FOLDER_USED_ABOVE/nginx
```