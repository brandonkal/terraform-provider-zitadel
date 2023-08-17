package default_login_texts_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/zitadel/zitadel-go/v2/pkg/client/zitadel/admin"

	"github.com/zitadel/terraform-provider-zitadel/zitadel/v2/helper/test_utils"
)

func TestAccDefaultLoginTexts(t *testing.T) {
	resourceName := "zitadel_default_login_texts"
	initialProperty := "initialtitle"
	updatedProperty := "updatedtitle"
	language := "en"
	frame, err := test_utils.NewInstanceTestFrame(resourceName)
	if err != nil {
		t.Fatalf("setting up test context failed: %v", err)
	}
	test_utils.RunLifecyleTest[string](
		t,
		frame.BaseTestFrame,
		func(configProperty, _ string) string {
			return fmt.Sprintf(`
resource "%s" "%s" {
  language    = "%s"

  email_verification_done_text = {
    cancel_button_text = "example"
    description        = "example"
    login_button_text  = "example"
    next_button_text   = "example"
    title              = "%s"
  }
  email_verification_text = {
    code_label         = "example"
    description        = "example"
    next_button_text   = "example"
    resend_button_text = "example"
    title              = "example"
  }
  external_registration_user_overview_text = {
    back_button_text      = "example"
    description           = "example"
    email_label           = "example"
    firstname_label       = "example"
    language_label        = "example"
    lastname_label        = "example"
    next_button_text      = "example"
    nickname_label        = "example"
    phone_label           = "example"
    privacy_link_text     = "example"
    title                 = "example"
    tos_and_privacy_label = "example"
    tos_confirm           = "example"
    tos_confirm_and       = "example"
    tos_link_text         = "example"
    username_label        = "example"
  }
  external_user_not_found_text = {
    auto_register_button_text = "example"
    description               = "example"
    link_button_text          = "example"
    privacy_link_text         = "example"
    title                     = "example"
    tos_and_privacy_label     = "example"
    tos_confirm               = "example"
    tos_confirm_and           = "example"
    tos_link_text             = "example"
  }
  footer_text = {
    help           = "example"
    privacy_policy = "example"
    tos            = "example"
  }
  init_mfa_done_text = {
    cancel_button_text = "example"
    description        = "example"
    next_button_text   = "example"
    title              = "example"
  }
  init_mfa_otp_text = {
    cancel_button_text = "example"
    code_label         = "example"
    description        = "example"
    description_otp    = "example"
    next_button_text   = "example"
    secret_label       = "example"
    title              = "example"
  }
  init_mfa_prompt_text = {
    description      = "example"
    next_button_text = "example"
    otp_option       = "example"
    skip_button_text = "example"
    title            = "example"
    u2f_option       = "example"
  }
  init_mfa_u2f_text = {
    description                = "example"
    error_retry                = "example"
    not_supported              = "example"
    register_token_button_text = "example"
    title                      = "example"
    token_name_label           = "example"
  }
  init_password_done_text = {
    cancel_button_text = "example"
    description        = "example"
    next_button_text   = "example"
    title              = "example"
  }
  init_password_text = {
    code_label                 = "example"
    description                = "example"
    new_password_confirm_label = "example"
    new_password_label         = "example"
    next_button_text           = "example"
    resend_button_text         = "example"
    title                      = "example"
  }
  initialize_done_text = {
    cancel_button_text = "example"
    description        = "example"
    next_button_text   = "example"
    title              = "example"
  }
  initialize_user_text = {
    code_label                 = "example"
    description                = "example"
    new_password_confirm_label = "example"
    new_password_label         = "example"
    next_button_text           = "example"
    resend_button_text         = "example"
    title                      = "example"
  }
  linking_user_done_text = {
    cancel_button_text = "example"
    description        = "example"
    next_button_text   = "example"
    title              = "example"
  }
  login_text = {
    description                 = "example"
    description_linking_process = "example"
    external_user_description   = "example"
    login_name_label            = "example"
    login_name_placeholder      = "example"
    next_button_text            = "example"
    register_button_text        = "example"
    title                       = "example"
    title_linking_process       = "example"
    user_must_be_member_of_org  = "example"
    user_name_placeholder       = "example"
  }
  logout_text = {
    description       = "example"
    login_button_text = "example"
    title             = "example"
  }
  mfa_providers_text = {
    choose_other = "example"
    otp          = "example"
    u2f          = "example"
  }
  password_change_done_text = {
    description      = "example"
    next_button_text = "example"
    title            = "example"
  }
  password_change_text = {
    cancel_button_text         = "example"
    description                = "example"
    new_password_confirm_label = "example"
    new_password_label         = "example"
    next_button_text           = "example"
    old_password_label         = "example"
    title                      = "example"
  }
  password_reset_done_text = {
    description      = "example"
    next_button_text = "example"
    title            = "example"
  }
  password_text = {
    back_button_text = "example"
    confirmation     = "example"
    description      = "example"
    has_lowercase    = "example"
    has_number       = "example"
    has_symbol       = "example"
    has_uppercase    = "example"
    min_length       = "example"
    next_button_text = "example"
    password_label   = "example"
    reset_link_text  = "example"
    title            = "example"
  }
  passwordless_prompt_text = {
    description              = "example"
    description_init         = "example"
    next_button_text         = "example"
    passwordless_button_text = "example"
    skip_button_text         = "example"
    title                    = "example"
  }
  passwordless_registration_done_text = {
    cancel_button_text = "example"
    description        = "example"
    description_close  = "example"
    next_button_text   = "example"
    title              = "example"
  }
  passwordless_registration_text = {
    description                = "example"
    error_retry                = "example"
    not_supported              = "example"
    register_token_button_text = "example"
    title                      = "example"
    token_name_label           = "example"
  }
  passwordless_text = {
    description                = "example"
    error_retry                = "example"
    login_with_pw_button_text  = "example"
    not_supported              = "example"
    title                      = "example"
    validate_token_button_text = "example"
  }
  registration_option_text = {
    description                = "example"
    external_login_description = "example"
    title                      = "example"
    user_name_button_text      = "example"
  }
  registration_org_text = {
    description            = "example"
    email_label            = "example"
    firstname_label        = "example"
    lastname_label         = "example"
    orgname_label          = "example"
    password_confirm_label = "example"
    password_label         = "example"
    privacy_link_text      = "example"
    save_button_text       = "example"
    title                  = "example"
    tos_and_privacy_label  = "example"
    tos_confirm            = "example"
    tos_confirm_and        = "example"
    tos_link_text          = "example"
    username_label         = "example"
  }
  registration_user_text = {
    back_button_text         = "example"
    description              = "example"
    description_org_register = "example"
    email_label              = "example"
    firstname_label          = "example"
    gender_label             = "example"
    language_label           = "example"
    lastname_label           = "example"
    next_button_text         = "example"
    password_confirm_label   = "example"
    password_label           = "example"
    privacy_link_text        = "example"
    title                    = "example"
    tos_and_privacy_label    = "example"
    tos_confirm              = "example"
    tos_confirm_and          = "example"
    tos_link_text            = "example"
    username_label           = "example"
  }
  select_account_text = {
    description                 = "example"
    description_linking_process = "example"
    other_user                  = "example"
    session_state_active        = "example"
    session_state_inactive      = "example"
    title                       = "example"
    title_linking_process       = "example"
    user_must_be_member_of_org  = "example"
  }
  success_login_text = {
    auto_redirect_description = "example"
    next_button_text          = "example"
    redirected_description    = "example"
    title                     = "example"
  }
  username_change_done_text = {
    description      = "example"
    next_button_text = "example"
    title            = "example"
  }
  username_change_text = {
    cancel_button_text = "example"
    description        = "example"
    next_button_text   = "example"
    title              = "example"
    username_label     = "example"
  }
  verify_mfa_otp_text = {
    code_label       = "example"
    description      = "example"
    next_button_text = "example"
    title            = "example"
  }
  verify_mfa_u2f_text = {
    description         = "example"
    error_retry         = "example"
    not_supported       = "example"
    title               = "example"
    validate_token_text = "example"
  }
}`, resourceName, frame.UniqueResourcesID, language, configProperty)
		},
		initialProperty, updatedProperty,
		"", "",
		true,
		checkRemoteProperty(frame, language),
		regexp.MustCompile(`^en$`),
		// When deleted, the default should be returned
		checkRemoteProperty(frame, language)(""),
		nil, nil, "", "",
	)
}

func checkRemoteProperty(frame *test_utils.InstanceTestFrame, lang string) func(string) resource.TestCheckFunc {
	return func(expect string) resource.TestCheckFunc {
		return func(state *terraform.State) error {
			remoteResource, err := frame.GetCustomLoginTexts(frame, &admin.GetCustomLoginTextsRequest{Language: lang})
			if err != nil {
				return err
			}
			actual := remoteResource.GetCustomText().GetEmailVerificationDoneText().GetTitle()
			if actual != expect {
				return fmt.Errorf("expected %s, but got %s", expect, actual)
			}
			return nil
		}
	}
}
