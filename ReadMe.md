<h1>Создание простого сервиса аутентификации</h1>

<h2>Функционал сервиса</h2>
<ul>
    <li>
    ✅ Регистрация пользователя (email + пароль) </li>
    <li> 
    ✅ Вход (генерация JWT-токена)
    </li>
    <li>
    ✅ Проверка токена (middleware)
    </li>
    <li>
    ✅ Обновление токена
    </li>
    <li>
    ✅ Выход (удаление refresh-токена)
    </li>
    <li>
    ✅ Подключение к PostgreSQL для хранения пользователей
    </li>
</ul>

<h2>Стек технологий</h2>
<ul>
    <li>
    🛠 Golang — основной язык
    <li> 
    🛠 PostgreSQL — база данных
    </li>
    <li>
    🛠 GORM — ORM для работы с БД
    </li>
    <li>
    🛠 JWT (github.com/golang-jwt/jwt/v5) — генерация токенов
    </li>
    <li>
    🛠 Docker — контейнеризация
    </li>
</ul>

<h3>Скопировать данный репозиторий</h3>

<h3>Установка зависимостей</h3>
```
go mod init auth-service
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go get -u github.com/golang-jwt/jwt/v5
go get -u github.com/spf13/viper
```