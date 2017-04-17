# Trivial smart contract that will run in IBM Hyperledger Fabric V0.6 
A trivial smart contract using HLF v 0.6. 

## Register user request
User id and secret would be available in the membersrvc.yaml file.
membersrvc.yaml should be copied using the following command 

```
docker cp <container id>:/opt/gopath/src/github.com/hyperledger/fabric/membersrvc/membersrvc.yaml .

```
Container id could be found using the command

```

docker ps

```
Register using the following json request

```
{
 "enrollId": "<<user id for your chain code peer >>",
 "enrollSecret": "<<User secret>>"
}

```
## Deployment request 
```
{
 "jsonrpc": "2.0",
 "method": "deploy",
 "params": {
   "type": 1,
   "chaincodeID": {
     "path": "https://github.com/suddutt1/trivialsmartcontractv6"
   },
   "ctorMsg": {
	 "function": "init",
     "args": ["10"]
   },
   "secureContext": "<<user id for your chain code peer >>" 
},
 "id": 1
}


```
Above invocation will retun a chain code in the format 

```
{
	"jsonrpc": "2.0",
	"result": {
		"status": "OK",
		"message": "<<Chain code hash would be present here>>"
	},
	"id": 1
}
```

## Invoke example
With the following invoke , we will send 100 to store against sudip. As per the business logic, it will store (100-10)= 90

```

{
 "jsonrpc": "2.0", 
 "method": "invoke",
 "params": {
   "type": 1,
   "chaincodeID":{ "name":"<<Your chain code hash>>"
    },
   "ctorMsg": {
       
           "function": "deposite",
            "args": [
                "SUDIP",
                "100"
            ]
   },

   "secureContext": "<<user id for your chain code peer >>"
 },
 "id": 1
}


```
## Query example
To query what is stored against SUDIP

```
{
  "jsonrpc": "2.0",
  "method": "query",
  "params": {
    "type": 1,
    "chaincodeID":{
 "name": "<<Your code code hash>>"
     },
    "ctorMsg": {
        
     "function": "read",
      "args": [
                "SUDIP"
            ]
   },

    "secureContext": "<<user id for your chain code peer >>"
  },
  "id": 1
}


```
Above query should result the following output

```

{"jsonrpc":"2.0","result":{"status":"OK","message":"990"},"id":1}

```


