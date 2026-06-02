package vantage

import (
	"fmt"
	"testing"

	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/vantage-sh/terraform-provider-vantage/vantage/acctest"
)

func TestAccVantageFolderByName_basic(t *testing.T) {
	rName := sdkacctest.RandStringFromCharSet(10, sdkacctest.CharSetAlphaNum)
	folderTitle := fmt.Sprintf("tf-test-folder-%s", rName)
	resourceName := "vantage_folder.test"
	dataSourceName := "data.vantage_folder_by_name.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFolderByNameDataSourceConfig(folderTitle),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "token", resourceName, "token"),
					resource.TestCheckResourceAttr(dataSourceName, "title", folderTitle),
					resource.TestCheckResourceAttrPair(dataSourceName, "workspace_token", resourceName, "workspace_token"),
				),
			},
		},
	})
}

func TestAccVantageFolderByName_withWorkspaceFilter(t *testing.T) {
	rName := sdkacctest.RandStringFromCharSet(10, sdkacctest.CharSetAlphaNum)
	folderTitle := fmt.Sprintf("tf-test-folder-ws-%s", rName)
	resourceName := "vantage_folder.test"
	dataSourceName := "data.vantage_folder_by_name.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFolderByNameWithWorkspaceFilterConfig(folderTitle),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "token", resourceName, "token"),
					resource.TestCheckResourceAttr(dataSourceName, "title", folderTitle),
					resource.TestCheckResourceAttrPair(dataSourceName, "workspace_token", resourceName, "workspace_token"),
				),
			},
		},
	})
}

func testAccFolderByNameDataSourceConfig(title string) string {
	return fmt.Sprintf(`
resource "vantage_folder" "test" {
  title = %[1]q
}

data "vantage_folder_by_name" "test" {
  title      = %[1]q
  depends_on = [vantage_folder.test]
}
`, title)
}

func testAccFolderByNameWithWorkspaceFilterConfig(title string) string {
	return fmt.Sprintf(`
resource "vantage_folder" "test" {
  title = %[1]q
}

data "vantage_folder_by_name" "test" {
  title           = %[1]q
  workspace_token = vantage_folder.test.workspace_token
  depends_on      = [vantage_folder.test]
}
`, title)
}
