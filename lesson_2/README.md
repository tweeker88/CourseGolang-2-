# API

## Маршрутизатор и исполнители
***Маршрутизатор (router)*** - это экземпляр, который имеет внутренний функционал , заключающийся в следующем:
* принимает на вход адрес запроса (по сути это строка ```http://localhost:8080/resource```) и вызывает исполнителя, который будет ассоциирован с этим запросом.  

***Исполнитель (handler)*** - это функция/метод, котоырй вызывает маршрутизатором.