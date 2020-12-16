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

## TODO
- http server handler refactoring
- 스케줄러 자동 시작(오전 9시 ~ 오후 6시)
