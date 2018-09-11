package cmd

import (
	"fmt"
	"io"
	"log"

	"github.com/dimitrovvlado/grafctl/grafana"
	"github.com/gosuri/uitable"

	"github.com/spf13/cobra"
)

type datasourceListCmd struct {
	client *grafana.Client
	out    io.Writer
	output string
}

func newDatasourceListCommand(client *grafana.Client, out io.Writer) *cobra.Command {
	get := &datasourceListCmd{
		client: client,
		out:    out,
	}
	getDatasourcesCmd := &cobra.Command{
		Use:     "datasources",
		Aliases: []string{"ds"},
		Short:   "Display one or many datasources",
		RunE: func(cmd *cobra.Command, args []string) error {
			ensureClient(get.client)
			return get.run()
		},
	}
	f := getDatasourcesCmd.Flags()
	f.StringVarP(&get.output, "output", "o", "", "Output the specified format (|json)")
	return getDatasourcesCmd
}

func (i *datasourceListCmd) run() error {
	ds, err := i.client.ListDatasources()
	if err != nil {
		log.Fatalln(err)
	}

	//TODO extract as flag
	var colWidth uint = 60
	formatter := func() string {
		if ds == nil || len(ds) == 0 {
			return fmt.Sprintf("No datasources found.")
		}
		table := uitable.New()
		table.MaxColWidth = colWidth
		table.AddRow("ID", "NAME", "TYPE", "ACCESS", "URL")
		for _, lr := range ds {
			table.AddRow(lr.ID, lr.Name, lr.Type, lr.Access, lr.URL)
		}
		return fmt.Sprintf("%s%s", table.String(), "\n")
	}

	result, err := formatResult(i.output, ds, formatter)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintln(i.out, result)

	return nil
}
