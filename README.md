# proxy

## Примеры использования
curl --location 'http://localhost:80/api/v1/proxy' \
--header 'Content-Type: application/json' \
--data '{
    "method": "GET",
    "url": "http://google.com",
    "headers": {"Authentication": "Basic bG9naW46cGFzc3dvcmQ="}
}'

response
{
    "success": true,
    "data": {
        "data": {
            "id": "6adca40a-1b25-46ff-ba24-a423ff0d4616",
            "status": "200 OK",
            "headers": {
                "Cache-Control": "private, max-age=0",
                ...
            },
            "length": 19082
        }
    }
}


# Запустить приложение
go run cmd/main.go
