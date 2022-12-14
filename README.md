[![ci](https://github.com/TOIIIA86/image-previewer/actions/workflows/tests.yml/badge.svg)](https://github.com/TOIIIA86/image-previewer/actions/workflows/tests.yml)
# Image Previewer

## Описание сервиса
Сервис предназначен для изготовления preview (создания из

## Команды для работы с проектом
1. make build-server - билдим бинарник приложения
2. make test-single - запуск одного прохода тестов с таймаутом 1м
3. make test-race - запускает 100 проходов теста с таймаутом 7м
4. make test-coverage - посмотреть % покрытия тестами проекта 
5. make test-integration - запуск интеграционных тестов
6. make lint - запуск линтера в проекте

### Перед стартом докера:
1. Скопировать `.env.dist` в `.env`
2. Проставить желаемый параметр для кеша в файле `.env`, по дефолту 100.

### Запуск в докере

Докер запускается с флагом `--remove-orphans`, чтобы не захламлять старыми контейнерами

1. `make start` - запускает проект 
2. `make stop` - остановка контейнеров

### Запуск

1. `make build-server` - генерим бинарный файл
2. `make run-server` - стартуем сервер из консоли

## Описание работы сервиса
Отправляем на url resize параметры для изменения размера изображения и ссылку для получения исходного изображения которое хотим изменить

## Параметры url
`host/resize/{width}/{height}/{imageUrl}`

1. width - ширина картинки, которую хотим получить
2. height - высота картинки, которую хотим получить
3. imageUrl - ссылка на исходное изображение для уменьшения по заданным параметрам

## Ответы сервиса

1. 200 - если все ок, то получаем нашу уменьшенную картинку
2. 400 - не верные параметры width/height/imageUrl
3. 500 - не удалось отдать картинку
4. 502 - в случае не доступности или ошибки удаленного сервиса

## Пример для получения картинки после запуска докера
**Пример для Докера:** 
`http://127.0.0.1/resize/200/200/raw.githubusercontent.com/OtusGolang/final_project/master/examples/image-previewer/gopher_1024x252.jpg`
**Пример для запуска из консоли**
`http://localhost:8080/resize/300/300/raw.githubusercontent.com/OtusGolang/final_project/master/examples/image-previewer/_gopher_original_1024x504.jpg`