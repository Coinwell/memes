# n2n2-memes

[How To Deploy](ops/ansible/README.md)

## Bash
```
ssh -i ~/.ssh/n2n2 ubuntu@memes-staging.n2n2.chat
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