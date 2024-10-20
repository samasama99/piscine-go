if [ -z "${1}" ]; then 
	echo "Argument is empty"; 
else 
	gofumpt -d . &&
	git add $1 &&
	git commit -m "$1" &&
	git push
fi
