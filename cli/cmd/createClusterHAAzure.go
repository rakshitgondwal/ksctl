/*
Kubesimplify
@maintainer: 	Dipankar Das <dipankardas0115@gmail.com>
Anurag Kumar <contact.anurag7@gmail.com>
Avinesh Tripathi <avineshtripathi1@gmail.com>
*/

package cmd

import (
	log "github.com/kubesimplify/ksctl/api/logger"

	"github.com/kubesimplify/ksctl/api/azure"
	util "github.com/kubesimplify/ksctl/api/utils"
	"github.com/spf13/cobra"
)

var createClusterHAAzure = &cobra.Command{
	Use:   "ha-azure",
	Short: "Use to create a HA k3s cluster in Azure",
	Long: `It is used to create cluster with the given name from user. For example:

	ksctl create-cluster ha-azure <arguments to civo cloud provider>
	`,
	Run: func(cmd *cobra.Command, args []string) {
		isSet := cmd.Flags().Lookup("verbose").Changed
		logger := log.Logger{Verbose: true}
		if !isSet {
			logger.Verbose = false
		}

		payload := &azure.AzureProvider{
			ClusterName: azhcclusterName,
			HACluster:   true,
			Region:      azhcregion,
			Spec: util.Machine{
				Disk:                azhcsize,
				HAControlPlaneNodes: azhcnodeCCP,
				HAWorkerNodes:       azhcnodeCWP,
			},
		}
		err := payload.CreateCluster(logger)
		if err != nil {
			logger.Err(err.Error())
			return
		}
		logger.Info("CREATED CLUSTER", "")
	},
}

var (
	azhcclusterName string
	azhcnodeCWP     int
	azhcnodeCCP     int
	azhcsize        string
	azhcregion      string
)

func init() {
	createClusterCmd.AddCommand(createClusterHAAzure)
	createClusterHAAzure.Flags().StringVarP(&azhcclusterName, "name", "n", "", "Cluster name")
	createClusterHAAzure.Flags().StringVarP(&azhcsize, "node-size", "s", "Standard_F2s", "Node size")
	createClusterHAAzure.Flags().StringVarP(&azhcregion, "region", "r", "eastus", "Region")
	createClusterHAAzure.Flags().IntVarP(&azhcnodeCWP, "worker-nodes", "w", 1, "Number of worker Nodes")
	createClusterHAAzure.Flags().IntVarP(&azhcnodeCCP, "control-nodes", "c", 3, "Number of control Nodes")
	createClusterHAAzure.Flags().BoolP("verbose", "v", true, "for verbose output")
	createClusterHAAzure.MarkFlagRequired("name")
}
