## Dockerfile practice

- simple practice self writing dockerfile images. 

- Without adding your site:
```
    docker build -t name:tag .
    docker run --rm -d -P --name name image:tag
```

- Adding your own site

```
    docker run -d --rm --name nginx -p 80:80 -v your_path/Dockerfile_practice/files:/var/www/html image:name
```

- To see results (curl -I http://localhost) or open localhost in browser