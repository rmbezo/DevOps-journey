# Same dockerfile with copy

- to run

```
    docker build -t name:tag .
    docker run -d --rm --name name -p 80:80 image
```

- If you want to put your site change /files/index.html ////
or when start add -e OWNER=yourname -e TYPE=yourtype

- Also if you want to use nginx instead apache change RUN line and CMD line(and volumes instead of /var/www/html/ use /usr/share/nginx/html/)