find -name "*.sh" | rev | cut -d '/' -f 1 | rev | cut -d '.' -f 1 | sort -r
