# New package
```bash
go get github.com/alexflint/go-arg
go mod tidy
go build
go mod vendor
```
# Git update from mainstream
1. Rebase
```bash
git fetch # скачиваем изменения с сервера
# надо находиться в ветке, которую вы хотите обновить
# вместо origin/master нужно вписать вашу mainstream-ветку
git rebase origin/master
# далее, если возник конфликт, исправляем его и делаем 
git add file-with-conflict.go
git rebase --continue
# повторяем до тех пор, пока не появится сообщение
# Successfully rebased and updated ...
```
2. Merge
```bash
git fetch # скачиваем изменения с сервера
# надо находиться в ветке, которую вы хотите обновить
# вместо origin/master нужно вписать вашу mainstream ветку
git merge origin/master
# далее, если возник конфликт
git add file-with-conflict.go
git merge --continue
```
