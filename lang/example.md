## Practices
- Think in terms of an application. Not in terms of programming.
- Make the right thing the easy thing


Therefore we need to be able to define our model:

```ruby
model JobApplication {
	company_name String
	company_address String
	resume File
}
```

Let's set some ground rules:
Constants are PascalCased
Variables are snake_cased

We can see from the model definition above that there 
are a few built-in types. These include `String`, `Int`, `Bool`,
`File`, `Password`.

`String`, `Int`, and `Bool` are self explanatory. 

`File` is an uploaded file.
The presence of a file field anywhere in the application model 
implies that there is a hard dependency on a filestore of some kind
(S3, etc).

`Password` is stored in the database using a salted hash. The mechanism
is encoded into the language to discourage people from storing passwords
as plaintext. 

Writing the above model definition gives me:

- A database schema which can store it.
- Actions for CRUD on that model. 
- The ability to use data in that model in my views.

For example, for `JobApplication`, I would get the following:

A db table:
```SQL
CREATE TABLE JobApplication (
	id SERIAL INT PRIMARY KEY,
	company_name TEXT,
	company_address TEXT
);
```

These typical REST endpoints:

```
GET /job_applications
GET /job_application/:id
POST /job_applications
PUT /job_application/:id
```

Which brings me to:

## Actions

The POST / PUT endpoints above are available as actions in the views.
The GET endpoints are available as data sources in views.

Other actions are made available my modules. Modules are written in code.
They give you access to things like email, texting, etc.

