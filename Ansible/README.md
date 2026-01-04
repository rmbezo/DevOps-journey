# Ansible

- for ansible practice i use VM's with debian

```
    ansible-playbook -i hosts playbook.yml -K
```

## что делает наш ansible?

- В файле playbook.yml описаны действия ansible на машине

должно быть что то вроде: 

![example](./images/image.png)



## fastcheck server alivness

```
    ansible virtuals -i hosts -m ping
```

![example](./images/image2.png)