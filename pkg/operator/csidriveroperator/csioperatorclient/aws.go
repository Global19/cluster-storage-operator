package csioperatorclient

import (
	"os"
	"strings"

	configv1 "github.com/openshift/api/config/v1"
)

const (
	AWSEBSCSIDriverName          = "ebs.csi.aws.com"
	envAWSEBSDriverOperatorImage = "AWS_EBS_DRIVER_OPERATOR_IMAGE"
	envAWSEBSDriverImage         = "AWS_EBS_DRIVER_IMAGE"
)

func GetAWSEBSCSIOperatorConfig() CSIOperatorConfig {
	pairs := []string{
		"${OPERATOR_IMAGE}", os.Getenv(envAWSEBSDriverOperatorImage),
		"${DRIVER_IMAGE}", os.Getenv(envAWSEBSDriverImage),
	}

	return CSIOperatorConfig{
		CSIDriverName:   AWSEBSCSIDriverName,
		ConditionPrefix: "AWSEBS",
		Platform:        configv1.AWSPlatformType,
		StaticAssets: []string{
			"csidriveroperators/aws-ebs/01_namespace.yaml",
			"csidriveroperators/aws-ebs/02_sa.yaml",
			"csidriveroperators/aws-ebs/03_role.yaml",
			"csidriveroperators/aws-ebs/04_rolebinding.yaml",
			"csidriveroperators/aws-ebs/05_clusterrole.yaml",
			"csidriveroperators/aws-ebs/06_clusterrolebinding.yaml",
		},
		CRAsset:         "csidriveroperators/aws-ebs/08_cr.yaml",
		DeploymentAsset: "csidriveroperators/aws-ebs/07_deployment.yaml",
		ImageReplacer:   strings.NewReplacer(pairs...),
		Optional:        false,
		/* For reference / experiments only. OpenShift does not support
		   update from OLM-based AWS EBS operator to CVO/CSO one.
		OLMOptions: &OLMOptions{
			OLMOperatorDeploymentName: "aws-ebs-csi-driver-operator",
			OLMPackageName:            "aws-ebs-csi-driver-operator",
			CRResource: schema.GroupVersionResource{
				Group:    "csi.openshift.io",
				Version:  "v1alpha1",
				Resource: "awsebsdrivers",
			},
		},
		*/
	}
}
