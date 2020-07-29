[![CircleCI](https://circleci.com/gh/kimeuichan/stock-to-slack.svg?style=shield)](https://circleci.com/gh/kimeuichan/stock-to-slack)
# stock-to-slack

## Build
```shell script
sh ./build.sh
```

## Run with docker
```shell script
docker run -d --name <container_name> \
-e SLACK_WEBHOOK_URL=<slack_url> \
-e STOCK_NUMBER=<stock_id> \
stock-to-slack:latest
```
