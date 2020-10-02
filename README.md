# cloud-devops-capstone

<h2>Project Overview</h2>

<p> This project applies various Cloud DevOps skills and knowledge in the development of gRPC microservices that run on Kubernetes(AWS EKS). These skills include:</p>

<ul>
	<li>AWS</li>
	<li>Using Jenkins to implement Continuous Integration and Continuous Deployment</li>
	<li>Building pipelines</li>
	<li>Working with CloudFormation to deploy clusters</li>
	<li>Building Kubernetes clusters</li>
	<li>Building Docker containers in pipelines</li>
</ul>

***

<p>This project uses a CI/CD pipeline for microservices applications with rolling deployment</p>

<h2>Environment Setup</h2>

<ul>
  <li>Create a new IAM profile for the deployment on AWS</li>
  <li>Create AWS infrastructure for the various CloudFormation scripts in the <code>infra/</code> folder using the <code>create-stack.sh</code> script</li>
  <li>Install Jenkins and the necessary plugins (BlueOcean, pipeline-aws). Install application-specific plugins(eksctl, kubectl, docker, awscli) in the running EC2 Instance</li>
  <li>Setup AWS and docker hub credentials in Jenkins</li>
  <li>Fork this repo and connect the forked repo in Jenkins while creating a Jenkins pipeline with BlueOcean</li>
</ul>

