# Сервис управления сотрудниками

## Запуск приложения

```shell
make run
````

## Точки доступа

### 1. Нанять сотрудника
**Точка доступа:** `http://localhost:8080/hire`  
**Метод:** `POST`  
**Пример тела запроса:**
```json
{
  "full_name": "Anton",
  "phone": "+79123456789",
  "gender": "man",
  "age": 30,
  "email": "ivanov@example.com",
  "address": "Moscow, Tverskaya St., b. 10"
}
```
### 2. Уволить сотрудника
**Точка доступа:** `http://localhost:8080/fire`  
**Метод:** `DELETE`  
**Параметры запроса:**  
`id`: ID сотрудника.  
Пример:  
`http://localhost:8080/fire?id=1`

### 3. Найти сотрудника по ФИО
**Точка доступа:** `http://localhost:8080/find`  
**Метод:** `GET`  
**Параметры запроса:**  
`name`: имя сотрудника  
Пример:  
`http://localhost:8080/find?name=Антон`

### 4. Получить количество дней отпуска
**Точка доступа:** `http://localhost:8080/vacation`  
**Метод:**  `GET`   
**Параметры запроса:**  
`id`: ID сотрудника   
Пример:  
`http://localhost:8080/vacation?id=1`


