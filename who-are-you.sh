curl -s https://oujda.01-edu.org/assets/superhero/all.json | jq '.[] | select(.id == 70) | .name'
