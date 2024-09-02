DATA=$(curl -s https://oujda.01-edu.org/assets/superhero/all.json)


echo $DATA | jq '.[] | select(.id == 170) | .name, .powerstats.power, .appearance.gender' | tr -d "\""
