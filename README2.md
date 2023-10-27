# Go REST Example 😛

Now that we aren't on the frontpage, I can finally be myself and use dumb emojis everywhere!!!

## So, understanding the Requests Lifecycle ↩️

**This is a big project**. If you dive in head-first without much knowledge you'll probably find your way around, but keep on reading just a little bit and you'll be able to really own this.


### First, a quick overview ⚡

`Server starts`. Router starts receiving HTTP Requests, mapping them to the corresponding Endpoint.

`HTTP Requests` are transformed into our custom Requests, for example `CreateUserRequest` (with fields like Username or Password). Those requests are validated.

Then our Request is transformed into a `Model`, the key structures of our program. They're our DB Schema and also where we perform operations. A `UserModel`, for example, has methods like .Create or .Update, taking the `Repository` as parameter to communicate with the DB.

A `CreateUserResponse` is generated with the operations' results, and finally transformed to an `HTTP Response`.

**`HTTP Request ➡️ CreateUserRequest ➡️ UserModel ➡️ Repository ➡️ Database ➡️ CreateUserResponse ➡️ HTTP Response`**

---

### Then, a not so quick one 🐌

`You start the server`. All of its modules are also initialized and if there are no errors, the server will run. 

A `Router` will be in place listening for any wandering HTTP Request that happens to fall beneath its path. Oh wow! A new `HTTP Request` came in, and it's pointed towards the Create User Endpoint. I'll run the `Middleware` and send it on its way.

The HTTP Request's data is used to create our own custom `CreateUserRequest` (with fields like Username or Password). Validations are executed here.

If they pass, then that CreateUserRequest will be passed along the way and eventually converted into a `Model`, in this case a `UserModel`.

Models are the key players of our app, they define the DB Schema and also serve as Business Objects, which means that much of the logic of the program will be inside of the methods of this UserModel. 

Methods like `.Create` or `.Update` which take the `Repository` as parameter and use it to talk to the `Database`, or also methods like HashPassword that just modify the fields inside of it.

After communicating with the Database the UserModel is filled with the resulting data, and then transformed into a `CreateUserResponse`. Middleware executes again, and finally our custom Response is transformed into our beautiful `HTTP Response`.

**`HTTP Request ➡️ Middleware ➡️ CreateUserRequest ➡️ Validations ➡️ UserModel ➡️ Repository ➡️ Database ➡️ CreateUserResponse ➡️ Middleware HTTP ➡️ Response`**

Easy. As. Fvck.

---

## What about the Architecture?

Transport & Service & Repository. TODO.