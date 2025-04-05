# Zadanie 4 | Go

Należy stworzyć projekt w echo w Go. Należy wykorzystać gorm do stworzenia 5 modeli, gdzie pomiędzy dwoma musi być relacja. Należy zaimplementować proste endpointy do dodawania oraz wyświetlania danych za pomocą modeli. Jako bazę danych można wybrać dowolną, sugerowałbym jednak pozostać przy sqlite.

- [X] 3.0 - należy stworzyć aplikację we frameworki echo w j. Go, która będzie miała kontroler Produktów zgodny z CRUD
- [X] 3.5 - należy stworzyć model Produktów wykorzystując gorm oraz wykorzystać model do obsługi produktów (CRUD) w kontrolerze (zamiast listy)
- [X] 4.0 - należy dodać model Koszyka oraz dodać odpowiedni endpoint
- [X] 4.5 - należy stworzyć model kategorii i dodać relację między kategorią, a produktem
- [X] 5.0 - należy pogrupować zapytania w gorm’owe scope'y

Przykładowe dane są automatycznie wypełniane dzięki database/seeds.go.

