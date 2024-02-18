# kro-backend

install golang with this link 
https://go.dev/doc/install

install NoSQL Workbench with this link 
https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/workbench.html


## setup dynamoDB local server with docker
```
cd deployment/docker && docker-compose up
```
### If you run success will show
![](https://cdn.discordapp.com/attachments/880831085431390301/1208794410247127070/2567-02-18_22.17.09.png?ex=65e49485&is=65d21f85&hm=43386978b9c04b2b12519b6dc81a709f3bfc124f6848885c2cb329d499cda7e0&)

## Next step import file Kro-backend.json To NoSQL Workbench
- In your NoSQL Workbench you select
  - **Import Data Model**
  - select file *Kro-backend.json* in project


## you can run golang with cmd
```
go run cmd/myapp/main.go
```
