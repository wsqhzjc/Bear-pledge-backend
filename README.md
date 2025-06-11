# pledge-backend

The project is divided into two parts, one is API and the other is scheduled task

API

    cd api
    export plgr_admin_private_key=ab6110cb1aca11bc7460843326af7d620138e69a9c5117700125e68bc8911bd5
    go run pledge_api.go

pool task

    cd schedule
    export plgr_admin_private_key=ab6110cb1aca11bc7460843326af7d620138e69a9c5117700125e68bc8911bd5
    go run pledge_task.go


```
abigen --abi=abi/PledgePool.abi.json --pkg=PledgePoolToken --out=bindings/pledgePoolToken.go

```

singer1  db612dee6e7108560a08bee6686491cb354d04920ec4a157b104f16d83a2aab3