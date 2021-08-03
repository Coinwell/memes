# zion-memes

[How To Deploy](ops/ansible/README.md)

## Bash
```
ssh -i ~/.ssh/zion ubuntu@memes.getzion.com
sudo -s
docker exec -it $(docker ps --latest --quiet) bash
```


## Database
```
truncate table memes;
```

## Logs

```
docker ps -q | xargs -L 1 docker logs -f
```

[Memes](doc/memes.md)