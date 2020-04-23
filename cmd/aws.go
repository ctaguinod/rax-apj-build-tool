/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func getVpc(environment, region, name, cidr string) {

	svc := ec2.New(session.New(&aws.Config{
		Region: aws.String(region),
	}))

	input := &ec2.DescribeVpcsInput{
		//VpcIds: []*string{
		//	aws.String("vpc-0850a46e6c4438e0e"),

		Filters: []*ec2.Filter{
			&ec2.Filter{
				Name: aws.String("tag:Name"),
				Values: []*string{
					aws.String(name),
				},
			},
			&ec2.Filter{
				Name: aws.String("tag:Environment"),
				Values: []*string{
					aws.String(environment),
				},
			},
			&ec2.Filter{
				Name: aws.String("cidrBlock"),
				Values: []*string{
					aws.String(cidr),
				},
			},
		},
	}

	describe, err := svc.DescribeVpcs(input)
	result, err := awsutil.ValuesAtPath(describe, "Vpcs[0]")

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	if result != nil {
		//fmt.Printf("DescribeVpcs Output: %s\n", result)
		VpcID := describe.Vpcs[0].VpcId
		fmt.Printf("###### AWS Environment ######\n")
		fmt.Printf("SubnetId: %s\n", *VpcID)

		fmt.Println("QC: PASSED")
	} else {
		fmt.Printf("DescribeVpcs Output: %s\n", result)
		fmt.Printf("QC: FAILED - Please review VPC %s\n", name)

	}

}

func getSubnets(environment, region, name, cidr, az string) {

	svc := ec2.New(session.New(&aws.Config{
		Region: aws.String(region),
	}))

	input := &ec2.DescribeSubnetsInput{
		Filters: []*ec2.Filter{
			&ec2.Filter{
				Name: aws.String("tag:Name"),
				Values: []*string{
					aws.String(name),
				},
			},
			&ec2.Filter{
				Name: aws.String("tag:Environment"),
				Values: []*string{
					aws.String(environment),
				},
			},
			&ec2.Filter{
				Name: aws.String("cidrBlock"),
				Values: []*string{
					aws.String(cidr),
				},
			},
			&ec2.Filter{
				Name: aws.String("availabilityZone"),
				Values: []*string{
					aws.String(az),
				},
			},
		},
	}

	describe, err := svc.DescribeSubnets(input)
	result, err := awsutil.ValuesAtPath(describe, "Subnets[0]")

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	if result != nil {
		//fmt.Printf("DescribeSubnets Output: %s\n", result)
		SubnetID := describe.Subnets[0].SubnetId
		fmt.Printf("###### AWS Environment ######\n")
		fmt.Printf("SubnetId: %s\n", *SubnetID)

		fmt.Println("QC: PASSED")
	} else {
		fmt.Printf("DescribeSubnets Output: %s\n", result)
		fmt.Printf("QC: FAILED - Please review subnet %s\n", name)
	}

}

func getEC2(environment, region, name, instanceType, rootVolume, rootVolumeEncrypt string) {

	svc := ec2.New(session.New(&aws.Config{
		Region: aws.String(region),
	}))

	input := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			&ec2.Filter{
				Name: aws.String("tag:Name"),
				Values: []*string{
					aws.String(name),
				},
			},
			&ec2.Filter{
				Name: aws.String("tag:Environment"),
				Values: []*string{
					aws.String(environment),
				},
			},
			&ec2.Filter{
				Name: aws.String("instance-type"),
				Values: []*string{
					aws.String(instanceType),
				},
			},
			//&ec2.Filter{
			//	Name: aws.String("subnet-id"),
			//	Values: []*string{
			//		aws.String(subnetID),
			//	},
			//},
			//{
			//	Name: aws.String("image-id"),
			//	Values: []*string{
			//		aws.String("ami-002ea395007afeafb"),
			//	},
			//},
		},
	}

	// Change encrypted value: No==false, Yes==true
	if rootVolumeEncrypt == "No" {
		rootVolumeEncrypt = fmt.Sprintf("false")
	} else if rootVolumeEncrypt == "Yes" {
		rootVolumeEncrypt = fmt.Sprintf("true")
	}

	inputVolumes := &ec2.DescribeVolumesInput{
		Filters: []*ec2.Filter{
			&ec2.Filter{
				Name: aws.String("tag:Name"),
				Values: []*string{
					aws.String(name),
				},
			},
			&ec2.Filter{
				Name: aws.String("size"),
				Values: []*string{
					aws.String(rootVolume),
				},
			},
			&ec2.Filter{
				Name: aws.String("encrypted"),
				Values: []*string{
					aws.String(rootVolumeEncrypt),
				},
			},
			&ec2.Filter{
				Name: aws.String("volume-type"),
				Values: []*string{
					aws.String("gp2"),
				},
			},
		},
	}

	describe, err := svc.DescribeInstances(input)
	result, err := awsutil.ValuesAtPath(describe, "Reservations[0]")
	describeVolumes, err := svc.DescribeVolumes(inputVolumes)
	resultVolumes, err := awsutil.ValuesAtPath(describeVolumes, "Volumes[0]")

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	if result != nil && resultVolumes != nil {
		//fmt.Printf("DescribeInstances Output: %s\n", result)
		//fmt.Printf("DescribeVolumes Output: %s\n", resultVolumes)

		ec2 := describe.Reservations[0].Instances[0]
		InstanceID := ec2.InstanceId
		SubnetID := ec2.SubnetId
		ImageID := ec2.ImageId
		KeyName := ec2.KeyName
		RootVolume := ec2.BlockDeviceMappings[0].Ebs.VolumeId
		RootVolumeType := describeVolumes.Volumes[0].VolumeType
		RootVolumeEncrypted := describeVolumes.Volumes[0].Encrypted

		fmt.Printf("###### AWS Environment ######\n")
		fmt.Printf("InstanceId: %s\n", *InstanceID)
		fmt.Printf("SubnetId: %s\n", *SubnetID)
		fmt.Printf("ImageId: %s\n", *ImageID)
		fmt.Printf("KeyName: %s\n", *KeyName)
		fmt.Printf("RootVolume: %s\n", *RootVolume)
		fmt.Printf("RootVolumeType: %s\n", *RootVolumeType)
		fmt.Printf("RootVolumeEncrypted: %v\n", *RootVolumeEncrypted)

		fmt.Println("QC: PASSED")

	} else {
		if result == nil {
			//fmt.Printf("DescribeInstances Output: %s\n", result)
			fmt.Printf("QC: FAILED - Please review EC2 Instance configs for %s\n", name)
		}
		if resultVolumes == nil {
			//fmt.Printf("DescribeVolumes Output: %s\n", resultVolumes)
			fmt.Printf("QC: FAILED - Please review EBS Volume configs for EC2 Instance %s\n", name)
		}
	}
}
