# Apache simple dockerfile example

- To start with basic html:

```
    docker build -t name:tag .
    docker run -d --rm --name name -p 80:80 test
```

- With your own (./files/index.html)

```
    docker run -d --rm --name apache -p 80:80 -v your_path/Dockerfile_practice/files:/var/www/html image:name
```


- Then open `localhost`

