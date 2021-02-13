package rollbar

// TODO: comment out code after figuring out a better way to define ID for the PAT resource.
//func TestAccRollbarProjectAccessToken_importBasic(t *testing.T) {
//	projectName := fmt.Sprintf("project-tftest-%s", acctest.RandString(10))
//	tokenName := fmt.Sprintf("token-tftest-%s", acctest.RandString(10))
//
//	resource.Test(t, resource.TestCase{
//		PreCheck: func() {
//			testAccPreCheck(t)
//		},
//		Providers: testAccProviders,
//		Steps: []resource.TestStep{
//			{
//				Config: testAccCheckRollbarProjectAccessToken_basic(projectName, tokenName),
//			},
//			{
//				ResourceName:            "rollbar_project_access_token.foobar",
//				ImportStateIdFunc:       testAccRollbarProjectAccessTokenImportStateIdFunc("rollbar_project_access_token.foobar"),
//				ImportState:             true,
//				ImportStateVerify:       false,
//				ImportStateVerifyIgnore: []string{"id"},
//			},
//		},
//	})
//}
//
//func testAccRollbarProjectAccessTokenImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
//	return func(s *terraform.State) (string, error) {
//		rs, ok := s.RootModule().Resources[resourceName]
//		if !ok {
//			return "", fmt.Errorf("Not found: %s", resourceName)
//		}
//
//		return fmt.Sprintf("%s:%s", rs.Primary.Attributes["project_id"],
//			rs.Primary.Attributes["access_token"]), nil
//	}
//}
