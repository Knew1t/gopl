#!/bin/bash

# Проверка, выполняется ли скрипт в репозитории Git
if ! git rev-parse --is-inside-work-tree &>/dev/null; then
  echo "Ошибка: Скрипт нужно запускать внутри Git-репозитория."
  exit 1
fi

# Поиск всех файлов с расширением .mod и выполнение команды git rm --cached
find . -type f -name "*.mod" | while read -r file; do
  echo "Удаление из индекса: $file"
  git rm --cached "$file"
done

echo "Удаление завершено. Не забудьте закоммитить изменения."
