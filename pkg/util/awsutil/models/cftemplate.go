package awsmodels

const K8SCFTemplate = `# Copyright 2017 by the contributors
#
#    Licensed under the Apache License, Version 2.0 (the "License");
#    you may not use this file except in compliance with the License.
#    You may obtain a copy of the License at
#
#        http://www.apache.org/licenses/LICENSE-2.0
#
#    Unless required by applicable law or agreed to in writing, software
#    distributed under the License is distributed on an "AS IS" BASIS,
#    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#    See the License for the specific language governing permissions and
#    limitations under the License.

---
AWSTemplateFormatVersion: '2010-09-09'
Description: 'QS(5042) Kubernetes AWS CloudFormation Template: Create a Kubernetes
  cluster in an existing VPC. This template is for users who want to add
  a Kubernetes cluster to existing AWS infrastructure. The master node is
  an auto-recovering Amazon EC2 instance. 1-20 additional EC2 instances in an
  AutoScalingGroup join the Kubernetes cluster as nodes. An ELB provides
  configurable external access to the Kubernetes API. If you choose a private
  subnet, make sure it has a bastion host for SSH access to your cluster. If you
  choose a public subnet, you can connect directly to the master node. The stack
  is suitable for development and small single-team clusters. **WARNING** This
  template creates four Amazon EC2 instances with default settings. You will
  be billed for the AWS resources used if you create a stack from this template.
  **SUPPORT** Please visit http://jump.heptio.com/aws-qs-help for support.
  **NEXT STEPS** Please visit http://jump.heptio.com/aws-qs-next.'

# The Metadata tells AWS how to display the parameters during stack creation
Metadata:
  AWS::CloudFormation::Interface:
    ParameterGroups:
    - Label:
        default: Amazon EC2 Configuration
      Parameters:
      - VPCID
      - AvailabilityZone
      - InstanceType
      - DiskSizeGb
      - ClusterSubnetId
      - LoadBalancerSubnetId
      - LoadBalancerType
    - Label:
        default: Access Configuration
      Parameters:
      - SSHLocation
      - ApiLbLocation
      - KeyName
    - Label:
        default: Kubernetes Configuration
      Parameters:
      - K8sNodeCapacity
      - NetworkingProvider
      - ClusterDNSProvider
    - Label:
        default: Advanced
      Parameters:
      - QSS3BucketName
      - QSS3KeyPrefix
      - ClusterAssociation

    ParameterLabels:
      KeyName:
        default: SSH Key
      VPCID:
        default: VPC
      AvailabilityZone:
        default: Availability Zone
      ClusterSubnetId:
        default: Subnet
      SSHLocation:
        default: SSH Ingress Location
      ApiLbLocation:
        default: API Ingress Location
      InstanceType:
        default: Instance Type
      DiskSizeGb:
        default: Disk Size (GiB)
      K8sNodeCapacity:
        default: Node Capacity
      QSS3BucketName:
        default: S3 Bucket
      QSS3KeyPrefix:
        default: S3 Key Prefix
      ClusterAssociation:
        default: Cluster Association
      NetworkingProvider:
        default: Networking Provider
      LoadBalancerSubnetId:
        default: Load Balancer Subnet
      LoadBalancerType:
        default: Load Balancer Type
      ClusterDNSProvider:
        default: Cluster DNS Provider


# The Parameters allow the user to pass custom settings to the stack before creation
Parameters:
  # Required. Calls for the name of an existing EC2 KeyPair, to enable SSH access to the instances
  # http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html
  KeyName:
    Description: Existing EC2 KeyPair for SSH access.
    Type: AWS::EC2::KeyPair::KeyName
    ConstraintDescription: must be the name of an existing EC2 KeyPair.

  VPCID:
    Description: Existing VPC to use for this cluster.
    Type: AWS::EC2::VPC::Id

  ClusterSubnetId:
    Description: Existing subnet to use for this cluster. Must belong to the Availability Zone above.
    Type: AWS::EC2::Subnet::Id

  LoadBalancerSubnetId:
    Description: Existing subnet to use for load balancing HTTPS access to the Kubernetes API server.
      Must be a public subnet. Must belong to the Availability Zone above.
    Type: AWS::EC2::Subnet::Id

  LoadBalancerType:
    Description:  Create Internet-facing (public, external) or Internal-facing Load Balancers
    Type: String
    Default: internet-facing
    AllowedValues: [ "internet-facing", "internal" ]

  ClusterAssociation:
    Description: Enter a string, unique within your AWS account, to associate resources in
     this Kubernetes cluster with each other. This adds a tag, with the key KubernetesCluster
     and the value of this parameter, to resources created as part of this stack. Leave blank
     to use this Quick Start Stack name.
    Type: String

  # https://aws.amazon.com/ec2/instance-types/
  InstanceType:
    Description: EC2 instance type for the cluster.
    Type: String
    Default: m4.large
    AllowedValues:
    - m5.large
    - m5.xlarge
    - m5.2xlarge
    - m5.4xlarge
    - m5.12xlarge
    - m5.24xlarge
    - c5.large
    - c5.xlarge
    - c5.2xlarge
    - c5.4xlarge
    - c5.9xlarge
    - c5.18xlarge
    - r4.large
    - r4.xlarge
    - r4.2xlarge
    - r4.4xlarge
    - r4.8xlarge
    - r4.16xlarge
    - x1.16xlarge
    - x1.32xlarge
    - i3.large
    - i3.xlarge
    - i3.2xlarge
    - i3.4xlarge
    - i3.8xlarge
    - i3.16xlarge
    - p2.xlarge
    - p2.8xlarge
    - p2.16xlarge
    - p3.2xlarge
    - p3.8xlarge
    - p3.16xlarge
    - m4.large
    - m4.xlarge
    - m4.2xlarge
    - m4.4xlarge
    - m4.10xlarge
    - m4.16xlarge
    - c4.large
    - c4.xlarge
    - c4.2xlarge
    - c4.4xlarge
    - c4.8xlarge
    - g3.4xlarge
    - g3.8xlarge
    - g3.16xlarge
    - r3.large
    - r3.xlarge
    - r3.2xlarge
    - r3.4xlarge
    - r3.8xlarge
    ConstraintDescription: must be a valid Current Generation (non-burstable) EC2 instance type.

  # Specifies the size of the root disk for all EC2 instances, including master
  # and nodes.
  DiskSizeGb:
    Description: 'Size of the root disk for the EC2 instances, in GiB.  Default: 40'
    Default: 40
    Type: Number
    MinValue: 8
    MaxValue: 1024

  # Required. This is an availability zone from your region
  # http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html
  AvailabilityZone:
    Description: The Availability Zone for this cluster. Heptio recommends
      that you run one cluster per AZ and use tooling to coordinate across AZs.
    Type: AWS::EC2::AvailabilityZone::Name
    ConstraintDescription: must be the name of an AWS Availability Zone

  # Specifies the IP range from which you will have SSH access over port 22
  # Used in the allow22 SecurityGroup
  SSHLocation:
    Description: CIDR block (IP address range) to allow SSH access to the
      instances. Use 0.0.0.0/0 to allow access from all locations.
    Type: String
    MinLength: '9'
    MaxLength: '18'
    AllowedPattern: "(\\d{1,3})\\.(\\d{1,3})\\.(\\d{1,3})\\.(\\d{1,3})/(\\d{1,2})"
    ConstraintDescription: must be a valid IP CIDR range of the form x.x.x.x/x.

  # Specifies the IP range from which you will have HTTPS access to the Kubernetes API server load balancer
  # Used in the ApiLoadBalancerSecGroup SecurityGroup
  ApiLbLocation:
    Description: CIDR block (IP address range) to allow HTTPS access to
      the Kubernetes API. Use 0.0.0.0/0 to allow access from all locations.
    Type: String
    MinLength: '9'
    MaxLength: '18'
    AllowedPattern: "(\\d{1,3})\\.(\\d{1,3})\\.(\\d{1,3})\\.(\\d{1,3})/(\\d{1,2})"
    ConstraintDescription: must be a valid IP CIDR range of the form x.x.x.x/x.

  # Default 2. Choose 1-20 initial nodes to run cluster workloads (in addition to the master node instance)
  # You can scale up your cluster later and add more nodes
  K8sNodeCapacity:
    Default: '2'
    Description: Initial number of Kubernetes nodes (1-20).
    Type: Number
    MinValue: '1'
    MaxValue: '20'
    ConstraintDescription: must be between 1 and 20 EC2 instances.

  # S3 Bucket configuration: allows users to use their own downstream snapshots
  # of the quickstart-aws-vpc and quickstart-linux-bastion templates
  QSS3BucketName:
    AllowedPattern: "^[0-9a-zA-Z]+([0-9a-zA-Z-]*[0-9a-zA-Z])*$"
    ConstraintDescription: Quick Start bucket name can include numbers, lowercase
      letters, uppercase letters, and hyphens (-). It cannot start or end with a hyphen
      (-).

    Default: aws-quickstart
    Description: Only change this if you have set up assets, like your own networking
      configuration, in an S3 bucket. This and the S3 Key Prefix parameter let you access
      scripts from the scripts/ and templates/ directories of your own fork of the Heptio
      Quick Start assets, uploaded to S3 and stored at
      ${bucketname}.s3.amazonaws.com/${prefix}/scripts/somefile.txt.S3. The bucket name
      can include numbers, lowercase letters, uppercase letters, and hyphens (-).
      It cannot start or end with a hyphen (-).
    Type: String

  QSS3KeyPrefix:
    AllowedPattern: ^[0-9a-zA-Z-/]*$
    ConstraintDescription: Quick Start key prefix can include numbers, lowercase letters, uppercase
      letters, hyphens (-), and forward slash (/).
    Default: quickstart-heptio/
    Description: Only change this if you have set up assets in an S3 bucket, as explained
      in the S3 Bucket parameter. S3 key prefix for the Quick Start assets.
      Quick Start key prefix can include numbers, lowercase letters, uppercase
      letters, hyphens (-), and forward slash (/).
    Type: String

  NetworkingProvider:
    AllowedValues:
    - calico
    - weave
    ConstraintDescription: 'Currently supported values are "calico" and "weave"'
    Default: calico
    Description: Choose the networking provider to use for communication between
      pods in the Kubernetes cluster. Supported configurations are calico
      (https://docs.projectcalico.org/v2.6/getting-started/kubernetes/installation/hosted/kubeadm/)
      and weave (https://github.com/weaveworks/weave/blob/master/site/kubernetes/kube-addon.md).
    Type: String

  ClusterDNSProvider:
    AllowedValues:
    - CoreDNS
    - KubeDNS
    ConstraintDescription: 'Currently supported values are "CoreDNS" and "KubeDNS"'
    Default: CoreDNS
    Description: Choose the cluster DNS provider to use for internal cluster DNS. Supported
      configurations are CoreDNS and KubeDNS
    Type: String

# http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/mappings-section-structure.html
Mappings:
  # http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html
  RegionMap:
    ap-northeast-1:
      '64': ami-0f159b60cdff05086
    ap-northeast-2:
      '64': ami-0680ca62fb23e2d9e
    ap-south-1:
      '64': ami-023b820d038deb341
    ap-southeast-1:
      '64': ami-00477ca666574012f
    ap-southeast-2:
      '64': ami-010afcf89fcf332b7
    ca-central-1:
      '64': ami-0e7fa16f04b747613
    eu-central-1:
      '64': ami-0954aa00742537446
    eu-west-1:
      '64': ami-0d8ce815b6762bf57
    eu-west-2:
      '64': ami-046dc2dacfa0ae854
    eu-west-3:
      '64': ami-0b159220317af0751
    sa-east-1:
      '64': ami-07691c625b83f4021
    us-east-2:
      '64': ami-006c052f934cd9329
    us-west-1:
      '64': ami-07022640a9cef0f13
    us-west-2:
      '64': ami-057c58c0eca3e6fe3
    us-east-1:
      '64': ami-0c0803fb7e8ca30d3
# Helper Conditions which help find the right values for resources
Conditions:
  AssociationProvidedCondition:
    Fn::Not:
    - Fn::Equals:
      - !Ref ClusterAssociation
      - ''
  LoadBalancerSubnetProvidedCondition:
    Fn::Not:
    - Fn::Equals:
      - !Ref LoadBalancerSubnetId
      - ''

# Resources are the AWS services we want to actually create as part of the Stack
Resources:
  ClusterInfoBucket:
    Type: AWS::S3::Bucket
    Properties:
      AccessControl: Private
      Tags:
      - Key: KubernetesCluster
        Value:
          Fn::If:
          - AssociationProvidedCondition
          - !Ref ClusterAssociation
          - !Ref AWS::StackName

  # Install a CloudWatch logging group for system logs for each instance
  KubernetesLogGroup:
    Type: AWS::Logs::LogGroup
    Properties:
      LogGroupName: !Ref AWS::StackName
      RetentionInDays: 14

  # This is an EC2 instance that will serve as our master node
  K8sMasterInstance:
    Type: AWS::EC2::Instance
    DependsOn: ApiLoadBalancer
    Metadata:
      AWS::CloudFormation::Init:
        configSets:
          master-setup: master-setup
        master-setup:
          files:
            # Script that will allow for development kubernetes binaries to replace the pre-packaged AMI binaries.
            "/tmp/kubernetes-override-binaries.sh":
              source: !Sub "https://${QSS3BucketName}.s3.amazonaws.com/${QSS3KeyPrefix}scripts/kubernetes-override-binaries.sh.in"
              mode: '000755'
              context:
                BaseBinaryUrl: !Sub "https://${QSS3BucketName}.s3.amazonaws.com/${QSS3KeyPrefix}bin/"

            # Configuration file for the cloudwatch agent. The file is a Mustache template, and we're creating it with
            # the below context (mainly to substitute in the AWS Stack Name for the logging group.)
            "/tmp/kubernetes-awslogs.conf":
              source: !Sub "https://${QSS3BucketName}.s3.amazonaws.com/${QSS3KeyPrefix}scripts/kubernetes-awslogs.conf"
              context:
                StackName: !Ref AWS::StackName

            # Installation script for the Cloudwatch agent
            "/usr/local/aws/awslogs-agent-setup.py":
              source: https://s3.amazonaws.com/aws-cloudwatch/downloads/latest/awslogs-agent-setup.py
              mode: '000755'

            # systemd init script for the Cloudwatch logs agent
            "/etc/systemd/system/awslogs.service":
              source: !Sub "https://${QSS3BucketName}.s3.amazonaws.com/${QSS3KeyPrefix}scripts/awslogs.service"

            # setup kubelet hostname
            "/tmp/setup-kubelet-hostname.sh":
              source: !Sub "https://${QSS3BucketName}.s3.amazonaws.com/${QSS3KeyPrefix}scripts/setup-kubelet-hostname.sh"
              mode: '000755'

            # Setup script for initializing the Kubernetes master instance.  This is where most of the cluster
            # initialization happens.  See scripts/setup-k8s-master.sh in the Quick Start repo for details.
            "/tmp/setup-k8s-master.sh":
              source: !Sub "https://${QSS3BucketName}.s3.amazonaws.com/${QSS3KeyPrefix}scripts/setup-k8s-master.sh.in"
              mode: '000755'
              context:
                LoadBalancerDns: !GetAtt ApiLoadBalancer.DNSName
                LoadBalancerName: !Ref ApiLoadBalancer
                ClusterToken: !GetAtt KubeadmToken.Token
                ClusterDNSProvider: !Ref ClusterDNSProvider
                NetworkingProvider: !Ref NetworkingProvider
                NetworkingProviderUrl: !Sub "https://${QSS3BucketName}.s3.amazonaws.com/${QSS3KeyPrefix}scripts/${NetworkingProvider}.yaml"
                DashboardUrl: !Sub "https://${QSS3BucketName}.s3.amazonaws.com/${QSS3KeyPrefix}scripts/dashboard.yaml"
                StorageClassUrl: !Sub "https://${QSS3BucketName}.s3.amazonaws.com/${QSS3KeyPrefix}scripts/default.storageclass.yaml"
                NetworkPolicyUrl: !Sub "https://${QSS3BucketName}.s3.amazonaws.com/${QSS3KeyPrefix}scripts/network-policy.yaml"
                ClusterInfoBucket: !Ref ClusterInfoBucket
                Region: !Ref AWS::Region

            # patch kube proxy
            "/tmp/patch-kube-proxy.sh":
              source: !Sub "https://${QSS3BucketName}.s3.amazonaws.com/${QSS3KeyPrefix}scripts/patch-kube-proxy.sh"
              mode: '000755'

          commands:
            # Override the AMI binaries with any kubelet/kubeadm/kubectl binaries in the S3 bucket
            "00-kubernetes-override-binaries":
              command: "/tmp/kubernetes-override-binaries.sh"
            # Install the Cloudwatch agent with configuration for the current region and log group name
            "01-cloudwatch-agent-setup":
              command: !Sub "python /usr/local/aws/awslogs-agent-setup.py -n -r ${AWS::Region} -c /tmp/kubernetes-awslogs.conf"
            # Enable the Cloudwatch service and launch it
            "02-cloudwatch-service-config":
              command: "systemctl enable awslogs.service && systemctl start awslogs.service"
            # Setup kubelet hostname
            "03-setup-kubelet-hostname":
              command: "/tmp/setup-kubelet-hostname.sh"
            # Run the master setup
            "04-master-setup":
              command: "/tmp/setup-k8s-master.sh"
            "05-patch-kube-proxy":
              command: "/tmp/patch-kube-proxy.sh"
    Properties:
      # Where the EC2 instance gets deployed geographically
      AvailabilityZone: !Ref AvailabilityZone
      # Refers to the MasterInstanceProfile resource, which applies the IAM role for the master instance
      # The IAM role allows us to create further AWS resources (like an EBS drive) from the cluster
      # This is needed for the Kubernetes-AWS cloud-provider integration
      IamInstanceProfile: !Ref MasterInstanceProfile
      # Type of instance; the default is m3.medium
      InstanceType: !Ref InstanceType
      # Adds our SSH key to the instance
      KeyName: !Ref KeyName
      NetworkInterfaces:
      - DeleteOnTermination: true
        DeviceIndex: 0
        SubnetId: !Ref ClusterSubnetId
        # Joins the ClusterSecGroup Security Group for cluster communication and SSH access
        # The ClusterSecGroupCrossTalk rules allow all instances in the same stack to communicate internally
        # The ClusterSecGroupAllow22 rules allow external communication on port 22 from a chosen CIDR range
        # The ClusterSecGroupAllow6443FromLB rules allow HTTPS access to the load balancer on port 6443
        GroupSet:
        - !Ref ClusterSecGroup
      # Designates a name for this EC2 instance that will appear in the instances list (k8s-master)
      # Tags it with KubernetesCluster=<stackname> or chosen value (needed for cloud-provider's IAM roles)
      Tags:
      - Key: Name
        Value: k8s-master
      - Key: KubernetesCluster
        Value:
          Fn::If:
          - AssociationProvidedCondition
          - !Ref ClusterAssociation
          - !Ref AWS::StackName
        # Also tag it with kubernetes.io/cluster/clustername=owned, which is the newer convention for cluster resources
      - Key:
          Fn::Sub:
          - "kubernetes.io/cluster/${ClusterID}"
          - ClusterID:
              Fn::If:
              - AssociationProvidedCondition
              - !Ref ClusterAssociation
              - !Ref AWS::StackName
        Value: 'owned'
      # http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-imageid
      ImageId:
        Fn::FindInMap:
        - RegionMap
        - !Ref AWS::Region
        - '64'
      BlockDeviceMappings:
      - DeviceName: '/dev/sda1'
        Ebs:
          VolumeSize: !Ref DiskSizeGb
          VolumeType: gp2
      # The userdata script is launched on startup, but contains only the commands that call out to cfn-init, which runs
      # the commands in the metadata above, and cfn-signal, which signals when the initialization is complete.
      UserData:
        Fn::Base64:
          Fn::Sub: |
            #!/bin/bash
            set -o xtrace

            CFN_INIT=$(which cfn-init)
            CFN_SIGNAL=$(which cfn-signal)

            ${!CFN_INIT} \
              --verbose \
              --stack '${AWS::StackName}' \
              --region '${AWS::Region}' \
              --resource K8sMasterInstance \
              --configsets master-setup

            ${!CFN_SIGNAL} \
              --exit-code $? \
              --stack '${AWS::StackName}' \
              --region '${AWS::Region}' \
              --resource K8sMasterInstance
    CreationPolicy:
      ResourceSignal:
        Timeout: PT10M

  # IAM role for Lambda function for generating kubeadm token
  LambdaExecutionRole:
    Type: "AWS::IAM::Role"
    Properties:
      AssumeRolePolicyDocument:
        Version: "2012-10-17"
        Statement:
        - Effect: "Allow"
          Principal:
            Service: ["lambda.amazonaws.com"]
          Action: "sts:AssumeRole"
      Path: "/"
      Policies:
      - PolicyName: "lambda_policy"
        PolicyDocument:
          Version: "2012-10-17"
          Statement:
          - Effect: "Allow"
            Action:
            - "logs:CreateLogGroup"
            - "logs:CreateLogStream"
            - "logs:PutLogEvents"
            Resource: "arn:aws:logs:*:*:*"

  # Lambda Function for generating the kubeadm token
  GenKubeadmToken:
    Type: "AWS::Lambda::Function"
    Properties:
      Code:
        ZipFile: |
          import random
          import string
          import cfnresponse
          def id_generator(size, chars=string.ascii_lowercase + string.digits):
            return ''.join(random.choice(chars) for _ in range(size))
          def handler(event, context):
            if event['RequestType'] == 'Delete':
              cfnresponse.send(event, context, cfnresponse.SUCCESS, {})
            if event['RequestType'] == 'Create':
              token = ("%s.%s" % (id_generator(6), id_generator(16)))
              responseData = {}
              responseData['Token'] = token
              cfnresponse.send(event, context, cfnresponse.SUCCESS, responseData)
              return token
      Handler: "index.handler"
      Runtime: "python2.7"
      Timeout: "5"
      Role: !GetAtt LambdaExecutionRole.Arn

  # A Custom resource that uses the lambda function to generate our cluster token
  KubeadmToken:
    Type: "Custom::GenerateToken"
    Version: "1.0"
    Properties:
      ServiceToken: !GetAtt GenKubeadmToken.Arn

  # This is a CloudWatch alarm https://aws.amazon.com/cloudwatch/
  # If the master node is unresponsive for 5 minutes, AWS will attempt to recover it
  # It will preserve the original IP, which is important for Kubernetes networking
  # Based on http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/quickref-cloudwatch.html#cloudwatch-sample-recover-instance
  RecoveryTestAlarm:
    Type: AWS::CloudWatch::Alarm
    Properties:
      AlarmDescription: Trigger a recovery when instance status check fails for 5
        consecutive minutes.
      Namespace: AWS/EC2
      MetricName: StatusCheckFailed_System
      Statistic: Minimum
      # 60-second periods (1 minute)
      Period: '60'
      # 5-minute check-ins
      EvaluationPeriods: '5'
      ComparisonOperator: GreaterThanThreshold
      Threshold: '0'
      # This is the call that actually tries to recover the instance
      AlarmActions:
        - !Sub "arn:aws:automate:${AWS::Region}:ec2:recover"
      # Applies this alarm to our K8sMasterInstance
      Dimensions:
      - Name: InstanceId
        Value: !Ref K8sMasterInstance

  # This is the Auto Scaling Group that contains EC2 instances that are Kubernetes nodes
  # http://docs.aws.amazon.com/autoscaling/latest/userguide/AutoScalingGroup.html
  K8sNodeGroup:
    Type: AWS::AutoScaling::AutoScalingGroup
    DependsOn: K8sMasterInstance
    CreationPolicy:
      ResourceSignal:
        # Ensure at least <K8sNodeCapacity> nodes have signaled success before
        # this resource is considered created.
        Count: !Ref K8sNodeCapacity
        Timeout: PT10M
    Properties:
      # Where the EC2 instance gets deployed geographically
      AvailabilityZones:
      - !Ref AvailabilityZone
      # Refers to the K8sNodeCapacity parameter, which specifies the number of nodes (1-20)
      DesiredCapacity: !Ref K8sNodeCapacity
      # Refers to the LaunchConfig, which has specific config details for the EC2 instances
      LaunchConfigurationName: !Ref LaunchConfig
      # More cluster sizing
      MinSize: '1'
      MaxSize: '20'
      # VPC Zone Identifier is the subnets to put the hosts in
      VPCZoneIdentifier:
        - !Ref ClusterSubnetId
      # Designates names for these EC2 instances that will appear in the instances list (k8s-node)
      # Tags each node with KubernetesCluster=<stackname> or chosen value (needed for cloud-provider's IAM roles)
      Tags:
      - Key: Name
        Value: k8s-node
        PropagateAtLaunch: 'true'
      - Key: KubernetesCluster
        Value:
          Fn::If:
          - AssociationProvidedCondition
          - !Ref ClusterAssociation
          - !Ref AWS::StackName
        PropagateAtLaunch: 'true'
        # Also tag it with kubernetes.io/cluster/clustername=owned, which is the newer convention for cluster resources
      - Key:
          Fn::Sub:
          - "kubernetes.io/cluster/${ClusterID}"
          - ClusterID:
              Fn::If:
              - AssociationProvidedCondition
              - !Ref ClusterAssociation
              - !Ref AWS::StackName
        Value: 'owned'
        PropagateAtLaunch: 'true'
    # Tells the group how many instances to update at a time, if an update is applied
    UpdatePolicy:
      AutoScalingRollingUpdate:
        MinInstancesInService: '1'
        MaxBatchSize: '1'

  # This tells AWS what kinds of servers we want in our Auto Scaling Group
  LaunchConfig:
    Type: AWS::AutoScaling::LaunchConfiguration
    Metadata:
      AWS::CloudFormation::Init:
        configSets:
          node-setup: node-setup
        node-setup:
          # (See comments in the master instance Metadata for details.)
          files:
            "/tmp/kubernetes-override-binaries.sh":
              source: !Sub "https://${QSS3BucketName}.s3.amazonaws.com/${QSS3KeyPrefix}scripts/kubernetes-override-binaries.sh.in"
              mode: '000755'
              context:
                BaseBinaryUrl: !Sub "https://${QSS3BucketName}.s3.amazonaws.com/${QSS3KeyPrefix}bin/"
            "/tmp/kubernetes-awslogs.conf":
              source: !Sub "https://${QSS3BucketName}.s3.amazonaws.com/${QSS3KeyPrefix}scripts/kubernetes-awslogs.conf"
              context:
                StackName: !Ref AWS::StackName
            "/usr/local/aws/awslogs-agent-setup.py":
              source: https://s3.amazonaws.com/aws-cloudwatch/downloads/latest/awslogs-agent-setup.py
              mode: '000755'
            "/etc/systemd/system/awslogs.service":
              source: !Sub "https://${QSS3BucketName}.s3.amazonaws.com/${QSS3KeyPrefix}scripts/awslogs.service"
            "/tmp/setup-kubelet-hostname.sh":
              source: !Sub "https://${QSS3BucketName}.s3.amazonaws.com/${QSS3KeyPrefix}scripts/setup-kubelet-hostname.sh"
              mode: '000755'
            "/tmp/setup-k8s-node.sh":
              mode: '000755'
              source: !Sub "https://${QSS3BucketName}.s3.amazonaws.com/${QSS3KeyPrefix}scripts/setup-k8s-node.sh.in"
              context:
                K8sMasterPrivateIp: !GetAtt K8sMasterInstance.PrivateIp
                ClusterToken: !GetAtt KubeadmToken.Token
                ClusterInfoBucket: !Ref ClusterInfoBucket

          commands:
            "00-kubernetes-override-binaries":
              command: "/tmp/kubernetes-override-binaries.sh"
            "01-cloudwatch-agent-setup":
              command: !Sub "python /usr/local/aws/awslogs-agent-setup.py -n -r ${AWS::Region} -c /tmp/kubernetes-awslogs.conf"
            "02-cloudwatch-service-config":
              command: "systemctl enable awslogs.service && systemctl start awslogs.service"
            "03-setup-kubelet-hostname":
              command: "/tmp/setup-kubelet-hostname.sh"
            "04-k8s-setup-node":
              command: "/tmp/setup-k8s-node.sh"
    Properties:
      # Refers to the NodeInstanceProfile resource, which applies the IAM role for the nodes
      # The IAM role allows us to create further AWS resources (like an EBS drive) from the cluster
      # This is needed for the Kubernetes-AWS cloud-provider integration
      IamInstanceProfile: !Ref NodeInstanceProfile
      # http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-imageid
      ImageId:
        Fn::FindInMap:
        - RegionMap
        - !Ref AWS::Region
        - '64'
      BlockDeviceMappings:
      - DeviceName: '/dev/sda1'
        Ebs:
          VolumeSize: !Ref DiskSizeGb
          VolumeType: gp2
      # Type of instance; the default is m3.medium
      InstanceType: !Ref InstanceType
      # Adds our SSH key to the instance
      KeyName: !Ref KeyName
      # Join the cluster security group so that we can customize the access
      # control (See the ClusterSecGroup resource for details)
      SecurityGroups:
      - !Ref ClusterSecGroup
      # The userdata script is launched on startup, but contains only the commands that call out to cfn-init, which runs
      # the commands in the metadata above, and cfn-signal, which signals when the initialization is complete.
      UserData:
        Fn::Base64:
          Fn::Sub: |
            #!/bin/bash
            set -o xtrace

            /usr/local/bin/cfn-init \
              --verbose \
              --stack '${AWS::StackName}' \
              --region '${AWS::Region}' \
              --resource LaunchConfig \
              --configsets node-setup

            /usr/local/bin/cfn-signal \
              --exit-code $? \
              --stack '${AWS::StackName}' \
              --region '${AWS::Region}' \
              --resource K8sNodeGroup

  # Define the (one) security group for all machines in the cluster.  Keeping
  # just one security group helps with k8s's cloud-provider=aws integration so
  # that it knows what security group to manage.
  ClusterSecGroup:
    Type: AWS::EC2::SecurityGroup
    Properties:
      GroupDescription: Security group for all machines in the cluster
      VpcId: !Ref VPCID
      # Security Groups must be tagged with KubernetesCluster=<cluster> so that
      # they can coexist in the same VPC
      Tags:
      - Key: KubernetesCluster
        Value:
          Fn::If:
          - AssociationProvidedCondition
          - !Ref ClusterAssociation
          - !Ref AWS::StackName
      - Key:
          Fn::Sub:
          - "kubernetes.io/cluster/${ClusterID}"
          - ClusterID:
              Fn::If:
              - AssociationProvidedCondition
              - !Ref ClusterAssociation
              - !Ref AWS::StackName
        Value: 'owned'
      - Key: Name
        Value: k8s-cluster-security-group

  # Permissions we add to the main security group:
  # - Ensure cluster machines can talk to one another
  ClusterSecGroupCrossTalk:
    Type: AWS::EC2::SecurityGroupIngress
    Properties:
      GroupId: !Ref ClusterSecGroup
      SourceSecurityGroupId: !Ref ClusterSecGroup
      IpProtocol: '-1'
      FromPort: '0'
      ToPort: '65535'

  # - Open up port 22 for SSH into each machine
  # The allowed locations are chosen by the user in the SSHLocation parameter
  ClusterSecGroupAllow22:
    Metadata:
      Comment: Open up port 22 for SSH into each machine
    Type: AWS::EC2::SecurityGroupIngress
    Properties:
      GroupId: !Ref ClusterSecGroup
      IpProtocol: tcp
      FromPort: '22'
      ToPort: '22'
      CidrIp: !Ref SSHLocation

  # Allow the apiserver load balancer to talk to the cluster on port 6443
  ClusterSecGroupAllow6443FromLB:
    Metadata:
      Comment: Open up port 6443 for load balancing the API server
    Type: AWS::EC2::SecurityGroupIngress
    Properties:
      GroupId: !Ref ClusterSecGroup
      IpProtocol: tcp
      FromPort: '6443'
      ToPort: '6443'
      SourceSecurityGroupId: !Ref ApiLoadBalancerSecGroup

  # IAM role for nodes http://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html
  NodeRole:
    Type: AWS::IAM::Role
    Properties:
      AssumeRolePolicyDocument:
        Version: '2012-10-17'
        Statement:
        - Effect: Allow
          Principal:
            Service:
            - ec2.amazonaws.com
          Action:
          - sts:AssumeRole
      Path: "/"
      # IAM policy for nodes that allows specific AWS resource listing and creation
      # http://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html
      Policies:
      - PolicyName: node
        PolicyDocument:
          Version: '2012-10-17'
          Statement:
          - Effect: Allow
            Action:
            - ec2:Describe*
            - ecr:GetAuthorizationToken
            - ecr:BatchCheckLayerAvailability
            - ecr:GetDownloadUrlForLayer
            - ecr:GetRepositoryPolicy
            - ecr:DescribeRepositories
            - ecr:ListImages
            - ecr:BatchGetImage
            Resource: "*"

      - PolicyName: cwlogs
        PolicyDocument:
          Version: '2012-10-17'
          Statement:
          - Effect: Allow
            Action:
            - logs:CreateLogGroup
            - logs:CreateLogStream
            - logs:PutLogEvents
            - logs:DescribeLogStreams
            Resource: !Sub ["${LogGroupArn}:*", LogGroupArn: !GetAtt KubernetesLogGroup.Arn]

      - PolicyName: discoverBucketWrite
        PolicyDocument:
          Version: '2012-10-17'
          Statement:
          - Effect: Allow
            Action:
            - s3:GetObject
            Resource: !Sub "arn:aws:s3:::${ClusterInfoBucket}/cluster-info.yaml"

  # Resource that creates the node IAM role
  NodeInstanceProfile:
    Type: AWS::IAM::InstanceProfile
    Properties:
      Path: "/"
      Roles:
      - !Ref NodeRole

  # IAM role for the master node http://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html
  MasterRole:
    Type: AWS::IAM::Role
    Properties:
      AssumeRolePolicyDocument:
        Version: '2012-10-17'
        Statement:
        - Effect: Allow
          Principal:
            Service:
            - ec2.amazonaws.com
          Action:
          - sts:AssumeRole
      Path: "/"
      # IAM policy for the master node that allows specific AWS resource listing and creation
      # More permissive than the node role (it allows load balancer creation)
      # http://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html
      Policies:
      - PolicyName: master
        PolicyDocument:
          Version: '2012-10-17'
          Statement:
          - Effect: Allow
            Action:
            - ec2:*
            - elasticloadbalancing:*
            - ecr:GetAuthorizationToken
            - ecr:BatchCheckLayerAvailability
            - ecr:GetDownloadUrlForLayer
            - ecr:GetRepositoryPolicy
            - ecr:DescribeRepositories
            - ecr:ListImages
            - ecr:BatchGetImage
            - autoscaling:DescribeAutoScalingGroups
            - autoscaling:UpdateAutoScalingGroup
            Resource: "*"

      - PolicyName: cwlogs
        PolicyDocument:
          Version: '2012-10-17'
          Statement:
          - Effect: Allow
            Action:
            - logs:CreateLogGroup
            - logs:CreateLogStream
            - logs:PutLogEvents
            - logs:DescribeLogStreams
            Resource: !Sub ["${LogGroupArn}:*", LogGroupArn: !GetAtt KubernetesLogGroup.Arn]

      - PolicyName: discoverBucketWrite
        PolicyDocument:
          Version: '2012-10-17'
          Statement:
          - Effect: Allow
            Action:
            - s3:PutObject
            Resource: !Sub "arn:aws:s3:::${ClusterInfoBucket}/cluster-info.yaml"

  # Bind the MasterRole to a profile for the VM instance.
  MasterInstanceProfile:
    Type: AWS::IAM::InstanceProfile
    Properties:
      Path: "/"
      Roles:
      - !Ref MasterRole

  # Create a placeholder load balancer for the API server. Backend instances will be added by the master itself on the
  # first boot in the startup script above. The load balancer listens on both port 443 and 6443. 
  # The port 6443 is the api port specified in kubeadm init and it's used by clients to connet to the master.
  ApiLoadBalancer:
    Type: AWS::ElasticLoadBalancing::LoadBalancer
    Properties:
      Scheme: !Ref LoadBalancerType
      Listeners:
      - Protocol: TCP
        InstancePort: 6443
        InstanceProtocol: TCP
        LoadBalancerPort: 443
      - Protocol: TCP
        InstancePort: 6443
        InstanceProtocol: TCP
        LoadBalancerPort: 6443
      HealthCheck:
        Target: TCP:6443
        HealthyThreshold: '3'
        UnhealthyThreshold: '2'
        Interval: '10'
        Timeout: '5'
      ConnectionSettings:
        IdleTimeout: 3600
      Subnets:
      - Fn::If:
        - LoadBalancerSubnetProvidedCondition
        - !Ref LoadBalancerSubnetId
        - !Ref ClusterSubnetId
      SecurityGroups:
      - !Ref ApiLoadBalancerSecGroup
      Tags:
      - Key: KubernetesCluster
        Value:
          Fn::If:
          - AssociationProvidedCondition
          - !Ref ClusterAssociation
          - !Ref AWS::StackName
      - Key:
          Fn::Sub:
          - "kubernetes.io/cluster/${ClusterID}"
          - ClusterID:
              Fn::If:
              - AssociationProvidedCondition
              - !Ref ClusterAssociation
              - !Ref AWS::StackName
        Value: 'owned'
      - Key: 'kubernetes.io/service-name'
        Value: 'kube-system/apiserver-public'

  # Security group to allow public access to port 443 and 6443
  ApiLoadBalancerSecGroup:
    Type: AWS::EC2::SecurityGroup
    Properties:
      GroupDescription: Security group for API server load balancer
      VpcId: !Ref VPCID
      SecurityGroupIngress:
      - CidrIp: !Ref ApiLbLocation
        FromPort: 443
        ToPort: 443
        IpProtocol: tcp
      - CidrIp: !Ref ApiLbLocation
        FromPort: 6443
        ToPort: 6443
        IpProtocol: tcp
      Tags:
      - Key: Name
        Value: apiserver-lb-security-group

  # Give the cleanup lambda function the necessary policy.
  CleanupClusterInfoRole:
    Type: "AWS::IAM::Role"
    Properties:
      AssumeRolePolicyDocument:
        Version: "2012-10-17"
        Statement:
        - Effect: "Allow"
          Principal:
            Service: ["lambda.amazonaws.com"]
          Action: "sts:AssumeRole"
      Path: "/"
      Policies:
      - PolicyName: s3upload
        PolicyDocument:
          Version: '2012-10-17'
          Statement:
          - Effect: Allow
            Action:
            - logs:CreateLogGroup
            - logs:CreateLogStream
            - logs:PutLogEvents
            Resource: arn:aws:logs:*:*:*
          - Effect: Allow
            Action: ['s3:DeleteObject']
            Resource: [!Sub "arn:aws:s3:::${ClusterInfoBucket}/cluster-info.yaml"]

  # CleanupClusterInfo backs the custom resource defined below.
  # When the custom resource gets created, deleted or updated this function will be executed.
  CleanupClusterInfo:
    Type: "AWS::Lambda::Function"
    Properties:
      Code:
        ZipFile:
          Fn::Sub: |
            import boto3
            import cfnresponse

            def lambda_handler(event, context):
                try:
                    s3 = boto3.client('s3')
                    bucket = '${ClusterInfoBucket}'
                    key = 'cluster-info.yaml'

                    if event['RequestType'] == 'Delete':
                        s3.delete_object(Bucket=bucket, Key=key)

                    cfnresponse.send(event, context, cfnresponse.SUCCESS, {})
                    return
                except Exception as e:
                    print(e)
                cfnresponse.send(event, context, cfnresponse.FAILED, {})
      Handler: "index.lambda_handler"
      Runtime: "python3.6"
      Timeout: "5"
      Role: !GetAtt CleanupClusterInfoRole.Arn

  CleanupClusterInfoOnDelete:
    Type: "Custom::CleanupClusterInfo"
    Properties:
      ServiceToken: !GetAtt CleanupClusterInfo.Arn

# Outputs are what AWS will show you after stack creation
# Generally they let you easily access some information about the stack
# like what IP address is assigned to your master node
# Read Descriptions below for more detail
Outputs:
  MasterInstanceId:
    Description: InstanceId of the master EC2 instance.
    Value: !Ref K8sMasterInstance

  MasterPrivateIp:
    Description: Private IP address of the master.
    Value: !GetAtt K8sMasterInstance.PrivateIp

  NodeGroupInstanceId:
    Description: InstanceId of the newly-created NodeGroup.
    Value: !Ref K8sNodeGroup

  JoinNodes:
    Description: Command to join more nodes to this cluster.
    Value: !Sub "aws s3 cp s3://${ClusterInfoBucket}/cluster-info.yaml /tmp/cluster-info.yaml && kubeadm join --node-name=\"$(hostname -f 2>/dev/null || curl http://169.254.169.254/latest/meta-data/local-hostname)\" --token=${KubeadmToken.Token} --discovery-file=/tmp/cluster-info.yaml ${K8sMasterInstance.PrivateIp}:6443"

  NextSteps:
    Description: Verify your cluster and deploy a test application. Instructions -
      http://jump.heptio.com/aws-qs-next
    Value: http://jump.heptio.com/aws-qs-next
`