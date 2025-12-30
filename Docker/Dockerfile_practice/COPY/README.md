## Same dockerfile with copy

- to run

```
    docker build -t name:tag .
    docker run -d --rm --name name -p 80:80 image
```

- If you want to put your site change /files/index.html

- Also if you want to use nginx instead apache change RUN line and CMD line 