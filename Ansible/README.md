# Ansible

- for ansible practice i use VM's with debian

```
    ansible-playbook -i hosts playbook.yml -K
```

## что делает наш ansible?

- В файле playbook.yml описаны действия ansible на машине

должно быть что то вроде: 


![example](./images/image3.png)

## Commands:

- fastcheck servers alivness:

```
    ansible virtuals -i hosts -m ping
```

![example](./images/image2.png)


- copy local file to servers:

```
ansible all -i hosts -m copy -a "src=file dest=/home mode=777" -b
```

- Start playbook:

```
ansible virtuals -i hosts playbook.yml -K
```