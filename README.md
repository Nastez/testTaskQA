# testTaskQA


Инструкция по запуску программы для отправки сообщений в Slack

Предусловия:

1. Зарегистрирован аккаунт в Slack
2. Создано workspace
3. Создано три открытых канала "test1", "test2", "test3"
4. Создано приложение (бот) для отправки сообщений из json-файла в каналы (см https://api.slack.com/)
5. Получен и сохранен Bot User OAuth Token (бот-токен)
6. Загружен zip-архив с программой (https://github.com/Nastez/testTaskQA)
7. Из zip-архива "testTaskQA-master.zip" выгружена папка "testTaskQA-master"

Шаги для запуска программы на Windows OC:

1. Открыть папку "testTaskQA-master"
2. Открыть файл example.json в текстовом редакторе
3. Изменить строку "your_bot_token", вставив вместо your_bot_token скопированный Bot User OAuth Token (см предусловие 5, json-формат должен быть сохранен)
4. Сохранить изменения в файле example.json 
5. Запустить testTaskQA.exe
