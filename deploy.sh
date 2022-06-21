# build docker image on Mac m1 os
docker build -t linebot . --no-cache --platform linux/amd64

# tag the docker image with /web to let heroku get the image
docker tag linebot registry.heroku.com/line-bot-stepn/web

# push to heroku registry
docker push registry.heroku.com/line-bot-stepn/web

# run the image in heroku container
heroku container:release web -a line-bot-stepn 
