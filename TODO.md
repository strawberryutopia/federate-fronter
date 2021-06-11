# TODO

## Vault Integration

Auth with Vault

For members:
https://vault.lmhd.me/ui/vault/secrets/kv/show/lmhd/api/federate/slack/members

```
{
  "members": {
    // _default is used if no member-specific config is set
    "_default": {
      "use_api_avatar": true,
      "use_api_name": true,
      // i.e. by default, just use whichever
      // avatar and display name the API returns

      "workspace_overrides": {
        "work": {
          // Fallback avatar if no workspace-specific avatar set
          "avatar": "https://lmhd.keybase.pub/system.png",

          // Fallback avatar if no workspace-specific name
          "name": "Lucy Davinhart (Maybe)"
        }
      }
    },
    "hol": {
      // No name or avatar set, so use the one from _default

      "workspace_overrides": {
        "work": {
          "name": "Hol (Lucy Davinhart)"
        }
      }
    },
    "ivie": {
      // No name or avatar set, so use the one from _default

      "workspace_overrides": {
        "work": {
      	  "avatar": "https://lmhd.keybase.pub/ivie.jpg",
          "name": "Ivie (Lucy Davinhart)"
        }
      }
    },
    "jesper": {
      // No name or avatar set, so use the one from _default

      "workspace_overrides": {
        "work": {
          "name": "Jesper (Lucy Davinhart)"
        }
      }
    },
    "lucy": {
      // No name or avatar set, so use the one from _default

      "workspace_overrides": {
        "work": {
      	  "avatar": "https://lmhd.keybase.pub/lucy.jpg",
          "name": "Lucy (Lucy Davinhart)"
        }
      }
    }
  }
}
```

When determining what name and avatar to use:
```
Check if there is a key for a specific member (e.g. Lucy)
	Yes: Check if there is a specific workspace override for that member
		Yes: Use workspace override keys, with fallback to _default for missing keys (e.g. avatar)
	No: use _default

_default
	Check if there is a specific workspace override
		Yes: Use workspace override keys, with fallback to _default for missing keys
	No: use _default
```


Tokens:
https://vault.lmhd.me/ui/vault/secrets/kv/show/lmhd/api/federate/slack/tokens

```
{
  "pcw": "xoxp-123467",
  "strawb": "xoxp-123467"
}
```

and UIDs:
https://vault.lmhd.me/ui/vault/secrets/kv/show/lmhd/api/federate/slack/user_ids

```
{
  "pcw": "U1234567",
  "strawb": "U1234567"
}
```

## Private Fronters

e.g. when PK Front is cleared, 


## Pronouns

* Unclear if https://github.com/slack-go/slack supports setting custom fields


## Integrate into api.lmhd.me

* Require an X.509 Cert from Vault, with certain parameters
* Convert this whole thing into a package which can be imported into API
	* Include client library for API?