# Часть сервиса аутентификации

## Технологии
- **Go**
- **PostgreSQL**
- **Go Fiber** 
- **GORM** 
- **JWT** 

## Запуск
1. Клонируйте репозиторий:
    ```bash
    git clone https://github.com/usernameisavailablee/TestTaskBackDev
    ```

2. Убедитесь, что Docker Compose установлен:  
   [Инструкция по установке Docker Compose](https://docs.docker.com/compose/install/)

3. Перейдите в директорию проекта:
    ```bash
    cd TestTaskBackDev
    ```

4. Создайте файл `.env`, скопировав пример `.env-example`:
    ```bash
    cp .env-example .env
    ```

5. Соберите и запустите приложение при помощи:
    ```bash
    docker compose up --build
    ```

## Примеры использования

### 1. Добавление пользователя
**Endpoint:POST http://127.0.0.1:3000/user**
**Тело запроса:**
```json
{
    "email": "e321x123111@gmail.com"
}
```
Пример запроса через curl:

```bash
curl -X POST http://127.0.0.1:3000/user \
-H "Content-Type: application/json" \
-d '{
    "email": "e321x123111@gmail.com"
}'
```
### 2. Генерация пары токенов
**Endpoint:POST http://127.0.0.1:3000/auth/generate-pair**
```json
{
    "user_id": "796e357f-b7c8-4768-b9f6-bf32fe4202ce"
}
```
Пример запроса через curl:
```bash
curl -X POST http://127.0.0.1:3000/auth/generate-pair \
-H "Content-Type: application/json" \
-d '{
    "user_id": "796e357f-b7c8-4768-b9f6-bf32fe4202ce"
}'
```

### 3. Обновление токена
**Endpoint:POST http://127.0.0.1:3000/auth/refresh**
```json
{
    "user_id": "796e357f-b7c8-4768-b9f6-bf32fe4202ce",
    "refresh_token": "REzs3+xIAawEz723Dwtxss90Yi/4X6OgJ55FmPZH98o="
}
```
Пример запроса через curl:
```bash
curl -X POST http://127.0.0.1:3000/auth/refresh \
-H "Content-Type: application/json" \
-H "X-Forwarded-For: 192.168.1.124" \
-d '{
    "user_id": "796e357f-b7c8-4768-b9f6-bf32fe4202ce",
    "refresh_token": "REzs3+xIAawEz723Dwtxss90Yi/4X6OgJ55FmPZH98o="
}'

```



