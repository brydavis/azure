# Azure

Experimental SQL Terminal using Go + Windows Azure.

Program requires a `config.json` in root directory:

```javascript
{
	"password" : "MyPassword!123",
	"port"     : 1433,
	"server"   : "myserver.database.windows.net",
	"user"     : "UserMe",
	"database" : "demoDB"
}

```

Build and run the program.

When you see terminal prompt `azure ~>`, proceed using SQL.

Special commands:

1. `azure ~> run query.sql` executes code in `./sql/query.sql`
2. `azure ~> exit` or `azure ~> quit` exits the program and returns to the command line

<br>
<br>

<hr>
<small>
<strong>&copy; 2015 MIT License</strong>
</small>