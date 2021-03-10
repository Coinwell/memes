```
docker build -t n2n2/n2n2-meme .
docker run n2n2/n2n2-meme

docker run -p 5000:5000 --env-file .env n2n2/n2n2-meme
```
