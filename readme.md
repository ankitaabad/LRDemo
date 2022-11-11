# LR API Demo

There are 3 APIs
1. Create Account: Create a new Account
    - Post http://127.0.0.1:3000/account
    - Query Params: apikey, apisecret
    - Body : ``` {
    "Email": [{
        "Type": "Primary",
        "Value": "hamaffeicruddou-8816@yopmail.com"
    }],
    "Password": "L@1234"
}```

2. Login: Login to your account, You will get accesstoken that can be used to update the profile
    - Post http://127.0.0.1:3000/login
    - Query Params: apikey
    - Body: ```{
     "email":"hamaffeicruddou-8816@yopmail.com",
    "password":"L@1234"
}```

3. Update Profile: update the profile
    - Post http://127.0.0.1:3000/account
    - Header : Authorization : Bearer accesstoken
    - Query Params: apikey
    - Body: fields you want to update (FirstName, LastName, Gender) ```{
    "FirstName":"Rohit"
}```

