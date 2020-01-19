# Time Exporter

## Motivation
A time exporter provides a simple method of modifying time-based alert routing - essentially providing the functionality of querying metric values and deciding if the respective alert should be fired during the time period.

### Use-Case
Modifying alerts to only fire during office hours [See: https://github.com/prometheus/alertmanager/issues/876].
