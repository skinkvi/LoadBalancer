Привет это мой просто балансировщик нагрузок на GO

Я тут решил использовать алгоритм Least Connections потому что из всех считаю его более эффективным. К примеру тот же Round Robin не смотря на свою простоту не может предоставить балансировку. Да раньше он был очень хорош, но сейчас цикличная балансировка не может называться хорошей.
Как минимум нужно учитывать нагрузку сервером и их состояние, слышал было такое что даже когда сервер выключен или он не  исправен Round Robin все равно будет отправлять ему запросы, что является серьезной ошибкой. 
Балансирощик получился неплохим, да скорее всего он не подойдет на 10+ серверов но для просто проекта где не нужен мощный и забирающий произовидельность балансировщик, подойдет мой простой легкий легко внендряемый. 
Надеюсь тут будет что то полезное, только добра!