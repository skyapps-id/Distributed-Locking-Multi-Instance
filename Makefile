env:
	source .env && export $(cut -d= -f1 .env)