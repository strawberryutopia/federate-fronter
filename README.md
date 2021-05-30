# Proof of Concept Standalone Fronter Federation


<a href="https://slack.com/oauth/v2/authorize?client_id=4525723959.1446028368468&scope=&user_scope=users.profile:read,users.profile:write"><img alt="Add to Slack" height="40" width="139" src="https://platform.slack-edge.com/img/add_to_slack.png" srcSet="https://platform.slack-edge.com/img/add_to_slack.png 1x, https://platform.slack-edge.com/img/add_to_slack@2x.png 2x" /></a>

See: https://api.slack.com/start/distributing#multi_workspace_apps



## Exchanging a temporary authorization code for an access token

From: https://api.slack.com/authentication/oauth-v2#exchanging

If all's well, a user goes through the Slack app installation UI and okays your app with all the scopes that it requests. After that happens, Slack redirects the user back to your specified Redirect URL.

Parse the HTTP request that lands at your Redirect URL for a code field. That's your temporary authorization code, which expires after ten minutes.

Check the state parameter if you sent one along with your initial user redirect. If it doesn't match what you sent, consider the authorization a forgery.

Now you just have to exchange the code for an access token. You'll do this by calling the oauth.v2.access method:

```
curl -F code=1234 -F client_id=4525723959.1446028368468 -F client_secret=ABCDEFGH https://slack.com/api/oauth.v2.access
```

Note: If you initiate an install with the v2/authorize URL, it must be completed with oauth.v2.access, not the old oauth.access method.