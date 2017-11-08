package main

import (
	"fmt"
	//"encoding/json"
	//"database/sql"
	//"io/ioutil"

	//_ "github.com/lib/pq"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/redshift"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/awserr"
	//"github.com/user/docker/vendor/github.com/docker/go/canonical/json"
)

type ClusterResponse struct {
	Response *redshift.DescribeClustersOutput `json:"json_string"`
}

func main() {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials("", "", ""),
	})
	if err != nil {
		fmt.Println("Error creating session", err)
	}

	svc := redshift.New(sess)

	params := &redshift.CreateClusterInput{
		ClusterIdentifier:                aws.String("xyz"), // Required
		MasterUserPassword:               aws.String("xyz"), // Required
		MasterUsername:                   aws.String("xyz"), // Required
		NodeType:                         aws.String("xyz"), // Required
		//AdditionalInfo:                   aws.String("String"),
		AllowVersionUpgrade:              aws.Bool(true),
		AutomatedSnapshotRetentionPeriod: aws.Int64(1),
		//AvailabilityZone:                 aws.String("String"),
		//ClusterParameterGroupName:        aws.String("String"),
		//ClusterSecurityGroups: []*string{
		//	aws.String("String"), // Required
		//	// More values...
		//},
		//ClusterSubnetGroupName:         aws.String("String"),
		ClusterType:                    aws.String("single-node"),
		//ClusterVersion:                 aws.String("String"),
		//DBName:                         aws.String("String"),
		//ElasticIp:                      aws.String("String"),
		Encrypted:                      aws.Bool(false),
		EnhancedVpcRouting:             aws.Bool(true),
		//HsmClientCertificateIdentifier: aws.String("String"),
		//HsmConfigurationIdentifier:     aws.String("String"),
		//IamRoles: []*string{
		//	aws.String("coolsb"), // Required
		//	// More values...
		//},
		//KmsKeyId:      aws.String("String"),
		//NumberOfNodes: aws.Int64(1),
		//Port:          aws.Int64(1),
		//PreferredMaintenanceWindow: aws.String("Mon:"),
		PubliclyAccessible:         aws.Bool(true),
		//Tags: []*redshift.Tag{
		//	{ // Required
		//		Key:   aws.String("String"),
		//		Value: aws.String("String"),
		//	},
		//	// More values...
		//},
		//VpcSecurityGroupIds: []*string{
		//	aws.String("String"), // Required
		//	// More values...
		//},
	}
	//fmt.Println(params.ClusterIdentifier)
	resp, err := svc.CreateCluster(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	//b, _ := json.Marshal(resp)
	//fmt.Printf("%s", b)
	fmt.Println(resp)


/*

	params := &redshift.DescribeClustersInput{
		ClusterIdentifier:                aws.String("testredshift"),
	}
	//resp, _ := svc.DescribeClusters(params)
	//p := ClusterResponse{resp}
	//b, _ := json.Marshal(p)
	//fmt.Println(string(b))
	resp, _ := describeRedshiftCluster(*params.ClusterIdentifier, svc)
	//fmt.Printf("%s	%d\n", *resp.Clusters[0].Endpoint.Address, *resp.Clusters[0].Endpoint.Port)
	fmt.Println(resp)
*/
	/*
	if yes := DidAwsCallSucceed(err); yes {
		fmt.Println(*resp.Clusters[0].ClusterStatus)
	} else {
		if awsErr, ok := err.(awserr.Error); ok {
			if awsErr.Code() == "ClusterNotFound" {
				fmt.Println("Cluster has been deleted")
			}
		} else {
			fmt.Println(awsErr.Code())
		}
	}
/*
	//fmt.Println(*resp.Clusters[0].ClusterStatus)
	// Pretty-print the response data.
	//fmt.Println("Response Status Code: ", resp )
	//fmt.Println(resp.String())


	params := &redshift.DeleteClusterInput{
		ClusterIdentifier:                aws.String("coolsb-cluster01"),
		SkipFinalClusterSnapshot:          aws.Bool(true),
	}
	resp, err := svc.DeleteCluster(params)
	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Printf("%T", resp)
*/

/*
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
		Credentials: credentials.NewStaticCredentials("", "",""),
	})
	if err != nil {
		fmt.Println("Error creating session", err)
	}

	svc := redshift.New(sess)

	params := &redshift.CreateClusterInput{
		//ClusterIdentifier:                aws.String(""), // Required
		ClusterIdentifier:                aws.String(""), // Required
		MasterUserPassword:               aws.String(""), // Required
		MasterUsername:                   aws.String(""), // Required
		NodeType:                         aws.String(""), // Required
		//AdditionalInfo:                   aws.String("String"),
		AllowVersionUpgrade:              aws.Bool(true),
		AutomatedSnapshotRetentionPeriod: aws.Int64(1),
		AvailabilityZone:                 aws.String("us-west-2a"),
		//ClusterParameterGroupName:        aws.String("String"),
		//ClusterSecurityGroups: []*string{
		//	aws.String("String"), // Required
		//	// More values...
		//},
		ClusterSubnetGroupName:         aws.String(""),
		ClusterType:                    aws.String("multi-node"),
		//ClusterVersion:                 aws.String("String"),
		//DBName:                         aws.String("String"),
		//ElasticIp:                      aws.String("String"),
		Encrypted:                      aws.Bool(false),
		EnhancedVpcRouting:             aws.Bool(true),
		//HsmClientCertificateIdentifier: aws.String("String"),
		//HsmConfigurationIdentifier:     aws.String("String"),
		//IamRoles: []*string{
		//	aws.String("coolsb"), // Required
		//	// More values...
		//},
		//KmsKeyId:      aws.String("String"),
		NumberOfNodes: aws.Int64(2),
		//Port:          aws.Int64(1),
		PreferredMaintenanceWindow: aws.String("Mon:17:00-Mon:17:30"),
		PubliclyAccessible:         aws.Bool(false),
		//Tags: []*redshift.Tag{
		//	{ // Required
		//		Key:   aws.String("String"),
		//		Value: aws.String("String"),
		//	},
		//	// More values...
		//},
		//VpcSecurityGroupIds: []*string{
		//	aws.String("String"), // Required
		//	// More values...
		//},
	}
	//fmt.Println(params.ClusterIdentifier)
	resp, err := svc.CreateCluster(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Printf("%T", resp)
	*/
}


func DidAwsCallSucceed(err error) bool {
	// TODO Eventually return a formatted error object.
	if err != nil {
		if _, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			//logger.Error(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if _, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				//logger.Error(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should always return an
			// error which satisfies the awserr.Error interface.
			//logger.Error(err.Error())
		}
		return false
	}
	return true
}



func describeRedshiftCluster(clusterID string, svc *redshift.Redshift) (resp *redshift.DescribeClustersOutput, err error) {
	params := &redshift.DescribeClustersInput{
		ClusterIdentifier:         aws.String(clusterID),
	}
	resp, err = svc.DescribeClusters(params)
	// Decide if AWS service call was successful
	if yes := DidAwsCallSucceed(err); yes {
		return
	} else {
		return
	}
}

func t2() {
	s3Access()
}

func s3Access() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
		Credentials: credentials.NewStaticCredentials("", "", ""),
	})
	if err != nil {
		fmt.Println("Error creating session", err)
	}

	svc := s3.New(sess)

	params := &s3.GetBucketAclInput{
		Bucket: aws.String(""), // Required
	}
	resp, err := svc.GetBucketAcl(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

