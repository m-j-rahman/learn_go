find . -name '*sh'| sed 's/\.[^.]*s//' | sed 's/.*\///' | sort -r 