terraform {
  required_providers {
    speakeasy = {
      source  = "speakeasy-api/speakeasy"
      version = "0.0.4"
    }
  }
}

variable "api_key" {
  type = string
}


provider "speakeasy" {
  api_key = var.api_key
}

resource "speakeasy_plugin" "test_plugin" {
  plugin_id    = "test_plugin"
  code         = <<EOM

/**
 * @param {RequestData} data
 */
const handler = (data) => {
    // Documentation for plugins can be found at
    // https://docs.speakeasyapi.dev/docs/using-speakeasy/api-portal-tools/plugins/
    
    if(data.http.info.status >= 400) {
      data.addMetadata("error", "true")
    }
}
EOM
  title        = "test plugin via terraform"
  workspace_id = "tr-test"
}
