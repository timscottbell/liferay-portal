package cmd

import (
	"time"

	"github.com/spf13/cobra"
)

const (
	flagAWSBackupServiceAssumedIamRoleArn = "aws-backup-service-assumed-iam-role-arn"
	flagBackupVaultName                   = "backup-vault-name"
	flagRdsSnapshotIdOutputPath           = "rds-snapshot-id-output-path"
	flagRecoveryPointArn                  = "recovery-point-arn"
	flagS3RecoveryPointArnOutputPath      = "s3-recovery-point-arn-output-path"
	flagTargetBucketName                  = "target-bucket-name"
	flagWait                              = "wait"
	flagWaitTimeout                       = "wait-timeout"
)

func addAwsBackupServiceAssumedIamRoleArnFlag(cmd *cobra.Command) {
	usage := "The IAM role ARN that AWS backup service will assume for the " +
		"restore"

	cmd.Flags().StringVar(
		&awsBackupServiceAssumedIamRoleARN,
		flagAWSBackupServiceAssumedIamRoleArn,
		"",
		usage)

	cobra.CheckErr(cmd.MarkFlagRequired(flagAWSBackupServiceAssumedIamRoleArn))
}

func addBackupVaultNameFlag(cmd *cobra.Command) {
	usage := "The name of the AWS backup vault where the recovery point is " +
		"stored"

	cmd.Flags().StringVar(
		&backupVaultName,
		flagBackupVaultName,
		"",
		usage)

	cobra.CheckErr(cmd.MarkFlagRequired(flagBackupVaultName))
}

func addRDSSnapshotIdOutputPathFlag(cmd *cobra.Command) {
	cmd.Flags().StringVar(
		&rdsSnapshotIdOutputPath,
		flagRdsSnapshotIdOutputPath,
		"",
		"The path to write the RDS snapshot Id to")
}

func addRecoveryPointArnFlag(cmd *cobra.Command) {
	cmd.Flags().StringVar(
		&recoveryPointARN,
		flagRecoveryPointArn,
		"",
		"The ARN of the AWS backup recovery point to restore from")

	cobra.CheckErr(cmd.MarkFlagRequired(flagRecoveryPointArn))
}

func addS3RecoveryPointArnOutputPathFlag(cmd *cobra.Command) {
	cmd.Flags().StringVar(
		&s3RecoveryPointArnOutputPath,
		flagS3RecoveryPointArnOutputPath,
		"",
		"The path to write the S3 recovery point ARN to")
}

func addTargetBucketNameFlag(cmd *cobra.Command) {
	cmd.Flags().StringVar(
		&targetBucketName,
		flagTargetBucketName,
		"",
		"The name of the S3 bucket to restore the data to")

	cobra.CheckErr(cmd.MarkFlagRequired(flagTargetBucketName))
}

func addWaitFlag(cmd *cobra.Command) {
	cmd.Flags().BoolVar(
		&wait,
		flagWait,
		false,
		"Whether to wait for the restore job to complete before exiting")
}

func addWaitTimeoutFlag(cmd *cobra.Command) {
	cmd.Flags().DurationVar(
		&waitTimeout,
		flagWaitTimeout,
		2*time.Hour,
		"The maximum time to wait for the restore job to complete")
}

var (
	awsBackupServiceAssumedIamRoleARN string
	backupVaultName                   string
	recoveryPointARN                  string
	rdsSnapshotIdOutputPath           string
	s3RecoveryPointArnOutputPath      string
	targetBucketName                  string
	wait                              bool
	waitTimeout                       time.Duration
)