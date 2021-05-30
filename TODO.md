# TODO

## Vault Integration

Auth with Vault

Iterate through all Secrets with specified prefix
e.g. kv/api.lmhd.me/federate/slack/

For each secret with that prefix...

Secret will have schema like

```
{
	"user_id": "U1234567",
	"api_token": "xoxp-123467",
	"name_suffix": " (Lucy Davinhart)",
	"avatars": {
		"Lucy": "https://lmhd.keybase.pub/lucy.jpg",
		"Ivie": "https://lmhd.keybase.pub/ivie.jpg",
	},
	"default_avatar": {
		"https://lmhd.keybase.pub/system.png"
	}
}
```

* If "avatars.${fronter}" key present, use specific avatar for specific fronter
* else
	* If default_avatar key present, use that
	* Else, use avatar from api.lmhd.me/v1/front


* If "name_suffix" key present, suffix slack name with that

## Private Fronters

e.g. when PK Front is cleared, 


## Pronouns

* Unclear if https://github.com/slack-go/slack supports setting custom fields


## Integrate into api.lmhd.me

* Require an X.509 Cert from Vault, with certain parameters
* Convert this whole thing into a package which can be imported into API
	* Include client library for API