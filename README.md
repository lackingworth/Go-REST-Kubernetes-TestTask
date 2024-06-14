# Go-REST-K8S-Test
___
#### Русский:
___
## Установка
* У вас должен быть установлен [Go v1.21.5](https://go.dev/doc/install) (или выше)
* У вас должен быть установлена актуальная версия [Docker](https://www.docker.com/) и активирован Kubernetes

## Запуск с использованием Docker

* Клонируйте данный репозиторий в выбранную вами локацию
* *(Необязательно)* Измените предоставленные ```Dockerfile``` и ```.yml``` файлы для кастомизации сетевых подключений
* Откройте консоль разработчика / Bash / PowerShell
* Зайдите в корневую директорию склонированного репозитория (там, где находится файл ```docker-compose.yml```)
  с помощью ```cd "путь/к/папке"```
* Используйте команду ```docker-compose up -d``` для формирования и запуска контейнеров

> [!NOTE]
> Сервис можно использовать уже на данном этапе, адрес доступа API - ```localhost:3000/```

## Деплой приложения через Kubernetes
* После контейнирезации приложения в Docker необходимо выложить образец сервиса на Docker hub. Используйте следующие команды:
 ```
 docker login

 docker push ваш-логин-докер/go-rest-app
 ```
* Примените деплоймент манифест Kubernetes командой ```kubectl apply -f deployment.yml```
* Примените сервис манифест Kubernetes командой ```kubectl apply -f service.yml```
* Найдите IP-адрес балансировщика загрузки командой ```kubectl get services``` - он должен быть в столбце ```EXTERNAL-IP``` напротив названия сервиса ```go-rest-service```

> [!NOTE]
> После деплоймента, адрес доступа API - ```EXTERNAL-IP:80/```

## Использование (Запросы)

> [!NOTE]  
> Имеются 4 API эндпоинта:
> 
> ```*DOMAIN*/api/create_book``` - POST запрос, добавляет книгу в базу данных. Тело JSON:
> ```
> {
>   "author":"имя автора",
>   "title":"название книги",
>   "publisher":"название издателя"
> }
> ```
> ```*DOMAIN*/api/delete_book/:id``` - DELETE запрос, удаляет книгу по уникальному идентификатору (ID) из базы данных. ID начинается с 1 по возрастанию с единичным шагом
> 
> ```*DOMAIN*/api/get_book/:id``` - GET запрос, возвращает книгу по уникальному идентификатору (ID) из базы данных. ID начинается с 1 по возрастанию с единичным шагом
> 
> ```*DOMAIN*/api/books``` - GET запрос, возвращает все книги из базы данных

>[!NOTE]
> Протестировано с помощью Postman

___
#### English:
___

## Installing
* You must have [Go v1.21.5](https://go.dev/doc/install) (or higher) installed on your system
* You must have up-to-date version of [Docker](https://www.docker.com/) installed on your system with Kubernetes enabled

## Starting containers / executing program

* Clone this repository to the location of your choosing
* *(Optional)* Change provided ```Dockerfile``` and ```.yml``` files for networking customization
* Open Bash / PowerShell terminal
* Navigate to the root directory of the cloned repository (where ```docker-compose.yml``` is located)
  using ```cd "path/to/directory"```
* Run ```docker-compose up -d``` command for building and starting up containers

> [!NOTE]
> Service can be accessed already, API address - ```localhost:3000/```

## Kubernetes deployment
* After Docker containerization you need to push api image to Docker hub. Run following commands:
 ```
 docker login

 docker push your-docker-login/go-rest-app
 ```
* Apply Kubernetes deployment manifest by running ```kubectl apply -f deployment.yml```
* Apply Kubernetes service manifest by running ```kubectl apply -f service.yml```
* Find load balancer IP by running ```kubectl get services``` - it must be in ```EXTERNAL-IP``` column across ```go-rest-service``` service name

> [!NOTE]
> After deployment API address - ```EXTERNAL-IP:80/```

## Usage (Requests)

> [!NOTE]  
> There are 4 API endpoints:
>
> ```*DOMAIN*/api/create_book``` - POST request, adds a book to the db. JSON  body:
> ```
> {
>   "author":"author name",
>   "title":"book title",
>   "publisher":"publisher name"
> }
> ```
> ```*DOMAIN*/api/delete_book/:id``` - DELETE request, deletes the book by ID from the database. ID begins with 1 in ascending order with the step of 1
> 
> ```*DOMAIN*/api/get_book/:id``` - GET request, returns the book by ID from the database. ID begins with 1 in ascending order with the step of 1
> 
> ```*DOMAIN*/api/books``` - GET request, returns all books from database

>[!NOTE]
> Tested w/ Postman

## Version History

* v.0.0.1:

    * Initial Release
