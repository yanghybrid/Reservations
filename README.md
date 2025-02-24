#Setup Go Enviornment

brew install go
go version
echo 'export PATH="/usr/local/go/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc

go version

go get github.com/gin-gonic/gin
go get github.com/gin-contrib/cors

#Test Restaurant Reservations

go run main.go

npm start

http://10.0.0.139:3001

#Trouble Shooting
lsof -i :8080
http://localhost:8080/reservations
curl -X POST http://localhost:8080/reservations -H "Content-Type: application/json" -d '{"name": "John Doe", "email": "john@example.com", "guests": 2}'
curl -X GET http://localhost:8080/reservations
