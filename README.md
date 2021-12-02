#Infrastructure

For this project, please think about how you would architect a scalable and secure static web
application in AWS.

• Create and deploy a running instance of a web server using a configuration management
tool of your choice. The web server should serve one page with the following content.

	<html>
	<head>
	<title>Hello World</title>
	</head>
	<body>
	<h1>Hello World!</h1>
	</body>
	</html>

• Secure this application and host such that only appropriate ports are publicly exposed and
any http requests are redirected to https. This should be automated using a configuration
management tool of your choice and you should feel free to use a self-signed certificate for
the web server.

Add Inbound rules (HTTP, HTTPs on 80, 443)
Security Group: sg-0277d5a51a3433bea 

Commands to install certificate - Certbot:

	sudo amazon-linux-extras install nginx1 -y
	sudo systemctl enable nginx
	sudo systemctl start nginx
	sudo systemctl stop nginx
	sudo wget -r --no-parent -A 'epel-release-*.rpm' https://dl.fedoraproject.org/pub/epel/7/x86_64/Packages/e/
	sudo rpm -Uvh dl.fedoraproject.org/pub/epel/7/x86_64/Packages/e/epel-release-*.rpm
	sudo yum-config-manager --enable epel*
	sudo yum install -y certbot 
	sudo yum install -y python-certbot-nginx
	sudo certbot certonly --standalone --debug -d your.domain.here
	sudo systemctl restart nginx


• Develop and apply automated tests to validate the correctness of the server configuration.

• Express everything in code.

• Provide your code in a Git repository named <FIRSTNAME>_Challenge on GitHub.com
  
Be prepared to walk though your code, discuss your thought process, and talk through how you
might monitor and scale this application. You should also be able to demo a running instance of the
host.
  
  
#Coding
	
Please solve the below problem in python or go. You don't need to do the submission on the web
site. Include the program in the repo above.
  
1. https://www.hackerrank.com/challenges/validating-credit-card-number/problem
	
	
Reference:
	HTTPS on Amazon Linux with Nginx - Certbot
	https://sammeechward.com/https-on-amazon-linux-with-nginx/
	
	Configure HTTPS on AWS EC2 without a Custom Domain
	https://bansalanuj.com/https-aws-ec2-without-custom-domain
