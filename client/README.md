# Form3 Take Home Exercise

Engineers at Form3 build highly available distributed systems in a microservices environment. Our take home test is designed to evaluate real world activities that are involved with this role. We recognise that this may not be as mentally challenging and may take longer to implement than some algorithmic tests that are often seen in interview exercises. Our approach however helps ensure that you will be working with a team of engineers with the necessary practical skills for the role (as well as a diverse range of technical wizardry). 

## How to create a Services
To create a new Account Services is requere to import the client
```
	f3 "github.com/ferchuhill/form3-client/client"
```
and then call the create method, this will create a client with the default configuration
```
	acService := f3.NewAccountService(nil)
```


## How to use the service
first is requiere to import the module (for the create)
```
  model "github.com/ferchuhill/form3-client/client/models"
```
To use the serives:
```
  // Create
  acData, err := acService.Create(accountData)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Printf("Create : %s - %d\n", acData.ID, acData.Version)

  //Fetch by Id
  acDataFectch, err := acService.Fetch(acData.ID)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(acDataFectch.ID)

  //Delete
  del, errDel := acService.Delete(acDataFectch.ID, int64(acDataFectch.Version))
  if errDel != nil {
    fmt.Println(errDel)
  } else {
    fmt.Println(del)
  }
```

## How to run the test

There are two way to run test, one with `docker-compose up` or running `make testDocker`  
