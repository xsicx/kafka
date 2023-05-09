package main

import (
	"github.com/spf13/cobra"
	"github.com/xsicx/kafka/internal/consumer"
	"github.com/xsicx/kafka/internal/producer"
)

func main() {
	var cobraCmd = &cobra.Command{}
	cobraCmd.AddCommand(produceMsgCmd(), consumeMsgCmd())

	must(cobraCmd.Execute())
}

func produceMsgCmd() *cobra.Command {
	var server, topic string
	cmd := &cobra.Command{
		Use:   "produce",
		Short: "Kafka message producer",
		Run: func(cmd *cobra.Command, args []string) {
			producer.Run(server, topic)
		},
	}

	cmd.Flags().StringVarP(&server, "server", "s", "", "bootstrap server")
	must(cmd.MarkFlagRequired("server"))
	cmd.Flags().StringVarP(&topic, "topic", "t", "", "topic name")
	must(cmd.MarkFlagRequired("topic"))

	return cmd
}

func consumeMsgCmd() *cobra.Command {
	var server, topic, group string
	cmd := &cobra.Command{
		Use:   "consume",
		Short: "Kafka message consumer",
		Run: func(cmd *cobra.Command, args []string) {
			consumer.Run(server, topic, group)
		},
	}

	cmd.Flags().StringVarP(&server, "server", "s", "", "bootstrap server")
	must(cmd.MarkFlagRequired("server"))
	cmd.Flags().StringVarP(&topic, "topic", "t", "", "topic name")
	must(cmd.MarkFlagRequired("topic"))
	cmd.Flags().StringVarP(&topic, "group", "g", "group1", "group")

	return cmd
}

func must(err error) {
	if err == nil {
		return
	}

	panic(err)
}
