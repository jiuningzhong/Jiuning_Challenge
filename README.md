#Infrastructure

For this project, please think about how you would architect a scalable and secure static web
application in AWS.

S3 

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

alb-redirect-only.yaml sets up inbound rules

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

----------------------------
	
	https://jiuning-challenge.s3.us-east-1.amazonaws.com/index.html?response-content-disposition=inline&X-Amz-Security-Token=IQoJb3JpZ2luX2VjEEQaCXVzLWVhc3QtMSJGMEQCIHG2z%2FB7JS32QjDz20mRE0D9ur0PJgYVcipIEt9rPtEnAiBnV6w4mF3qYHUmKm2gAXYbsdOZTwY1qjLgzeiFc3z4hSr2AggcEAAaDDQwNDE4NDI5NTIzNiIMZWLZ7hyk4OfMO%2Bb9KtMCfDVfpULTCcWKwI3teilnHJYT0f1VhtL5Lr6XRG9kcku0NrbfEo4ha8yr%2F2spiTXlzmSBsRc%2FNEgsz3Wh8rni7tckGdAk7TvhLJUmT1%2BhLgX%2FfVccIZLJ04VOBnXjBnYAn23jFvbd%2B1o54%2FTES%2Bg3K0N0skOlStsdcRtyxMHCxvmVLhVrUbLc7htjrMBc%2BnirODtpP7PPyPGV2QhkHjGmVw8bDtR3da7%2F74wXNSAko6EenW32%2FwH8Nhw84gtPC86tYqA2VHJ1KIiHrYqB6R8zvMpPvoBpj4J5D%2FPYVe2gd1thq%2F7j2t5gBo3oQ0mp96c7l%2F1Y%2FIyE1gmSym2dnSjjitXQbC9AeM4SRYy%2Bhne%2FfCiHKreQyurl4gfXMN6WT7k8rv4P%2B1K8lZ8jXEYLhmBzh0R5%2F6Za%2BskQVvntxfzyU7X8HBUTBTOMsXyQ4Yza0n2KFT6aMJiGn40GOrQC%2B6rHM8jajvZwMlR8Ngv%2FHojJsjBRCDHtDyJbQTJr8BI5sy9Pk8RqZ8%2FObCBhCmy%2FHSRi3nrCDVv7AT3NP9veTwHdhpUin8xDNgzxxM6jckIqn2H%2BNU7kgNo19CXlL1SdA383xjD%2FQKddEg6vn3eGmD06R2LVFL8xxJ5kRAhfl7DEe3mRX6i%2F9QR%2FjY7CP18n62nfR1qmM1DAQ3Jwj2kD30Apdes4BPK7aWpzaBvFFV5hKQ4WNt2U2mr3mewdV6r6ktgVavKpwqh54NvZBxoQsxHg8gzTyYXZ%2Bpr5WYHj6JSpyzelGZpbrFmt8Myl56WMZeyYnbNNKyObBZRTp7fvlO9S1zKJrfVWWAlN0PKNcnUXVANLC%2FrYI%2BCIkD%2Fre14XP72iKfrBh9LIcY4AjUjwSZzIi7U%3D&X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20211202T045130Z&X-Amz-SignedHeaders=host&X-Amz-Expires=300&X-Amz-Credential=ASIAV4G2C45CAZFD2DWE%2F20211202%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Signature=04951ff4d82919ab3d74143a2eec877419647a8484089fb7a3fdaac8019083d9
