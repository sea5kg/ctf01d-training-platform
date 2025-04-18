# Contributing to ctf01d-training-platform

Спасибо за ваш интерес к развитию проекта!

---

## Общие рекомендации

- Перед созданием PR ознакомьтесь с [README.md](./README.md) и документацией в папке `docs/`.
- Следуйте принятому стилю кода и структуре проекта.
- Для багов и предложений используйте Issues.


---

## Как внести вклад

1. Форкните репозиторий и создайте новую ветку для вашей задачи.
2. Оформляйте коммиты понятно и лаконично.
3. Перед PR убедитесь, что проект собирается и проходят тесты.
4. Оформите Pull Request с описанием изменений.

---

## Контакты

Если возникли вопросы — пишите в Issues.

---

## SSL/TLS сертификаты (Let's Encrypt)

Для генерации или обновления сертификатов используйте:

```sh
cd build
docker run --rm --name temp_certbot \
    -v ./nginx/certbot/conf:/etc/letsencrypt \
    -v ./nginx/certbot/www:/tmp/letsencrypt \
    -v ./nginx/certbot/log:/var/log \
    certbot/certbot:latest \
    certonly --webroot --agree-tos --renew-by-default \
    --preferred-challenges http-01 --server https://acme-v02.api.letsencrypt.org/directory \
    --text --email hotorcelexo@gmail.com \
    -w /tmp/letsencrypt -d ctf01d.ru
```
