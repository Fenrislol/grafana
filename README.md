![Grafana](docs/logo-horizontal.png)

The open-source platform for monitoring and observability.

[![License](https://img.shields.io/github/license/grafana/grafana)](LICENSE)
[![Circle CI](https://img.shields.io/circleci/build/gh/grafana/grafana)](https://circleci.com/gh/grafana/grafana)
[![Go Report Card](https://goreportcard.com/badge/gitlab.com/digitalizm/grafana)](https://goreportcard.com/report/gitlab.com/digitalizm/grafana)


## Update grafana
go get -u  gitlab.com/digitalizm/grafana
cd /home/webmaster/go/src/gitlab.com/digitalizm/grafana

go run build.go setup
go run build.go build

# Build the Frontend Assets (first time)
# For this you need nodejs (v.6+).

npm install -g yarn
yarn install --pure-lockfile
yarn start


## Run
cd /home/webmaster/go/src/gitlab.com/digitalizm/grafana
./bin/linux-amd64/grafana-server --config=./conf/grafana.ini

================

## Installation
Head to [docs.grafana.org](http://docs.grafana.org/installation/) for documentation or [download](https://grafana.com/get) to get the latest release.
=======
Grafana allows you to query, visualize, alert on and understand your metrics no matter where they are stored. Create, explore, and share dashboards with your team and foster a data driven culture:

- **Visualize:** Fast and flexible client side graphs with a multitude of options. Panel plugins for many different way to visualize metrics and logs.
- **Dynamic Dashboards:** Create dynamic & reusable dashboards with template variables that appear as dropdowns at the top of the dashboard.
- **Explore Metrics:** Explore your data through ad-hoc queries and dynamic drilldown. Split view and compare different time ranges, queries and data sources side by side.
- **Explore Logs:** Experience the magic of switching from metrics to logs with preserved label filters. Quickly search through all your logs or streaming them live.
- **Alerting:** Visually define alert rules for your most important metrics. Grafana will continuously evaluate and send notifications to systems like Slack, PagerDuty, VictorOps, OpsGenie.
- **Mixed Data Sources:** Mix different data sources in the same graph! You can specify a data source on a per-query basis. This works for even custom datasources.

## Get started

- [Get Grafana](https://grafana.com/get)
- [Installation guides](http://docs.grafana.org/installation/)

Unsure if Grafana is for you? Watch Grafana in action on [play.grafana.org](https://play.grafana.org/)!

## Documentation

The Grafana documentation is available at [grafana.com/docs](https://grafana.com/docs/).

## Contributing

If you're interested in contributing to the Grafana project:

- Start by reading the [Contributing guide](/CONTRIBUTING.md).
- Learn how to set up your local environment, in our [Developer guide](/contribute/developer-guide.md).
- Explore our [beginner-friendly issues](https://gitlab.com/digitalizm/grafana/issues?q=is%3Aopen+is%3Aissue+label%3A%22beginner+friendly%22).

## Get involved

- Follow [@grafana on Twitter](https://twitter.com/grafana/)
- Read and subscribe to the [Grafana blog](https://grafana.com/blog/)
- If you have a specific question, check out our [discussion forums](https://community.grafana.com/).
- For general discussions, join us on the [official Slack](http://slack.raintank.io/) team.

## License

Grafana is distributed under the [Apache 2.0 License](https://gitlab.com/digitalizm/grafana/blob/master/LICENSE).
