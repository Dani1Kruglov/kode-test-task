<h2>Project Description</h2>
<p>This project is a REST API service implemented in the Go programming language. The service provides the following functions:</p>
<ul>
   <li>Adding notes</li>
   <li>Listing a user's notes</li>
    <li>Spelling validation when saving notes using the Yandex.Speller service</li>
    <li>Authentication (login)</li>
    <li>Registration (register)</li>
    <li>Authorization using middleware</li>
    <li>Users have access only to their own notes</li>
    <li>Data storage in PostgreSQL</li>
</ul>
<p>Before using it, change the example-config.hcl to config.hcl, fill it with your data, and docker-compose.dev.yml with your information.</p>
<h2>Requirements</h2>
<ul>
    <li>Programming Language: Go</li>
    <li>Data Transmission Format: JSON</li>
    <li>Logging of events in a unified format</li>
    <li>It is recommended to use the standard Go library, as well as libraries from golang.org. The choice of a logging library is at the discretion of the developer, and for the web server, you can use chi, gorilla, or the standard library.</li>
    <li>The service and its infrastructure should be run in Docker containers.</li>
</ul>
<h2>Testing API functionality is done in Postman (below are the necessary data for each method, the method's link, and the request method)</h2>
<ul>
    <li>Get all user notes (GET method): http://localhost:8080/user/notes</li>
    <li>Create a user note (POST method): http://localhost:8080/user/note/store
    <code>
{
    "note_title": "some note title",
    "note_content": "some note text"
}
</code>
    </li>
    <li>Delete a user note (DELETE method): http://localhost:8080/user/note/delete
<code>
{
    "note_id": 10 //user note ID
}
</code>
    </li>
    <li>User registration (POST method): http://localhost:8080/user/register
<code>
{
    "user_name": "user name",
    "user_email": "email",
    "user_password": "password"
}
</code>
    </li>
    <li>User login (POST method): http://localhost:8080/user/login
<code>
{
    "user_email": "email",
    "user_password": "password"
}
</code>
    </li>
    <li>User logout (POST method): http://localhost:8080/user/logout</li>
</ul>
<h2>Results</h2>
<p>The solution to the task is provided in the form of a link to a GitHub repository. The project implements all the requirements, including note addition and listing, spelling validation, authentication and authorization, registration, and data storage in PostgreSQL. The project enables interaction through a JSON-based REST API.</p>
<p>This project is developed in accordance with the requirements and recommendations specified in the task and is ready for review.</p>



<h2>Описание проекта</h2>
<p>Этот проект представляет собой REST API сервис, реализованный на языке программирования Go. Сервис предоставляет следующие функции:</p>
<ul>
   <li>Добавление заметок</li>
   <li>Вывод списка заметок пользователя</li>
    <li>Валидация орфографии при сохранении заметок с использованием сервиса Яндекс.Спеллер</li>
    <li>Аутентификация (login)</li>
    <li>Регистрация (register)</li>
    <li>Авторизация с использованием middleware</li>
    <li>Пользователи имеют доступ только к своим заметкам</li>
    <li>Хранение данных в PostgreSQL</li>
</ul>
<p>Перед использованием измените example-config.hcl на config.hcl, заполните его и docker-compose.dev.yml своими данными.</p>
<h2>Условия</h2>
<ul>
    <li>Язык программирования: Go</li>
    <li>Формат передачи данных: JSON</li>
    <li>Логирование событий в едином формате</li>
    <li>Рекомендуется использовать стандартную библиотеку Go, а также библиотеки из golang.org. Выбор библиотеки для логгирования остается на усмотрение разработчика, а для веб-сервера можно использовать chi, gorilla или стандартную библиотеку.</li>
    <li>Запуск сервиса и его инфраструктуры должен осуществляться в Docker-контейнерах.</li>
</ul>
<h2>проверка работоспособности API рпоизводится в Postman (ниже приведены необходимые данные для каждого метода, ссылка метода и метод запроса)</h2>
<ul>
    <li>Получение всех заметок пользователя (метод GET): http://localhost:8080/user/notes</li>
    <li>Создание заметки пользователя (метод POST): http://localhost:8080/user/note/store
    <code>
{
    "note_title": "какое-то название заметки",
    "note_content": "какой-то текст заметки"
}
</code>
    </li>
    <li>Удаление заметки пользователя (метод DELETE): http://localhost:8080/user/note/delete
<code>
{
    "note_id": 10 //id заметки пользователя
}
</code>
    </li>
    <li>Регистрация пользователя (метод POST): http://localhost:8080/user/register
<code>
{
    "user_name": "имя пользователя",
    "user_email": "email",
    "user_password": "password"
}
</code>
    </li>
    <li>Вход пользователя (метод POST): http://localhost:8080/user/login
<code>
{
    "user_email": "email",
    "user_password": "password"
}
</code>
    </li>
    <li>Выход пользователя (метод POST): http://localhost:8080/user/logout</li>
       
</ul>

<h2>Результаты</h2>
<p>Решение задачи предоставлено в виде ссылки на репозиторий GitHub. В проекте реализованы все требования, включая функции добавления и вывода заметок, валидацию орфографии, аутентификацию и авторизацию,регистрацию, а также хранение данных в PostgreSQL. Проект обеспечивает работу через REST API с форматом JSON.</p>
<p>Этот проект разработан в соответствии с требованиями и рекомендациями, указанными в задаче, и готов к ревью.</p>
