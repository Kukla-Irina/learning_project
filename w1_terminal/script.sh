#!/bin/bash
echo "Введите ваше имя:"
read name
mkdir "$name"
echo "Привет, $name! Это твоя первая папка." > "$name/welcome.txt"
