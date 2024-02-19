# Kro-backend

Install golang with this link 
https://go.dev/doc/install

Download NoSQL Workbench with this link 
https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/workbench.settingup.html

## Next step import file Kro-backend.json To NoSQL Workbench
- When you download NoSQL Workbench successfully open the app.
  - Select **Import Data Model**.
  - Select file *Kro-backend.json* in project.
  - Next step click on ddb local.
  > ![](https://cdn.discordapp.com/attachments/880831085431390301/1208857410807275581/2567-02-19_02.27.36.png?ex=65e4cf32&is=65d25a32&hm=0e778af7e8022c526cac623f37fca4c9c36743dc29ab2c2cd9c19b68fea9bbfc&)
  - Go to Operation builder and create Dynamodb local.
  - View the credentials Keys when you create connection success.
  > ![](https://cdn.discordapp.com/attachments/880831085431390301/1208860132939669624/2567-02-19_02.38.28.png?ex=65e4d1bb&is=65d25cbb&hm=4fcc48235ee18cf11a86473cc36a77b6cd6684eff902c4bdd8a15298ab4c30a1&)
  - Go to file db.go and edit the AccessKeyId and SecretAccessKey.

## you can run golang with cmd
```
go run cmd/myapp/main.go
```
### http://localhost:8080/kro-games