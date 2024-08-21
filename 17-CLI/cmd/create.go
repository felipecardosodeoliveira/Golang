/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/felipecardosodeoliveira/Golang/17-CLI/internal/database"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
// var createCmd1 = &cobra.Command{
// Use:   "create",
// Short: "A brief description of your command",
// Long:  `A longer description that spans multiple`,
// Run: func(cmd *cobra.Command, args []string) {
// 	db := GetDB()
// 	categoryDB := GetCategoryDB(db)
// 	name, _ := cmd.Flags().GetString("name")
// 	description, _ := cmd.Flags().GetString("description")
// 	categoryDB.Create(name, description)
// },
// RunE: createCategory(GetCategoryDB(GetDB())),
// }

func newCreateCmd(categoryDB database.Category) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "A brief description of your command",
		Long:  `A longer description that spans multiple`,
		RunE:  runCreate(categoryDB),
	}
}

func runCreate(categoryDB database.Category) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		_, err := categoryDB.Create(name, description)
		if err != nil {
			return err
		}
		return nil
	}
}

func init() {
	createCmd := newCreateCmd(GetCategoryDB(GetDB()))
	categoryCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("name", "n", "", "category name")
	createCmd.Flags().StringP("description", "d", "", "category description")
	createCmd.MarkFlagsRequiredTogether("name", "description")
}
