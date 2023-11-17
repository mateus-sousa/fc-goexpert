printf '// +build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools.go

go mod init github.com/mateus-sousa/fc-goexpert/11-graphql

go mod tidy

go run github.com/99designs/gqlgen init

go mod tidy

"modificar os schemas"

go run github.com/99designs/gqlgen generate


OBS: Entender uma forma melhor de fazer uma busca personalizada com relacionamentos no banco, pois da forma que foi ensinada no curso, gera um problema de N+1 no banco.