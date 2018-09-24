# docker-volume
reader_1  | 2018/09/24 13:55:52 got content hello world
```
reader_1  | 2018/09/24 13:52:25 listening to port *:8080. press ctrl + c to cancel.
reader_1  | 2018/09/24 13:52:26 got content open ./data/test.txt: no such file or directory
reader_1  | 2018/09/24 13:52:42 got content open ./data/test.txt: no such file or directory
reader_1  | 2018/09/24 13:53:13 got content hello world
```

## Test fail 
```
curl http://localhost:8080?dir=data/test.txt
curl http://localhost:8080?dir=./data/test.txt
```
Output:
```
reader_1  | 2018/09/24 13:56:01 got content open ./data/test.txt: no such file or directory
```

## Test success
```
curl http://localhost:8080?dir=/data/test.txt
```

Output:
```
reader_1  | 2018/09/24 13:56:54 got content hello world
```
