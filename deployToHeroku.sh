

docker build -t linebot . --no-cache --platform linux/amd64

docker tag linebot registry.heroku.com/line-bot-stepn/web

docker push registry.heroku.com/line-bot-stepn/web

heroku container:release web -a line-bot-stepn 
