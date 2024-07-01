name: {{ .ActionName }}
on: [push]
jobs:
  promotedChannel:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: deployix/deployed-github-actions@:{{ .Version }}
        if: always()
        with:
          promotionName: {{ .PromotionName }}
          githubPAT: {{ .GithubPATSecret }}