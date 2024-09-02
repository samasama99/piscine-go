curl -s https://oujda.01-edu.org/assets/superhero/all.json | jq ".[] | select(.id==${HERO_ID}) | .connections.relatives" | tr -d '"'
