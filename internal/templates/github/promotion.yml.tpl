on: [push]

jobs:
  new_push_job:
    runs-on: ubuntu-latest
    name: test
    steps:
    - name: Publish to slack channel via bot token
      id: slack
      uses: slackapi/slack-github-action@v1.24.0
      with:
        channel-id: 'SLACK_CHANNEL_ID' # ID of Slack Channel you want to post to
        slack-message: 'posting from a github action!' # The message you want to post
      env:
        SLACK_BOT_TOKEN: 