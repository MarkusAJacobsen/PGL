app-name=ancient-stream-43747

buildh:
	heroku container:push web -a ${app-name}

release:
	heroku container:release web -a ${app-name}

open:
	heroku open -a ${app-name}