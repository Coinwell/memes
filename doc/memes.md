# zion-meme (Go)

### Bash
```
sudo -s
docker exec -it $(docker ps --latest --quiet) sh
```

```
docker kill $(docker ps -q)
```

aws s3 sync s3://zion-memes s3://zion-memes-staging