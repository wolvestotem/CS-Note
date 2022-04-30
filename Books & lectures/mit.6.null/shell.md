stream
redirect input and output of the files
one command
```shell
echo hello > hello.txt
cat hello.txt
cat < hello.txt
cat < hello.txt > hello2.txt
cat hello.txt >> hello2.txt // append
```

pipe
connect std output of the left to the std input of the right
several commands
```shell
ls -l | tail -n1
```

ls -l
权限：用户 组 else
