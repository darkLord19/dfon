# dfon
This program checks for auto incrementing fileds of tables of database and notifies on given emails if any of the fields is about to run out of its datatypes range.

# Config file format
```
{
    "databases": [
        {
            "host": "localhost",
            "port: 1234,
            "user": "test",
            "password": "passwd",
            "notify_list": ["abc@xyz.com", "test@test.com"],
            "threshold": 92.1,
            "type": "postgres"
        },
        {
            "host": "localhost",
            "port: 1234,
            "user": "test",
            "password": "passwd",
            "notify_list": ["abc@xyz.com", "test@test.com"],
            "threshold": 92.1,
            "type": "mysql"
        }
    ]
}
```