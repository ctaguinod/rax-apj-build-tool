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
		fmt.Println("DescribeVpcs Output:")
		fmt.Println(result)
		fmt.Println("QC Status: PASS")
		fmt.Println()
	} else {
		fmt.Println("DescribeVpcs Output:")
		fmt.Println(result)
		fmt.Println("QC Status: FAILED - please review VPC")
		fmt.Println()
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
		fmt.Println("DescribeSubnets Output:")
		fmt.Println(result)
		fmt.Println("QC Status: PASS")
		fmt.Println()
	} else {
		fmt.Println("DescribeSubnets Output:")
		fmt.Println(result)
		fmt.Println("QC Status: FAILED - please review Subnets")
		fmt.Println()
	}

}
