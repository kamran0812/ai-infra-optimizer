package cloud

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type AWSProvider struct {
	session *session.Session
}

func NewAWSProvider() (*AWSProvider, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"), // Change this to your desired region
	})
	if err != nil {
		return nil, err
	}
	return &AWSProvider{session: sess}, nil
}

func (p *AWSProvider) GetName() string {
	return "AWS"
}

func (p *AWSProvider) GetResourceUsage() ([]ResourceUsage, error) {
	ec2Svc := ec2.New(p.session)
	cwSvc := cloudwatch.New(p.session)

	instances, err := ec2Svc.DescribeInstances(nil)
	if err != nil {
		return nil, err
	}

	var usages []ResourceUsage
	for _, reservation := range instances.Reservations {
		for _, instance := range reservation.Instances {
			cpu, err := p.getInstanceCPUUtilization(*instance.InstanceId, cwSvc)
			if err != nil {
				return nil, err
			}

			usages = append(usages, ResourceUsage{
				Provider:   "AWS",
				ResourceID: *instance.InstanceId,
				Type:       *instance.InstanceType,
				CPU:        cpu,
				Memory:     0, // AWS doesn't provide memory utilization out of the box
				Timestamp:  time.Now(),
			})
		}
	}

	return usages, nil
}

func (p *AWSProvider) getInstanceCPUUtilization(instanceID string, cwSvc *cloudwatch.CloudWatch) (float64, error) {
	input := &cloudwatch.GetMetricStatisticsInput{
		Namespace:  aws.String("AWS/EC2"),
		MetricName: aws.String("CPUUtilization"),
		Dimensions: []*cloudwatch.Dimension{
			{
				Name:  aws.String("InstanceId"),
				Value: aws.String(instanceID),
			},
		},
		StartTime:  aws.Time(time.Now().Add(-1 * time.Hour)),
		EndTime:    aws.Time(time.Now()),
		Period:     aws.Int64(3600),
		Statistics: []*string{aws.String("Average")},
	}

	result, err := cwSvc.GetMetricStatistics(input)
	if err != nil {
		return 0, err
	}

	if len(result.Datapoints) == 0 {
		return 0, fmt.Errorf("no datapoints found for instance %s", instanceID)
	}

	return *result.Datapoints[0].Average, nil
}
