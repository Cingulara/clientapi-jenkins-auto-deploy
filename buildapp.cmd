set CGO_ENABLED=0
set GOOS=linux
c:\\go\\bin\\go get -u github.com/go-sql-driver/mysql
c:\\go\\bin\\go get -u github.com/gorilla/mux
c:\\go\\bin\\go build -a -installsuffix cgo -o clientapipipeline .