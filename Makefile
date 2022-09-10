build:
	docker build -t httptester:latest .
	docker tag httptester:latest public.ecr.aws/zinclabs/httptester:latest
push:
	aws ecr-public get-login-password --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws/zinclabs
	docker push public.ecr.aws/zinclabs/httptester:latest