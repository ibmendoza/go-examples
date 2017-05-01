### Tracing / Debugging

- honeycomb.io
- backtrace.io
- https://www.circonus.com/2016/07/rapid-resolution-right-tools/

### Monitoring

[Turn expected exceptions to metrics](http://yellerapp.com/posts/2015-06-01-getting-to-exception-zero.html)

**Logging**

What to log according to Dave Cheney?

via http://dave.cheney.net/2015/11/05/lets-talk-about-logging

    Things that developers care about when they are developing or debugging software.
    Things that users care about when using your software.
 
What to log according to Peter Bourgon?

via 

- http://peter.bourgon.org/go-in-production/#logging-and-telemetry
- http://peter.bourgon.org/blog/2016/02/07/logging-v-instrumentation.html

In the end, we settled on plain package log. It works because we only log actionable information. That means serious, panic-level errors that need to be addressed by humans, or structured data that will be consumed by other machines. 

- https://github.com/keep94/weblogs

**Metrics**

- http://peter.bourgon.org/go-in-production/#logging-and-telemetry

Everything else emitted by a running process we consider telemetry. Request response times, QPS, runtime errors, queue depths, and so on. And telemetry basically operates in one of two models: push and pull.

- Push means emitting metrics to a known external system. For example, Graphite, Statsd, and AirBrake work this way.

- Pull means exposing metrics at some known location, and allowing an external system to scrape them. For example, expvar and Prometheus work this way. (Maybe there are others?)
 
**Metrics Packages**

- https://github.com/sourcegraph/appdash
- https://github.com/bosun-monitor/bosun
- https://github.com/mackerelio
- https://github.com/VividCortex
- https://github.com/ekanite/ekanite
- https://github.com/adg/sched
- https://github.com/VividCortex/pm
- https://github.com/VividCortex/pm-web
- https://github.com/VividCortex/trace
- https://github.com/vrischmann/go-metrics-influxdb
- https://github.com/rcrowley/go-metrics
- https://github.com/bmhatfield/go-runtime-metrics
