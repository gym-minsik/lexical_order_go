# Lexical Order

A string generator designed to simplify real-time editing of ordered sequences. It facilitates efficient reordering, sorting, and interleaving of items by generating lexicographically ordered keys.

It is intended to pair with the [**lexicographical_order**](https://pub.dev/packages/lexicographical_order) package on pub.dev. If you’re developing with Flutter and Go, you can easily implement real-time transaction reordering using these companion libraries.

## Features
### Between

Generates a lexicographically ordered key between two given strings (prev and next). The resulting key is ideal for maintaining sort order in real-time scenarios.
```go
var service := NewAlphabetOrderService()
mid := service.Between("A", "Z") // mid == "N", lexically between "A" and "Z"
```
### Generate

Generates multiple lexicographical keys at once, typically used to assign initial sort keys for a collection.
```go
var service := NewAlphabetOrderService()
keys := service.Generate(1000) // generates 1000 keys
```
**Common use-cases include:**
1.	Inserting into an empty collection (when Between cannot be used):
    ```go
    func AddTodo(command CreateTodo, todos []Todo, repo TodoRepository) error {
        var orderKey string
        if len(todos) == 0 {
            orderKeys := service.Generate(1)
            orderKey = orderKeys[0]
        } else {
            orderKey = service.Between(todos[len(todos)-1].OrderKey, "")
        }

        todo, err := repo.Create(command, orderKey)
        if err != nil {
            return err
        }

        todos = append(todos, todo)
        return nil
    }
    ```go  

2.	Migrating an existing ordering system to Lexical Order:
    ```go
    func MigrateToLexicalOrderSystem(table Table) error {
        itemCount := table.Count()
        orderKeys := service.Generate(itemCount)
        // Migration logic omitted for brevity
        return nil
    }
    ```